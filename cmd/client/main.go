package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
  ctx := context.Background()
  const connString = "amqp://guest:guest@localhost:5672/"
  fmt.Println("RabbitMQ connection string:", connString)
  conn, err := amqp.Dial(connString)
  if err != nil {
    // handle error
    fmt.Println("Error opening connection: ", err)
    return
  }
  defer conn.Close()
  chann, err := conn.Channel()
  if err != nil {
    errf := fmt.Errorf("%v", err)

    fmt.Println("Error creating channel")
    fmt.Printf("Error: %s\n", errf)
    return
  }
}
