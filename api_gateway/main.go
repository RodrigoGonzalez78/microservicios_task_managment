package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"handler", "method", "path"},
	)
)

func main() {
	// Crear el enrutador para el API Gateway
	router := mux.NewRouter()

	// Ruta para el microservicio de usuarios
	usersProxy := NewReverseProxy("http://172.17.0.5:5000", "users", "/users")
	router.PathPrefix("/users").Handler(usersProxy)

	// Ruta para el microservicio de tareas
	tasksProxy := NewReverseProxy("http://172.17.0.4:4000", "tasks", "/tasks")
	router.PathPrefix("/tasks").Handler(tasksProxy)

	// Registra el contador Prometheus
	prometheus.MustRegister(requestCount)

	// Configurar el servidor HTTP para métricas Prometheus
	router.Handle("/metrics", promhttp.Handler())

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
func NewReverseProxy(target string, handlerName string, path string) *httputil.ReverseProxy {
	url, err := url.Parse(target)

	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Configura el contador Prometheus para esta ruta
	proxy.Director = func(req *http.Request) {
		req.URL.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))

		// Incrementa el contador Prometheus con la información de la ruta
		requestCount.WithLabelValues(handlerName, req.Method, path).Inc()
	}

	return proxy
}
