package controller

import (
	"app/config"
	"fmt"
	"app/model/domain"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitTestDB() *echo.Echo {
	config.ConnectDBTest()
	e := echo.New()
	return e
}

// Register User Controller Valid
func TestRegisterUserControllerValid(t *testing.T) {
	user := domain.User{
		Fullname: "User Test",
		Email:    "usertest@upi.edu",
		Password: "passwordtest",
	}
	requestData, _ := json.Marshal(user)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := RegisterUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

// Register User Controller Invalid
func TestRegisterUserControllerInvalid(t *testing.T) {
	user := domain.User{
		Fullname: "User Test",
		Email:    "usertest@gmail.com",
		Password: "passwordtest",
	}
	requestData, _ := json.Marshal(user)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/users/register/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := RegisterUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, rec.Code)
}

// Login User Controller Valid
func TestLoginUserControllerValid(t *testing.T) {
	loginUser := `{"Email": "testloginuser@gmail.com", "Password": "passwordtest"}`
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/users/login/", strings.NewReader(loginUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) 	
	rec := httptest.NewRecorder()
	c := e.NewContext(req,rec)
	
	err := LoginUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Login User Controller Invalid
func TestLoginUserControllerInvalid(t *testing.T) {
	loginUser := `{"Username": "", "Password": ""}`
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/users/login/", strings.NewReader(loginUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) 	
	rec := httptest.NewRecorder()
	c := e.NewContext(req,rec)
	
	err := LoginUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Get All Users Controller Valid
func TestGetAllUsersControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetAllUsersController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Get All Users Controller Invalid
func TestGetAllUsersControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetAllUsersController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Get User By ID Controller Valid
func TestGetUserControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := GetUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Get User By ID Controller Invalid
func TestGetUserControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := GetUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Update User by ID Controller Valid
func TestUpdateUserControllerValid(t *testing.T) {
	user := domain.User{
		Fullname:     "UpdateTest",
		Email:    "TestingUpdate@email.com",
		Password: "PasswordUpdateTest",
	}
	requestData, _ := json.Marshal(user)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/users/:id/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("11")

	err := UpdateUserController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Update User by ID Controller Invalid
func TestUpdateUserControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/users/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := UpdateUserController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Delete User by ID Controller Valid
func TestDeleteUserControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := DeleteUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Delete User by ID Controller Invalid
func TestDeleteUserControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := DeleteUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

