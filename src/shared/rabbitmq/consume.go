package rabbitmq

import "errors"

type Consumer interface {
	StartConsuming(consumerTag string) error
	StartConsumingByHandle(consumerTag string, handle Handle) error
}

func (c *context) StartConsuming(consumerTag string) error {
	if c.channel == nil {
		return errors.New("No channel found. Please connect to exchange first!")
	}

	err := c.bindQueue()
	if err != nil {
		return err
	}

	deliveries, err := c.channel.Consume(
		c.queueName, // name
		consumerTag, // consumerTag,
		false,       // noAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	go handle(c, deliveries, c.done)

	return err
}

func (c *context) StartConsumingByHandle(consumerTag string, handle Handle) error {
	c.WithHandle(handle)
	return c.StartConsuming(consumerTag)
}