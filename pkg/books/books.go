package books

import "github.com/leangeder/calibre/pkg/author"

type Books struct {
	ID          uint            `json:"id"`
	Title       string          `json:"title"`
	Path        string          `json:"path"`
	Format      string          `json:"format"`
	DocDate     string          `json:"docdate"`
	Lang        string          `json:"lang"`
	Isbn        string          `json:"isbn""`
	Rating      uint            `json:"rating"`
	Rank        uint            `json:"rank"`
	Avail       uint            `json:"avail"`
	Authors     []author.Author `json:"authors"`
	Annotations []Annotation    `json:"annotations"`
	Genres      []Genre         `json:"genres"`
	Covers      []Cover         `json:"covers"`
	Series      []Serie         `json:"series"`
}

// var author author.Author
//
// // Get Single Books
// func (e Books) Get(id int) {
// }
//
// // Set Single Books
// func (e Books) Set(book Ebooks) {
// }
//
// // Delete Single Books
// func (e Books) Delete(id int) {
// }
//
// // Get All Books
// func (e Books) List() {
// }
