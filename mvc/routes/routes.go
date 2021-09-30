package routes

import (
	"cookly/mvc/controllers/products"
	recipecategory "cookly/mvc/controllers/recipe_category"
	"cookly/mvc/controllers/recipes"
	"cookly/mvc/controllers/reviews"
	"cookly/mvc/controllers/users"

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

	// Recipe Categories
	baseRoute.POST("/categories", recipecategory.CreateCategoryController)
	baseRoute.GET("/categories", recipecategory.GetAllCategories)

	// Products
	baseRoute.POST("/products", products.CreateProduct)
	baseRoute.GET("/products", products.GetAllProducts)

	// Review
	baseRoute.POST("/reviews", reviews.CreateReview)
	baseRoute.GET("/reviews/:id", reviews.GetReviewByRecipeID)

	return e
}
