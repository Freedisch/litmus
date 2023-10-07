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
	"github.com/google/uuid"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/handlers/rest"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/entities"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/services"
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

func TestResetPassword(t *testing.T){
	// given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	ctx.Set("role", "user")
	request := &entities.UserPassword{
		Username: uuid.NewString(),
		OldPassword: uuid.NewString(),
		NewPassword: uuid.NewString(),
	}

	jsonBytes, _ := json.Marshal(request)

    // Set the request body to the mock user password request
    ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(jsonBytes))
	service := new(services.ApplicationService)
	var result entities.UserPassword
    err := ctx.BindJSON(&result)
    if err != nil {
        t.Fatal(err)
    }
	// when
	rest.ResetPassword(*service)(ctx)
	// then
	assert.Equal(t, 401, w.Code)
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