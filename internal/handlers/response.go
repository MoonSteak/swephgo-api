package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Data  interface{}    `json:"data,omitempty"`
	Error *ResponseError `json:"error,omitempty"`
}

func NewDataResponse(data interface{}) Response {
	return Response{Data: data}
}

func NewErrorResponse(code string) Response {
	log.Debugf("NewErrorResponse %s", code)
	return Response{Error: &ResponseError{Code: ErrorCode(code)}}
}

type ErrorCode string

type ResponseError struct {
	Code ErrorCode `json:"code,omitempty"`
}

func writeResponse(w http.ResponseWriter, status int, response Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	b, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func WriteResponse(w http.ResponseWriter, status int, response Response) {
	writeResponse(w, status, response)
}

func WriteDataResponse(w http.ResponseWriter, data interface{}) {
	WriteResponse(w, http.StatusOK, NewDataResponse(data))
}

func WriteErrorResponse(w http.ResponseWriter, status int, err string) {
	WriteResponse(w, status, NewErrorResponse(err))
}
