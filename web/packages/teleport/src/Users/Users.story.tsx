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

import { Users } from './Users';

export default {
  title: 'Teleport/Users',
};

export const Processing = () => {
  const attempt = {
    isProcessing: true,
    isFailed: false,
    isSuccess: false,
    message: '',
  };
  return <Users {...sample} attempt={attempt} />;
};

export const Loaded = () => {
  return <Users {...sample} />;
};

export const Failed = () => {
  const attempt = {
    isProcessing: false,
    isFailed: true,
    isSuccess: false,
    message: 'some error message',
  };
  return <Users {...sample} attempt={attempt} />;
};

const users = [
  {
    name: 'jane.doe@gmail.com',
    roles: [
      'apple',
      'code-api-dev',
      'code-api-staging',
      'donut-prod-apple-access',
      'elephant-dev-access',
      'golf-dev-apple-access',
      'golf-prod-apple-access',
      'golf-staging-apple-access',
      'legacy-access',
      'kangaroo-staging-access',
      'kangaroo-staging-elephant-access',
      'requester',
      'integration',
      'integration-dev',
      'integration-infra-dev',
      'integration-infra-prod',
      'integration-infra-stage',
      'integration-stage',
      'automation-prod-impersonator',
      'automation-staging-impersonator',
      'pipe-staging-apple-access',
      'internal-app-dev-apple-access',
      'internal-app-dev-db-access',
      'internal-app-prod-apple-access',
      'internal-app-prod-db-access',
      'internal-app-staging-apple-access',
      'internal-app-staging-db-access',
      'comp-event-handler-impersonator',
      'table-impersonator',
    ],
    authType: 'teleport local user',
    isLocal: true,
  },
  {
    name: 'teri+dactyl@code.co',
    roles: ['dino', 'apple'],
    authType: 'teleport local user',
    isLocal: true,
  },
  {
    name: 'alliegggrater@swamp.io',
    roles: ['editor', 'alligator', 'apple'],
    authType: 'github',
    isLocal: false,
  },
  {
    name: 'MACEYGOODWIN@macy.org',
    roles: [
      'donut-prod-apple-access',
      'golf-prod-apple-access',
      'golf-staging-apple-access',
      'opal-requester',
      'integration',
      'integration-dev',
      'integration-stage',
      'pipe-staging-apple-access',
      'internal-app-prod-apple-access',
      'internal-app-staging-apple-access',
    ],
    authType: 'saml',
    isLocal: false,
  },
  {
    name: 'marlowrichardstormolivertaylorparksaverymoorethethird@longer.com',
    roles: [
      'apple',
      'code-api-dev',
      'code-api-staging',
      'd-prod-apple-access',
      'elephant-dev-access',
      'golf-dev-apple-access',
      'golf-prod-apple-access',
      'golf-staging-apple-access',
      'kangaroo-staging-access',
      'kangaroo-staging-elephant-access',
      'requester',
      'integration',
      'integration-dev',
      'integration-infra-dev',
      'integration-infra-prod',
      'integration-infra-stage',
      'integration-stage',
      'automation-prod-impersonator',
      'automation-staging-impersonator',
      'pipe-staging-apple-access',
      'internal-app-dev-apple-access',
      'internal-app-dev-db-access',
      'internal-app-prod-apple-access',
      'internal-app-prod-db-access',
      'internal-app-staging-apple-access',
      'internal-app-staging-db-access',
      'comp-event-handler-impersonator',
      'table-impersonator',
    ],
    authType: 'oidc',
    isLocal: false,
  },
  {
    name: 'short@sweet.org',
    roles: [],
    authType: 'teleport local user',
    isLocal: true,
  },
];

const roles = ['apple', 'testrole'];

const sample = {
  attempt: {
    isProcessing: false,
    isFailed: false,
    isSuccess: true,
    message: '',
  },
  users: users,
  roles: roles,
  operation: {
    type: 'none',
    user: null,
  } as any,
  inviteCollaboratorsOpen: false,
  emailPasswordResetOpen: false,
  onStartCreate: () => null,
  onStartDelete: () => null,
  onStartEdit: () => null,
  onStartReset: () => null,
  onStartInviteCollaborators: () => null,
  onStartEmailResetPassword: () => null,
  onClose: () => null,
  onCreate: () => null,
  onDelete: () => null,
  onUpdate: () => null,
  onReset: () => null,
  onInviteCollaboratorsClose: () => null,
  InviteCollaborators: null,
  onEmailPasswordResetClose: () => null,
  EmailPasswordReset: null,
};
