package animals

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	Database "petclinic/database"
	models "petclinic/models"

	"github.com/gorilla/mux"
)

func GetAnimals(w http.ResponseWriter, r *http.Request) {
	db := Database.Instance

	var animals []models.Animal
	result := db.Find(&animals)

	if result.Error != nil {
			log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(animals)
}

func GetAnimal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := Database.Instance

	var animal models.Animal
	result := db.First(&animal, params["id"])

	if result.Error != nil {
			log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(animal)
}

func CreateAnimal(w http.ResponseWriter, r *http.Request) {
	db := Database.Instance

	var animal models.Animal
	json.NewDecoder(r.Body).Decode(&animal)

	result := db.Create(&animal)

	if result.Error != nil {
			log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(animal)
}

func UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := Database.Instance

	var animal models.Animal
	result := db.First(&animal, params["id"])

	if result.Error != nil {
			log.Fatal(result.Error)
	}

	json.NewDecoder(r.Body).Decode(&animal)

	result = db.Save(&animal)

	if result.Error != nil {
			log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(animal)
}

func DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := Database.Instance
	
	var animal models.Animal
	result := db.Delete(&animal, params["id"])

	if result.Error != nil {
			log.Fatal(result.Error)
	}

	fmt.Fprintf(w, "Animal supprimé")
}
