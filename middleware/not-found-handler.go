package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotFoundHandlers(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				if he.Code == http.StatusNotFound {
					errorMessage := "Invalid Endpoint"
					return c.JSON(http.StatusNotFound, map[string]interface{}{
						"message": errorMessage,
					})
				}
			}

			fmt.Println("Error:", err)
		}

		return err
	}
}

func NotFoundHandler(e *echo.Echo) {
	e.Use(NotFoundHandlers)
}