package main

import (
	"cookly/configs"
	"cookly/routes"
)
func main() {
	configs.InitDB()
	e := routes.InitRoutes()
	e.Start(":8000")
}