package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks/db"
	"github.com/RodrigoGonzalez78/tasks/models"
	"github.com/gorilla/mux"
)

func GetTaksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	//Verificamos si existe el id en la tabla
	//Golang devuelve 0 por defecto, es decir todos los campos con ZERO value
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)
	taskCreated := db.DB.Create(&task)
	err := taskCreated.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Cambia el valor de deleted_at, no elmina el elemento en si
	//db.DB.Delete(&task) //igual la libreria se encarga de no mostrar mas el elemento

	//Remueve totalamente de la tabla
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
}

func GetTasksByUserHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	params := mux.Vars(r)
	userId := params["userId"]

	// Consulta la base de datos para obtener todas las tareas del usuario con el ID especificado
	db.DB.Where("user_id = ?", userId).Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)
}
