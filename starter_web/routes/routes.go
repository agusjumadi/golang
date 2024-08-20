package routes

import (
	"go-starter-webapp/app/controllers/productcontroller"

	"github.com/gorilla/mux"
)

func Getroutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/products", productcontroller.Index)
	return r
}
