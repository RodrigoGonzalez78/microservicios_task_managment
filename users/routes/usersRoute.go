package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/users/db"
	"github.com/RodrigoGonzalez78/users/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	//Verificamos si existe el id en la tabla
	//Golang devuelve 0 por defecto, es decir todos los campos con ZERO value
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	userCreated := db.DB.Create(&user)
	err := userCreated.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Cambia el valor de deleted_at, no elmina el elemento en si
	//db.DB.Delete(&user) igual la libreria se encarga de no mostrar mas el elemento

	//Remueve totalamente de la tabla
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
