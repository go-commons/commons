package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetJSONContent(t *testing.T) {

	// struct to marshall
	customError := RESTError{
		"myCustomErrorMessage",
	}

	// json marshalling struct
	content, err := json.Marshal(&customError)
	if err != nil {
		t.Errorf("failed to marshall struct as a JSON body, %v", err)
	}
	reader := bytes.NewReader(content)

	// building new request
	request, err := http.NewRequest(http.MethodPost, "localhost:8020", reader)
	if err != nil {
		t.Errorf("failed to build test resquest, %v", err)
	}

	customErrorRsp := RESTError{}
	err = GetJSONContent(&customErrorRsp, request)
	if err != nil {
		t.Errorf("failed to unmarshall request body, %v", err)
	}

	if customError != customErrorRsp {
		t.Errorf("Expected body %v is different from decoded one %v", customError, customErrorRsp)
	}

}
