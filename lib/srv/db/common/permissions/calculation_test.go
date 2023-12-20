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
	"github.com/gravitational/teleport/api/utils"
)

type mockGetter struct {
	allow []types.DatabasePermission
	deny  []types.DatabasePermission
}

func (m mockGetter) GetDatabasePermissions() (allow types.DatabasePermissions, deny types.DatabasePermissions) {
	return m.allow, m.deny
}

func TestCalculatePermissions(t *testing.T) {
	mkDatabaseObject := func(name string, labels map[string]string) types.DatabaseObject {
		out, err := types.NewDatabaseObject(
			types.Metadata{Name: name},
			types.DatabaseObjectSpec{
				Name:       "name",
				Protocol:   "postgres",
				ObjectKind: "table",
			},
		)
		require.NoError(t, err)
		if len(labels) > 0 {
			out.SetStaticLabels(labels)
		}
		return out
	}

	tests := []struct {
		name    string
		getter  GetDatabasePermissions
		objs    []types.DatabaseObject
		want    PermissionSet
		summary string
	}{
		{
			name: "Some permissions for some objects",
			getter: &mockGetter{
				allow: []types.DatabasePermission{
					{
						Permissions: []string{"SELECT"},
						Match:       map[string]utils.Strings{"kind": []string{"table"}},
					},
					{
						Permissions: []string{"DELETE"},
						Match:       map[string]utils.Strings{"kind": []string{"schema"}},
					},
				},
				deny: nil,
			},
			objs: []types.DatabaseObject{
				mkDatabaseObject("foo", map[string]string{"kind": "table"}),
				mkDatabaseObject("bar", map[string]string{"kind": "schema"}),
			},
			want: PermissionSet{
				"SELECT": {
					mkDatabaseObject("foo", map[string]string{"kind": "table"}),
				},
				"DELETE": {
					mkDatabaseObject("bar", map[string]string{"kind": "schema"}),
				},
			},
			summary: "\"DELETE\": 1 objects (table:1), \"SELECT\": 1 objects (table:1)",
		},
		{
			name: "Deny removes permissions",
			getter: &mockGetter{
				allow: []types.DatabasePermission{
					{
						Permissions: []string{"SELECT", "INSERT"},
						Match:       map[string]utils.Strings{"kind": []string{"table"}},
					},
					{
						Permissions: []string{"SELECT", "DELETE"},
						Match:       map[string]utils.Strings{"kind": []string{"schema"}},
					},
				},
				deny: []types.DatabasePermission{
					{
						Permissions: []string{"*"},
						Match:       map[string]utils.Strings{"kind": []string{"table"}},
					},
					{
						Permissions: []string{"DELETE"},
						Match:       map[string]utils.Strings{"kind": []string{"schema"}},
					},
				},
			},
			objs: []types.DatabaseObject{
				mkDatabaseObject("foo", map[string]string{"kind": "table"}),
				mkDatabaseObject("bar", map[string]string{"kind": "schema"}),
			},
			want: PermissionSet{
				"SELECT": {
					mkDatabaseObject("bar", map[string]string{"kind": "schema"}),
				},
			},
			summary: "\"SELECT\": 1 objects (table:1)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculatePermissions(tt.getter, tt.objs)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.summary, SummarizePermissions(got))
		})
	}
}

func TestDatabasePermissionMatch(t *testing.T) {
	mkDatabaseObject := func(labels map[string]string) types.DatabaseObject {
		out, err := types.NewDatabaseObject(types.Metadata{Name: "foo"}, types.DatabaseObjectSpec{ObjectKind: "table", Protocol: "postgres"})
		require.NoError(t, err)
		if len(labels) > 0 {
			out.SetStaticLabels(labels)
		}
		return out
	}

	tests := []struct {
		name      string
		obj       types.DatabaseObject
		labels    types.Labels
		wantMatch bool
	}{
		{
			name: "object with some labels",
			obj:  mkDatabaseObject(map[string]string{"owner": "scrooge", "env": "dev"}),
			labels: types.Labels{
				"env":   {"production", "dev"},
				"owner": {"john_doe", "scrooge"},
			},
			wantMatch: true,
		},
		{
			name: "labels and glob patterns",
			obj:  mkDatabaseObject(map[string]string{"owner": "scrooge", "env": "dev"}),
			labels: types.Labels{
				"env":   {"*"},
				"owner": {"john_doe", "scrooge"},
			},
			wantMatch: true,
		},

		{
			name: "mismatch: object without labels",
			obj:  mkDatabaseObject(nil),
			labels: types.Labels{
				"env":   {"production", "dev"},
				"owner": {"john_doe", "scrooge"},
			},
			wantMatch: false,
		},
		{
			name:      "mismatch: no labels",
			obj:       mkDatabaseObject(map[string]string{"env": "dev"}),
			labels:    types.Labels{},
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := databasePermissionMatch(types.DatabasePermission{Permissions: []string{"SELECT"}, Match: tt.labels}, tt.obj)
			require.Equal(t, tt.wantMatch, matched)
		})
	}
}
