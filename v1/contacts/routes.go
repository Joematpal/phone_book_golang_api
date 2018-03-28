package contacts

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// Routes for contacts
func Routes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/contact/{id}", GetContact(db)).Methods("GET")
	router.HandleFunc("/contact", CreateContact(db)).Methods("PUT")
	router.HandleFunc("/contact/{id}", UpdateContact(db)).Methods("POST")
	router.HandleFunc("/contact/{id}", DeleteContact(db)).Methods("DELETE")
	router.HandleFunc("/contacts", GetContacts(db)).Methods("GET")
	router.HandleFunc("/csv/contacts", GetCSVContacts(db)).Methods("GET")
	router.HandleFunc("/csv/contacts", UpdateCSVContacts(db)).Methods("POST")
}
