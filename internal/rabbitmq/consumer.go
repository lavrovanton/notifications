package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageHandler func(message []byte)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error
	mh      MessageHandler
}

func NewConsumer(uri, queueName string, messageHandler func(message []byte),
) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		done:    make(chan error),
		mh:      messageHandler,
	}

	var err error

	c.conn, err = amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := c.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	deliveries, err := c.channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	go c.handle(deliveries, c.done)

	return c, nil
}

func (c *Consumer) handle(deliveries <-chan amqp.Delivery, done chan error) {
	cleanup := func() {
		done <- nil
	}

	defer cleanup()

	for d := range deliveries {
		c.mh(d.Body)
	}
}

func (c *Consumer) Shutdown() error {
	if err := c.channel.Cancel("", true); err != nil {
		return err
	}
	if err := c.conn.Close(); err != nil {
		return err
	}
	return <-c.done
}
