package v1

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/Joematpal/phone_book_golang_api/v1/version"
	"github.com/gorilla/mux"
	// postgres
	_ "github.com/lib/pq"
)

// route struct
type V1 struct {
	DB     *sql.DB
	Router *mux.Router
}

// Initialize this thing
func (v1 *V1) Initialize(host, port, user, password, dbname string, newRouter *mux.Router) {
	// sslmode=disable this stuff below if wrong
	// certPath := "cert/server.pem"
	// keyPath := "cert/server.key"

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)

	var err error
	v1.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	v1.Router = newRouter.PathPrefix("/api/v1").Subrouter()
}
