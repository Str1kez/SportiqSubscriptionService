package redisjson

import (
	"errors"
	"fmt"
	"strings"
)

func (r *ReJSONDB) Subscribe(userId, eventId string) error {
	documentName := fmt.Sprintf("events:%s", eventId)
	index, err := r.handler.JSONArrIndex(documentName, ".users", userId)
	if err != nil {
		return errors.New("не смог рассмотреть подписчиков события")
	}
	if index.(int64) != -1 {
		return errors.New("вы уже подписаны на это событие")
	}
	event, err := r.handler.JSONGet(documentName, ".status")
	if err != nil {
		return errors.New("не смог найти событие")
	}

	var status string
	if eventBytes, ok := event.([]byte); ok {
		status = string(eventBytes)
		status = strings.Trim(status, "\"")
	} else {
		return errors.New("couldn't parse status")
	}

	switch status {
	case r.config.EventStatus.Deleted:
		return errors.New("нельзя подписаться на удаленное событие")
	case r.config.EventStatus.Completed:
		return errors.New("нельзя подписаться на завершенное событие")
	}

	_, err = r.handler.JSONArrAppend(documentName, ".users", userId)
	if err != nil {
		return errors.New("не получилось подписаться")
	}
	return nil
}

func (r *ReJSONDB) Unsubscribe(userId, eventId string) error {
	documentName := fmt.Sprintf("events:%s", eventId)
	index, err := r.handler.JSONArrIndex(documentName, ".users", userId)
	if err != nil {
		return errors.New("не смог рассмотреть подписчиков события")
	}
	indexConverted := index.(int64)
	switch indexConverted {
	case -1:
		return errors.New("вы не подписаны на это событие")
	case 0:
		return errors.New("нельзя отписаться от своего события")
	}
	event, err := r.handler.JSONGet(documentName, ".status")
	if err != nil {
		return errors.New("не смог найти событие")
	}

	var status string
	if eventBytes, ok := event.([]byte); ok {
		status = string(eventBytes)
		status = strings.Trim(status, "\"")
	} else {
		return errors.New("couldn't parse status")
	}

	switch status {
	case r.config.EventStatus.Completed:
		return errors.New("нельзя отписаться от завершенного события")
	}

	_, err = r.handler.JSONArrPop(documentName, ".users", int(indexConverted))
	if err != nil {
		return errors.New("не получилось отписаться")
	}
	return nil
}
