package main

import (
	"encoding/json"
	"fmt"
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

var people []Person

func getPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getPeople")
	json.NewEncoder(w).Encode(people)
}

func getPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getPerson")

	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createPerson")

	params := mux.Vars(r)

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]

	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func deletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deletePerson")

	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Ryan", LastName: "Ray", Address: &Address{City: "Dublin", State: "California"}})
	people = append(people, Person{ID: "2", FirstName: "Joe", LastName: "DiMagio"})

	// endpoints
	router.HandleFunc("/api/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/api/people/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/api/people/{id}", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/api/people/{id}", deletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
