package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Server(r *mux.Router) error {
	log.Println("server running on  port 9001...")
	log.Fatal(http.ListenAndServe(":9001", r))
	return nil
}
