package api

import (
	"github.com/gorilla/mux"
)

func InitHandlers() (*mux.Router, error) {

	router := mux.NewRouter()

	router.HandleFunc("/api/createParkingLot", createParkingLot()).Methods("POST")
	router.HandleFunc("/api/park-car", parkCar()).Methods("POST")
	router.HandleFunc("/api/status-all", getStatusOfParkingLot()).Methods("GET")
	router.HandleFunc("/api/remove-car", removeCar()).Methods("DELETE")
	router.HandleFunc("/api/same-color-car", getSameColorCar()).Methods("POST")
	router.HandleFunc("/api/parked-slot", getParkedSlot()).Methods("POST")

	return router, nil
}
