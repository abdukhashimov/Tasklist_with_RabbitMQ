package event

import (
	"encoding/json"
	"log"

	"bitbucket.org/alien_soft/TaskListRabbitMQ/task"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Channel *amqp.Channel
}

//Connection
//Declare Exchange
//Declare queue
func NewRabbitMQ() RabbitMQ {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	defer conn.Close()

	if err != nil {
		log.Panic("Failed to connection RabbitMQ", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Panic("Failed to create a new channel", err)
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		"course",
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Panic("Failed to declare exchange", err)
	}

	queue1, err := ch.QueueDeclare(
		"course.update",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		log.Panic("Error to declare queue", err)
	}

	queue2, err := ch.QueueDeclare(
		"course.create",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		log.Panic("Error to declare queue", err)
	}

	err = ch.QueueBind(
		queue1.Name,
		"course.#",
		"course",
		false,
		nil,
	)

	if err != nil {
		log.Panic("Error while binding to exchange", err)
	}

	err = ch.QueueBind(
		queue2.Name,
		"course.create",
		"course",
		false,
		nil,
	)

	if err != nil {
		log.Panic("Error while binding to exchange", err)
	}

	return RabbitMQ{
		Channel: ch,
	}
}

func (r *RabbitMQ) Publish(exchangeName, route string, body task.Task) error {
	bodyByte, err := json.Marshal(body)
	if err != nil {
		return nil
	}

	err = r.Channel.Publish(
		exchangeName,
		route,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bodyByte,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
