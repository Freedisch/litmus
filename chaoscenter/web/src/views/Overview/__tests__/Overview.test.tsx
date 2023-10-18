/* eslint-disable jest/no-commented-out-tests */
import { render, screen } from '@testing-library/react';
import React from 'react';
import type { ApolloQueryResult } from '@apollo/client';
import type { ListExperimentRequest, ListExperimentResponse } from '@api/core';
import { TestWrapper } from 'utils/testUtils';
import OverviewView from '../Overview';
// import type { ListExperimentRequest, ListExperimentResponse } from '@api/core';
// import type { ApolloQueryResult } from '@apollo/client';
// import type { ListExperimentRequest, ListExperimentResponse } from '@api/core';
// import type { ApolloQueryResult } from '@apollo/client';
// import type { ListExperimentRequest, ListExperimentResponse } from '@api/core';
// import type { ApolloQueryResult } from '@apollo/client';
// import type { ListExperimentRequest, ListExperimentResponse } from '@api/core';
// import type { ApolloQueryResult } from '@apollo/client';

describe('OverviewView Component', () => {
  //   it('renders without crashing', () => {
  //     render(
  //       <TestWrapper>
  //         <OverviewView loading={chaosHubStats={undefined} infraStats={undefined} experimentStats={undefined} experimentDashboardTableData={undefined} refetchExperiments={function (variables?: Partial<ListExperimentRequest> | undefined): Promise<ApolloQueryResult<ListExperimentResponse>> {
  //                 throw new Error('Function not implemented.');
  //             } }} />
  //       </TestWrapper>
  //     );
  //   });

  test('shows loading state', () => {
    render(
      <TestWrapper>
        <OverviewView
          loading={{ chaosHubStats: true, infraStats: true, experimentStats: true, recentExperimentsTable: true }}
          chaosHubStats={undefined}
          infraStats={undefined}
          experimentStats={undefined}
          experimentDashboardTableData={undefined}
          refetchExperiments={function (
            _variables?: Partial<ListExperimentRequest> | undefined
          ): Promise<ApolloQueryResult<ListExperimentResponse>> {
            throw new Error('Function not implemented.');
          }}
        />
      </TestWrapper>
    );
    expect(screen.getByText('Loading...')).toBeVisible();
  });

  //   it('opens the chaos modal if no infrastructures', () => {
  //     render(
  //         <OverviewView infraStats={{ totalInfrastructures: 0 }} ... />
  //     );
  //     expect(screen.getByText('enableChaosInfrastructure')).toBeVisible();
  //   });

  //   it('does not open the chaos modal if there are infrastructures', () => {
  //     render(
  //         <OverviewView infraStats={{ totalInfrastructures: 1 }} ... />
  //     );
  //     const modal = screen.queryByText('enableChaosInfrastructure');
  //     expect(modal).toBeNull();
  //   });

  //   it('renders MemoisedExperimentDashboardV2Table if data is present', () => {
  //     render(
  //       <TestWrapper>
  //         <OverviewView experimentDashboardTableData={{ content: ['someData'] }} ... />
  //       </TestWrapper>
  //     );
  //     expect(screen.getByText('recentExperimentRuns')).toBeVisible();  // Or another text/content from the table.
  //   });

  test('renders NewUserLanding if no experiment data is present', () => {
    render(
      <TestWrapper>
        <OverviewView
          experimentDashboardTableData={{ content: [] }}
          chaosHubStats={undefined}
          infraStats={undefined}
          experimentStats={undefined}
          loading={{
            chaosHubStats: false,
            infraStats: false,
            experimentStats: false,
            recentExperimentsTable: false
          }}
          refetchExperiments={function (
            _variables?: Partial<ListExperimentRequest> | undefined
          ): Promise<ApolloQueryResult<ListExperimentResponse>> {
            throw new Error('Function not implemented.');
          }}
        />
      </TestWrapper>
    );
    expect(screen.getByText('Loading...')).toBeVisible();
  });

  //   test('navigates correctly when RbacButton is clicked', () => {
  //     render(
  //       <OverviewView
  //         chaosHubStats={undefined}
  //         infraStats={undefined}
  //         experimentStats={undefined}
  //         experimentDashboardTableData={undefined}
  //         loading={{
  //           chaosHubStats: false,
  //           infraStats: false,
  //           experimentStats: false,
  //           recentExperimentsTable: false
  //         }}
  //         refetchExperiments={function (
  //           _variables?: Partial<ListExperimentRequest> | undefined
  //         ): Promise<ApolloQueryResult<ListExperimentResponse>> {
  //           throw new Error('Function not implemented.');
  //         }}
  //       />
  //     );
  //     const button = screen.getByText('viewAllExperiments'); // Assuming 'viewAllExperiments' is the button's text.
  //     fireEvent.click(button);
  //     // Assertion for navigation, perhaps using a mock/spy on the 'useHistory' hook.
  //   });

  //   test('navigates correctly when Enable Chaos Infrastructure button in the modal is clicked', () => {
  //     render(<OverviewView infraStats={{ totalInfrastructures: 0 }} />);
  //     const button = screen.getByText('enableChaosInfraButton'); // Assuming 'enableChaosInfraButton' is the button's text.
  //     fireEvent.click(button);
  //     // Assertion for navigation, perhaps using a mock/spy on the 'useHistory' hook.
  //   });
});
