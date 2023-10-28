package controller

import (
	"app/config"
	"app/helper"
	"app/model/domain"
	"app/model/web"
	"app/utils/request"
	"app/utils/response"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create MedicalID Controller
func CreateMedicalIDController(c echo.Context) error {
	// Parse the user ID from the request parameters
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid User ID"))
	}

	// Check if the user with the provided ID exists
	var user domain.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
	}

	// Check if a MedicalID already exists for the same user
	var existingMedicalID domain.MedicalID
	if err := config.DB.Where("user_id = ?", userID).First(&existingMedicalID).Error; err == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("MedicalID for this user already exists"))
	}

	// Parse the MedicalID request from the request body
	var medicalIDRequest web.MedicalIDRequest
	if err := c.Bind(&medicalIDRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid MedicalID Input"))
	}

	// Create the MedicalID record and associate it with the user
	medicalID := request.ConvertToMedicalIDRequest(medicalIDRequest)
	medicalID.UserID = uint(userID) // Set the User ID

	if err := config.DB.Create(&medicalID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create MedicalID"))
	}

	response := response.ConvertToGetMedicalID(medicalID)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("MedicalID Created Successfully", response))
}


// Get All Medical ID
func GetAllMedicalIDController(c echo.Context) error {
	var medicalIDs []domain.MedicalID
	// maunya ada data usernya, sesuai input id
	err := config.DB.Find(&medicalIDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medical ID"))
	}

	if len(medicalIDs) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Medical ID Data"))
	}

	response := response.ConvertToGetAllMedicalIDs(medicalIDs)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medical ID Data Successfully Retrieved", response))
}

// Get Medical by ID
func GetMedicalIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var medicalID domain.MedicalID

	if err := config.DB.First(&medicalID, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medical ID"))
	}

	response := response.ConvertToGetMedicalID(&medicalID)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medical ID Data Successfully Retrieved", response))
}

// Update Medical ID by ID
func UpdateMedicalIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var updatedMedicalID domain.MedicalID

	if err := c.Bind(&updatedMedicalID); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input"))
	}

	var existingMedicalID domain.MedicalID
	result := config.DB.First(&existingMedicalID, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medical ID"))
	}

	config.DB.Model(&existingMedicalID).Updates(updatedMedicalID)

	response := response.ConvertToGetMedicalID(&existingMedicalID)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medical ID Data Successfully Updated", response))
}

// Delete Medical ID by ID
func DeleteMedicalIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var existingMedicalID domain.MedicalID
	result := config.DB.First(&existingMedicalID, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medical ID"))
	}

	config.DB.Delete(&existingMedicalID)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medical ID Data Successfully Deleted", nil))
}
