package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Server(r *mux.Router) error {
	log.Println("server running on  port 9003...")
	log.Fatal(http.ListenAndServe(":9003", r))
	return nil
}
