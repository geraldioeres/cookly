package main

import (
	"cookly/mvc/configs"
	"cookly/mvc/routes"
)

func main() {
	configs.InitDB()
	e := routes.InitRoutes()
	e.Start(":8000")
}
