package api

import (
	"encoding/json"
	"my/v1/api/wrapper"
	"my/v1/app"
	"my/v1/errors"
	"my/v1/model/validation"
	"net/http"
)

func createParkingLot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestCTX wrapper.RequestContext
		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		var form validation.ValidateCreateParkingLot

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			requestCTX.SetErrs(errs)
			wrapper.Response(requestCTX, w, r)
			return
		}

		res, err := app.NewApp().CreateParkingLot(form.NoOfSlot)
		if err != nil {
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		requestCTX.SetAppResponse(res, 200)
		wrapper.Response(requestCTX, w, r)
	}
}

func parkCar() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestCTX wrapper.RequestContext

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		var form validation.ValidateParkCar

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			requestCTX.SetErrs(errs)
			wrapper.Response(requestCTX, w, r)
			return
		}

		res, err := app.NewApp().ParkCar(form.Code, form.RegistrationNo, form.Color)
		if err != nil {
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return

		}

		requestCTX.SetAppResponse(res, 200)
		wrapper.Response(requestCTX, w, r)
	}
}

func getStatusOfParkingLot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestCTX wrapper.RequestContext
		res, err := app.NewApp().GetStatusOfParkingLot()
		if err != nil {
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		requestCTX.SetAppResponse(res, 200)
		wrapper.Response(requestCTX, w, r)

	}
}

func removeCar() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestCTX wrapper.RequestContext

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		var form validation.ValidateLeaveParking

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			requestCTX.SetErrs(errs)
			wrapper.Response(requestCTX, w, r)
			return
		}

		res, err := app.NewApp().RemoveCar(form.RegistrationNo)
		if err != nil {
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		requestCTX.SetAppResponse(res, 200)
		wrapper.Response(requestCTX, w, r)
	}
}

func getSameColorCar() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestCTX wrapper.RequestContext

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		var form validation.ValidateReturnSameColorVehicles

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			requestCTX.SetErrs(errs)
			wrapper.Response(requestCTX, w, r)
			return
		}

		res, err := app.NewApp().GetSameColorCar(form.Color)
		if err != nil {
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		requestCTX.SetAppResponse(res, 200)
		wrapper.Response(requestCTX, w, r)
	}
}

func getParkedSlot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestCTX wrapper.RequestContext

		if l := r.ContentLength; l == 0 {
			var err error
			err = errors.BadRequest.Wrapf(err, "Empty request data")
			err = errors.AddErrorContext(err, "schema", "Recieved empty json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		var form validation.ValidateParkedSlot

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			err = errors.BadRequest.Wrapf(err, "unable to read json schema")
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		if errs := validation.NewValidation().Validate(&form); errs != nil {
			requestCTX.SetErrs(errs)
			wrapper.Response(requestCTX, w, r)
			return
		}

		res, err := app.NewApp().GetParkedSlot(form.RegistrationNo)
		if err != nil {
			requestCTX.SetErr(err)
			wrapper.Response(requestCTX, w, r)
			return
		}

		requestCTX.SetAppResponse(res, 200)
		wrapper.Response(requestCTX, w, r)
	}
}
