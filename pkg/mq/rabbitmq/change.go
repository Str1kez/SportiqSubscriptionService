package rabbitmq

import (
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type eventChangeRequest struct {
	Status string   `json:"status"`
	Events []string `json:"events"`
}

func eventChangeValidation(request *eventChangeRequest) error {
	if request.Status == "" || request.Events == nil || len(request.Events) == 0 {
		return errors.New("message on event.create invalid")
	}
	return nil
}

func (r *RabbitMQ) eventChange(message *amqp.Delivery) error {
	event := eventChangeRequest{}
	if err := json.Unmarshal(message.Body, &event); err != nil {
		return err
	}
	if err := eventChangeValidation(&event); err != nil {
		return err
	}
	for _, e := range event.Events {
	  if err := r.subscriptionDB.UpdateEventStatus(e, event.Status); err != nil {
      log.Errorf("Error in event [%s] updating: %v\n", e, err)
	  }
	}
	log.Debugf("Handled event.change message: %+v\n", event)
	return nil
}
