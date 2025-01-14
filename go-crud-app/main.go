package main

import (
	"encoding/json"
	"go-crud-app/models" // Import the models package
	"net/http"
)

var books = []models.Book{}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/books", getBooks)

	http.ListenAndServe(":8080", nil)
}
