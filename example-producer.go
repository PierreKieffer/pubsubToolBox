package main

import (
	"context"
	"github.com/PierreKieffer/pubsubToolBox/client"
	"github.com/PierreKieffer/pubsubToolBox/producer"
	"log"
)

func main() {
	ctx := context.Background()

	projectID := "ai-datalake"
	topicID := "testTopic"

	pubsubClient, err := client.InitPubSubClient(ctx, projectID, "private_key.json")
	if err != nil {
		log.Println(err)
	}

	message := `{"Message" : "Hello world"}`
	producer.Publish(ctx, pubsubClient, topicID, message)

}
