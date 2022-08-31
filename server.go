package main

import (
	"net/http"

	"timbul.io/routes"
)

func main() {
	routes := routes.CreateRoute()
	http.ListenAndServe(":3000", routes)
}
