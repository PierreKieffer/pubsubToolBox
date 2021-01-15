package main

func main() {
	ctx := context.Background()

	projectID := "PROJECT_ID"
	topicID := "TOPIC_ID"

	pubsubClient, err := client.InitPubSubClient(ctx, projectID, "private_key.json")

	// If GOOGLE_APPLICATION_CREDENTIALS is set :
	pubsubClient, err := client.InitPubSubClient(ctx, projectID)

	if err != nil {
		log.Println(err)
	}

	message := `{"Message" : "Hello world"}`
	publisher.Publish(ctx, pubsubClient, topicID, message)

}
