package controllers

import (
	"cleancode/lib/databases"
	"cleancode/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsersController(c echo.Context) error {
	users, err := databases.GetAllUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse("can't get all users"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("success get all users", users))
}
