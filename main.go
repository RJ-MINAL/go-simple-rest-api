package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

func getPeopleEndpoint(w http.ResponseWriter, r *http.Request) {

}

func getPersonEndpoint(w http.ResponseWriter, r *http.Request) {

}

func createPersonEndpoint(w http.ResponseWriter, r *http.Request) {

}

func deletePersonEndpoint(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/api/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/api/people/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/api/people/{id}", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/api/people/{id}", deletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
