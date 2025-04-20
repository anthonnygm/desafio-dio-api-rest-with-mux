package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anthonnygm/desafio-dio-api-rest-with-mux/models"
	"github.com/gorilla/mux"
)

// GetPeople mostra todos os contatos da variável people
func GetPeople(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(app.People)
	}
}

// GetPerson Mostra apenas um contato
func GetPerson(app *models.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		for _, item := range app.People {
			if item.ID == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		json.NewEncoder(w).Encode(&models.Person{})
	}
}

// CreatePerson cria um novo contato e retorna a lista de people cadastrados
func CreatePerson(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var person models.Person
		_ = json.NewDecoder(r.Body).Decode(&person)
		person.ID = params["id"]
		app.People = append(app.People, person)
		json.NewEncoder(w).Encode(app.People)
	}
}

// DeletePerson deleta um contato
func DeletePerson(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		for i, item := range app.People {
			if item.ID == params["id"] {
				app.People = append(app.People[:i], app.People[i+1:]...)
				break
			}
		}

		json.NewEncoder(w).Encode(app.People)
	}
}
