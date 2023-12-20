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
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/services"
)

type PermissionSet map[string][]types.DatabaseObject

type GetDatabasePermissions interface {
	GetDatabasePermissions() (allow types.DatabasePermissions, deny types.DatabasePermissions)
}

func databasePermissionMatch(perm types.DatabasePermission, obj types.DatabaseObject) bool {
	ok, _, _ := services.MatchLabels(perm.Match, obj.GetAllLabels())
	return ok
}

// SummarizePermissions summarizes permissions for logging.
func SummarizePermissions(perms PermissionSet) string {
	var fragments []string

	permNames := maps.Keys(perms)
	slices.Sort(permNames)
	for _, perm := range permNames {
		objects := perms[perm]
		objKinds := map[string]int{}
		for _, object := range objects {
			kind := object.GetSpec().ObjectKind
			objKinds[kind] += 1
		}
		kinds := maps.Keys(objKinds)
		slices.Sort(kinds)
		var kindCounts []string
		for _, kind := range kinds {
			kindCounts = append(kindCounts, fmt.Sprintf("%v:%v", kind, objKinds[kind]))
		}

		fragments = append(fragments, fmt.Sprintf("%q: %v objects (%v)", perm, len(objects), strings.Join(kindCounts, ", ")))
	}

	if len(fragments) == 0 {
		return "(none)"
	}

	return strings.Join(fragments, ", ")
}

// CalculatePermissions calculates the effective permissions for a set of database objects based on the provided allow and deny permissions.
func CalculatePermissions(getter GetDatabasePermissions, objs []types.DatabaseObject) (PermissionSet, error) {
	allow, deny := getter.GetDatabasePermissions()

	out := map[string][]types.DatabaseObject{}

	for _, obj := range objs {
		permsToAdd := map[string]string{}

		// add allowed permissions
		for _, perm := range allow {
			if databasePermissionMatch(perm, obj) {
				for _, permission := range perm.Permissions {
					permsToAdd[strings.TrimSpace(strings.ToUpper(permission))] = permission
				}
			}
		}

		// remove denied permissions
		for _, perm := range deny {
			if databasePermissionMatch(perm, obj) {
				for _, permission := range perm.Permissions {
					// wildcard clears the permissions
					if permission == types.Wildcard {
						clear(permsToAdd)
						break
					}

					delete(permsToAdd, strings.TrimSpace(strings.ToUpper(permission)))
				}
			}
		}

		for _, perm := range permsToAdd {
			out[perm] = append(out[perm], obj)
		}
	}

	return out, nil
}
