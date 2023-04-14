/*
Copyright 2023 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import React from 'react';

import { SVGIcon } from './SVGIcon';

import type { SVGIconProps } from './common';

export function ServersIcon({ size = 20, fill }: SVGIconProps) {
  return (
    <SVGIcon size={size} fill={fill} viewBox="0 0 20 20">
      <path d="M14.6875 9.99984C14.6875 9.47775 15.1072 9.05452 15.625 9.05452C16.1428 9.05452 16.5625 9.47775 16.5625 9.99984C16.5625 10.5219 16.1428 10.9451 15.625 10.9451C15.1072 10.9451 14.6875 10.5219 14.6875 9.99984ZM13.125 10.9451C13.6428 10.9451 14.0625 10.5219 14.0625 9.99984C14.0625 9.47775 13.6428 9.05452 13.125 9.05452C12.6072 9.05452 12.1875 9.47775 12.1875 9.99984C12.1875 10.5219 12.6072 10.9451 13.125 10.9451ZM20 5.90348C20 6.3878 19.8192 6.82938 19.5222 7.1639C19.8192 7.49842 20 7.94 20 8.42432V11.5754C20 12.0597 19.8192 12.5013 19.5222 12.8358C19.8192 13.1703 20 13.6119 20 14.0962V17.2472C20 18.2914 19.1605 19.1379 18.125 19.1379H1.875C0.839453 19.1379 0 18.2914 0 17.2472V14.0962C0 13.6119 0.18082 13.1703 0.477813 12.8358C0.18082 12.5013 0 12.0597 0 11.5754V8.42432C0 7.94 0.18082 7.49842 0.477813 7.1639C0.18082 6.82938 0 6.3878 0 5.90348V2.75244C0 1.70826 0.839453 0.861816 1.875 0.861816H18.125C19.1605 0.861816 20 1.70826 20 2.75244V5.90348ZM1.25 5.90348C1.25 6.25096 1.53039 6.53369 1.875 6.53369H18.125C18.4696 6.53369 18.75 6.25096 18.75 5.90348V2.75244C18.75 2.40496 18.4696 2.12223 18.125 2.12223H1.875C1.53039 2.12223 1.25 2.40496 1.25 2.75244V5.90348ZM18.125 7.79411H1.875C1.53039 7.79411 1.25 8.07683 1.25 8.42432V11.5754C1.25 11.9228 1.53039 12.2056 1.875 12.2056H18.125C18.4696 12.2056 18.75 11.9228 18.75 11.5754V8.42432C18.75 8.07683 18.4696 7.79411 18.125 7.79411ZM18.75 14.0962C18.75 13.7487 18.4696 13.466 18.125 13.466H1.875C1.53039 13.466 1.25 13.7487 1.25 14.0962V17.2472C1.25 17.5947 1.53039 17.8774 1.875 17.8774H18.125C18.4696 17.8774 18.75 17.5947 18.75 17.2472V14.0962ZM15.625 5.27327C16.1428 5.27327 16.5625 4.85005 16.5625 4.32796C16.5625 3.80587 16.1428 3.38265 15.625 3.38265C15.1072 3.38265 14.6875 3.80587 14.6875 4.32796C14.6875 4.85005 15.1072 5.27327 15.625 5.27327ZM13.125 5.27327C13.6428 5.27327 14.0625 4.85005 14.0625 4.32796C14.0625 3.80587 13.6428 3.38265 13.125 3.38265C12.6072 3.38265 12.1875 3.80587 12.1875 4.32796C12.1875 4.85005 12.6072 5.27327 13.125 5.27327ZM15.625 14.7264C15.1072 14.7264 14.6875 15.1496 14.6875 15.6717C14.6875 16.1938 15.1072 16.617 15.625 16.617C16.1428 16.617 16.5625 16.1938 16.5625 15.6717C16.5625 15.1496 16.1428 14.7264 15.625 14.7264ZM13.125 14.7264C12.6072 14.7264 12.1875 15.1496 12.1875 15.6717C12.1875 16.1938 12.6072 16.617 13.125 16.617C13.6428 16.617 14.0625 16.1938 14.0625 15.6717C14.0625 15.1496 13.6428 14.7264 13.125 14.7264Z" />
    </SVGIcon>
  );
}
