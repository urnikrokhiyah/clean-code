package routes

import (
	"cleancode/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetAllUsersController)
	e.GET("/users/:id", controllers.GetSingleUserController)
	e.POST("/users", controllers.CreateUserControllers)
	return e
}
