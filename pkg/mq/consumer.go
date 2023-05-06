package mq

import (
	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/pkg/mq/rabbitmq"
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

func InitMQConsumer(config *config.MQConfig) *MQConsumer {
	return &MQConsumer{
		rabbitmq.NewRabbitMQ(config),
	}
}
