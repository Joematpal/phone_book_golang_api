package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Joematpal/phone_book_golang_api/v1"
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	host := os.Getenv("TEST_DB_HOST")
	port := os.Getenv("TEST_DB_PORT")
	username := os.Getenv("TEST_DB_USERNAME")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbname := os.Getenv("TEST_DB_NAME")
	router := mux.NewRouter()

	v = v1.V1{}

	v.Initialize(
		host,
		port,
		username,
		password,
		dbname,
		router,
	)

	// ensureTableExists()

	code := m.Run()

	// clearTable()

	os.Exit(code)
}

func TestTableExists(t *testing.T) {
	ensureTableExists()
}

// func TestEmptyTable(t *testing.T) {
// 	clearTable()
//
// 	req, _ := http.NewRequest("GET", "/person", nil)
// 	response := executeRequest(req)
//
// 	checkResponseCode(t, http.StatusOK, response.Code)
//
// 	if body := response.Body.String(); body != "[]" {
// 		t.Errorf("Expected an empty array. Got %s", body)
// 	}
// }

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func ensureTableExists() {
	if _, err := v.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	v.DB.Exec("DELETE FROM contact")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	// v1.Router.ServeHTTP(rr, req)

	return rr
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS contacts
(
	id VARCHAR(64),
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(50),
	phone VARCHAR(50),
	CONSTRAINT contact_pkey PRIMARY KEY (id),
	CONSTRAINT email_id UNIQUE("email")
)`
