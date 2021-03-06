package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_routes "cookly/app/routes"
	_mysqlDriver "cookly/drivers/mysql"

	_userUsecase "cookly/business/users"
	_userController "cookly/controllers/users"
	_userRepository "cookly/drivers/databases/users"

	_categoryUsecase "cookly/business/categories"
	_categoryController "cookly/controllers/categories"
	_categoryRepository "cookly/drivers/databases/categories"

	_productUsecase "cookly/business/products"
	_productController "cookly/controllers/products"
	_productRepository "cookly/drivers/databases/products"

	_recipeUsecase "cookly/business/recipes"
	_recipeController "cookly/controllers/recipes"
	_recipeRepository "cookly/drivers/databases/recipes"

	_reviewUsecase "cookly/business/reviews"
	_reviewController "cookly/controllers/reviews"
	_reviewRepository "cookly/drivers/databases/reviews"

	_middleware "cookly/app/middleware"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN or DEBUG mode")
	}
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:      viper.GetString(`jwt.secret`),
		ExiresDuration: viper.GetInt(`jwt.expired`),
	}

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(db)
	userUseCase := _userUsecase.NewUserUseCase(userRepository, timeoutContext, &configJWT)
	userController := _userController.NewUserController(userUseCase)

	categoryRepository := _categoryRepository.NewMysqlCategoryRepository(db)
	categoryUseCase := _categoryUsecase.NewCategoryUseCase(categoryRepository, timeoutContext)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	productRepository := _productRepository.NewMysqlProductRepository(db)
	productUseCase := _productUsecase.NewProductUseCase(productRepository, timeoutContext)
	productController := _productController.NewProductController(productUseCase)

	recipeRepository := _recipeRepository.NewMysqlRecipeRepository(db)
	recipeUseCase := _recipeUsecase.NewRecipeUseCase(recipeRepository, timeoutContext)
	recipeController := _recipeController.NewRecipeController(recipeUseCase)

	reviewRepository := _reviewRepository.NewMysqlReviewRepository(db)
	reviewUseCase := _reviewUsecase.NewReviewUseCase(reviewRepository, recipeRepository ,timeoutContext)
	reviewController := _reviewController.NewReviewController(reviewUseCase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userController,
		CategoryController: *categoryController,
		ProductController:  *productController,
		RecipeController:   *recipeController,
		ReviewController:   *reviewController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
