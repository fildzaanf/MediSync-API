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


// Create Schedule Controller Valid
func TestCreateScheduleControllerValid(t *testing.T) {
	schedule := domain.Schedule{
		Name:    "Schedule Test",
		Details: "Detail Test",
		Minute:  "07",
		Hour:    "07",
		Day:     "1",
		Email:   "hanisahfildza@upi.edu",
		Subject: "Reminder",
	}
	requestData, _ := json.Marshal(schedule)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/schedules/:id/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := CreateScheduleController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

// Create Schedule Controller Invalid
func TestCreateScheduleControllerInvalid(t *testing.T) {
	schedule := domain.Schedule{
		Name:    "Schedule Test",
		Details: "Detail Test",
		Minute:  "07",
		Hour:    "07",
		Day:     "1",
		Email:   "hanisahfildza@upi.edu",
		Subject: "Reminder",
	}
	requestData, _ := json.Marshal(schedule)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/schedules/:id/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := CreateMedicalIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// Get All Schedules Controller Valid
func TestGetAllSchedulesControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/schedules/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetAllSchedulesController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Get All Schedules Controller Invalid
func TestGetAllSchedulesControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/schedules/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetAllSchedulesController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// Get Schedule by ID Controller Valid
func TestGetScheduleControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/schedules/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := GetScheduleController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Get Schedule by ID Controller Invalid
func TestGetScheduleControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/schedules/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := GetScheduleController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Update Schedule by ID Controller Valid
func TestUpdateScheduleControllerValid(t *testing.T) {
	schedule := domain.Schedule{
		Name:    "Schedule Test",
		Details: "Detail Test",
		Minute:  "07",
		Hour:    "07",
		Day:     "1",
		Email:   "hanisahfildza@upi.edu",
		Subject: "Reminder",
	}
	requestData, _ := json.Marshal(schedule)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/schedules/:id/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := UpdateScheduleController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Update Schedule by ID Controller Invalid
func TestUpdateScheduleControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/schedules/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := UpdateScheduleController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// Delete Schedule by ID Controller Valid
func TestDeleteScheduleControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/schedules/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := DeleteScheduleController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Delete Schedule by ID Controller Invalid
func TestDeleteScheduleControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/schedules/:id/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := DeleteScheduleController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
