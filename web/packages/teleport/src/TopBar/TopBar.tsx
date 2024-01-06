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

import React, { lazy, Suspense, useState } from 'react';
import styled, { useTheme } from 'styled-components';
import { Link } from 'react-router-dom';
import { Flex, Image, Text, TopNav } from 'design';

import { matchPath, useHistory } from 'react-router';

import { BrainIcon } from 'design/SVGIcon';

import { ArrowLeft, ChevronRight, SlidersVertical } from 'design/Icon';
import { HoverTooltip } from 'shared/components/ToolTip';

import useTeleport from 'teleport/useTeleport';
import { UserMenuNav } from 'teleport/components/UserMenuNav';
import { useFeatures } from 'teleport/FeaturesContext';
import { NavigationCategory } from 'teleport/Navigation/categories';

import { useLayout } from 'teleport/Main/LayoutContext';
import { getFirstRouteForCategory } from 'teleport/Navigation/Navigation';

import { Notifications } from './Notifications';
import { ButtonIconContainer } from './Shared';
import logoLight from './logoLight.svg';
import logoDark from './logoDark.svg';

const Assist = lazy(() => import('teleport/Assist'));

export function TopBar({ navigationItems }: NavigationProps) {
  const ctx = useTeleport();
  const history = useHistory();
  const features = useFeatures();
  const assistEnabled = ctx.getFeatureFlags().assist && ctx.assistEnabled;

  const [showAssist, setShowAssist] = useState(false);

  const { hasDockedElement } = useLayout();

  // find active feature
  const feature = features
    .filter(feature => Boolean(feature.route))
    .find(f =>
      matchPath(history.location.pathname, {
        path: f.route.path,
        exact: f.route.exact ?? false,
      })
    );

  function handleBack() {
    const firstRouteForCategory = getFirstRouteForCategory(
      features,
      feature.category
    );

    history.push(firstRouteForCategory);
  }

  return (
    <TopBarContainer navigationHidden={feature?.hideNavigation}>
      <TeleportLogo />
      {feature?.hideNavigation && (
        <ButtonIconContainer onClick={handleBack}>
          <ArrowLeft size="medium" />
        </ButtonIconContainer>
      )}
      <NavigationButton
        selected={feature?.category === NavigationCategory.Management}
        to="/web/users"
        title="Access Management"
      >
        <>
          <SlidersVertical
            css={`
              display: none;
              @media screen and (max-width: ${p =>
                  p.theme.breakpoints.medium}px) {
                display: inline-flex;
              }
            `}
            color="text.main"
          />
          <Text
            css={`
              display: none;
              @media screen and (min-width: ${p =>
                  p.theme.breakpoints.medium}px) {
                display: block;
              }
            `}
            fontSize={18}
            fontWeight={500}
            color="text.muted"
          >
            Access Management
          </Text>
          {feature?.category !== NavigationCategory.Management && (
            <ChevronRight
              css={`
                align-self: center;
                height: 100%;
                @media screen and (max-width: ${p =>
                    p.theme.breakpoints.medium}px) {
                  display: none;
                }
              `}
              color="text.muted"
            />
          )}
        </>
      </NavigationButton>

      <Flex ml="auto" height="100%" alignItems="center">
        {navigationItems.map(({ Icon, path, title }) => (
          <NavigationButton
            key={path}
            to={path}
            selected={history.location.pathname.includes(path)}
            title={title}
          >
            {Icon}
          </NavigationButton>
        ))}
        {!hasDockedElement && assistEnabled && (
          <ButtonIconContainer onClick={() => setShowAssist(true)}>
            <BrainIcon />
          </ButtonIconContainer>
        )}
        <Notifications />
        <UserMenuNav username={ctx.storeUser.state.username} />
      </Flex>

      {showAssist && (
        <Suspense fallback={null}>
          <Assist onClose={() => setShowAssist(false)} />
        </Suspense>
      )}
    </TopBarContainer>
  );
}

export const TopBarContainer = styled(TopNav)`
  background: ${p => p.theme.colors.levels.surface};
  overflow-y: initial;
  flex-shrink: 0;
  z-index: 1000;
  border-bottom: 1px solid ${({ theme }) => theme.colors.spotBackground[0]};

  height: 48px;
  @media screen and (min-width: ${p => p.theme.breakpoints.small}px) {
    height: 56px;
  }
  @media screen and (min-width: ${p => p.theme.breakpoints.large}px) {
    height: 72px;
  }

  box-shadow: 0px 1px 3px 0px rgba(0, 0, 0, 0.12),
    0px 1px 1px 0px rgba(0, 0, 0, 0.14), 0px 2px 1px -1px rgba(0, 0, 0, 0.2);
`;

const TeleportLogo = () => {
  const theme = useTheme();

  return (
    <HoverTooltip
      anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      transformOrigin={{ vertical: 'top', horizontal: 'center' }}
      tipContent="Teleport Resources Home"
    >
      <Link
        css={`
          cursor: pointer;
          height: 100%;
          display: flex;
          width: 256px;

          transition: background-color 0.1s linear;
          &:hover {
            background-color: ${p =>
              p.theme.colors.interactive.tonal.primary[0]};
          }
          align-items: center;
        `}
        to="/web"
      >
        <Image
          src={theme.type === 'dark' ? logoDark : logoLight}
          alt="teleport logo"
          css={`
            padding-left: ${props => props.theme.space[4]}px;
          `}
        />
      </Link>
    </HoverTooltip>
  );
};

const NavigationButton = ({
  to,
  selected,
  children,
  title,
}: {
  to: string;
  selected: boolean;
  children: JSX.Element;
  title: string;
}) => {
  const theme = useTheme();
  const selectedBorder = `2px solid ${theme.colors.brand}`;
  const selectedBackground = theme.colors.interactive.tonal.primary[0];

  return (
    <HoverTooltip
      anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      transformOrigin={{ vertical: 'top', horizontal: 'center' }}
      tipContent={title}
    >
      <Link
        to={to}
        css={`
          text-decoration: none;
          color: rgba(0, 0, 0, 0.54);
          height: 100%;
          padding-left: 24px;
          padding-right: 24px;
          border-bottom: ${selected ? selectedBorder : 'none'};
          background-color: ${selected ? selectedBackground : 'inherit'};
          &:hover {
            background-color: ${selected
              ? selectedBackground
              : theme.colors.buttons.secondary.default};
          }
        `}
      >
        <Flex
          css={`
            height: 100%;
          `}
          justifyContent="center"
          alignItems="center"
        >
          {children}
        </Flex>
      </Link>
    </HoverTooltip>
  );
};

export type NavigationProps = {
  CustomLogo?: () => React.ReactElement;
  showPoweredByLogo?: boolean;
  navigationItems?: NavigationItem[];
};

export type NavigationItem = {
  title: string;
  path: string;
  Icon: JSX.Element;
};
