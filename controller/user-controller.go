package controller

import (
	"app/config"
	"app/helper"
	"app/middleware"
	"app/model/domain"
	"app/model/web"
	"app/utils/request"
	"app/utils/response"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Register User
func RegisterUserController(c echo.Context) error {
	var userCreateRequest web.UserCreateRequest

	if err := c.Bind(&userCreateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Register Input"))
	}

	user := request.ConvertToUserCreateRequest(userCreateRequest)

	if existingUser := config.DB.Where("email = ?", user.Email).First(&user).Error; existingUser == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))
	}

	user.Password = helper.HashPassword(user.Password)
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Account"))
	}

	response := response.ConvertToGetUser(user)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success Created Account", response))
}

// Login User
func LoginUserController(c echo.Context) error {
	var userLoginRequest web.UserLoginRequest

	if err := c.Bind(&userLoginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Login Input"))
	}

	var user domain.User
	if err := config.DB.Where("email = ?", userLoginRequest.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Email Not Registered"))
	}

	if err := helper.ComparePassword(user.Password, userLoginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Incorrect Email or Password"))
	}

	userLoginRequest.Password = helper.HashPassword(userLoginRequest.Password)

	token := middleware.CreateToken(int(user.ID), user.Email)

	response := web.UserLoginResponse{
		Email:    user.Email,
		Password: userLoginRequest.Password,
		Token:    token,
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Welcome to MediSync! Login Successful", response))
}

// Get All Users
func GetAllUsersController(c echo.Context) error {
	var users []domain.User

	err := config.DB.Preload("MedicalID").Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Users"))
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Users Data"))
	}

	response := response.ConvertToGetAllUsers(users)

	return c.JSON(http.StatusOK, helper.SuccessResponse("User Data Successfully Retrieved", response))
}

// Get User by ID
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid User ID"))
	}

	var user domain.User

	if err := config.DB.Preload("MedicalID").First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve User"))
	}

	response := response.ConvertToGetUser(&user)

	return c.JSON(http.StatusOK, helper.SuccessResponse("User Data Successfully Retrieved", response))
}

// Update User by ID
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid User ID"))
	}

	var updatedUser domain.User

	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid User Input"))
	}

	var existingUser domain.User
	result := config.DB.Preload("MedicalID").First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve User"))
	}

	updatedUser.Password = helper.HashPassword(updatedUser.Password)

	config.DB.Model(&existingUser).Updates(updatedUser)

	response := response.ConvertToGetUser(&existingUser)

	return c.JSON(http.StatusOK, helper.SuccessResponse("User Data Successfully Updated", response))
}

// Delete User by ID
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var existingUser domain.User
	result := config.DB.Preload("MedicalID").First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve User"))
	}

	config.DB.Delete(&existingUser)

	return c.JSON(http.StatusOK, helper.SuccessResponse("User Data Successfully Deleted", nil))
}
