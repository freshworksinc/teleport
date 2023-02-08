/**
 * Copyright 2023 Gravitational, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import {
  ChannelCredentials,
  credentials,
  ServerCredentials,
} from '@grpc/grpc-js';

import { RuntimeSettings } from 'teleterm/mainProcess/types';

export function createClientCredentials(
  clientKeyPair: { cert: Buffer; key: Buffer },
  serverCert: Buffer
): ChannelCredentials {
  return credentials.createSsl(
    serverCert,
    clientKeyPair.key,
    clientKeyPair.cert
  );
}

export function createServerCredentials(
  serverKeyPair: { cert: Buffer; key: Buffer },
  clientCert: Buffer
): ServerCredentials {
  return ServerCredentials.createSsl(
    clientCert,
    [
      {
        cert_chain: serverKeyPair.cert,
        private_key: serverKeyPair.key,
      },
    ],
    true
  );
}

export function createInsecureClientCredentials(): ChannelCredentials {
  return credentials.createInsecure();
}

export function createInsecureServerCredentials(): ServerCredentials {
  return ServerCredentials.createInsecure();
}

/**
 * Checks if the gRPC connection should be encrypted.
 * The only source of truth is the type of tshd protocol.
 * Any protocol other than `unix` should be encrypted.
 * The same check is performed on the tshd side.
 */
export function shouldEncryptConnection(
  runtimeSettings: RuntimeSettings
): boolean {
  return (
    new URL(runtimeSettings.tshd.requestedNetworkAddress).protocol !== 'unix:'
  );
}
