package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetJSONContent returns the JSON unmarshalled content of a http.Request
func GetJSONContent(v interface{}, r *http.Request) error {
	if r == nil && r.Body == nil {
		return errors.New("cannot get content out of a nil response or body")
	}

	// try to decode body
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}

	// close body
	defer r.Body.Close()

	return nil
}
