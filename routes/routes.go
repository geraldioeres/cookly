package routes

import (
	"cookly/controllers/recipes"
	"cookly/controllers/users"

	"github.com/labstack/echo/v4"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	baseRoute := e.Group("/api/v1")

	// Users
	baseRoute.POST("/register", users.RegisterController)
	baseRoute.POST("/login", users.LoginController)
	baseRoute.GET("/users/:id", users.GetUserByID)

	// Recipes
	recipeRoute := baseRoute.Group("/recipes")
	recipeRoute.GET("", recipes.GetRecipesController)
	return e
}
