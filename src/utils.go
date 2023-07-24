package src

import (
	"encoding/json"
	"log"
	"net/http"
)

func internalError(err error, w http.ResponseWriter) {
	log.Default().Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	resp := Response{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
		Data:    nil,
	}
	val, _ := json.Marshal(resp)
	_, _ = w.Write(val)
}

func respondWithData(data interface{}, message string, w http.ResponseWriter) {
	resp := Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
	write(&resp, w)
}

func badRequest(message string, w http.ResponseWriter) {
	resp := Response{
		Status:  http.StatusBadRequest,
		Message: message,
	}
	w.WriteHeader(http.StatusBadRequest)
	write(&resp, w)
}

func notFound(message string, w http.ResponseWriter) {
	resp := Response{
		Status:  http.StatusNotFound,
		Message: message,
	}
	w.WriteHeader(http.StatusNotFound)
	write(&resp, w)
}

func write(resp *Response, w http.ResponseWriter) {
	val, err := json.Marshal(resp)
	if err != nil {
		internalError(err, w)
		return
	}
	_, err = w.Write(val)
	if err != nil {
		internalError(err, w)
		return
	}
}
