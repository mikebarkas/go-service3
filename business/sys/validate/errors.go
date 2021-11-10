package validate

import (
	"encoding/json"
	"errors"
)

// ErrorInvalidID.
var ErrorInvalidID = errors.New("ID is not valid")

// ErrorResponse is used for API failure responses.
type ErrorResponse struct {
	Error  string `json:"error"`
	Fields string `json:"fields,omitempty"`
}

// RequestError is used to pass an error during the request throught the app.
type RequestError struct {
	Err    error
	Status int
	Fields error
}

func NewRequestError(err error, status int) error {
	return &RequestError{err, status, nil}
}

func (err *RequestError) Error() string {
	return err.Err.Error()
}

// FieldError is used to indicate an error with a specific request field.
type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// FieldErrors represents a collection of field errors.
type FieldErrors []FieldError

// Error implments the error interface.
func (fe FieldErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

// Fields returns the fields that failed validation
func (fe FieldErrors) Fields() map[string]string {
	m := make(map[string]string)
	for _, fld := range fe {
		m[fld.Field] = fld.Error
	}
	return m
}

// IsFieldErrors checks if an error of type FieldErrors exists.
func IsFieldErrors(err error) bool {
	var fe FieldErrors
	return errors.As(err, &fe)
}

// GetFieldErrors returns a copy of the FieldErrors pointer.
func GetFieldErrors(err error) FieldErrors {
	var fe FieldErrors
	if !errors.As(err, &fe) {
		return nil
	}
	return fe
}

// Cause iterates through all the wrapped errors.
func Cause(err error) error {
	root := err
	for {
		if err := errors.Unwrap(root); err == nil {
			return root
		}
		root = err
	}
}