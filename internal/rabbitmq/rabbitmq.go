package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Service interface {
	Connect() error
	Publish(message string) error
}

type RabbitMQ struct {
	Conn *amqp.Connection
	Channel *amqp.Channel
}

func (rmq *RabbitMQ) Connect() error {
	fmt.Println("Connecting to rabbitmq...")
	var err error
	rmq.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to open a channel")

	if err != nil{
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully rabbitmq connected...")
	rmq.Channel, err = rmq.Conn.Channel()
	if err!= nil{
		return err
	}
	_,err = rmq.Channel.QueueDeclare("TestQueue",false,false,false, false, nil)
	return nil
}

func (rmq *RabbitMQ) Publish(message string)error  {
	err:= rmq.Channel.Publish("","TestQueue", false, false,amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(message),
	})
	if err != nil{
		return err
	}
	fmt.Println("Successfully published message to queue")
	return nil
}

func (rmq *RabbitMQ) Consume() {
	megs,err:= rmq.Channel.Consume("TestQueue","", true, false, false, false, nil)
	if err!=nil{
		fmt.Println(err)
		return
	}
	for msg:= range megs {
		fmt.Printf("Received message: %s\n", msg.Body)
	}
}
func NewRabbitMqService() *RabbitMQ {
	return &RabbitMQ{}
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}