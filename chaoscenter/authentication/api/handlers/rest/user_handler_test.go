package rest_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/api/handlers/rest"
	"github.com/litmuschaos/litmus/chaoscenter/authentication/pkg/services"
	"github.com/stretchr/testify/assert"
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

func TestCreateUser(t *testing.T) {
	//given
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	ctx.Set("role", "admin")
	service := new(services.ApplicationService)
	//userResponse, err := service.CreateUser(nil)
	//when
	rest.CreateUser(*service)(ctx) // pass the interface to the CreateUser function
	//then
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

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
	service := new(services.ApplicationService)
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