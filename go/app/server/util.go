package server

import (
	"net/http"
)

func responseWrite(aResponseWriter http.ResponseWriter, aStatus int, aBody []byte) (int, error) {
	aResponseWriter.Header().Set("Content-Type", "application/json")
	aResponseWriter.WriteHeader(aStatus)
	if _, tError := aResponseWriter.Write(aBody); tError != nil {
		return http.StatusInternalServerError, tError
	}
	return 0, nil
}
