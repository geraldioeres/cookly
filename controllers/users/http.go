package users

import (
	"cookly/business/users"
	"cookly/controllers"
	"cookly/controllers/users/requests"
	"cookly/controllers/users/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.UseCase
}

func NewUserController(userUseCase users.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (loginController *UserController) Login(c echo.Context) error {
	var userLogin requests.User
	c.Bind(&userLogin)

	ctx := c.Request().Context()

	user, err := loginController.UserUseCase.Login(ctx, userLogin.Email, userLogin.Password)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.LoginFromDomain(user))
}

func (registerController *UserController) Register(c echo.Context) error {
	userRegister := requests.User{}
	c.Bind(&userRegister)

	ctx := c.Request().Context()

	err := registerController.UserUseCase.Register(ctx, userRegister.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully registered")
}

func (getUserByIDController *UserController) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := getUserByIDController.UserUseCase.GetUserByID(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(user))
}
