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

// Create Category
func CreateCategoryController(c echo.Context) error {
	var categoryRequest web.CategoryRequest

	if err := c.Bind(&categoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input"))
	}

	category := request.ConvertToCategoryRequest(categoryRequest)

	if existingCategory := config.DB.Where("name = ?", category.Name).First(&category).Error; existingCategory == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("Category Already Exist"))
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Category"))
	}

	response := response.ConvertToGetCategory(category)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success Created Category", response))
}

// Get All Categories
func GetAllCategoriesController(c echo.Context) error {
	var categories []domain.Category

	err := config.DB.Find(&categories).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Categories"))
	}

	if len(categories) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Categories Data"))
	}

	response := response.ConvertToGetAllCategories(categories)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Categories Data Successfully Retrieved", response))
}

// Get Category by ID
func GetCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var category domain.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Category"))
	}

	response := response.ConvertToGetCategory(&category)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Category Data Successfully Retrieved", response))
}

// Update Category by ID
func UpdateCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var updatedCategory domain.Category

	if err := c.Bind(&updatedCategory); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input"))
	}

	var existingCategory domain.Category
	result := config.DB.Model(&existingCategory).First(&existingCategory, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Category"))
	}

	config.DB.Model(&existingCategory).Updates(updatedCategory)

	response := response.ConvertToGetCategory(&existingCategory)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Category Data Successfully Updated", response))
}

// Delete Category by ID
func DeleteCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var existingCategory domain.Category
	result := config.DB.Model(&existingCategory).First(&existingCategory, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Category"))
	}
	config.DB.Delete(&existingCategory)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Category Data Successfully Deleted", nil))
}
