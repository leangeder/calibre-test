package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leangeder/calibre/pkg/server/handler"
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
	// // Route for Books
	// r.HandleFunc(apiVersion+"/accounts/{name}", accounts.Get).Methods("GET")
	// r.HandleFunc(apiVersion+"/accounts/{name}/books", books.List).Methods("GET")
	// r.HandleFunc(apiVersion+"/accounts", accounts.List).Methods("GET")

	// r.HandleFunc(apiVersion+"/users", users.Add).Methods("POST")
	// r.HandleFunc(apiVersion+"/users", users.List).Methods("GET")
	// r.HandleFunc(apiVersion+"/users/{id}", users.Add).Methods("DELETE")
	// r.HandleFunc(apiVersion+"/users/{id}", users.Delete).Methods("GET")
	// r.HandleFunc(apiVersion+"/users/{id}", users.Update).Methods("PUT")
	// r.HandleFunc(apiVersion+"/users/me", users.Get).Methods("GET")
	// r.HandleFunc(apiVersion+"/users/me", users.Get).Methods("PUT")
	// r.HandleFunc(apiVersion+"/users/me/avatar/pick", account.Get).Methods("POST")
	// r.HandleFunc(apiVersion+"/users/register", account.Get).Methods("POST")

	r.HandleFunc(apiVersion+"/books", account.Get).Methods("POST")
	r.HandleFunc(apiVersion+"/books/{id}", account.Update).Methods("PUT")
	r.HandleFunc(apiVersion+"/books/{id}", account.Delete).Methods("POST")
	r.HandleFunc(apiVersion+"/books/{id}", account.Get).Methods("POST")
	r.HandleFunc(apiVersion+"/books/categories", account.Get).Methods("POST")
	r.HandleFunc(apiVersion+"/books/licences", account.Get).Methods("POST")
	r.HandleFunc(apiVersion+"/books/languages", account.Get).Methods("POST")

	r.Use(loggingMiddleware)
	return r
}
