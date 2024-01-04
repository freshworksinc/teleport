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

import React, { useState, useEffect } from 'react';
import { useHistory } from 'react-router';
import { ButtonSecondary, Flex, Menu, MenuItem, Text } from 'design';
import { ChevronDown } from 'design/Icon';
import cfg from 'teleport/config';
import { Cluster } from 'teleport/services/clusters';
import useTeleport from 'teleport/useTeleport';

import { HoverTooltip } from 'shared/components/ToolTip';

interface ClusterDropdownProps {
  value: string;
  // onLoad is an optional prop. If onLoad is not passed, it will use the built in "loadCluters" function
  onLoad?: () => Promise<Cluster[]>;
  // onChange is an optional prop. If onChange is not passed, it will use the built in "changeCluster" function
  onChange?: (newValue: string) => void;
  // onError is not optional. Because this dropdown can be placed on any page, it does not display it's own error
  // messages. Even if using the internal "loadClusters", we will pass the error back to be consumed by the parent.
  onError: (errorMessage: string) => void;
}

export function ClusterDropdown({
  value,
  onChange,
  onLoad,
  onError,
}: ClusterDropdownProps) {
  const ctx = useTeleport();
  const [options, setOptions] = React.useState<Option[]>([]);
  const history = useHistory();
  const [anchorEl, setAnchorEl] = useState(null);

  const selectedOption = {
    value,
    label: value,
  };

  function loadClusters() {
    onError('');
    try {
      return ctx.clusterService.fetchClusters();
    } catch (err) {
      onError(err.message);
    }
  }

  function changeCluster(value: string) {
    const newPrefix = cfg.getClusterRoute(value);

    const oldPrefix = cfg.getClusterRoute(selectedOption.value);

    const newPath = history.location.pathname.replace(oldPrefix, newPrefix);

    // keep current view just change the clusterId
    history.push(newPath);
  }

  function onChangeOption(clusterId: string) {
    if (onChange) {
      onChange(clusterId);
    } else {
      changeCluster(clusterId);
    }
    handleClose();
  }

  useEffect(() => {
    const loadFunc = onLoad || loadClusters;
    async function getOptions() {
      try {
        const res = await loadFunc();
        setOptions(
          res.map(cluster => ({
            value: cluster.clusterId,
            label: cluster.clusterId,
          }))
        );
      } catch (err) {
        onError(err.message);
      }
    }

    getOptions();
  }, []);

  const handleOpen = event => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  if (options.length < 1) {
    return null;
  }

  return (
    <Flex textAlign="center" alignItems="center">
      <HoverTooltip tipContent={'Select cluster'}>
        <ButtonSecondary
          px={2}
          css={`
            border-color: ${props => props.theme.colors.spotBackground[0]};
          `}
          textTransform="none"
          size="small"
          onClick={handleOpen}
        >
          {selectedOption.label}
          <ChevronDown ml={2} size="small" color="text.slightlyMuted" />
        </ButtonSecondary>
      </HoverTooltip>
      <Menu
        popoverCss={() => `margin-top: 36px;`}
        transformOrigin={{
          vertical: 'top',
          horizontal: 'left',
        }}
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'left',
        }}
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        {options.map(cluster => (
          <MenuItem
            disabled={false}
            px={2}
            key={cluster.value}
            onClick={() => onChangeOption(cluster.value)}
          >
            <Text ml={2} fontWeight={300} fontSize={2}>
              {cluster.label}
            </Text>
          </MenuItem>
        ))}
      </Menu>
    </Flex>
  );
}

type Option = { value: string; label: string };
