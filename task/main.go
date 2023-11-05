package main

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks/db"
	"github.com/RodrigoGonzalez78/tasks/models"
	"github.com/RodrigoGonzalez78/tasks/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	//Se crean las tablas task
	db.DB.AutoMigrate(models.Task{})

	router := mux.NewRouter()

	//Ruta home
	router.HandleFunc("/", routes.HomeHandler)

	//Rutas del crud de tareas
	router.HandleFunc("/tasks", routes.GetTaksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/tasks/user/{userId}", routes.GetTasksByUserHandler).Methods("GET")

	//Iniciamos el servidor
	http.ListenAndServe(":4000", router)
}
