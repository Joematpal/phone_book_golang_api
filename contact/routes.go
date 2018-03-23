package contacts

import (
	"github.com/Joematpal/test-golang-api/src/v1/version"
)

// Routes for contacts
func Routes(v version.V1) {
	v.Subrouter.HandleFunc("/contact/{id}", GetContact(v.DB)).Methods("GET")
	v.Subrouter.HandleFunc("/contact", CreateContact(v.DB)).Methods("POST")
	v.Subrouter.HandleFunc("/contact/{id}", UpdateContact(v.DB)).Methods("PUT")
	v.Subrouter.HandleFunc("/contact/{id}", DeleteContact(v.DB)).Methods("DELETE")
	v.Subrouter.HandleFunc("/contacts", GetContacts(v.DB)).Methods("GET")
}
