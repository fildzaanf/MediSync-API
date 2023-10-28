package routes

import (
	"app/controller"
	"app/middleware"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	middlewares "github.com/labstack/echo/v4/middleware"
)

func Routes() *echo.Echo {

	e := echo.New()

	e.GET("/medisync/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to MediSync API")
	})

	middleware.AddTrailingSlash(e)
	middleware.Logger(e)
	middleware.RateLimiter(e)
	middleware.Recover(e)
	middleware.CORS(e)
	middleware.NotFoundHandler(e)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	JWT := middlewares.JWT([]byte(os.Getenv("SECRET_KEY")))

	gUsers := e.Group("/users")
	gUsers.POST("/register/", controller.RegisterUserController)
	gUsers.POST("/login/", controller.LoginUserController)
	gUsers.GET("/", controller.GetAllUsersController, JWT)
	gUsers.GET("/:id/", controller.GetUserController, JWT)
	gUsers.PUT("/:id/", controller.UpdateUserController, JWT)
	gUsers.DELETE("/:id/", controller.DeleteUserController, JWT)

	gMedicalIDs := e.Group("/medicalids")
	gMedicalIDs.POST("/:id/", controller.CreateMedicalIDController, JWT)
	gMedicalIDs.GET("/", controller.GetAllMedicalIDController, JWT)
	gMedicalIDs.GET("/:id/", controller.GetMedicalIDController, JWT)
	gMedicalIDs.PUT("/:id/", controller.UpdateMedicalIDController, JWT)
	gMedicalIDs.DELETE("/:id/", controller.DeleteMedicalIDController, JWT)

	gMedicines := e.Group("/medicines")
	gMedicines.POST("/:id/", controller.CreateMedicineController, JWT)
	gMedicines.GET("/", controller.GetAllMedicinesController, JWT)
	gMedicines.GET("/:id/", controller.GetMedicineController, JWT)
	gMedicines.PUT("/:id/", controller.UpdateMedicineController, JWT)
	gMedicines.DELETE("/:id/", controller.DeleteMedicineController, JWT)

	gSchedules := e.Group("/schedules")
	gSchedules.POST("/:id/", controller.CreateScheduleController, JWT)
	gSchedules.GET("/", controller.GetAllSchedulesController, JWT)
	gSchedules.GET("/:id/", controller.GetScheduleController, JWT)
	gSchedules.PUT("/:id/", controller.UpdateScheduleController, JWT)
	gSchedules.DELETE("/:id/", controller.DeleteScheduleController, JWT)

	gCategories := e.Group("/categories")
	gCategories.POST("/", controller.CreateCategoryController, JWT)
	gCategories.GET("/", controller.GetAllCategoriesController, JWT)
	gCategories.GET("/:id/", controller.GetCategoryController, JWT)
	gCategories.PUT("/:id/", controller.UpdateCategoryController, JWT)
	gCategories.DELETE("/:id/", controller.DeleteCategoryController, JWT)

	return e

}
