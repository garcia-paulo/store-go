package repositories

import (
	"github.com/garcia-paulo/store-go/db"
	"github.com/garcia-paulo/store-go/models"
)

func GetProducts() []models.Product {
	db := db.DBConnect()
	query, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	products := []models.Product{}

	for query.Next() {
		product := models.Product{}

		err = query.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	defer db.Close()
	return products
}

func CreateProducts(name string, description string, price float64, quantity int) {

}
