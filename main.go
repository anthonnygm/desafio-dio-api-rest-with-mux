package main

import (
	"log"
	"net/http"

	"github.com/anthonnygm/desafio-dio-api-rest-with-mux/handlers"
	"github.com/anthonnygm/desafio-dio-api-rest-with-mux/models"
	"github.com/gorilla/mux"
)

func main() {
	app := &models.App{
		People: []models.Person{
			{ID: "1", Firstname: "Peter", Lastname: "Quill", Address: &models.Address{City: "St. Charles", State: "Missouri"}},
			{ID: "2", Firstname: "Matthew", Lastname: "Murdok", Address: &models.Address{City: "New York City", State: "NY"}},
		},
	}
	router := mux.NewRouter()

	router.HandleFunc("/contato", handlers.GetPeople(app)).Methods("GET")
	router.HandleFunc("/contato/{id}", handlers.GetPerson(app)).Methods("GET")
	router.HandleFunc("/contato/{id}", handlers.CreatePerson(app)).Methods("POST")
	router.HandleFunc("/contato/{id}", handlers.DeletePerson(app)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}
