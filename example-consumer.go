package main

import (
	"context"
	"github.com/PierreKieffer/pubsubToolBox/client"
	"github.com/PierreKieffer/pubsubToolBox/consumer"
	"log"
)

var exit = make(chan bool)

func main() {

	ctx := context.Background()

	projectID := "platform-hubstairs-test"
	topicID := "airflow-trigger-notification"
	subscriberName := "airflow-trigger-subscriber"

	pubsubClient, _ := client.InitPubSubClient(ctx, projectID, "private_key.json")

	// Init message buffer to receive pulled messages
	var buffer = consumer.InitBuffer(10)

	// Launch local buffer consumer to process messages
	go ProcessBuffer(buffer)

	// Launch the pubsub consumer to pull messages
	go consumer.Pull(ctx, pubsubClient, subscriberName, topicID, buffer)

	<-exit

}

func ProcessBuffer(messageBuffer chan string) {
	for {
		// ... Process received messages
		log.Println("Message consumed : ", <-messageBuffer)
	}
}
