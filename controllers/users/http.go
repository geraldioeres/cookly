package users

import (
	"cookly/business/users"
	"cookly/controllers"
	"cookly/controllers/users/requests"
	"cookly/controllers/users/responses"
	"net/http"

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

func (userController UserController) Login(c echo.Context) error {
	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()

	user, error := userController.UserUseCase.Login(ctx, userLogin.Email, userLogin.Password)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(user))
}
