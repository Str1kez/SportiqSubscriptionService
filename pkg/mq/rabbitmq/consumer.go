package rabbitmq

import (
	log "github.com/sirupsen/logrus"
)

func (r *RabbitMQ) Consume(autoAck bool) {
	messages, err := r.channel.Consume(r.queue.Name, "", autoAck, false, false, false, nil)
	if err != nil {
		log.Panicf("Couldn't init consumer: %v\n", err)
	}

	var worker chan struct{}
	log.Infof("Consumer №%d is receiving messages on queue [%s]\n", r.id, r.queue.Name)

	go func() {
		for d := range messages {
			log.Debugf("Consumer №%d received a message: %s\n", r.id, d.Body)
			switch d.Type {
			case "event.create":
				if err := r.eventCreate(&d); err != nil {
					log.Errorf("Couldn't create event: %v\n", err)
				}
			case "event.change":
				if err := r.eventChange(&d); err != nil {
					log.Errorf("Couldn't change events: %v\n", err)
				}
			case "event.complete":
				if err := r.eventComplete(&d); err != nil {
					log.Errorf("Couldn't complete events: %v\n", err)
				}
			case "event.delete":
				if err := r.eventDelete(&d); err != nil {
					log.Errorf("Couldn't delete events: %v\n", err)
				}
			default:
				log.Errorf("Unrecognized type of message: %s\n", d.Type)
			}
			log.Debugf("Done by consumer №%d\n", r.id)
			d.Ack(false)
		}
	}()

	<-worker
}
