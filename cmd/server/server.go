package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Server(r *mux.Router) error {

	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Autherization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Println("server running on  port 9000...")
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(headers, methods, origins)(r)))
	return nil
}
