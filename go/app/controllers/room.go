package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mizutanimeen/mandatory-declaration-of-intent/models"
)

func ParseRequestRoom(aRequest *http.Request) (*models.Room, int, error) {
	var tRoom *models.Room
	switch aRequest.Header["Content-Type"][0] {
	case "application/json":
		var tStatus int
		var tError error
		tRoom, tStatus, tError = jsonToRoom(aRequest)
		if tError != nil {
			return nil, tStatus, tError
		}
	default:
		return nil, http.StatusPreconditionFailed, fmt.Errorf("%s is bad content-type", aRequest.Header["Content-Type"][0])
	}

	return tRoom, 0, nil
}

// bodyがjsonのrequestからstruct userを作る
func jsonToRoom(aRequest *http.Request) (*models.Room, int, error) {
	tBody := make([]byte, aRequest.ContentLength)
	aRequest.Body.Read(tBody)

	tRoom := &models.Room{}
	if tError := json.Unmarshal(tBody, tRoom); tError != nil {
		return nil, http.StatusInternalServerError, tError
	}

	return tRoom, 0, nil
}
