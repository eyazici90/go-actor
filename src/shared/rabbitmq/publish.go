package rabbitmq

import (
	"time"

	"github.com/streadway/amqp"
)

type Publisher interface {
	Publish(message Message) error
}

func (c *context) Publish(msg Message) error {
	if c.channel == nil {
		return NoChannelFound
	}

	payload, err := toJsonDefault(msg)

	if err != nil {
		return err
	}

	err = c.channel.Publish(
		c.exchangeName, // exchange
		c.bindingKey,   // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			DeliveryMode: defaultDeliveryMode,
			ContentType:  defaultContentType,
			Body:         payload,
			Timestamp:    time.Now(),
		})

	return err
}
