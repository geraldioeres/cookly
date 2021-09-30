package routes

import (
	"cookly/controllers/categories"
	"cookly/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	baseRoute := e.Group("/api/v1")

	// User
	userRoute := baseRoute.Group("/users")
	userRoute.POST("/login", cl.UserController.Login)
	userRoute.POST("/register", cl.UserController.Register)
	userRoute.GET("/:id", cl.UserController.GetUserByID)

	// Category
	baseRoute.POST("/categories", cl.CategoryController.Create)
}
