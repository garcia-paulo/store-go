package routes

import (
	"net/http"

	"github.com/garcia-paulo/store-go/controllers"
)

func Load() {
	http.HandleFunc("/", controllers.Index)

}
