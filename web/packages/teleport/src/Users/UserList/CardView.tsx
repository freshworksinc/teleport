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
import styled from 'styled-components';

import { Flex } from 'design';
import {
  FETCH_MORE_SIZE,
  PinningSupport,
} from 'shared/components/UnifiedResources';
import { ResourceCard } from 'shared/components/UnifiedResources/CardsView/ResourceCard';
import { LoadingCard } from 'shared/components/UnifiedResources/CardsView/LoadingCard';
import { LoadingSkeleton } from 'shared/components/UnifiedResources/shared/LoadingSkeleton';

import {UserCardViewProps} from "teleport/Users/UserList/types";


export function UserCardView({
  mappedUsers,
  onLabelClick,
  selectedResources,
  onSelectResource,
  isProcessing,
}: UserCardViewProps) {
  return (
    <CardsContainer className="CardsContainer" gap={2}>
      {mappedUsers.map(({ item, key }) => (
        <ResourceCard
          key={key}
          name={item.name}
          ActionButton={item.ActionButton}
          primaryIconName={item.primaryIconName}
          onLabelClick={onLabelClick}
          SecondaryIcon={item.SecondaryIcon}
          cardViewProps={item.cardViewProps}
          selected={selectedResources.includes(key)}
          selectResource={() => onSelectResource(key)}
          labels={[]}
          pinned={false}
          pinResource={() => {}}
          pinningSupport={PinningSupport.NotSupported}
        />
      ))}
      {isProcessing && (
        <LoadingSkeleton count={FETCH_MORE_SIZE} Element={<LoadingCard />} />
      )}
    </CardsContainer>
  );
}

const CardsContainer = styled(Flex)`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
`;
