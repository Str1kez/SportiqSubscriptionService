package rabbitmq

import (
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type eventCreationRequest struct {
	UserId  string `json:"user_id"`
	EventId string `json:"event_id"`
}

func eventCreationValidate(request *eventCreationRequest) error {
	if request.EventId == "" || request.UserId == "" {
		return errors.New("message on event.create invalid")
	}
  return nil
}

func (r *RabbitMQ) eventCreate(message *amqp.Delivery) error {
	event := eventCreationRequest{}
	if err := json.Unmarshal(message.Body, &event); err != nil {
		return err
	}
  if err := eventCreationValidate(&event); err != nil {
    return err
  }
  if err := r.subscriptionDB.CreateEvent(event.EventId, event.UserId); err != nil {
    return err
  }
	log.Debugf("Handled event.create message: %+v\n", event)
	return nil
}
