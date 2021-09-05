package controllers

import (
	"cleancode/lib/databases"
	"cleancode/models"
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

func CreateUserControllers(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	newUser, err := databases.CreateNewUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, response.ErrorResponse("can't create new user"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("success create new user", newUser))
}
