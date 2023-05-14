package redisjson

import (
	"context"
	"fmt"
	"strings"

	"github.com/Str1kez/SportiqSubscriptionService/internal/dto"
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
	res, err := r.handler.JSONSet(documentName, ".", event) // ! need flag NX
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
	res, err := r.handler.JSONSet(documentName, ".status", eventStatus) // ! need flag XX
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

func (r *ReJSONDB) GetEvents(userId string) ([]dto.SubscriptionResponse, error) {
	escapedUserId := tools.EscapeUUID(userId)
	query := fmt.Sprintf(`FT.SEARCH idx:events @user_id:{%s} RETURN 1 status LIMIT 0 10000`, escapedUserId)
	querySlice := strings.Split(query, " ")
	q := make([]interface{}, len(querySlice))
	for i, v := range querySlice {
		q[i] = v
	}
	cmd := r.client.Do(context.Background(), q...)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	cmdSlice, err := cmd.Slice()
	if err != nil {
		return nil, err
	}
	resultSlice := make([]dto.SubscriptionResponse, 0, len(cmdSlice)/2)
	var result dto.SubscriptionResponse
	for _, v := range cmdSlice {
		switch value := v.(type) {
		case string:
			result.EventId, _ = strings.CutPrefix(value, "events:")
		case []interface{}:
			if type_, ok := value[0].(string); ok && type_ == "status" && len(value) == 2 {
				if status, ok := value[1].(string); ok {
					result.Status = status
					resultSlice = append(resultSlice, result)
				}
			}
		}
	}
	return resultSlice, nil
}
