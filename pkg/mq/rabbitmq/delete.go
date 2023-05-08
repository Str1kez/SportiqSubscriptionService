package rabbitmq

import (
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type eventDeleteRequest struct {
	Events []string `json:"events"`
}

func eventDeleteValidation(request *eventDeleteRequest) error {
	if request.Events == nil || len(request.Events) == 0 {
		return errors.New("message on event.delete invalid")
	}
	return nil
}

func (r *RabbitMQ) eventDelete(message *amqp.Delivery) error {
	event := eventDeleteRequest{}
	if err := json.Unmarshal(message.Body, &event); err != nil {
		return err
	}
	if err := eventDeleteValidation(&event); err != nil {
		return err
	}
	var isCorrupted error = nil
	for _, e := range event.Events {
		users, err := r.subscriptionDB.GetUsers(e)
		if err != nil {
			log.Errorf("Couldn't find users: %v\n", err)
			isCorrupted = err
			continue
		}
		if err = r.historyDB.Create(e, nil, users, true); err != nil {
			log.Errorf("Couldn't save event %s into history: %v\n", e, err)
			isCorrupted = err
			continue
		}
		if err = r.subscriptionDB.DeleteEvent(e); err != nil {
			log.Errorf("Couldn't delete event %s from db: %v\n", e, err)
			isCorrupted = err
			continue
		}
	}
	log.Debugf("Handled event.change message: %+v\n", event)
	return isCorrupted
}
