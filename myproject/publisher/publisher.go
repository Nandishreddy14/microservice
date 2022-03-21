package publisher

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Publish() {

	amqpstr := "amqp://guest:guest@" + os.Getenv("amqpURL") + ":5672"
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(amqpstr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"amqpq", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to create a queue")

	fmt.Println(q)

	randint := rand.Int()

	body := strconv.Itoa(randint)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish to RabbitMQ")
	
	log.Printf("Published message %s",body)
}
