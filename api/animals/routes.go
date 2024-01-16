package animals

import (
	Authentication "petclinic/authentication"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	s := r.PathPrefix("/animals").Subrouter()

	s.Use(Authentication.TokenVerifyMiddleWare)
	s.HandleFunc("", GetAnimals).Methods("GET")
	s.HandleFunc("/{id}", GetAnimal).Methods("GET")
	s.HandleFunc("", CreateAnimal).Methods("POST")
	s.HandleFunc("/{id}", UpdateAnimal).Methods("PUT")
	s.HandleFunc("/{id}", DeleteAnimal).Methods("DELETE")
}