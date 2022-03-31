package controllers

import (
	"net/http"
	"text/template"

	"github.com/garcia-paulo/store-go/servicers"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", servicers.GetProducts())
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		servicers.CreateProducts(r)
	}
}
