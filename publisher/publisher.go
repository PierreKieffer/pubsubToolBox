package publisher

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func Publish(ctx context.Context, pubsubClient *pubsub.Client, topicID, message string) error {

	t := pubsubClient.Topic(topicID)

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})

	id, err := result.Get(ctx)
	if err != nil {
		log.Println("ERROR : publisher.Publish : " + err.Error())
		return err
	}

	log.Println("INFO : publisher.Publish : Message published : " + id)
	return nil
}
