package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting API server")
	router := mux.NewRouter() // Create New Router

	log.Println("Setting up routes")
	router.HandleFunc("/health-check", HealthCheckHandler).
		Methods("GET")

	router.HandleFunc("/students", StudentsHandler).
		Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
