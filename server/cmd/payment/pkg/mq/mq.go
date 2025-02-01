package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"strconv"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(conn *amqp.Connection) *RabbitMQ {
	channel, err := conn.Channel()
	if err != nil {
		klog.Fatal("Failed to create RabbitMQ instance:", err)
		return nil
	}
	return &RabbitMQ{conn: conn, channel: channel}
}

func (r *RabbitMQ) PublishWithDelay(queue string, body []byte, delayMs int) error {
	_, err := r.channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    "payment_exchange",
			"x-dead-letter-routing-key": "payment_cancel",
		},
	)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			Expiration:   strconv.Itoa(delayMs),
			DeliveryMode: amqp.Persistent,
		},
	)
}
func (r *RabbitMQ) Publish(queue string, event []byte) error {
	//Todo:
	return nil
}
func (r *RabbitMQ) Consume(queue string, handler func(string)) error {
	q, err := r.channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(string(d.Body))
		}
	}()
	return nil
}
func (r *RabbitMQ) ConsumeByte(queue string, handler func([]byte)) error {
	q, err := r.channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler([]byte(d.Body))
		}
	}()
	return nil
}
