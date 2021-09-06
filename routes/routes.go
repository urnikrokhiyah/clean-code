package routes

import (
	"cleancode/constants"
	"cleancode/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/login", controllers.LoginUserController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// user controller with auth
	r.GET("/users/:id", controllers.GetSingleUserController)
	r.GET("/users", controllers.GetAllUsersController)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)

	// user controller without auth
	e.POST("/users", controllers.CreateUserControllers)

	return e
}
