package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
  const connString = "amqp://guest:guest@localhost:5672/"
  fmt.Println("RabbitMQ connection string:", connString)
  conn, err := amqp.Dial(connString)
  if err != nil {
    // handle error
    fmt.Println("Error opening connection: ", err)
    return
  }
  defer conn.Close()
}
