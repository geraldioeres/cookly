package controllers

import (
	users "cookly/models/Users"
	"cookly/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRegiterController(c echo.Context) error {
	var UserRegister users.UserRegister
	c.Bind(&UserRegister)

	if UserRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code: http.StatusBadRequest,
			Message: "Name cannot blank",
			Data: nil,
		})
	}

	var userDB users.Users
	userDB.Name = UserRegister.Name
	userDB.Email = UserRegister.Email
	userDB.Password = UserRegister.Password

	result := 
}
