package routes

import (
	"go-starter-webapp/app/controllers/productcontroller"
	
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	fmt.Println("auth middleware")
        /*username, password, ok := r.BasicAuth()
        if !ok {
            w.Write([]byte(`something went wrong`))
            return
        }

        isValid := (username == "u") && (password == "p")
        if !isValid {
            w.Write([]byte(`wrong username/password`))
            return
        }*/

        next.ServeHTTP(w, r)
    })
}
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        fmt.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}
func Getroutes() *mux.Router {
	r := mux.NewRouter()
	r.Path("/api/v1/auth/login").Handler(http.HandlerFunc(productcontroller.Index))
	r.HandleFunc("/products", productcontroller.Index)
	
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(authMiddleware)
	//api.Path("/users").Handler(http.HandlerFunc(productcontroller.Index)).Methods("GET")
	api.HandleFunc("/users", productcontroller.Index).Methods("GET")
	api.Path("/users").Handler(http.HandlerFunc(productcontroller.Create)).Methods("POST")
	
	r.Use(loggingMiddleware)
	return r
}
