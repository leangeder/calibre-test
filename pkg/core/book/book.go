package book

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	// ID     uint   `json:"id"`
	ID     string `json:"id"`
	Title  string `json:"title"`
	Path   string `json:"path"`
	Format string `json:"format"`
	// DocDate     string           `json:"docdate"`
	// Lang        string           `json:"lang"`
	// Isbn        string           `json:"isbn""`
	// Rating      uint             `json:"rating"`
	// Rank        uint             `json:"rank"`
	// Avail       uint             `json:"avail"`
	// Authors     *[]author.Author `json:"authors"`
	// Annotations []Annotation     `json:"annotations"`
	// Genres      []Genre          `json:"genres"`
	// Covers      []Cover          `json:"covers"`
	// Series      []Serie          `json:"series"`
}

// var author author.Author
//
// // Get Single Books
// func (e Book) Get(id int) {
// }
//
// // Set Single Books
// func (e Book) Set(book Ebooks) {
// }
//
// // Delete Single Books
// func (e Book) Delete(id int) {
// }
//
// // Get All Books
// func (e Book) List() {
// }

var books []Book

// Get All Books

func GetBooksPerAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books = append(books, Book{ID: "1"})
	books = append(books, Book{ID: "2"})
	json.NewEncoder(w).Encode(books)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetBooksPerCategorie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetLicences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetBooksPerLicence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetLanguages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetBooksPerLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// // Get a Book
// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
//
// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&book.Book{})
// }
//
// // // Create a Book
// // func CreateBook(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")
// // 	params := mux.Vars(r)
// // }
// //
// // // Update a Book
// // func UpdateBook(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")
// // 	params := mux.Vars(r)
// // }
// //
// // // Delete a Book
// // func DeleteBook(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")
// // 	params := mux.Vars(r)
// // }
// //
// // // Delete All Books
// // func DeleteBooks(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")
// // 	params := mux.Vars(r)
// // }
