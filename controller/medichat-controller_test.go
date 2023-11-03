package controller

import (
	"app/model/domain"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Create Medichat Controller Valid
func TestCreateMediChatControllerValid(t *testing.T) {
	medichat := domain.MediChat{
		Message:        "Apa obat penghilang sakit kepala?",
	}
	requestData, _ := json.Marshal(medichat)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/medichats", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)


	err := CreateMediChatController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Create Medichat Controller Invalid
func TestCreateMediChatControllerInvalid(t *testing.T) {
	medichat := domain.MediChat{
		Message:        "Berikan saya film terbaik?",
	}
	requestData, _ := json.Marshal(medichat)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost,  "/medichats", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	err := CreateMediChatController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

