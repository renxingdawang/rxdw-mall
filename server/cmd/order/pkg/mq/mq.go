package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
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
func (r *RabbitMQ) Publish(queue string, event []byte) error {
	//Todo:
	_, err := r.channel.QueueDeclare(
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

	// 发布消息到队列
	return r.channel.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         event,
			DeliveryMode: amqp.Persistent,
		},
	)
}

func (r *RabbitMQ) Consume(queue string, handler func([]byte)) error {
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
