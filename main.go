package main

import (
	"net/http"
	"todo-list/routes"
)

func main() {
	routes.CarriesRoutes()
	http.ListenAndServe(":8000", nil)
}
