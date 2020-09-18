package rabbitmq

import (
	"github.com/streadway/amqp"
)

type context struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	done         chan error
	amqpURI      string
	exchangeName string
	exchangeType string
	queueName    string
	bindingKey   string
	handler      Handle
}

type ContextBuilder interface {
	WithExchange(name, exchangeType string) RabbitMqContext
	WithQueue(queueName, key string) RabbitMqContext
	WithHandle(handle Handle) RabbitMqContext
}

type RabbitMqContext interface {
	ContextBuilder
	Connector
	Publisher
	Consumer
}

func NewContext(amqpURI string) RabbitMqContext {
	return &context{
		conn:    nil,
		channel: nil,
		amqpURI: amqpURI,
	}
}

func WithContext(rCtx RabbitMqContext) RabbitMqContext {
	ctx := rCtx.(*context)

	c := &context{
		conn:         ctx.conn,
		channel:      ctx.channel,
		amqpURI:      ctx.amqpURI,
		bindingKey:   ctx.bindingKey,
		exchangeName: ctx.exchangeName,
		exchangeType: ctx.exchangeType,
		queueName:    ctx.queueName,
		done:         ctx.done,
		handler:      ctx.handler,
	}
	return c
}

func (c *context) WithExchange(name, exchangeType string) RabbitMqContext {
	c.exchangeName = name
	c.exchangeType = exchangeType
	return c

}

func (c *context) WithQueue(queueName, key string) RabbitMqContext {
	c.queueName = queueName
	c.bindingKey = key
	return c

}

func (c *context) WithHandle(handle Handle) RabbitMqContext {
	c.handler = handle
	return c
}
