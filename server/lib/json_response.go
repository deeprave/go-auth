package lib

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SetContentType(w http.ResponseWriter, content_type string, status int) {
	w.Header().Set("Content-Type", content_type)
	w.WriteHeader(status)
}

func SetJson(w http.ResponseWriter, status int) {
	SetContentType(w, "application/json", status)
}

const MaxBytes = 1024 * 1024

func JSONWriteResponse(w http.ResponseWriter, data interface{}, status int, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err == nil {
		if len(headers) > 0 {
			for key, value := range headers[0] {
				w.Header()[key] = value
			}
		}
		SetJson(w, status)
		_, err = w.Write(out)
	}
	return err
}

func JSONReadRequest(w http.ResponseWriter, r *http.Request, data interface{}) error {
	r.Body = http.MaxBytesReader(w, r.Body, MaxBytes)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(data)
	if err == nil {
		err = decoder.Decode(&struct{}{})
		if err == io.EOF {
			return nil
		}
		err = errors.New("body must contain a single JSON value")
	}
	return err
}

func JSONErrorResponse(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var report = JSONResponse{
		Error:   true,
		Message: err.Error(),
	}
	return JSONWriteResponse(w, report, statusCode)
}
