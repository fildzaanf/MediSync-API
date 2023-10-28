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

// Create Medicine
func CreateMedicineController(c echo.Context) error {
	var medicineRequest web.MedicineRequest

	if err := c.Bind(&medicineRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input"))
	}

	medicine := request.ConvertToMedicineRequest(medicineRequest)

	if existingBatchNumber := config.DB.Where("batch_number = ?", medicine.BatchNumber).First(&medicine).Error; existingBatchNumber == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("Batch Number Already Exist"))
	}

	if err := config.DB.Create(&medicine).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Medicine"))
	}

	response := response.ConvertToGetMedicine(medicine)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success Created Medicine", response))
}

// Get All Medicines
func GetAllMedicinesController(c echo.Context) error {
	var medicines []domain.Medicine

	err := config.DB.Find(&medicines).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicines"))
	}

	if len(medicines) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Medicines Data"))
	}

	response := response.ConvertToGetAllMedicines(medicines)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicines Data Successfully Retrieved", response))
}

// Get Medicine by ID
func GetMedicineController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var medicine domain.Medicine

	if err := config.DB.First(&medicine, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine"))
	}

	response := response.ConvertToGetMedicine(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Data Successfully Retrieved", response))
}

// Update Medicine by ID
func UpdateMedicineController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var updatedMedicine domain.Medicine

	if err := c.Bind(&updatedMedicine); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input"))
	}

	var existingMedicine domain.Medicine
	result := config.DB.First(&existingMedicine, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine"))
	}

	config.DB.Model(&existingMedicine).Updates(updatedMedicine)

	response := response.ConvertToGetMedicine(&existingMedicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Data Successfully Updated", response))
}

// Delete Medicine by ID
func DeleteMedicineController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var existingMedicine domain.Medicine
	result := config.DB.First(&existingMedicine, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine"))
	}

	config.DB.Delete(&existingMedicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Data Successfully Deleted", nil))
}
