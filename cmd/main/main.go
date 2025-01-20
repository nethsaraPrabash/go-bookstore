package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nethsaraPrabash/go-bookstore/pkg/routes"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running!")
}

func main() {
	// Create a new mux router
	r := mux.NewRouter()

	// Register the '/hello' path
	r.HandleFunc("/hello", handler)

	// Register other routes
	routes.RegisterBookstoreRoutes(r)

	// Start the server on port 5050
	log.Println("Server started on: http://localhost:5050")
	log.Fatal(http.ListenAndServe(":5050", r)) // Ensure the router handles requests
}
