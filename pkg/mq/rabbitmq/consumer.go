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

	go func() {
		for d := range messages {
			log.Debugf("Received a message: %s", d.Body)
			d.Ack(false)
		}
	}()

	<-worker
}
