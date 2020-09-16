package client

import (
	"cloud.google.com/go/pubsub"
	"context"
	"google.golang.org/api/option"
	"log"
)

func InitPubSubClient(ctx context.Context, projectID, credFile string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credFile))
	if err != nil {
		log.Println("ERROR : client.InitPubSubClient : " + err.Error())
		return nil, err
	}
	return client, nil
}
