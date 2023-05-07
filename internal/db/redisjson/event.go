package redisjson

import (
	"context"
	"fmt"

	"github.com/Str1kez/SportiqSubscriptionService/tools"
	log "github.com/sirupsen/logrus"
)

type creationModel struct {
	Status string   `json:"status"`
	Users  []string `json:"users"`
}

func (r *ReJSONDB) CreateEvent(eventId, userId string) error {
	documentName := fmt.Sprintf("events:%s", eventId)
	event := creationModel{Status: r.config.EventStatus.Planned, Users: []string{userId}}
	res, err := r.handler.JSONSet(documentName, ".", event)
	if err != nil {
		log.Errorf("Failed to JSONSet: %v\n", err)
		return err
	}
	if res.(string) == "OK" {
		log.Debugf("Success: %s\n", res)
	} else {
		log.Errorln("Failed to Set Event")
		return fmt.Errorf("failed to create event: %s", eventId)
	}
	return nil
}

func (r *ReJSONDB) DeleteEvent(eventId string) error {
	documentName := fmt.Sprintf("events:%s", eventId)
	cmd := r.client.Del(context.Background(), documentName)
	return cmd.Err()
}

func (r *ReJSONDB) UpdateEventStatus(eventId, eventStatus string) error {
	documentName := fmt.Sprintf("events:%s", eventId)
	if err := tools.EventStatusChecker(&r.config.EventStatus, eventStatus); err != nil {
		log.Panicf("Wrong event status in args: %v\n", err)
		return err
	}
	res, err := r.handler.JSONSet(documentName, ".status", eventStatus)
	if err != nil {
		log.Errorf("Failed to JSONSet: %v\n", err)
		return err
	}
	if res.(string) == "OK" {
		log.Debugf("Success: %s\n", res)
	} else {
		log.Errorln("Failed to Update Event Status")
		return fmt.Errorf("failed to update event: %s", eventId)
	}
	return nil
}
