package routes

import (
	"cookly/controllers/categories"
	"cookly/controllers/products"
	"cookly/controllers/recipes"
	"cookly/controllers/reviews"
	"cookly/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	CategoryController categories.CategoryController
	ProductController  products.ProductController
	RecipeController   recipes.RecipeController
	ReviewController   reviews.ReviewController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	baseRoute := e.Group("/api/v1")

	// Users
	userRoute := baseRoute.Group("/users")
	userRoute.POST("/login", cl.UserController.Login)
	userRoute.POST("/register", cl.UserController.Register)
	userRoute.GET("/:id", cl.UserController.GetUserByID)

	// Categories
	baseRoute.POST("/categories", cl.CategoryController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	baseRoute.GET("/categories", cl.CategoryController.GetAll)
	baseRoute.PUT("/categories/:id", cl.CategoryController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	// Products
	baseRoute.POST("/products", cl.ProductController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	baseRoute.GET("/products", cl.ProductController.GetAll)
	baseRoute.PUT("/products/:id", cl.ProductController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	// Recipes
	baseRoute.POST("/recipes", cl.RecipeController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	baseRoute.GET("/recipes", cl.RecipeController.GetAll)
	baseRoute.GET("/recipes/:id", cl.RecipeController.RecipeByID)
	baseRoute.PUT("/recipes/:id", cl.RecipeController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	// Reviews
	baseRoute.POST("/reviews", cl.ReviewController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	baseRoute.GET("/reviews/:id", cl.ReviewController.GetReviewsByRecipeID)

}
