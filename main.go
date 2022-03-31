package main

import (
	"net/http"

	"github.com/garcia-paulo/store-go/routes"
)

func main() {
	routes.Load()
	http.ListenAndServe(":8080", nil)
}
