/**
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

import React, { useState } from 'react';

import { Server as ServerIcon } from 'design/Icon';

import { ViewMode } from 'shared/services/unifiedResourcePreferences';

import { FilterPanel } from 'shared/components/UnifiedResources/FilterPanel';

import { UnifiedResourcesQueryParams } from 'shared/components/UnifiedResources';

import {
  UserListProps,
  UserListViewProps,
} from 'teleport/Users/UserList/types';
import { UserActionButton } from 'teleport/Users/UserList/UserActionButton';
import { UserListView } from 'teleport/Users/UserList/ListView';
import { UserCardView } from 'teleport/Users/UserList/CardView';

export default function UserList({
  users = [],
  onEdit,
  onDelete,
  onReset,
}: UserListProps) {
  const [viewMode, setViewMode] = useState<ViewMode>(ViewMode.VIEW_MODE_LIST);
  const [params, setParams] = useState<UnifiedResourcesQueryParams>({
    query: '',
    search: '',
    sort: {
      fieldName: 'name',
      dir: 'ASC',
    },
    pinnedOnly: false,
    kinds: [],
  });

  const ViewComponent =
    viewMode === ViewMode.VIEW_MODE_LIST ? UserListView : UserCardView;

  function renderAuthType(authType: string) {
    switch (authType) {
      case 'github':
        return 'GitHub';
      case 'saml':
        return 'SAML';
      case 'oidc':
        return 'OIDC';
      case 'teleport local user':
        return 'Teleport Local User';
    }
    return authType;
  }

  const viewProps: UserListViewProps = {
    onLabelClick: () => {},
    selectedResources: [],
    onSelectResource: () => {},
    isProcessing: false,
    mappedUsers: users.map(u => ({
      item: {
        name: u.name,
        roles: u.roles,
        primaryIconName: 'User',
        SecondaryIcon: ServerIcon,
        ActionButton: UserActionButton({
          user: u,
          onEdit: onEdit,
          onReset: onReset,
          onDelete: onDelete,
        }),
        cardViewProps: {
          primaryDesc: renderAuthType(u.authType),
          secondaryDesc: '',
        },
        listViewProps: {
          description: '',
          addr: u.roles.toString(),
          resourceType: renderAuthType(u.authType),
        },
      },
      key: 'foo',
    })),
    expandAllLabels: false,
  };

  return (
    <>
      <FilterPanel
        availableKinds={[]}
        params={params}
        setParams={setParams}
        selectVisible={() => {}}
        selected={false}
        // BulkActions?: React.ReactElement
        currentViewMode={viewMode}
        setCurrentViewMode={setViewMode}
        expandAllLabels={false}
        setExpandAllLabels={() => {}}
      />
      <ViewComponent {...viewProps} />
    </>
  );
}
