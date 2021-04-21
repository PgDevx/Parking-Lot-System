package model

import (
	"encoding/json"

	"my/v1/errors"
)

type ResponseType string

const (
	HTMLResp     ResponseType = "html"
	JSONResp     ResponseType = "json"
	RedirectResp ResponseType = "redirect"
	FileResp     ResponseType = "file"
	ErrorResp    ResponseType = "error"
)

// AppErr := app error struct
type AppErr struct {
	Error     []error
	RequestID *string
}

// AppResponse := app response struct
type AppResponse struct {
	Payload interface{}
}

// MarshalJSON := marshalling error
func (err *AppErr) MarshalJSON() ([]byte, error) {
	var errJSONArr []map[string]interface{}
	for _, e := range err.Error {
		errJSON := errors.GetMap(e)
		errJSONArr = append(errJSONArr, errJSON)
	}
	return json.Marshal(&struct {
		Error     []map[string]interface{} `json:"error"`
		Sucess    bool                     `json:"success"`
		RequestID *string                  `json:"request_id"`
	}{
		Error:     errJSONArr,
		Sucess:    false,
		RequestID: err.RequestID,
	})
}

// MarshalJSON := marshalling error
func (r *AppResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Sucess  bool        `json:"success"`
		Payload interface{} `json:"payload"`
	}{
		Sucess:  true,
		Payload: &r.Payload,
	})
}
