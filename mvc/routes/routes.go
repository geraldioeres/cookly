package routes

// import (
// 	"cookly/controllers/products"
// 	recipecategory "cookly/controllers/recipe_category"
// 	"cookly/controllers/recipes"
// 	"cookly/controllers/users"
// 	"cookly/controllers/reviews"

// 	"github.com/labstack/echo/v4"
// )

// func InitRoutes() *echo.Echo {
// 	e := echo.New()
// 	baseRoute := e.Group("/api/v1")

// 	// Users
// 	baseRoute.POST("/register", users.RegisterController)
// 	baseRoute.POST("/login", users.LoginController)
// 	baseRoute.GET("/users/:id", users.GetUserByID)

// 	// Recipes
// 	baseRoute.POST("/recipes", recipes.CreateRecipeController)
// 	baseRoute.GET("/recipes", recipes.GetRecipesController)
// 	baseRoute.GET("/recipes/:id", recipes.GetRecipeByID)

// 	// Recipe Categories
// 	baseRoute.POST("/categories", recipecategory.CreateCategoryController)
// 	baseRoute.GET("/categories", recipecategory.GetAllCategories)

// 	// Products
// 	baseRoute.POST("/products", products.CreateProduct)
// 	baseRoute.GET("/products", products.GetAllProducts)

// 	// Review
// 	baseRoute.POST("/reviews", reviews.CreateReview)
// 	// baseRoute.GET("/reviews/:id", reviews.ReviewsByRecipeID)

// 	return e
// }
