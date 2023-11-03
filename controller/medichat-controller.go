package controller

import (
	"app/utils/response"
	// "app/utils/request"
	"app/config"
	"app/helper"
	"app/model/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create Chat
func CreateMediChatController(c echo.Context) error {

	var input web.MediChatRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input"))
	}

	medichat, err := OpenAIMediChatResponseController(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.MediChatResponse{
			Response: "Failed to Provide Response MediChat",
		})
	}

	if err := config.DB.Create(&medichat).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Provide Response MediChat"))
	}

	response := response.ConvertToMediChatResponse(medichat)

	return c.JSON(http.StatusOK, helper.SuccessResponse("MediChat Successful", response))

}
