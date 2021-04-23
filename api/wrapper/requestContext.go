package wrapper

import (
	"my/v1/model"
)

// RequestContext := contains request/reponse related data
type RequestContext struct {
	RequestID    string
	Response     *model.AppResponse
	Err          *model.AppErr
	ResponseType model.ResponseType
	ResponseCode int
}

// SetErr := setting Err reponse in request context
func (requestCTX *RequestContext) SetErr(err error) {
	appErr := requestCTX.Err
	requestCTX.ResponseType = model.ErrorResp
	if appErr == nil {
		appErr = &model.AppErr{}
	}
	appErr.Error = append(appErr.Error, err)
	requestCTX.Err = appErr
}

// SetAppResponse := setting app response in request context
func (requestCTX *RequestContext) SetAppResponse(message interface{}, statusCode int) {
	requestCTX.ResponseType = model.JSONResp
	requestCTX.ResponseCode = statusCode
	requestCTX.Response = &model.AppResponse{
		Payload: message,
	}
}

// SetErrs adds slice of errors in requestCTX
func (requestCTX *RequestContext) SetErrs(errs []error) {
	for _, e := range errs {
		requestCTX.SetErr(e)
	}
}

// SetHTMLResponse := setting app html response in request context
func (requestCTX *RequestContext) SetHTMLResponse(message []byte, statusCode int) {
	requestCTX.ResponseType = model.HTMLResp
	requestCTX.ResponseCode = statusCode
	requestCTX.Response = &model.AppResponse{
		Payload: message,
	}
}
