package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book ...
type Book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author ...
type Author struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Books int `json:"books"`
}

// MyBooks ...
var MyBooks[]Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MyBooks)
}
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, book := range MyBooks{
		if(book.ID == params["id"]){
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode("404 Not Found")
}
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = strconv.Itoa(rand.Intn(100000000))
	MyBooks = append(MyBooks, newBook)
	json.NewEncoder(w).Encode(newBook)
}

func main(){
	// Init router
	router := mux.NewRouter()

	for i := 0; i < 10; i++ {
		MyBooks = append(MyBooks, Book{ID: "1", Title: "Book1", Author: &Author{Name: "test", Email: "test@test.com", Books: 2}})
	}

	// routes
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")

	PORT := "5000"

	fmt.Println("Server starting on PORT",PORT)
	// log errors
	log.Fatal(http.ListenAndServe(":" + PORT, router))
}
