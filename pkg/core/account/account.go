package account

import (
	"encoding/json"
	"net/http"
)

type Account struct {
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

var accounts []Account

// Get All Books

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	accounts = append(accounts, Account{ID: "1"})
	accounts = append(accounts, Account{ID: "2"})
	json.NewEncoder(w).Encode(accounts)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	accounts = append(accounts, Account{ID: "1"})
	accounts = append(accounts, Account{ID: "2"})
	json.NewEncoder(w).Encode(accounts)
}
