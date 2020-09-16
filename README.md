# pubsubToolBox 
Google Cloud Platform Pub/Sub toolbox 

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

```


### Consumer 
To consume messages : 

```go
import (
        "context"
        "github.com/PierreKieffer/pubsubToolBox/client"
        "github.com/PierreKieffer/pubsubToolBox/consumer"
        "log"
)

ctx := context.Background()

        projectID := "PROJECT_ID"
        topicID := "TOPIC_ID"
        subscriberName := "SUBSCRIBER_NAME"

        pubsubClient, err:= client.InitPubSubClient(ctx, projectID, "private_key.json")
        if err != nil {
                log.Println(err)
        }
```
A message buffer must be instantiated in order to store messages at the application level.

The consumer.Pull method receives messages from the Pub/Sub broker and adds them to the message buffer.

This allows consumption of the Pub/Sub broker and processing of messages in parallel.

The initialization takes the size of the buffer as a parameter :

```go 
var buffer = consumer.InitBuffer(10)
```

Once a message is consumed, it is dropped from the buffer. 

Consume messages from Pub/Sub broker : 

If the subscriber doesn't exist, it will be created. 

```go 
go consumer.Pull(ctx, pubsubClient, subscriberName, topicID, buffer)

```
The message buffer is a classic channel, It can be consumed through a goroutine : 

```go
package main

import (
        "context"
        "github.com/PierreKieffer/pubsubToolBox/client"
        "github.com/PierreKieffer/pubsubToolBox/consumer"
        "log"
)

var exit = make(chan bool)

func main() {

        ctx := context.Background()

        projectID := "PROJECT_ID"
        topicID := "TOPIC_ID"
        subscriberName := "SUBSCRIBER_NAME"

        pubsubClient, _ := client.InitPubSubClient(ctx, projectID, "private_key.json")

        // Init message buffer to receive pulled messages
        var buffer = consumer.InitBuffer(10)

        // Launch local buffer consumer to process messages
        go ProcessBuffer(buffer)

        // Launch the pubsub consumer to pull messages
        go consumer.Pull(ctx, pubsubClient, subscriberName, topicID, buffer)

        <-exit

}

func ProcessBuffer(messageBuffer chan string) {
        for {
                // ... Process received messages
                log.Println("Message consumed : ", <-messageBuffer)
        }
}

```






