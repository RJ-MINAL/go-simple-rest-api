package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Person holds mock data
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address holds address mock data
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func getPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getPeople")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(people)
}

func getPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getPerson")
	w.Header().Set("Content-Type", "application/json")

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
	w.Header().Set("Content-Type", "application/json")

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe

	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func updatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: updatePerson")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			person.ID = params["id"]
			people = append(people, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func deletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deletePerson")
	w.Header().Set("Content-Type", "application/json")

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

	//mock data
	people = append(people, Person{ID: "1", FirstName: "Ryan", LastName: "Ray", Address: &Address{City: "Dublin", State: "California"}})
	people = append(people, Person{ID: "2", FirstName: "Joe", LastName: "DiMagio"})

	// endpoints
	router.HandleFunc("/api/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/api/people/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/api/people/", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/api/people/{id}", updatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/api/people/{id}", deletePersonEndpoint).Methods("DELETE")

	//start server
	log.Fatal(http.ListenAndServe(":3000", router))
}

// Request sample
// {
// 	"firstname":"Daniel",
// 	"lastname":"Potter",
// 	"address":{"city":"Dublin","state":"California"}
// }
