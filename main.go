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
	return nil
}
func main() {
	if err:=Run();err!=nil{
		fmt.Println("Error Setting of our application")
		fmt.Println(err)
	}
}