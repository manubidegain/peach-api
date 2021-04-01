package io

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	fmt.Println("tamoactivo")
	router := mux.NewRouter()

	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products", createProduct).Methods("POST")
	/*router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")*/

	return router
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yatelasabes")
	GetProducts(w, r)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	CreateProduct(w, r)
}
