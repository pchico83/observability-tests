package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

type operator struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

var (
	amqpURI = "amqp://guest:guest@amqp:5672/"
)

var conn *amqp.Connection
var ch *amqp.Channel
var replies <-chan amqp.Delivery

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func start() error {
	client := getRedisClient()
	err := client.Set("value", 1, 0).Err()
	if err != nil {
		return err
	}

	val, err := client.Get("value").Result()
	if err != nil {
		return err
	}
	fmt.Println("value", val)
	return nil
}

func configureAMQP() {
	log.Printf("getting Connection...")
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		panic(err)
	}

	log.Printf("got Connection, getting Channel...")
	ch, err = conn.Channel()
	if err != nil {
		panic(err)
	}

	log.Printf("got Channel, declaring Exchange (%s)", "go-test-exchange")
	err = ch.ExchangeDeclare(
		"go-test-exchange", // name of the exchange
		"direct",           // type
		true,               // durable
		false,              // delete when complete
		false,              // internal
		false,              // noWait
		nil,                // arguments
	)
	if err != nil {
		panic(err)
	}

	log.Printf("declared Exchange, declaring Queue (%s)", "go-test-queue")
	q, err := ch.QueueDeclare(
		"go-test-queue", // name, leave empty to generate a unique name
		true,            // durable
		false,           // delete when usused
		false,           // exclusive
		false,           // noWait
		nil,             // arguments
	)
	if err != nil {
		panic(err)
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		q.Name, q.Messages, q.Consumers, "go-test-key")
	err = ch.QueueBind(
		q.Name,             // name of the queue
		"go-test-key",      // bindingKey
		"go-test-exchange", // sourceExchange
		false,              // noWait
		nil,                // arguments
	)
	if err != nil {
		panic(err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", "go-amqp-example")

	replies, err = ch.Consume(
		q.Name,            // queue
		"go-amqp-example", // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Waiting...")
	time.Sleep(10 * time.Second)
	fmt.Println("Ready!")
	name := os.Getenv("NAME")
	configureAMQP()
	defer conn.Close()
	defer ch.Close()

	log.Println("Start consuming the Queue...")
	count := 1

	for r := range replies {
		log.Printf("Consuming reply number %d by %s", count, name)
		op := operator{}
		json.Unmarshal(r.Body, &op)
		fmt.Printf("Type: %s, Value: %d\n", op.Type, op.Value)
		count++
		r.Ack(false)
	}
}
