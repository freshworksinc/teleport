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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport/api/types"
)

func TestApplyDatabaseObjectImportRules(t *testing.T) {
	mkDatabase := func(name string, labels map[string]string) *types.DatabaseV3 {
		db, err := types.NewDatabaseV3(types.Metadata{
			Name:   name,
			Labels: labels,
		}, types.DatabaseSpecV3{
			Protocol: "postgres",
			URI:      "localhost:5252",
		})
		require.NoError(t, err)
		return db
	}

	type option func(db types.DatabaseObject) error

	mkDatabaseObject := func(name string, spec types.DatabaseObjectSpec, options ...option) types.DatabaseObject {
		out, err := types.NewDatabaseObject(types.Metadata{Name: name}, spec)
		require.NoError(t, err)
		for _, opt := range options {
			require.NoError(t, opt(out))
		}

		return out
	}

	mkImportRule := func(name string, spec types.DatabaseObjectImportRuleSpec) types.DatabaseObjectImportRule {
		out, err := types.NewDatabaseObjectImportRule(types.Metadata{Name: name}, spec)
		require.NoError(t, err)
		return out
	}

	tests := []struct {
		name     string
		rules    []types.DatabaseObjectImportRule
		database types.Database
		objs     []types.DatabaseObject
		want     []types.DatabaseObject
	}{
		{
			name:     "empty inputs",
			rules:    []types.DatabaseObjectImportRule{},
			database: mkDatabase("dummy", map[string]string{"env": "prod"}),
			objs:     nil,
			want:     nil,
		},
		{
			name: "database labels are matched by the rules",
			rules: []types.DatabaseObjectImportRule{
				mkImportRule("foo", types.DatabaseObjectImportRuleSpec{
					Priority:       10,
					DatabaseLabels: map[string]string{"env": "dev"},
					Mappings: []types.DatabaseObjectImportRuleMapping{
						{
							AddLabels: map[string]string{
								"dev_access":    "rw",
								"flag_from_dev": "dummy",
							},
						},
					},
				}),
				mkImportRule("bar", types.DatabaseObjectImportRuleSpec{
					Priority:       20,
					DatabaseLabels: map[string]string{"env": "prod"},
					Mappings: []types.DatabaseObjectImportRuleMapping{
						{
							AddLabels: map[string]string{
								"dev_access":     "ro",
								"flag_from_prod": "dummy",
							},
						},
					},
				}),
			},
			database: mkDatabase("dummy", map[string]string{"env": "prod"}),
			objs: []types.DatabaseObject{
				mkDatabaseObject("foo", types.DatabaseObjectSpec{ObjectKind: "table"}),
			},
			want: []types.DatabaseObject{
				mkDatabaseObject("foo", types.DatabaseObjectSpec{ObjectKind: "table"}, func(db types.DatabaseObject) error {
					db.SetStaticLabels(map[string]string{
						"dev_access":     "ro",
						"flag_from_prod": "dummy",
					})
					return nil
				}),
			},
		},
		{
			name: "rule priorities are applied",
			rules: []types.DatabaseObjectImportRule{
				mkImportRule("foo", types.DatabaseObjectImportRuleSpec{
					Priority: 10,
					Mappings: []types.DatabaseObjectImportRuleMapping{
						{
							AddLabels: map[string]string{
								"dev_access":    "rw",
								"flag_from_dev": "dummy",
							},
						},
					},
				}),

				mkImportRule("bar", types.DatabaseObjectImportRuleSpec{
					Priority: 20,
					Mappings: []types.DatabaseObjectImportRuleMapping{
						{
							AddLabels: map[string]string{
								"dev_access":     "ro",
								"flag_from_prod": "dummy",
							},
						},
					},
				}),
			},
			database: mkDatabase("dummy", map[string]string{}),
			objs: []types.DatabaseObject{
				mkDatabaseObject("foo", types.DatabaseObjectSpec{ObjectKind: "table"}),
			},
			want: []types.DatabaseObject{
				mkDatabaseObject("foo", types.DatabaseObjectSpec{ObjectKind: "table"}, func(db types.DatabaseObject) error {
					db.SetStaticLabels(map[string]string{
						"dev_access":     "ro",
						"flag_from_dev":  "dummy",
						"flag_from_prod": "dummy",
					})
					return nil
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := ApplyDatabaseObjectImportRules(tt.rules, tt.database, tt.objs)
			require.Equal(t, tt.want, out)
		})
	}
}

func TestMatchPattern(t *testing.T) {
	testCases := []struct {
		name     string
		pattern  string
		value    string
		expected bool
	}{
		{
			name:     "plain text match",
			pattern:  "exactMatch",
			value:    "exactMatch",
			expected: true,
		},
		{
			name:     "substring mismatch",
			pattern:  "exact",
			value:    "exactMatch",
			expected: false,
		},
		{
			name:     "plain text mismatch",
			pattern:  "exactMatch",
			value:    "noMatch",
			expected: false,
		},
		{
			name:     "glob match",
			pattern:  "gl*b*",
			value:    "globMatch",
			expected: true,
		},
		{
			name:     "glob mismatch",
			pattern:  "glob*",
			value:    "noMatch",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := matchPattern(tc.pattern, tc.value)
			require.Equal(t, tc.expected, result, "Mismatch in test case: %s", tc.name)
		})
	}
}

func TestDatabaseObjectImportMatch(t *testing.T) {
	testCases := []struct {
		name     string
		match    types.DatabaseObjectImportMatch
		spec     types.DatabaseObjectSpec
		expected bool
	}{
		{
			name:     "empty",
			match:    types.DatabaseObjectImportMatch{},
			spec:     types.DatabaseObjectSpec{},
			expected: true,
		},
		{
			name: "match",
			match: types.DatabaseObjectImportMatch{
				Databases:            []string{"db1", "db2"},
				DatabaseServiceNames: []string{"service1", "service2"},
				ProtocolNames:        []string{"postgres", "mysql"},
				ObjectKinds:          []string{"table", "view"},
				Names:                []string{"object1", "object2"},
				Schemas:              []string{"schema1", "schema2"},
			},
			spec: types.DatabaseObjectSpec{
				Database:            "db1",
				DatabaseServiceName: "service1",
				Protocol:            "postgres",
				ObjectKind:          "table",
				Name:                "object1",
				Schema:              "schema1",
			},
			expected: true,
		},
		{
			name: "glob",
			match: types.DatabaseObjectImportMatch{
				Databases:            []string{"*"},
				DatabaseServiceNames: []string{"service*"},
				ProtocolNames:        []string{"post*"},
				ObjectKinds:          []string{"*"},
				Names:                []string{"*"},
				Schemas:              []string{"*"},
			},
			spec: types.DatabaseObjectSpec{
				Database:            "db1",
				DatabaseServiceName: "service1",
				Protocol:            "postgres",
				ObjectKind:          "table",
				Name:                "object1",
				Schema:              "schema1",
			},
			expected: true,
		},
		{
			name: "different",
			match: types.DatabaseObjectImportMatch{
				Databases:            []string{"db1", "db2"},
				DatabaseServiceNames: []string{"service1", "service2"},
				ProtocolNames:        []string{"postgres", "mysql"},
				ObjectKinds:          []string{"table", "view"},
				Names:                []string{"object1", "object2"},
				Schemas:              []string{"schema1", "schema2"},
			},
			spec: types.DatabaseObjectSpec{
				Database:            "db3",
				DatabaseServiceName: "service3",
				Protocol:            "postgres",
				ObjectKind:          "view",
				Name:                "object3",
				Schema:              "schema3",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := databaseObjectImportMatch(tc.match, tc.spec)
			require.Equal(t, tc.expected, result)
		})
	}
}
