import { Cell } from 'design/DataTable';
import { MenuButton, MenuItem } from 'shared/components/MenuAction';
import React from 'react';

import { User } from 'teleport/services/user';

export type UserActionButtonProps = {
  user: User;
  onEdit: (user: User) => void;
  onReset: (user: User) => void;
  onDelete: (user: User) => void;
};

export const UserActionButton = ({
  user,
  onEdit,
  onReset,
  onDelete,
}: UserActionButtonProps) => {
  if (!user.isLocal) {
    return <MenuButton disabled />;
  }

  return (
    <Cell align="right">
      <MenuButton>
        <MenuItem onClick={() => onEdit(user)}>Edit...</MenuItem>
        <MenuItem onClick={() => onReset(user)}>
          Reset Authentication...
        </MenuItem>
        <MenuItem onClick={() => onDelete(user)}>Delete...</MenuItem>
      </MenuButton>
    </Cell>
  );
};
