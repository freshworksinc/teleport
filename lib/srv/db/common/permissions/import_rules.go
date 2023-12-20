// Teleport
// Copyright (C) 2023  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package permissions

import (
	"regexp"
	"sort"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/api/utils"
	utils2 "github.com/gravitational/teleport/lib/utils"
)

// ApplyDatabaseObjectImportRules applies the given set of rules onto a set of objects coming from a same database.
// Returns a fresh copy of a subset of supplied objects, filtered and modified.
// For the object to be returned, it must match at least one rule.
// The modification consists of application of extra labels, per matching mappings.
func ApplyDatabaseObjectImportRules(rules []types.DatabaseObjectImportRule, database types.Database, objs []types.DatabaseObject) []types.DatabaseObject {
	// sort: rules with higher priorities are applied last.
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].Priority() < rules[j].Priority()
	})

	// filter rules: keep those with matching labels
	// we only need mappings from the rules, so extract those.
	var mappings []types.DatabaseObjectImportRuleMapping
	for _, rule := range rules {
		if types.MatchLabels(database, rule.DatabaseLabels()) {
			mappings = append(mappings, rule.Mappings()...)
		}
	}

	// anything to do?
	if len(mappings) == 0 {
		return nil
	}

	var out []types.DatabaseObject

	// find all objects that match any of the rules
	for _, obj := range objs {
		var objClone types.DatabaseObject

		// apply each mapping in order.
		for _, mapping := range mappings {
			// the matching is applied to the object spec; existing object labels does not matter
			if databaseObjectImportMatch(mapping.Match, obj.GetSpec()) {
				if objClone == nil {
					objClone = obj.Copy()
				}

				// mapping applies additional (static) labels
				labels := objClone.GetAllLabels()
				if labels == nil {
					labels = map[string]string{}
				}
				for k, v := range mapping.AddLabels {
					labels[k] = v
				}
				objClone.SetStaticLabels(labels)
			}
		}

		if objClone != nil {
			out = append(out, objClone)
		}
	}

	return out
}

func matchPattern(pattern, value string) bool {
	matched, _ := regexp.MatchString("^"+utils2.GlobToRegexp(pattern)+"$", value)
	return matched
}

func databaseObjectImportMatch(match types.DatabaseObjectImportMatch, spec types.DatabaseObjectSpec) bool {
	matchAny := func(patterns []string, value string) bool {
		return utils.Any(patterns, func(pattern string) bool {
			return matchPattern(pattern, value)
		})
	}

	return matchAny(match.Databases, spec.Database) &&
		matchAny(match.DatabaseServiceNames, spec.DatabaseServiceName) &&
		matchAny(match.ProtocolNames, spec.Protocol) &&
		matchAny(match.ObjectKinds, spec.ObjectKind) &&
		matchAny(match.Names, spec.Name) &&
		matchAny(match.Schemas, spec.Schema)
}
