package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/talles-morais/go-rest-api/database"
	"github.com/talles-morais/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var person []models.Personality
	database.DB.Find(&person)

	json.NewEncoder(w).Encode(person)
}

func OnePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Personality
	database.DB.First(&person, id)
	json.NewEncoder(w).Encode(person)

}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPerson models.Personality
	json.NewDecoder(r.Body).Decode(&newPerson)
	database.DB.Create(&newPerson)

	json.NewEncoder(w).Encode(newPerson)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Personality

	database.DB.Delete(&person, id)
	json.NewEncoder(w).Encode(person)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Personality

	database.DB.First(&person, id)
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Save(&person)
	json.NewEncoder(w).Encode(person)
}
