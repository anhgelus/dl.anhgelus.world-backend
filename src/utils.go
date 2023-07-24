package src

import (
	"encoding/json"
	"net/http"
)

func internalError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	resp := Response{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
		Data:    nil,
	}
	val, _ := json.Marshal(resp)
	w.Write(val)
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
