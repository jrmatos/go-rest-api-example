package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type Response struct {
	Persons []Person `json:"persons"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("/health-check")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "API is up and running")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	log.Println("/persons")
	var response Response

	persons := prepareResponse()

	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func prepareResponse() []Person {
	var persons []Person
	var person Person

	person.Id = 1
	person.FirstName = "Isaac Newton"
	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Albert"
	person.LastName = "Einstein"
	persons = append(persons , person)

	return persons
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	http.Handle("/", router)

	fmt.Println("Escutando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

