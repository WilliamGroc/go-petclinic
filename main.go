package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	Animals "petclinic/api/animals"
	Users "petclinic/api/users"
	Database "petclinic/database"
	Middleware "petclinic/middlewares"
	models "petclinic/models"
)

func Migrate(){	
	Database.Instance.AutoMigrate(
		&models.Animal{}, 
		&models.User{})

	log.Println("Database Migration Completed...")
}

func main() {
		r := mux.NewRouter()

		r.Use(Middleware.LoggingMiddleware)
		r.Use(Middleware.ContentTypeApplicationJsonMiddleware)

		Database.Connect()
		Migrate()

		Animals.RegisterRoutes(r)
		Users.RegisterRoutes(r)

		log.Println("Le serveur écoute sur le port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}