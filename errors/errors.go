package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorType is the type of an error
type ErrorType string

const (
	// InvalidField error
	InvalidField ErrorType = "InvalidField"
	// NoType error
	NoType ErrorType = "null"
	// BadRequest error
	BadRequest ErrorType = "BadRequest"
	// NotFound error
	NotFound ErrorType = "NotFound"
	// InvalidSchema error
	InvalidSchema ErrorType = "InvalidSchema"
	// DBError error
	DBError ErrorType = "DBError"
	// SessionError error
	SessionError ErrorType = "SessionError"
	// AuthenticationError error
	AuthenticationError ErrorType = "AuthenticationError"
	// PermissionDenied Error
	PermissionDenied ErrorType = "PermissionDenied"

	// ServerError Error
	ServerError ErrorType = "ServerError"
)

type customError struct {
	errorType     ErrorType
	originalError error
	context       errorContext
}

type errorContext struct {
	Field   interface{}
	Message interface{}
}

// New creates a new customError
func (errorType ErrorType) New(msg string) error {
	return customError{errorType: errorType, originalError: errors.New(msg)}
}

// Newf creates a new customError with formatted message
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrap creates a new wrapped error
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

// Wrapf creates a new wrapped error with formatted message
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: errors.Wrapf(err, msg, args...)}
}

// Error returns the mssage of a customError
func (error customError) Error() string {
	return error.originalError.Error()
}

// New creates a no type error
func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) error {
	return customError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrap an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			context:       customErr.context,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field interface{}, message interface{}) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(customError); ok {
		return customError{errorType: customErr.errorType, originalError: customErr.originalError, context: context}
	}
	return customError{errorType: NoType, originalError: err, context: context}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]interface{} {
	emptyError := errorContext{}
	if customErr, ok := err.(customError); ok || customErr.context != emptyError {
		return map[string]interface{}{"field": customErr.context.Field, "message": customErr.context.Message}
	}
	return nil
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}
	return NoType
}

// GetMap object of the error
func GetMap(err error) map[string]interface{} {
	emptyError := errorContext{}
	var errMap = make(map[string]interface{})
	if customErr, ok := err.(customError); ok && customErr.context != emptyError {
		errMap["field"] = customErr.context.Field
		errMap["message"] = customErr.context.Message
		errMap["type"] = GetType(customErr)
		return errMap
	}
	errMap["message"] = err.Error()
	errMap["type"] = GetType(err)
	fmt.Println(errMap)
	return errMap
}
