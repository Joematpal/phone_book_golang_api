package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// postgres
	_ "github.com/lib/pq"
)

//App is the constructor for the application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize starts the database connection and provides the application with a route.
func (a *App) Initialize(user, password, dbname string) {

	// certPath := "cert/server.pem"
	// keyPath := "cert/server.key"

	// connectionString :=
	// 	fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable sslcert=%s sslkey=%s", user, password, dbname, certPath, keyPath)

	connectionString :=
		fmt.Sprintf("postgres://%s:%s@localhost:%s?sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

//Run tells the http listener which port to server the application.
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
