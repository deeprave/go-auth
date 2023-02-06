package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/deeprave/go-testutils/test"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Record struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func TestJSONWriteResponse(t *testing.T) {
	// func (app *App) ModelToJson(w http.ResponseWriter, status int, data interface{}, wrap string) (int, error)
	records := []Record{
		{Id: 14, Name: "Alpha", Description: "This is Alpha"},
		{Id: 22, Name: "Beta", Description: "This is Beta"},
		{Id: 39, Name: "Gamma", Description: "This is Gamma"},
	}
	rw := httptest.NewRecorder()
	err := JSONWriteResponse(rw, records, http.StatusOK)
	test.ShouldBeNoError(t, err, "Unexpected error in JSONWriteResponse: %v", err)
	response := rw.Result()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		test.ShouldBeNoError(t, err, "Unexpected error in ReadAll(response.Body): %v", err)
	} else {
		test.ShouldBeFalse(t, len(string(responseBody)) <= 0, "response headers + body has invalid length")
		test.ShouldNotBeNil(t, response.Header.Get("Content-Type"), "Does not contain Content-Type")
		test.ShouldNotBeNil(t, response.Header.Get("Content-Length"), "Does not contain Content-Length")
	}

	rw = httptest.NewRecorder()
	err = JSONWriteResponse(rw, records, http.StatusOK)
	response = rw.Result()
	var data []Record
	err = json.NewDecoder(response.Body).Decode(&data)
	test.ShouldBeNoError(t, err, "error unmarshalling response: %v", err)
	test.ShouldBeEqual(t, records, data)
}

func TestJSONErrorResponse(t *testing.T) {
	errText := "this is an error message"
	customError := errors.New(errText)
	rw := httptest.NewRecorder()
	err := JSONErrorResponse(rw, customError, http.StatusInternalServerError)
	test.ShouldBeNoError(t, err, "error sending error response: %v", err)
	response := rw.Result()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		test.ShouldBeNoError(t, err, "Unexpected error in ReadAll(response.Body): %v", err)
	} else {
		var respdata = JSONResponse{
			Error:   false,
			Message: "test",
		}
		buffer := bytes.NewBuffer(responseBody)
		err = json.NewDecoder(buffer).Decode(&respdata)
		test.ShouldBeNoError(t, err, "error unmarshalling response: %v", err)
		test.ShouldBeEqual(t, errText, respdata.Message)
	}
}
