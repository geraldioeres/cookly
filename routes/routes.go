package routes

import (
	"cookly/controllers/ingredients"
	recipecategory "cookly/controllers/recipe_category"
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
	baseRoute.POST("/recipes", recipes.CreateRecipeController)
	baseRoute.GET("/recipes", recipes.GetRecipesController)
	baseRoute.GET("/recipes/:id", recipes.GetRecipeByID)

	// Recipe Category
	baseRoute.POST("/category", recipecategory.CreateCategoryController)
	baseRoute.GET("/category", recipecategory.GetAllCategories)

	// Ingredients
	baseRoute.POST("/ingredient", ingredients.CreateIngredient)
	baseRoute.GET("/ingredient", ingredients.GetAllIngredients)
	return e
}
