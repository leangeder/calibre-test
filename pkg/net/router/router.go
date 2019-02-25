package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leangeder/calibre-test/pkg/core/account"
	"github.com/leangeder/calibre-test/pkg/core/book"
	"github.com/leangeder/calibre-test/pkg/net/handler"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./web/static/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/", handler.Test).Methods("GET")

	apiVersion := "/api/v1"
	// Route for Books
	r.HandleFunc(apiVersion+"/accounts/{id}", account.GetAccount).Methods("GET")
	r.HandleFunc(apiVersion+"/accounts/{id}/books", book.GetBooksPerAccount).Methods("GET")
	r.HandleFunc(apiVersion+"/accounts", account.GetAccounts).Methods("GET")

	// r.HandleFunc(apiVersion+"/users", users.Add).Methods("POST")
	// r.HandleFunc(apiVersion+"/users", users.List).Methods("GET")
	// r.HandleFunc(apiVersion+"/users/{id}", users.Add).Methods("DELETE")
	// r.HandleFunc(apiVersion+"/users/{id}", users.Delete).Methods("GET")
	// r.HandleFunc(apiVersion+"/users/{id}", users.Update).Methods("PUT")
	// r.HandleFunc(apiVersion+"/users/me", users.Get).Methods("GET")
	// r.HandleFunc(apiVersion+"/users/me", users.Get).Methods("PUT")
	// r.HandleFunc(apiVersion+"/users/me/avatar/pick", account.Get).Methods("POST")
	// r.HandleFunc(apiVersion+"/users/register", account.Get).Methods("POST")

	r.HandleFunc(apiVersion+"/books", book.GetBooks).Methods("GET")
	r.HandleFunc(apiVersion+"/books/{id}", book.GetBook).Methods("GET")
	r.HandleFunc(apiVersion+"/books/categories", book.GetCategories).Methods("GET")
	r.HandleFunc(apiVersion+"/books/categories/{id}", book.GetBooksPerCategorie).Methods("GET")
	r.HandleFunc(apiVersion+"/books/licences", book.GetLicences).Methods("GET")
	r.HandleFunc(apiVersion+"/books/licences/{id}", book.GetBooksPerLicence).Methods("GET")
	r.HandleFunc(apiVersion+"/books/languages", book.GetLanguages).Methods("GET")
	r.HandleFunc(apiVersion+"/books/languages/{id}", book.GetBooksPerLanguage).Methods("GET")

	r.Use(loggingMiddleware)
	return r
}
