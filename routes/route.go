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

	return e

}
