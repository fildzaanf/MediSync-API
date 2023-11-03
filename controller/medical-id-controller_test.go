package controller

import (
	"app/model/domain"
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


// Create MedicalID Controller Valid
func TestCreateMedicalIDControllerValid(t *testing.T) {
	medical_id := domain.MedicalID{
		Birthdate:        "2002-09-03",
		Gender:           "male",
		BloodType:        "A",
		Height:           170,
		Weight:           50,
		MedicalCondition: "Good",
		EmergencyContact: "089",
	}
	requestData, _ := json.Marshal(medical_id)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/medicalids/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := CreateMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

// Create MedicalID Controller Invalid
func TestCreateMedicalIDControllerInvalid(t *testing.T) {
	medical_id := domain.MedicalID{
		Birthdate:        "2002-09-03",
		Gender:           "male",
		BloodType:        "A",
		Height:           170,
		Weight:           50,
		MedicalCondition: "Good",
		EmergencyContact: "089",
	}
	requestData, _ := json.Marshal(medical_id)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/medicalids/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("z")

	err := CreateMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// Get All Medical ID Controller Valid
func TestGetAllMedicalIDControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/medicalids/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetAllMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// // Get All Medical ID Controller Invalid
// func TestGetAllMedicalIDControllerInvalid(t *testing.T) {
// 	e := InitTestDB()
// 	req := httptest.NewRequest(http.MethodGet, "/medicalids/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	fmt.Println(rec.Code)

// 	err := GetAllMedicalIDController(c)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusNotFound, rec.Code)
// }

// Get Medical ID by ID Controller Valid
func TestGetMedicalIDControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/medicalids/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := GetMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Get Medical ID by ID Controller Invalid
func TestGetMedicalIDControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/medicalids/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("z")

	err := GetMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code) 
}

// Update Medical ID by ID Controller Valid
func TestUpdateMedicalIDControllerValid(t *testing.T) {
	medical_id := domain.MedicalID{
		Birthdate:        "2002-09-03",
		Gender:           "male",
		BloodType:        "A",
		Height:           170,
		Weight:           50,
		MedicalCondition: "Good",
		EmergencyContact: "0897",
	}
	requestData, _ := json.Marshal(medical_id)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/medicalids/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := UpdateMedicalIDController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Update Medical ID by ID Controller Invalid
func TestUpdateMedicalIDControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/medicalids/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("z")

	err := UpdateMedicalIDController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Delete Medical ID by ID Controller Valid
func TestDeleteMedicalIDControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/medicalids/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := DeleteMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}


// Delete Medical ID by ID Controller Invalid
func TestDeleteMedicalIDControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/medicalids/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("z")

	err := DeleteMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}