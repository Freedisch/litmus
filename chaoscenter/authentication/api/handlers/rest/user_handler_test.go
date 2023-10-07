package rest_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/handlers/rest"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/mocks"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/entities"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/services"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestMain is the entry point for testing
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}


func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/josn")
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("role", "admin")
	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}
type MockApplicationService interface {
	mock.Mock
}
// func TestCreateUser(t *testing.T) {
// 		// Create a mock service
// 		mockService := &MockApplicationService{}
	
// 		// Create a new HTTP request
// 		req, err := http.NewRequest("POST", "/users", nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
	
// 		// Create a new HTTP recorder
// 		rr := httptest.NewRecorder()
	
// 		// Create a new Gin context
// 		c, _ := gin.CreateTestContext(rr)
// 		c.Request = req
	
// 		// Call the CreateUser function
// 		rest.CreateUser(mockService)(c)
	
// 		// Check the response status code
// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("handler returned wrong status code: got %v want %v",
// 				status, http.StatusOK)
// 		}
	
// 		// Check the response body
// 		expected := `{"message":"User created successfully"}`
// 		if rr.Body.String() != expected {
// 			t.Errorf("handler returned unexpected body: got %v want %v",
// 				rr.Body.String(), expected)
// 		}
	
	
// }

func TestUpdateuser(t *testing.T) {
	// given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	service := new(services.ApplicationService)
	// when
	rest.UpdateUser(*service)(ctx) // pass the interface to the UpdateUser function
	// then
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// func TestGetUser(t *testing.T) {
// 	// given
// 	w := httptest.NewRecorder()
// 	ctx := GetTestGinContext(w)
// 	service := new(services.ApplicationService)
// 	key := uuid.NewString()
// 	ctx.Params = []gin.Param{
// 		{
// 			Key:  "uid",
// 			Value: key,
// 		},
// 	}
// 	// when
// 	rest.GetUser(*service)(ctx) // pass the interface to the GetUser function
// 	// then
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }

// func TestFetchUsers(t *testing.T) {
// 	// given
// 	w := httptest.NewRecorder()
// 	ctx := GetTestGinContext(w)
// 	ctx.Set("role", "admin")
// 	service := new(services.ApplicationService)
// 	// when
// 	rest.FetchUsers(*service)(ctx) // pass the interface to the FetchUsers function
// 	// then
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }

func TestLoginUser(t *testing.T) {
	// given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	service := new(services.ApplicationService)
	// when
	rest.LoginUser(*service)(ctx)
	// then
	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestLogoutUser(t *testing.T) {
	// given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	service := new(services.ApplicationService)
	// when
	rest.LogoutUser(*service)(ctx)
	// then
	assert.Equal(t, 401, w.Code)
}

func TestUpdatePassword(t *testing.T){
	// given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	service := new(services.ApplicationService)
	// when
	rest.LogoutUser(*service)(ctx)
	// then
	assert.Equal(t, 401, w.Code)
}



func TestResetPassword_Success(t *testing.T) {
	// Set up Gin engine for testing
	gin.SetMode(gin.TestMode)

	// Mock service methods
	service := new(mocks.MockedApplicationService)
	service.On("IsAdministrator", mock.AnythingOfType("*entities.User")).Return(nil)
	service.On("UpdatePassword", mock.AnythingOfType("*entities.UserPassword"), false).Return(nil)

	// Create a request with valid JSON payload for UserPassword
	userPass := &entities.UserPassword{
		Username:    "testUser",
		OldPassword: "",
		NewPassword: "validPassword123",
	}
	bodyBytes, _ := json.Marshal(userPass)

	// Mock context values
	w := httptest.NewRecorder()
	c := GetTestGinContext(w)
	c.Request.Method = http.MethodPost
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("role", "admin")
	c.Set("uid", "testUID")
	c.Set("username", "adminUser")

	// Use the handler function directly to serve the request
	rest.ResetPassword(service)(c)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "password has been reset successfully", response["message"])
}

func TestResetPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := new(mocks.MockedApplicationService)

	tests := []struct {
		name          string
		inputBody     string
		mockRole      string
		mockUID       string
		mockUsername  string
		mockService   func()
		expectedCode  int
	}{
		{
			name:         "Non-admin role",
			inputBody:    `{"Username": "john", "NewPassword": "password123"}`,
			mockRole:     "user",
			mockUID:      "testUID",
			mockUsername: "user",
			expectedCode: utils.ErrorStatusCodes[utils.ErrUnauthorized],
		},
		{
			name:         "Invalid Request Body",
			inputBody:    `{"invalid}`,
			mockRole:     "admin",
			mockUID:      "testUID",
			mockUsername: "adminUser",
			expectedCode: utils.ErrorStatusCodes[utils.ErrInvalidRequest],
		},
		// {
		// 	name:         "Strict Password Policy Violation",
		// 	inputBody:    `{"Username": "john", "NewPassword": "short"}`,
		// 	mockRole:     "admin",
		// 	mockUID:      "testUID",
		// 	mockUsername: "adminUser",
		// 	expectedCode: utils.ErrorStatusCodes[utils.ErrStrictPasswordPolicyViolation],
		// },
		{
			name:         "Empty Username or Password",
			inputBody:    `{"Username": "", "NewPassword": ""}`,
			mockRole:     "admin",
			mockUID:      "testUID",
			mockUsername: "adminUser",
			expectedCode: utils.ErrorStatusCodes[utils.ErrInvalidRequest],
		},
		{
			name:         "Non-administrator user attempting to change password",
			inputBody:    `{"Username": "john", "NewPassword": "password123"}`,
			mockRole:     "user",
			mockUID:      "testUID",
			mockUsername: "john",
			mockService: func() {
				service.On("IsAdministrator", mock.AnythingOfType("*entities.User")).Return(nil)
			},
			expectedCode: utils.ErrorStatusCodes[utils.ErrUnauthorized],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockService != nil {
				tt.mockService()
			}
			service.On("IsAdministrator", mock.AnythingOfType("*entities.User")).Return(nil)
			service.On("UpdatePassword", mock.AnythingOfType("*entities.UserPassword"), false).Return(nil)
			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			c.Request.Method = http.MethodPost
			c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(tt.inputBody)))
			c.Set("role", tt.mockRole)
			c.Set("uid", tt.mockUID)
			c.Set("username", tt.mockUsername)

			rest.ResetPassword(service)(c)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestUpdateUserState(t *testing.T){
	// given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	ctx.Set("role", "user")
	service := new(services.ApplicationService)
	// when
	rest.UpdateUserState(*service)(ctx)
	// then
	assert.Equal(t, 401, w.Code)
}

func TestCreateApiToken(t *testing.T){
	// given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	service := new(services.ApplicationService)
	// when
	rest.CreateApiToken(*service)(ctx)
	// then
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// func TestGetApiTokens(t *testing.T){
// 	// given
// 	w := httptest.NewRecorder()
// 	ctx := GetTestGinContext(w)
// 	service := new(services.ApplicationService)
// 	// when
// 	rest.GetApiTokens(*service)(ctx)
// 	// then
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }

func TestDeleteApiToken(t *testing.T){
	// given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	service := new(services.ApplicationService)
	// when
	rest.DeleteApiToken(*service)(ctx)
	// then
	assert.Equal(t, http.StatusBadRequest, w.Code)
}