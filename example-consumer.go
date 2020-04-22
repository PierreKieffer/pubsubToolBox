package main

import (
	"github.com/PierreKieffer/pubsubToolBox/consumer"
	"log"
)

var exit = make(chan bool)

func main() {

	// Init message buffer to receive pulled messages
	var buffer = consumer.InitBuffer(10)

	// Launch local buffer consumer to process messages
	go ProcessBuffer(buffer)

	// Launch the pubsub consumer to pull messages
	go consumer.Pull("PROJECT_ID", "SUBSCRIBER_NAME", "TOPIC_NAME", "private_key.json", buffer)

	<-exit

}

func ProcessBuffer(messageBuffer chan string) {
	for {
		// ... Process received messages
		log.Println("Message consumed : ", <-messageBuffer)
	}
}
