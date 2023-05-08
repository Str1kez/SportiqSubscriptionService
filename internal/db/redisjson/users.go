package redisjson

import (
	"encoding/json"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type usersResponse struct {
	Users []string `json:"users"`
}

func (r *ReJSONDB) GetUsers(eventId string) ([]string, error) {
	documentName := fmt.Sprintf("events:%s", eventId)
	res, err := r.handler.JSONGet(documentName, ".")
	if err != nil {
		log.Errorf("Couldn't get event [%s]: %v\n", eventId, err)
		return nil, err
	}

	if _, ok := res.([]byte); !ok {
		return nil, errors.New("couldn't parse users from document")
	}
	userJSON := res.([]byte)
	users := usersResponse{}
	if err := json.Unmarshal(userJSON, &users); err != nil {
		log.Errorf("Couldn't parse to JSON: %v\n", err)
		return nil, err
	}
	log.Debugf("%+v\n", users.Users)
	return users.Users, nil
}
