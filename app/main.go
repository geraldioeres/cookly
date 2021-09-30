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

	routesInit := _routes.ControllerList{
		JWTMiddleware:  configJWT.Init(),
		UserController: *userController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
