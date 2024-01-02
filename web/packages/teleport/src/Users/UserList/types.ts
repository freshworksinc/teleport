
import React from 'react';
import { ResourceIconName } from 'design/ResourceIcon';
import { Icon } from 'design/Icon';
import {
  CardViewSpecificProps,
  ListViewSpecificProps,
} from 'shared/components/UnifiedResources';

import { User } from 'teleport/services/user';

import { ResourceLabel } from 'teleport/services/agents';

export type UserViewItem = User & {
  name: string;
  primaryIconName: ResourceIconName;
  SecondaryIcon: typeof Icon;
  ActionButton: React.ReactElement;
  cardViewProps: CardViewSpecificProps;
  listViewProps: ListViewSpecificProps;
};

export type MappedUsers = { item: UserViewItem; key: string }[];

export type UserListViewProps = {
  onLabelClick: (label: ResourceLabel) => void;
  selectedResources: string[];
  onSelectResource: (resourceId: string) => void;
  isProcessing: boolean;
  mappedUsers: MappedUsers;
  expandAllLabels: boolean;
};

export type UserCardViewProps = {
  onLabelClick: (label: ResourceLabel) => void;
  selectedResources: string[];
  onSelectResource: (resourceId: string) => void;
  isProcessing: boolean;
  mappedUsers: { item: UserViewItem; key: string }[];
  expandAllLabels: boolean;
};

export type UserListProps = {
  users: User[];
  pageSize?: number;
  onEdit(user: User): void;
  onDelete(user: User): void;
  onReset(user: User): void;
};
