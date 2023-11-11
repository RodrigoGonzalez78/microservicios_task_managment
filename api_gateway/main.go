package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	// Crear el enrutador para el API Gateway
	router := mux.NewRouter()

	// Ruta para el microservicio de usuarios
	usersProxy := NewReverseProxy("http://172.17.0.5:5000")
	router.PathPrefix("/users").Handler(usersProxy)

	// Ruta para el microservicio de tareas
	tasksProxy := NewReverseProxy("http://172.17.0.4:4000")
	router.PathPrefix("/tasks").Handler(tasksProxy)

	// Configurar el servidor HTTP
	server := &http.Server{
		Addr:    ":8080", // Puerto en el que escuchará el API Gateway
		Handler: router,
	}

	fmt.Println("API Gateway en ejecución en el puerto 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// NewReverseProxy crea un proxy inverso para redirigir solicitudes a un servicio en un puerto determinado
func NewReverseProxy(target string) *httputil.ReverseProxy {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	return proxy
}
