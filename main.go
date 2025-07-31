package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Language string  `json:"language"`
	ISBN     string  `json:"isbn"`
	Pages    string  `json:"pages"`
	Author   *Author `json:"author"`
}
type Author struct {
	Publisher string `json:"publisher"`
	Date      string `json:"date"`
}

var books []Book

func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)

}
func getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(books)
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.IntN(10000000))
	books = append(books, book)

	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break

		}

	}

	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if params["id"] == item.ID {
			books = append(books[:index], books[index+1:]...)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	books = append(books, Book{
		ID:       "1",
		Title:    "The go programming",
		Language: "English",
		ISBN:     "35268465",
		Pages:    "400",
		Author: &Author{
			Publisher: "addishon",
			Date:      "2015",
		},
	})
	books = append(books, Book{
		ID:       "2",
		Title:    "Mastering Go",
		Language: "English",
		ISBN:     "9781492077213",
		Pages:    "520",
		Author: &Author{
			Publisher: "O'Reilly Media",
			Date:      "2021",
		},
	})

	r.HandleFunc("/books", getbooks).Methods("GET")
	r.HandleFunc("/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/books/id", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Print("server starting at port address:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
