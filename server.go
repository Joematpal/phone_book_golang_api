package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Joematpal/phone_book_golang_api/v1"
	"github.com/Joematpal/phone_book_golang_api/v1/contacts"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var v v1.V1

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading Enironment Variables")
	}

	router := mux.NewRouter()

	v.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		router,
	)

	contacts.Routes(v.Router, v.DB)

	run(":8080", router)
}

func run(addr string, route *mux.Router) {
	log.Fatal(http.ListenAndServe(addr, route))
}
