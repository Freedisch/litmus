package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/handlers/rest"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/mocks"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/entities"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/validations"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUserWithProject(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := new(mocks.MockedApplicationService)

	tests := []struct {
		name         string
		username     string
		given        func()
		expectedCode int
	}{
		{
			name:     "Successfully retrieve user with projects",
			username: "testUser",
			given: func() {
				user := &entities.User{
					ID:       "testUID",
					Username: "testUser",
					Email:    "test@example.com",
				}
				project := &entities.Project{ /* mock project data */ }

				service.On("FindUserByUsername", "testUser").Return(user, nil)
				service.On("GetProjectsByUserID", "testUID", false).Return([]*entities.Project{project}, nil)
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{
				{"username", tt.username},
			}
			// Mocking the context for MustGet
			c.Set("username", tt.username)

			tt.given()

			rest.GetUserWithProject(service)(c)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}


func TestGetProject(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := new(mocks.MockedApplicationService)
	user := &entities.User{ID: "testUserID"}
	expectedFilter := primitive.D{
		{Key: "_id", Value: "testProjectID"},
		{Key: "members", Value: primitive.D{
			{Key: "$elemMatch", Value: primitive.D{
				{Key: "user_id", Value: "testUID"},
				{Key: "role", Value: primitive.D{
					{Key: "$in", Value: []string{"Owner", "Viewer", "Editor"}},
				}},
				{Key: "invitation", Value: "Accepted"},
			}},
		}},
	}

	MockProject := []*entities.Project{
		{
			ID: "testProjectID",
			Name: "Test Project",
		},
	}
	
	tests := []struct {
		name         string
		uid          string
		projectID    string
		mockService  func()
		expectedCode int
	}{
		{
			name:      "Successfully retrieve project",
			uid:       "testUID",
			projectID: "testProjectID",
			mockService: func() {
				project := &entities.Project{
					ID: "testProjectID",
					Name: "Test Project",
				}
				service.On("GetUser", "testUID").Return(user, nil)
				service.On("GetProjects", expectedFilter).Return(MockProject, nil)
				service.On("RbacValidator", "testUID", "testProjectID", validations.MutationRbacRules["getProject"], string(entities.AcceptedInvitation)).Return(nil)
				service.On("GetProjectByProjectID", "testProjectID").Return(project, nil)
			},
			expectedCode: http.StatusOK,
		},
		// {
		// 	name:      "Unauthorized due to RBAC failure",
		// 	uid:       "testUID",
		// 	projectID: "testProjectID",
		// 	mockService: func() {
		// 		//service.On("GetProjectByProjectID", mock.Anything).Return(nil, utils.ErrUnauthorized)
		// 	},
		// 	expectedCode: http.StatusUnauthorized,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Setting params and context values
			c.Params = gin.Params{{"project_id", tt.projectID}}
			c.Set("uid", tt.uid)

			tt.mockService()

			rest.GetProject(service)(c)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
func TestGetProjectsByUserID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := new(mocks.MockedApplicationService)

	tests := []struct {
		name         string
		uid          string
		given        func()
		expectedCode int
	}{
		{
			name: "Successfully retrieve projects by user ID",
			uid:  "testUserID",
			given: func() {
				projects := []*entities.Project{
					{
						ID: "testProjectID",
						Name: "Test Project",
					},
				}
				service.On("GetProjectsByUserID", "testUserID", false).Return(projects, nil)
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			w := httptest.NewRecorder()
			ctx := GetTestGinContext(w)
			ctx.Set("uid", "testUserID")
			tt.given()
			rest.GetProjectsByUserID(service)(ctx)
			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

// func TestCreateProject(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	service := new(mocks.MockedApplicationService)

// 	tests := []struct {
// 		name         string
// 		request      entities.CreateProjectInput
// 		given        func()
// 		expectedCode int
// 	}{
// 		{
// 			name: "Successfully create project",
// 			request: entities.CreateProjectInput{
// 				ProjectName: "TestProject",
// 			},
// 			given: func() {
// 				// Mock the service functions as needed
// 				// Example: service.On("CreateProject", mock.AnythingOfType("*entities.Project")).Return(nil)
// 			},
// 			expectedCode: http.StatusOK,
// 		},
// 		// ... add more test cases for different scenarios
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T){
// 			w := httptest.NewRecorder()
// 			ctx := GetTestGinContext(w)
// 			ctx.Set("uid", "testUserID")
// 			tt.given()
// 			rest.GetProjectsByUserID(service)(ctx)
// 			assert.Equal(t, tt.expectedCode, w.Code)
// 		})
// 	}
// }

// // A helper function to return a pointer to a Role.
// func pointerToRole(role entities.MemberRole) *entities.MemberRole {
// 	r := role
// 	return &r
// }

// func TestSendInvitation(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	service := new(mocks.MockedApplicationService)

// 	tests := []struct {
// 		name         string
// 		request      entities.MemberInput
// 		given        func()
// 		expectedCode int
// 	}{
// 		{
// 			name: "Successfully send invitation",
// 			request: entities.MemberInput{
// 				UserID:    "someUserID",
// 				ProjectID: "someProjectID",
// 				Role:      pointerToRole(entities.RoleEditor),
// 			},
// 			given: func() {
// 				// Mock the service functions as needed
// 				// Example: service.On("AddMember", "someProjectID", mock.AnythingOfType("*entities.Member")).Return(nil)
// 			},
// 			expectedCode: http.StatusOK,
// 		},
// 		// ... add more test cases for different scenarios
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T){
// 			w := httptest.NewRecorder()
// 			ctx := GetTestGinContext(w)
// 			ctx.Set("uid", "testUserID")
// 			tt.given()
// 			rest.GetProjectsByUserID(service)(ctx)
// 			assert.Equal(t, tt.expectedCode, w.Code)
// 		})
// 	}
// }

// func TestAcceptInvitation(t *testing.T) {
//     mockService := new(mocks.MockedApplicationService)
    
    
//     testCases := []struct{
//         name string
//         input entities.MemberInput
//         mockServiceError error
//         expectedStatusCode int
//         expectedResponse string
//     }{
//         {
//             name: "Successfully accept invitation",
//             input: entities.MemberInput{
//                 UserID:    "testUserID",
//                 ProjectID: "testProjectID",
//             },
//             mockServiceError: nil,
//             expectedStatusCode: http.StatusOK,
//             expectedResponse: `{"message": "Successful"}`,
//         },
//         //... Add more test cases as required.
//     }
    
//     for _, tt := range testCases {
//         t.Run(tt.name, func(t *testing.T) {
//             mockService.On("UpdateInvite", tt.input.ProjectID, tt.input.UserID, entities.AcceptedInvitation, nil).Return(tt.mockServiceError)
            
//             w := httptest.NewRecorder()
//             c := GetTestGinContext(w)
            
//             body, _ := json.Marshal(tt.input)
//             req, _ := http.NewRequest(http.MethodPost, "/accept-invitation", bytes.NewBuffer(body))
//             c.Request = req
            
//         	rest.AcceptInvitation(mockService)(c)
            
//             assert.Equal(t, tt.expectedStatusCode, w.Code)
//             assert.Equal(t, tt.expectedResponse, w.Body.String())
            
//             mockService.AssertExpectations(t)
//         })
//     }
// }
