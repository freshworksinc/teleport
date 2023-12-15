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

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/backend"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/services/local"
	"github.com/gravitational/trace"
)

// hostUserDrop performs a migration which maps the deprecated 'drop' host
// user creation mode to 'insecure-drop'. Introduced in v15.
type hostUserDrop struct {
	accessServiceFn func(b backend.Backend) services.Access
}

func (d hostUserDrop) Version() int64 {
	return 2
}

func (d hostUserDrop) Name() string {
	return "host_user_drop"
}

// Up updates all roles that use the 'drop' host user creation mode to use
// 'insecure-drop' instead. Roles that have other host user creation modes
// or none are skipped.
func (d hostUserDrop) Up(ctx context.Context, b backend.Backend) error {
	ctx, span := tracer.Start(ctx, "hostUserDrop/Up")
	defer span.End()

	if d.accessServiceFn == nil {
		d.accessServiceFn = func(b backend.Backend) services.Access {
			return local.NewAccessService(b)
		}
	}

	return trace.Wrap(d.mapHostUserCreationMode(ctx, d.accessServiceFn(b),
		types.CreateHostUserMode_HOST_USER_MODE_DROP,          // from
		types.CreateHostUserMode_HOST_USER_MODE_INSECURE_DROP, // to
	))
}

// Down updates all roles that use the 'insecure-drop' host user creation mode
// to use 'drop' instead. Roles that have other host user creation modes
// or none are skipped.
func (d hostUserDrop) Down(ctx context.Context, b backend.Backend) error {
	ctx, span := tracer.Start(ctx, "hostUserDrop/Down")
	defer span.End()

	if d.accessServiceFn == nil {
		d.accessServiceFn = func(b backend.Backend) services.Access {
			return local.NewAccessService(b)
		}
	}

	return trace.Wrap(d.mapHostUserCreationMode(ctx, d.accessServiceFn(b),
		types.CreateHostUserMode_HOST_USER_MODE_INSECURE_DROP, // from
		types.CreateHostUserMode_HOST_USER_MODE_DROP,          // to
	))
}

func (d hostUserDrop) mapHostUserCreationMode(ctx context.Context, access services.Access, from, to types.CreateHostUserMode) error {
	roles, err := access.GetRoles(ctx)
	if err != nil {
		return trace.Wrap(err)
	}
	for _, role := range roles {
		options := role.GetOptions()
		if options.CreateHostUserMode == from {
			options.CreateHostUserMode = to
			role.SetOptions(options)
			if _, err := access.UpdateRole(ctx, role); err != nil {
				return trace.Wrap(err)
			}
		}
	}
	return nil
}
