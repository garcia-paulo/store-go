package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DBConnect() *sql.DB {
	connection := "postgres://postgres:postgres@localhost:5432/store-go?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
