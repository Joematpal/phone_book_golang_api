package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	username := os.Getenv("TEST_DB_USERNAME")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbname := os.Getenv("TEST_DB_NAME")

	a.Initialize(
		username,
		password,
		dbname,
	)

	ensureTableExists()

	code := m.Run()

	clearTable()

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
	var c contacts.Contact
	v, _ := a.DB.Exec("SELECT * from PEOPLE") //err != nil {
	//	log.Fatal(err)
	//}
	log.Fatal(v.scan(&c.FirstName, &c.LastName, &c.Email, &c.Phone))
}

func clearTable() {
	a.DB.Exec("DELETE FROM people")
	a.DB.Exec("ALTER SEQUENCE people_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS people
(
	id VARCHAR(64),
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(50),
	phone VARCHAR(50),
	CONSTRAINT people_pkey PRIMARY KEY (id)
)`
