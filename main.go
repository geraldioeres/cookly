package main

import (
	"cookly/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/login", controllers.UserLogin())
}
