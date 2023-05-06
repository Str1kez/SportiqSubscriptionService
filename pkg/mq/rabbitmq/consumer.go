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
			// TODO: business
			log.Debugf("Done by consumer №%d\n", r.id)
			d.Ack(false)
		}
	}()

	<-worker
}
