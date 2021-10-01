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

	_middleware "cookly/app/middleware"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
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

	routesInit := _routes.ControllerList{
		JWTMiddleware:  configJWT.Init(),
		UserController: *userController,
		CategoryController: *categoryController,
		ProductController: *productController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
