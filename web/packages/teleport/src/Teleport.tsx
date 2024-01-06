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

import React, { lazy, Suspense } from 'react';
import ThemeProvider from 'design/ThemeProvider';
import { ListAddCheck, Terminal } from 'design/Icon';

import { Route, Router, Switch } from 'teleport/components/Router';
import { CatchError } from 'teleport/components/CatchError';
import Authenticated from 'teleport/components/Authenticated';

import { getOSSFeatures } from 'teleport/features';

import { LayoutContextProvider } from 'teleport/Main/LayoutContext';
import { UserContextProvider } from 'teleport/User';
import { NewCredentials } from 'teleport/Welcome/NewCredentials';
import { NavigationItem } from 'teleport/TopBar/TopBar';

import TeleportContextProvider from './TeleportContextProvider';
import TeleportContext from './teleportContext';
import cfg from './config';
import useStickyClusterId from './useStickyClusterId';

import type { History } from 'history';

const AppLauncher = lazy(() => import('./AppLauncher'));

const Teleport: React.FC<Props> = props => {
  const { ctx, history } = props;
  const createPublicRoutes = props.renderPublicRoutes || publicOSSRoutes;
  const createPrivateRoutes = props.renderPrivateRoutes || privateOSSRoutes;

  return (
    <CatchError>
      <ThemeProvider>
        <LayoutContextProvider>
          <Router history={history}>
            <Suspense fallback={null}>
              <Switch>
                {createPublicRoutes()}
                <Route path={cfg.routes.root}>
                  <Authenticated>
                    <UserContextProvider>
                      <TeleportContextProvider ctx={ctx}>
                        <TeleportSwitch
                          renderPrivateRoutes={createPrivateRoutes}
                        />
                      </TeleportContextProvider>
                    </UserContextProvider>
                  </Authenticated>
                </Route>
              </Switch>
            </Suspense>
          </Router>
        </LayoutContextProvider>
      </ThemeProvider>
    </CatchError>
  );
};

const LoginFailed = lazy(() => import('./Login/LoginFailed'));
const LoginSuccess = lazy(() => import('./Login/LoginSuccess'));
const Login = lazy(() => import('./Login'));
const Welcome = lazy(() => import('./Welcome'));

const Console = lazy(() => import('./Console'));
const Player = lazy(() => import('./Player'));
const DesktopSession = lazy(() => import('./DesktopSession'));

const HeadlessRequest = lazy(() => import('./HeadlessRequest'));

const Main = lazy(() => import('./Main'));

const TeleportSwitch = (props: TeleportSwitchProps) => {
  const { clusterId } = useStickyClusterId();
  const createPrivateRoutes = props.renderPrivateRoutes || privateOSSRoutes;
  return (
    <Switch>
      <Route path={cfg.routes.appLauncher} component={AppLauncher} />
      <Route>{createPrivateRoutes(clusterId)}</Route>
    </Switch>
  );
};

function publicOSSRoutes() {
  return [
    <Route
      title="Login"
      path={cfg.routes.login}
      component={Login}
      key="login"
    />,
    ...getSharedPublicRoutes(),
  ];
}

export function getSharedPublicRoutes() {
  return [
    <Route
      key="login-failed"
      title="Login Failed"
      path={cfg.routes.loginError}
      component={LoginFailed}
    />,
    <Route
      key="login-failed-legacy"
      title="Login Failed"
      path={cfg.routes.loginErrorLegacy}
      component={LoginFailed}
    />,
    <Route
      key="success"
      title="Success"
      path={cfg.routes.loginSuccess}
      component={LoginSuccess}
    />,
    <Route
      key="invite"
      title="Invite"
      path={cfg.routes.userInvite}
      render={() => <Welcome NewCredentials={NewCredentials} />}
    />,
    <Route
      key="password-reset"
      title="Password Reset"
      path={cfg.routes.userReset}
      render={() => <Welcome NewCredentials={NewCredentials} />}
    />,
  ];
}

function privateOSSRoutes(clusterId: string) {
  const navigationItems: NavigationItem[] = [
    {
      title: 'Access Requests',
      path: '/web/accessrequest',
      Icon: <ListAddCheck color="text.main" />,
    },
    {
      title: 'Active Sessions',
      path: cfg.getSessionsRoute(clusterId),
      Icon: <Terminal color="text.main" />,
    },
  ];

  return (
    <Switch>
      {getSharedPrivateRoutes()}
      <Route
        path={cfg.routes.root}
        render={() => (
          <Main
            features={getOSSFeatures()}
            navigationProps={{ navigationItems }}
          />
        )}
      />
    </Switch>
  );
}

export function getSharedPrivateRoutes() {
  return [
    <Route
      key="desktop"
      path={cfg.routes.desktop}
      component={DesktopSession}
    />,
    <Route key="console" path={cfg.routes.console} component={Console} />,
    <Route key="player" path={cfg.routes.player} component={Player} />,
    <Route
      key="headlessSSO"
      path={cfg.routes.headlessSso}
      component={HeadlessRequest}
    />,
  ];
}

export default Teleport;

export type Props = {
  ctx: TeleportContext;
  history: History;
  renderPublicRoutes?: () => React.ReactNode[];
  renderPrivateRoutes?: (clusterId: string) => React.ReactNode;
};

export type TeleportSwitchProps = Pick<Props, 'renderPrivateRoutes'>;
