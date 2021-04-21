package api

import (
	"encoding/json"
	"my/v1/app"
	"my/v1/errors"
	"my/v1/model/validation"
	"net/http"
)

func createParkingLot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		var form validation.ValidateCreateParkingLot

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			json.NewEncoder(w).Encode(errs)
			return
		}

		res, err := app.CreateParkingLot(form.NoOfSlot)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func parkCar() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		var form validation.ValidateParkCar

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			json.NewEncoder(w).Encode(errs)
			return
		}

		res, err := app.ParkCar(form.RegistrationNo, form.Color)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func getStatusOfParkingLot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		res, err := app.GetStatusOfParkingLot()
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func removeCar() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		var form validation.ValidateLeaveParking

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			json.NewEncoder(w).Encode(errs)
			return
		}

		res, err := app.RemoveCar(form.RegistrationNo)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func getSameColorCar() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		var form validation.ValidateReturnSameColorVehicles

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			json.NewEncoder(w).Encode(errs)
			return
		}

		res, err := app.GetSameColorCar(form.Color)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func getParkedSlot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		var form validation.ValidateParkedSlot

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			json.NewEncoder(w).Encode(err)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			json.NewEncoder(w).Encode(errs)
			return
		}

		res, err := app.GetParkedSlot(form.RegistrationNo)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
