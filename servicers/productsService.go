package servicers

import (
	"net/http"
	"strconv"

	"github.com/garcia-paulo/store-go/models"
	"github.com/garcia-paulo/store-go/repositories"
)

func GetProducts() []models.Product {
	return repositories.GetProducts()
}

func CreateProducts(r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		panic(err.Error())
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		panic(err.Error())
	}

	repositories.CreateProducts(name, description, price, quantity)
}
