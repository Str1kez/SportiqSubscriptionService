package mq

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/db"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/mq/rabbitmq"
	log "github.com/sirupsen/logrus"
)

type Consumer interface {
	OpenChannel()
	DeclareQueue(name string, durable bool, prefetchCount int)
	Consume(autoAck bool)
	Close() error
}

type MQConsumer struct {
	Consumer
}

func InitMQConsumer(id uint8, config *config.MQConfig, subscriptionDB *db.SubscriptionDB) *MQConsumer {
	return &MQConsumer{rabbitmq.NewRabbitMQ(id, config, subscriptionDB)}
}

func InitMQConsumerSlice(config *config.MQConfig, subscriptionDBs []*db.SubscriptionDB) []*MQConsumer {
	if config.ConsumerCount == 0 {
		log.Errorln("Zero consumers declared")
	}
	var i uint8
	consumerSlice := make([]*MQConsumer, 0, config.ConsumerCount)
	for i = 0; i < config.ConsumerCount; i++ {
		consumerSlice = append(consumerSlice, InitMQConsumer(i, config, subscriptionDBs[i]))
	}
	return consumerSlice
}

func HandleMessages(consumers ...*MQConsumer) {
	var handle chan struct{}
	for _, c := range consumers {
		c.OpenChannel()
		c.DeclareQueue("events", true, 1)
		go c.Consume(false)
	}
	<-handle
}

func GracefulShutdown(consumers ...*MQConsumer) {
	for _, c := range consumers {
		err := c.Close()
		if err != nil {
			log.Errorf("Error in MQConsumer.Close(): %v\n", err)
		}
	}
}
