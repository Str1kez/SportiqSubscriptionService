package rabbitmq

import (
	"fmt"

	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
	"github.com/Str1kez/SportiqSubscriptionService/internal/db"
	"github.com/Str1kez/SportiqSubscriptionService/internal/history"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type RabbitMQ struct {
	id             uint8
	connection     *amqp.Connection
	channel        *amqp.Channel
	queue          *amqp.Queue
	subscriptionDB *db.SubscriptionDB
	historyDB      *history.HistoryDB
}

func NewRabbitMQ(id uint8, config *config.MQConfig, subscriptionDB *db.SubscriptionDB, historyDB *history.HistoryDB) *RabbitMQ {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.User, config.Password, config.Host, config.Port))
	if err != nil {
		log.Fatalf("Couldn't connect to RabbitMQ: %v\n", err)
	}
	log.Infof("Consumer №%d connected\n", id)
	return &RabbitMQ{connection: connection, id: id, subscriptionDB: subscriptionDB, historyDB: historyDB}
}

func (r *RabbitMQ) OpenChannel() {
	channel, err := r.connection.Channel()
	if err != nil {
		log.Panicf("Couldn't open channel: %v\n", err)
	}
	r.channel = channel
}

func (r *RabbitMQ) DeclareQueue(name string, durable bool, prefetchCount int) {
	queue, err := r.channel.QueueDeclare(name, durable, false, false, false, nil)
	if err != nil {
		log.Panicf("Couldn't declare queue [%s]: %v\n", name, err)
	}
	err = r.channel.Qos(prefetchCount, 0, false)
	if err != nil {
		log.Panicf("Couldn't declare params for queue [%s]: %v\n", name, err)
	}
	r.queue = &queue
}

func (r *RabbitMQ) Close() error {
	if r.channel != nil && !r.channel.IsClosed() {
		err := r.channel.Close()
		if err != nil {
			return err
		}
	}
	if r.connection != nil && !r.connection.IsClosed() {
		err := r.connection.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
