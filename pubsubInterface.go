package pubsubInterface

import (
	"context"
	"google.golang.org/api/option"
	"log"

	"cloud.google.com/go/pubsub"
)

func InitClient(ctx context.Context, projectID, credFile string) *pubsub.Client {

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credFile))
	if err != nil {
		log.Println("ERROR : pubsubInterface.InitClient", err)
	}
	return client
}

func Produce(ctx context.Context, pubsubClient *pubsub.Client, topicID string, event string) {

	t := pubsubClient.Topic(topicID)

	result := t.Publish(ctx, &pubsub.Message{Data: []byte(event)})

	id, err := result.Get(ctx)
	if err != nil {
		log.Println("ERROR : pubsubInterface.Produce")
		log.Println(err)
	}
	log.Println("INFO : pubsubInterface.Produce : Event published : " + id)

}
