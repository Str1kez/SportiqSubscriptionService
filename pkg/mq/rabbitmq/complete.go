package rabbitmq

import (
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type eventCompleteRequest struct {
	Events []completedEvents `json:"events"`
}

type completedEvents struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func eventCompleteValidation(request *eventCompleteRequest) error {
	if request.Events == nil || len(request.Events) == 0 {
		return errors.New("message on event.complete invalid")
	}
	return nil
}

func (r *RabbitMQ) eventComplete(message *amqp.Delivery) error {
	event := eventCompleteRequest{}
	if err := json.Unmarshal(message.Body, &event); err != nil {
		return err
	}
	if err := eventCompleteValidation(&event); err != nil {
		return err
	}
	var isCorrupted error = nil
	for _, e := range event.Events {
		users, err := r.subscriptionDB.GetUsers(e.Id)
		if err != nil {
			log.Errorf("Couldn't find users: %v\n", err)
			isCorrupted = err
			continue
		}
		if err = r.historyDB.Create(e.Id, e.Title, users, false); err != nil {
			log.Errorf("Couldn't save event %s into history: %v\n", e.Id, err)
			isCorrupted = err
			continue
		}
		if err = r.subscriptionDB.DeleteEvent(e.Id); err != nil {
			log.Errorf("Couldn't delete event %s from db: %v\n", e.Id, err)
			isCorrupted = err
			continue
		}
	}
	log.Debugf("Handled event.change message: %+v\n", event)
	return isCorrupted
}
