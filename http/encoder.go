package api

import (
	"encoding/json"
	"net/http"
)

// RESTError is used to send back an error as a JSON message over HTTP connection
type RESTError struct {
	Error string `json:"error"`
}

// MediaType is used to declare a specific content-type
type MediaType string

const (
	// TypeJSONUTF8 is media type used for UTF-8 JSON content
	TypeJSONUTF8 MediaType = "application/json; charset=UTF-8"
)

// String returns a string representation of a MediaType
func (m MediaType) String() string {
	return string(m)
}

const (
	// HeaderContentType is the key used for response content type
	HeaderContentType = "Content-Type"

	// Error messages
	resourceNotFoundMsg = "resource not found"
	errorMsg            = "error"
)

// SendJSONWithHTTPCode outputs JSON RESTError through a http.ResponseWriter with a given HTTP code
func SendJSONWithHTTPCode(w http.ResponseWriter, d interface{}, code int) error {
	w.Header().Set(HeaderContentType, TypeJSONUTF8.String())
	w.WriteHeader(code)
	if d != nil {
		err := json.NewEncoder(w).Encode(d)
		if err != nil {
			return err
		}
	}
	return nil
}

// SendJSONOk outputs JSON RESTError through a http.ResponseWriter with http.StatusOK code
func SendJSONOk(w http.ResponseWriter, d interface{}) {
	SendJSONWithHTTPCode(w, d, http.StatusOK)
}

// SendJSONError sends an JSON RESTError error through a http.ResponseWriter with a custom message and error code
func SendJSONError(w http.ResponseWriter, error string, code int) {
	SendJSONWithHTTPCode(w, map[string]string{errorMsg: error}, code)
}

// SendJSONNotFound sends an JSON RESTError error through a http.ResponseWriter with a "Resource not found" message
func SendJSONNotFound(w http.ResponseWriter) {
	SendJSONError(w, resourceNotFoundMsg, http.StatusNotFound)
}

// NotFoundHandler returns an http.HandlerFunc with a JSON answer implementation
func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendJSONNotFound(w)
	}
}
