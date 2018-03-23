package contacts

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Joematpal/test-golang-api/src/v1/utils/respond"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// GetContacts from ./model.go
func GetContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}

		p := Contact{}

		contacts, err := p.getContacts(db, start, count)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			return
		}

		respond.With(w, r, http.StatusOK, nil, contacts)
	}
}

// GetContact from ./model.go
func GetContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.FromString(vars["id"])
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, nil, "Invalid contact ID")
			return
		}

		p := Contact{ID: id}
		if err := p.getContact(db); err != nil {
			switch err {
			case sql.ErrNoRows:
				respond.With(w, r, http.StatusNotFound, nil, "Contact not found")
			default:
				respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			}
			return
		}

		respond.With(w, r, http.StatusOK, p, nil)
	}
}

// CreateContact from ./model.go
func CreateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// UpdateContact from ./model.go
func UpdateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// DeleteContact from ./model.go
func DeleteContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
