package main

import (
	"context"
	"github.com/PierreKieffer/pubsubToolBox/client"
	"github.com/PierreKieffer/pubsubToolBox/publisher"
	"log"
)

func main() {
	ctx := context.Background()

	projectID := "PROJECT_ID"
	topicID := "TOPIC_ID"

	pubsubClient, err := client.InitPubSubClient(ctx, projectID, "private_key.json")
	if err != nil {
		log.Println(err)
	}

	message := `{"Message" : "Hello world"}`
	publisher.Publish(ctx, pubsubClient, topicID, message)

}
