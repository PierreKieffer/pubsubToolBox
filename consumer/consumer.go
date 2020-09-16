package consumer

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func InitBuffer(bufferSize int) chan string {
	var buffer = make(chan string, bufferSize)
	return buffer
}

func Pull(ctx context.Context, pubsubClient *pubsub.Client, subName, topicID string, buffer chan string) error {

	topic := pubsubClient.Topic(topicID)

	// Create topic subscription if it does not yet exist.
	sub := pubsubClient.Subscription(subName)
	exists, err := sub.Exists(ctx)
	if err != nil {
		log.Println("ERROR : consumer.Pull : " + err.Error())
		return err
	}
	if !exists {
		if _, err = pubsubClient.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{Topic: topic}); err != nil {
			log.Println("ERROR : consumer.Pull : " + err.Error())
			return err
		}
	}

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		buffer <- string(msg.Data)
		log.Println("INFO : consumer.Pull : Message received and pushed to buffer")
		msg.Ack()
	})
	if err != nil {
		log.Println("ERROR : consumer.Pull : " + err.Error())
		return err
	}
	return nil
}
