package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Service interface {
	Connect() error
}

type RabbitMQ struct {
	Conn *amqp.Connection
}

func (rmq *RabbitMQ) Connect() error {
	fmt.Println("Connecting to rabbitmq...")
	var err error
	rmq.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to open a channel")
	defer rmq.Conn.Close()
	if err != nil{
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully rabbitmq connected...")
	return nil
}

func NewRabbitMqService() *RabbitMQ {
	return &RabbitMQ{}
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}