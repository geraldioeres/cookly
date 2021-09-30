package routes

import (
	"cookly/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	baseRoute := e.Group("/api/v1")
	
	userRoute := baseRoute.Group("/users")
	userRoute.POST("/login", cl.UserController.Login)
}
