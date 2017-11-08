package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestSendJSONWithHTTPCode test sending struct as JSON bodies
func TestSendJSONWithHTTPCode(t *testing.T) {
	// build a response recorder
	response := httptest.NewRecorder()
	customError := RESTError{
		"myCustomErrorMessage",
	}

	// send JSON body
	err := SendJSONWithHTTPCode(response, customError, http.StatusOK)
	if err != nil {
		t.Errorf("failed to send successfuly a struct as a JSON body, %v", err)
	}

	if http.StatusOK != response.Code {
		t.Errorf("Expected response code %d is different from decoded one %d", http.StatusOK, response.Code)
	}

	customErrorRsp := RESTError{}
	err = json.NewDecoder(response.Body).Decode(&customErrorRsp)
	if err != nil {
		t.Errorf("failed to unmarshall response body, %v", err)
	}

	if customError != customErrorRsp {
		t.Errorf("Expected body %v is different from decoded one %v", customError, customErrorRsp)
	}

}
