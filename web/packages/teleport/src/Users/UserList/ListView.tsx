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

import React from 'react';

import { Flex } from 'design';
import { ResourceListItem } from 'shared/components/UnifiedResources/ListView/ResourceListItem';
import { LoadingSkeleton } from 'shared/components/UnifiedResources/shared/LoadingSkeleton';
import {
  FETCH_MORE_SIZE,
  PinningSupport,
} from 'shared/components/UnifiedResources';
import { LoadingListItem } from 'shared/components/UnifiedResources/ListView/LoadingListItem';

import { UserListViewProps } from 'teleport/Users/UserList/types';

export function UserListView({
  mappedUsers,
  onLabelClick,
  selectedResources,
  onSelectResource,
  isProcessing,
  expandAllLabels,
}: UserListViewProps) {
  return (
    <Flex className="ListContainer">
      {mappedUsers.map(({ item, key }) => (
        <ResourceListItem
          ActionButton={item.ActionButton}
          SecondaryIcon={item.SecondaryIcon}
          key={key}
          name={item.name}
          primaryIconName={item.primaryIconName}
          onLabelClick={onLabelClick}
          listViewProps={item.listViewProps}
          selected={selectedResources.includes(key)}
          selectResource={() => onSelectResource(key)}
          expandAllLabels={expandAllLabels}
          labels={[]}
          pinResource={null}
          pinned={null}
          pinningSupport={PinningSupport.NotSupported}
        />
      ))}
      {isProcessing && (
        <LoadingSkeleton
          count={FETCH_MORE_SIZE}
          Element={<LoadingListItem />}
        />
      )}
    </Flex>
  );
}
