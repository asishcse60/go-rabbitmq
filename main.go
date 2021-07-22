package main

import (
	"fmt"
	"github.com/asiscse60/go-rabbitmq/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error{
	fmt.Println("Go RabbitMq Experiment...")
	rmq := rabbitmq.NewRabbitMqService()

	app:= App{Rmq: rmq}
	err := app.Rmq.Connect()
	if err != nil{
		return err
	}
	defer app.Rmq.Conn.Close()
	err = app.Rmq.Publish("Hello boys")
	if err != nil{
		return err
	}
	app.Rmq.Consume()
	return nil
}
func main() {
	if err:=Run();err!=nil{
		fmt.Println("Error Setting of our application")
		fmt.Println(err)
	}
}
