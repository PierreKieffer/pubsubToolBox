package consumer

import (
	"context"
	"google.golang.org/api/option"
	"log"

	"cloud.google.com/go/pubsub"
)

func InitBuffer(bufferSize int) chan string {
	var buffer = make(chan string, bufferSize)
	return buffer
}

func Pull(projectID, subName, topicID, credFile string, buffer chan string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credFile))
	if err != nil {
		log.Println("pubsub.NewClient: %v", err)
	}

	topic := client.Topic(topicID)

	// Create topic subscription if it does not yet exist.
	sub := client.Subscription(subName)
	exists, err := sub.Exists(ctx)
	if err != nil {
		log.Println("Error checking for subscription: %v", err)
	}
	if !exists {
		if _, err = client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{Topic: topic}); err != nil {
			log.Println("Failed to create subscription: %v", err)
		}
	}

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		buffer <- string(msg.Data)
		log.Println("INFO : pubsub.Pull : Message received and pushed to buffer")
		msg.Ack()
	})
	if err != nil {
		log.Println("ERROR : pull : %v", err)
	}
}
