package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mizutanimeen/mandatory-declaration-of-intent/models"
)

func ParseRequestUser(aRequest *http.Request) (*models.User, int, error) {
	var tUser *models.User
	switch aRequest.Header["Content-Type"][0] {
	case "application/json":
		var tStatus int
		var tError error
		tUser, tStatus, tError = jsonToUser(aRequest)
		if tError != nil {
			return nil, tStatus, tError
		}
	default:
		return nil, http.StatusPreconditionFailed, fmt.Errorf("%s is bad content-type", aRequest.Header["Content-Type"][0])
	}

	return tUser, 0, nil
}

// bodyがjsonのrequestからstruct userを作る
func jsonToUser(aRequest *http.Request) (*models.User, int, error) {
	tBody := make([]byte, aRequest.ContentLength)
	aRequest.Body.Read(tBody)

	tUser := &models.User{}
	if tError := json.Unmarshal(tBody, tUser); tError != nil {
		return nil, http.StatusInternalServerError, tError
	}

	return tUser, 0, nil
}
