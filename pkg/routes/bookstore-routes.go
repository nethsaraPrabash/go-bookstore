package routes

import (
	"github.com/gorilla/mux"
	"github.com/nethsaraPrabash/go-bookstore/pkg/controllers"

)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
}