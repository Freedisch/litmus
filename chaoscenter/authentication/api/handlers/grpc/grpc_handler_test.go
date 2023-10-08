package grpc_test

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/handlers/grpc"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/mocks"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/presenter/protos"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServer struct {
	grpc.ServerGrpc
}

func (m *MockServer) ValidateToken(token string) (*jwt.Token, error) {
	claims := jwt.MapClaims{
		"uid": "12345",
	}
	return &jwt.Token{
		Claims: claims,
	}, nil
}

// Mocking RbacValidator function
func (m *MockServer) RbacValidator(uid, projectId string, requiredRoles []string, invitation bool, service interface{}) error {
	return nil
}

// func TestValidateRequest_Success(t *testing.T) {
// 	s := &MockServer{}
// 	req := &protos.ValidationRequest{
// 		Jwt:           "valid-jwt",
// 		ProjectId:     "project-id",
// 		RequiredRoles: []string{"admin"},
// 		Invitation:    "true",
// 	}

// 	ctx := context.Background()

// 	// Call the function
// 	resp, err := s.ValidateRequest(ctx, req)

// 	// Assertions
// 	assert.Nil(t, err, "Expected no error")
// 	assert.NotNil(t, resp, "Expected non-nil response")
// 	assert.True(t, resp.IsValid, "Expected the request to be valid")
// 	assert.Empty(t, resp.Error, "Expected error message to be empty")
// }


// Mocking ApplicationService methods
type MockApplicationService struct {
	mock.Mock
}

func TestGetProjectById_Success(t *testing.T) {
	s := &grpc.ServerGrpc{
		ApplicationService: &mocks.MockedApplicationService{},
	}

	// Mocking ApplicationService methods
	mockService := s.ApplicationService.(*mocks.MockedApplicationService)
	mockService.On("GetProjectByProjectID", "project-id").Return(&entities.Project{
		ID:      "project-id",
		Name:    "test-project",
		Members: []*entities.Member{
			{
				UserID:    "user-1",
				Invitation: entities.PendingInvitation,  // adjust to your actual type
				JoinedAt:  1234567890,
			},
		},
	}, nil)
	mockService.On("FindUsersByUID", []string{"user-1"}).Return(&[]entities.User{
		{
			ID:       "user-1",
			Email:    "user1@email.com",
			Username: "user1",
		},
	}, nil)

	req := &protos.GetProjectByIdRequest{
		ProjectID: "project-id",
	}

	ctx := context.Background()

	resp, err := s.GetProjectById(ctx, req)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "test-project", resp.Name)
	assert.Equal(t, "user1@email.com", resp.Members[0].Email)
}