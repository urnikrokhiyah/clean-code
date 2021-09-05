package controllers

import (
	"cleancode/lib/databases"
	"cleancode/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsersController(c echo.Context) error {
	users, err := databases.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse("can't get all users"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("success get all users", users))
}

func GetSingleUserController(c echo.Context) error {
	userId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse("invalid user id"))
	}

	user, err := databases.GetSingleUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse("can't get single user"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("success get single user", user))
}
