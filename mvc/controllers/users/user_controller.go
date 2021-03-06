package users

import (
	"cookly/mvc/configs"
	"cookly/mvc/models/responses"
	"cookly/mvc/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	userLogin := users.UserLogin{}
	c.Bind(&userLogin)

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login Success",
		Data:    userLogin,
	})

}

func RegisterController(c echo.Context) error {
	var userRegister users.UserRegister
	c.Bind(&userRegister)

	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name cannot be blank!",
			Data:    nil,
		})
	}

	if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email cannot be blank!",
			Data:    nil,
		})
	}

	if userRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password cannot be blank!",
			Data:    nil,
		})
	}

	var userDB users.User
	userDB.Name = userRegister.Name
	userDB.Email = userRegister.Email
	userDB.Password = userRegister.Password

	result := configs.DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error when inserting user value to Database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Registration Success",
		Data:    userDB,
	})
}

func GetUserByID(c echo.Context) error {
	var usersById users.User

	id, _ := strconv.Atoi(c.Param("id"))
	result := configs.DB.First(&usersById, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get user data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success get user data",
		Data:    usersById,
	})
}
