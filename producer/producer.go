package producer

import (
	"context"
	"google.golang.org/api/option"
	"log"

	"cloud.google.com/go/pubsub"
)

func Publish(projectID, topicID, credFile, message string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credFile))
	if err != nil {
		log.Println("pubsub.NewClient: %v", err)
	}

	t := client.Topic(topicID)

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})

	// Block until the result is returned and a server-generated ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Println("ERROR : publish")
		log.Println(err)
	}
	log.Println("Message published : " + id)

}
