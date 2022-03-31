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
	db := db.DBConnect()
	insert, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id int) {
	db := db.DBConnect()
	delete, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.DBConnect()
	update, err := db.Prepare("update products set name = $1, description = $2, price = $3, quantity = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(name, description, price, quantity, id)
	defer db.Close()
}
