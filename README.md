# pubsubToolBox 
Google Cloud Platform Pub/Sub toolbox 

* [Prerequisities](#prerequisities)
* [Usage](#usage)
	* [Publisher](#publisher)
	* [Consumer](#consumer)


## Prerequisities 
- GCP project with a Pub/Sub up & running. 
- GCP service account and the credential file _private_key.json_ associated.
- GCP root package "cloud.google.com/go" (```go get -u cloud.google.com/go/pubsub```) 


## Usage
 
```bash
go get github.com/PierreKieffer/pubsubToolBox
```

### Publisher 
To publish messages : 

```go
import (
        "context"
        "github.com/PierreKieffer/pubsubToolBox/client"
        "github.com/PierreKieffer/pubsubToolBox/publisher"
        "log"
)

func main() {

        projectID := "PROJECT_ID"

        // If GOOGLE_APPLICATION_CREDENTIALS is set :
        pubsubClient, err := client.InitPubSubClient(ctx, projectID)

        // Else you need to pass the json key                                                                         
        pubsubClient, err := client.InitPubSubClient(ctx, projectID, "private_key.json")
        if err != nil {
                log.Println(err)
        }   

	p := publisher.Publisher{
		Context : context.Background(),
		TopicID : "TOPIC_ID",
		PubSubClient : pubsubClient
	}

        message := `{"Message" : "Hello world"}`
	attributes := map[string]string{"foo" : "bar"}
        p.Publish(message, attributes)
}


```


### Consumer 
To consume messages : 

A message buffer must be instantiated in order to store messages at the application level.

The consumer.Pull method receives messages from the Pub/Sub broker and adds them to the message buffer.

The message structure is : 
```go
type Message struct {
	Data       string
	Attributes map[string]string
}
```

This allows consumption of the Pub/Sub broker and processing of messages in parallel.

The initialization takes the size of the buffer as a parameter :

```go 
var buffer = consumer.InitBuffer(10)
```

Once a message is consumed, it is dropped from the buffer. 

Consume messages from Pub/Sub broker : 

If the subscriber doesn't exist, it will be created. 

Instanciate first a consumer.Consumer : 
```go
	c := &consumer.Consumer{
		Context:        context.Background(),
		PubSubClient:   pubsubClient,
		SubscriberName: "subscriber-name",
		TopicID:        "topic-id",
		Buffer:         buffer,
	}
```

Start message consumption : 
```go 
	go c.Pull()
```
The message buffer is a classic channel, It can be consumed through a goroutine : 

```go
package main

import (
	"context"
	"log"

	"github.com/PierreKieffer/pubsubToolBox/client"
	"github.com/PierreKieffer/pubsubToolBox/consumer"
)

var exit = make(chan bool)

func main() {

	projectID := "test-project-name"

	pubsubClient, _ := client.InitPubSubClient(ctx, projectID, "private_key.json")

	// Init message buffer to receive pulled messages
	var buffer = consumer.InitBuffer(10)

	c := &consumer.Consumer{
		Context:        context.Background(),
		PubSubClient:   pubsubClient,
		SubscriberName: "subscriber-name",
		TopicID:        "topic-id",
		Buffer:         buffer,
	}

	// Launch local buffer consumer to process messages
	go ProcessBuffer(buffer)

	// Launch the pubsub consumer to pull messages
	go c.Pull()

	<-exit
}

func ProcessBuffer(messageBuffer chan consumer.Message) {
	for {
		// ... Process received messages
		log.Println("Message consumed : ", <-messageBuffer)
	}
}
```
