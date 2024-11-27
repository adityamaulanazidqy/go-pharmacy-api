package routes

import (
	"go-pharmacy-api/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/reset", controllers.ResetProductsHandler).Methods("POST")
	return router
}
