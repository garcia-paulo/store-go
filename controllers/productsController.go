package controllers

import (
	"net/http"
	"text/template"

	"github.com/garcia-paulo/store-go/servicers"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func GetProducts(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", servicers.GetProducts())
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		servicers.CreateProducts(r)
	}
	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	servicers.DeleteProduct(r)
	http.Redirect(w, r, "/", 301)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	servicers.UpdateProduct(r)
	http.Redirect(w, r, "/", 301)
}
