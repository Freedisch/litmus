/* eslint-disable no-restricted-imports */

import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import { StringsContext } from '@strings';
import strings from 'strings/strings.en.yaml';
import '../bootstrap.scss';
import { AppStoreContext } from '@context';
import { LitmusAPIProvider } from '@api/LitmusAPIProvider';

interface TestWrapperProps {
  children: React.ReactElement;
}

export const findDialogContainer = (): HTMLElement | null => document.querySelector('.bp3-dialog');
export const findPopoverContainer = (): HTMLElement | null => document.querySelector('.bp3-popover-content');

export function TestWrapper({ children }: TestWrapperProps): React.ReactElement {
  const getString = (key: string): string => key;
  const getAPIEndpoints = ()

  return (
    <BrowserRouter>
      <AppStoreContext.Provider
        value={{
          projectID: 'litmuschaos-test-project',
          projectRole: 'Owner',
          currentUserInfo: {
            ID: 'uid',
            username: 'admin',
            userRole: 'admin'
          },
          renderUrl: `/account/uid`,
          matchPath: '/account/:accountID',
          updateAppStore: () => void 0
        }}
      >
        <StringsContext.Provider value={{ data: strings as any, getString }}>
          <LitmusAPIProvider config={{
            gqlEndpoints: {
              chaosManagerUri: `$api/query`
            },
            restEndpoints: {
              authUri: `/auth`,
              chaosManagerUri: `/api`
            }
          }}>
            {children}
          </LitmusAPIProvider>
        </StringsContext.Provider>
      </AppStoreContext.Provider>
    </BrowserRouter>
  );
}
