package main

import (
	"net/http"

	"github.com/RodrigoGonzalez78/users/db"
	"github.com/RodrigoGonzalez78/users/models"
	"github.com/RodrigoGonzalez78/users/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	//Se crean las tablas users
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	//Ruta home
	router.HandleFunc("/", routes.HomeHandler)

	//Ruta crud de usuarios
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Iniciamos el servidor
	http.ListenAndServe(":5000", router)
}
