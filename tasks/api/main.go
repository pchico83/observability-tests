package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func startHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Only POST is allowed\n"))
		return
	}
	err := start()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("KO\n"))
		return
	}
	w.Write([]byte("OK\n"))
}

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func start() error {
	client := getRedisClient()
	log.Printf("Setting value to 1...")
	err := client.Set("value", 1, 0).Err()
	if err != nil {
		return err
	}

	val, err := client.Get("value").Result()
	if err != nil {
		return err
	}
	fmt.Println("Checking value:", val)

	log.Printf("publishing message...")
	op := operator{
		Type:  "suma",
		Value: 3,
	}
	payload, err := json.Marshal(op)

	err = ch.Publish(
		"go-test-exchange", // exchange
		"go-test-key",      // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})
	if err != nil {
		panic(err)
	}
	log.Printf("message published!")

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
}

func main() {
	fmt.Println("Waiting...")
	time.Sleep(10 * time.Second)
	fmt.Println("Ready!")
	configureAMQP()
	defer conn.Close()
	defer ch.Close()
	http.HandleFunc("/start", startHTTP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
