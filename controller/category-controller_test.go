package controller

import (
	"app/model/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


// Create Category Controller Valid
func TestCreateCategoryControllerValid(t *testing.T) {
	category := domain.Category{
		Name:    "Category Test",
		Details: "Details Test",
	}
	requestData, _ := json.Marshal(category)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/categories/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateCategoryController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

// Create Category Controller Invalid
func TestCreateCategoryControllerInvalid(t *testing.T) {
	category := domain.Category{
		Name:    "Category Test",
		Details: "Details Test",
	}
	requestData, _ := json.Marshal(category)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/categories/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Get All Category Controller Valid
func TestGetAllCategoriesControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/categories/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetAllCategoriesController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Get All Medical ID Controller Invalid
func TestGetAllCategoriesControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/categories/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetAllCategoriesController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// Get Category by ID Controller Valid
func TestGetCategoryControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/categories/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := GetCategoryController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Get Category by ID Controller Invalid
func TestGetCategoryControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/categories/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := GetCategoryController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Update Category by ID Controller Valid
func TestUpdateCategoryControllerValid(t *testing.T) {
	category := domain.Category{
		Name:    "Category Test",
		Details: "Details Test",
	}
	requestData, _ := json.Marshal(category)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/categories/:id/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := UpdateCategoryController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Update Category by ID Controller Invalid
func TestUpdateCategoryControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut,  "/categories/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := UpdateCategoryController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Delete Category by ID Controller Valid
func TestDeleteCategoryControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/categories/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := DeleteCategoryController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Delete Category by ID Controller Invalid
func TestDeleteCategoryControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/categories/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := DeleteCategoryController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
