package contacts

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Joematpal/phone_book_golang_api/utils"
	"github.com/Joematpal/phone_book_golang_api/utils/respond"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	validator "gopkg.in/validator.v2"
)

var (
	createValidator = validator.NewValidator()
)

func init() {
	createValidator.SetTag("creating")
}

// GetContacts from ./model.go
func GetContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryCount := r.URL.Query().Get("count")
		queryStart := r.URL.Query().Get("start")
		count, _ := strconv.Atoi(queryCount)
		start, _ := strconv.Atoi(queryStart)

		if count > (start+50) || count < 1 {
			count = start + 50
		}

		if start < 0 {
			start = 0
		}

		c := &Contact{}

		contacts, err := c.getContacts(db, start, count)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			return
		}

		respond.With(w, r, http.StatusOK, contacts, nil)
	}
}

// GetContact from ./model.go
func GetContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, badString := uuid.FromString(vars["id"])
		if badString != nil {
			respond.With(w, r, http.StatusBadRequest, nil, "Invalid contact ID")
			return
		}

		c := &Contact{ID: id}
		if err := c.getContact(db); err != nil {
			switch err {
			case sql.ErrNoRows:
				respond.With(w, r, http.StatusNotFound, nil, "Contact not found")
			default:
				respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			}
			return
		}

		respond.With(w, r, http.StatusOK, c, nil)
	}
}

// CreateContact from ./model.go
func CreateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &Contact{}

		if err := utils.Decode(r, &c); err != nil {
			respond.With(w, r, http.StatusBadRequest, nil, "Err at decode")
			return
		}
		if errs := createValidator.Validate(c); errs != nil {
			respond.With(w, r, http.StatusBadRequest, nil, errs)
			return
		}
		c.ID = uuid.NewV4()

		if err := c.createContact(db); err != nil {
			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			return
		}

		respond.With(w, r, http.StatusOK, c, nil)
		return
	}
}

// UpdateContact from ./model.go
func UpdateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, badString := uuid.FromString(vars["id"])
		if badString != nil {
			respond.With(w, r, http.StatusBadRequest, nil, "Invalid contact ID")
			return
		}
		c := &Contact{ID: id}

		if err := utils.Decode(r, &c); err != nil {
			respond.With(w, r, http.StatusBadRequest, nil, "Err at decode")
			return
		}
		if err := createValidator.Validate(c); err != nil {
			respond.With(w, r, http.StatusBadRequest, nil, err)
			return
		}
		if err := c.updateContact(db); err != nil {
			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			return
		}
		respond.With(w, r, http.StatusOK, c, nil)
		return
	}
}

// DeleteContact from ./model.go
func DeleteContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, badString := uuid.FromString(vars["id"])
		if badString != nil {
			respond.With(w, r, http.StatusBadRequest, nil, "Invalid contact ID")
			return
		}
		c := &Contact{ID: id}

		if err := utils.Decode(r, &c); err != nil {
			respond.With(w, r, http.StatusBadRequest, nil, "Err at decode")
			return
		}
		if err := c.deleteContact(db); err != nil {
			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			return
		}
		respond.With(w, r, http.StatusNoContent, nil, nil)
		return
	}
}

// UpdateCSVContacts will upsert a csv to the database
func UpdateCSVContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("contacts")
		newFile, _ := ioutil.ReadAll(file)
		tempFile, _ := ioutil.TempFile("", "contacts")
		e := ioutil.WriteFile(tempFile.Name(), newFile, 0644)
		if e != nil {
			respond.With(w, r, http.StatusBadRequest, nil, e)
			return
		}
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, nil, err)
			return
		}
		// "INSERT INTO contacts() VAlUES () on DUPLICATE KEY UPDATE name, age"
		query := fmt.Sprintf("COPY contacts_imports(first_name, last_name, id, email, phone) FROM '%s' DELIMITER ',' CSV;", tempFile.Name())

		if _, er := db.Exec(`CREATE TEMP TABLE contacts_imports
				(
				id VARCHAR(64),
				first_name VARCHAR(50),
				last_name VARCHAR(50),
				email VARCHAR(50),
				phone VARCHAR(50),
				CONSTRAINT contact_pkey PRIMARY KEY (id),
				CONSTRAINT email_id UNIQUE("email")
				)`); er != nil {
			respond.With(w, r, http.StatusBadRequest, nil, er)
			return
		}
		if _, er := db.Exec(query); er != nil {
			respond.With(w, r, http.StatusBadRequest, nil, er)
			return
		}
		if _, er := db.Exec(`
			insert into posts(id, first_name, last_name, email, phone)
			select id, first_name, last_name, email, phone
			from contact_imports
			on DUPLICATE KEY UPDATE
			update first_name, last_name, email, phone, id
			`); er != nil {
			respond.With(w, r, http.StatusBadRequest, nil, er)
			return
		}
	}
}

// GetCSVContacts will return data as a csv
func GetCSVContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpfile, err := ioutil.TempFile("", "contacts.csv")
		if err != nil {
			log.Fatal(err)
		}
		query := fmt.Sprintf("COPY contacts TO '%s' DELIMITER ',' CSV HEADER;", tmpfile.Name())

		if _, er := db.Exec(query); er != nil {
			respond.With(w, r, http.StatusBadRequest, nil, er)
			return
		}

		Openfile, err := os.Open(tmpfile.Name())
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, nil, err)
			return
		}
		defer Openfile.Close()
		defer os.Remove(tmpfile.Name())
		FileHeader := make([]byte, 512)
		Openfile.Read(FileHeader)
		//Get content type of file
		FileContentType := http.DetectContentType(FileHeader)
		FileStat, _ := Openfile.Stat()
		FileSize := strconv.FormatInt(FileStat.Size(), 10)

		Openfile.Seek(0, 0)

		w.Header().Set("Content-Disposition", "attachment; filename=contacts.csv")
		w.Header().Set("Content-Type", FileContentType)
		w.Header().Set("Content-Length", FileSize)
		io.Copy(w, Openfile)
	}
}
