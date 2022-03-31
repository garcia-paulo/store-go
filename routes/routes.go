package routes

import (
	"net/http"

	"github.com/garcia-paulo/store-go/controllers"
)

func Load() {
	http.HandleFunc("/", controllers.GetProducts)
	http.HandleFunc("/insert", controllers.CreateProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
}
