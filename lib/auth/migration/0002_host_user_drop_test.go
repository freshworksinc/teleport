/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package migration

import (
	"context"
	"fmt"
	"testing"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/backend"
	"github.com/gravitational/teleport/lib/backend/memory"
	"github.com/gravitational/teleport/lib/services"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func TestHostUserDrop(t *testing.T) {
	t.Parallel()
	b, err := memory.New(memory.Config{EventsOff: true})
	require.NoError(t, err)
	access := newFakeAccess()
	migration := hostUserDrop{
		accessServiceFn: func(b backend.Backend) services.Access {
			return access
		},
	}

	tests := []struct {
		name          string
		existingModes []types.CreateHostUserMode
		expectedModes []types.CreateHostUserMode
		migrateFn     func(ctx context.Context, b backend.Backend) error
	}{
		{
			name: "up",
			existingModes: []types.CreateHostUserMode{
				types.CreateHostUserMode_HOST_USER_MODE_DROP,
				types.CreateHostUserMode_HOST_USER_MODE_INSECURE_DROP,
				types.CreateHostUserMode_HOST_USER_MODE_KEEP,
				types.CreateHostUserMode_HOST_USER_MODE_OFF,
			},
			expectedModes: []types.CreateHostUserMode{
				types.CreateHostUserMode_HOST_USER_MODE_INSECURE_DROP,
				types.CreateHostUserMode_HOST_USER_MODE_INSECURE_DROP,
				types.CreateHostUserMode_HOST_USER_MODE_KEEP,
				types.CreateHostUserMode_HOST_USER_MODE_OFF,
			},
			migrateFn: migration.Up,
		},
		{
			name: "down",
			existingModes: []types.CreateHostUserMode{
				types.CreateHostUserMode_HOST_USER_MODE_INSECURE_DROP,
				types.CreateHostUserMode_HOST_USER_MODE_INSECURE_DROP,
				types.CreateHostUserMode_HOST_USER_MODE_KEEP,
				types.CreateHostUserMode_HOST_USER_MODE_OFF,
			},
			expectedModes: []types.CreateHostUserMode{
				types.CreateHostUserMode_HOST_USER_MODE_DROP,
				types.CreateHostUserMode_HOST_USER_MODE_DROP,
				types.CreateHostUserMode_HOST_USER_MODE_KEEP,
				types.CreateHostUserMode_HOST_USER_MODE_OFF,
			},
			migrateFn: migration.Down,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			for i, mode := range tc.existingModes {
				role, err := types.NewRole(fmt.Sprintf("role-%d", i), types.RoleSpecV6{
					Options: types.RoleOptions{
						CreateHostUserMode: mode,
					},
				})
				require.NoError(t, err)
				access.roles[role.GetName()] = role
			}
			require.NoError(t, tc.migrateFn(ctx, b))

			for i, mode := range tc.expectedModes {
				role := access.roles[fmt.Sprintf("role-%d", i)]
				require.Equal(t, mode, role.GetOptions().CreateHostUserMode)
			}
		})
	}
}

type fakeAccess struct {
	services.Access
	roles map[string]types.Role
}

func newFakeAccess() *fakeAccess {
	return &fakeAccess{
		roles: make(map[string]types.Role),
	}
}

func (f *fakeAccess) GetRoles(_ context.Context) ([]types.Role, error) {
	return maps.Values(f.roles), nil
}

func (f *fakeAccess) UpdateRole(_ context.Context, role types.Role) (types.Role, error) {
	if _, ok := f.roles[role.GetName()]; ok {
		f.roles[role.GetName()] = role
	}
	return role, nil
}
