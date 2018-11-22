package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux" // import the package with:  go get -u github.com/gorilla/mux
)

// sample: {"ID":"1","Firstname":"Timmy","lastname":"Kerouac","address":{"city":"City X","state":"State X", "phone":"35387388888"}}
type Person struct {
	ID        string   `json:"id,omitempty`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Phone string `json:"phone,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	// the function mux.Vars(r) takes the http.Request as parameter and returns a map of the segments.
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

/*
  Tests the API with the following calls:
	- http://localhost:8000/people
	  returns:  [{"ID":"1","Firstname":"Timmy","lastname":"Kerouac","address":{"city":"City X","state":"State X","phone":"+35387388888"}},
	             {"ID":"2","Firstname":"Albert","lastname":"Einstein","address":{"city":"City Z","state":"State Y","phone":"+35385993751"}}]
	- http://localhost:8000/people/1
	  returns:  {"ID":"1","Firstname":"Timmy","lastname":"Kerouac","address":{"city":"City X","state":"State X","phone":"+35387388888"}}
	- test with postman
	    /people/{id} Methods POST
        /people/{id} Methods DELETE
*/

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Timmy", Lastname: "Kerouac", Address: &Address{City: "City X", State: "State X", Phone: "+35387388888"}})
	people = append(people, Person{ID: "2", Firstname: "Albert", Lastname: "Einstein", Address: &Address{City: "City Z", State: "State Y", Phone: "+35385993751"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	// router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("PUT") for the update
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
