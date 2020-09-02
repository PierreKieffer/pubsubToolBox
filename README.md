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

### Producer 
To publish messages : 

```go
import (
        "github.com/PierreKieffer/pubsubToolBox/producer"
)
```

```go
        message := `{"Message" : "Hello world"}`
        producer.Publish("PROJECT_ID", "TOPIC_NAME", "private_key.json", message)

```

### Consumer 
To consume messages : 

```go
import (
        "github.com/PierreKieffer/pubsubToolBox/consumer"
)
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
go consumer.Pull("PROJECT_ID", "SUBSCRIBER_NAME", "TOPIC_NAME", "private_key.json", buffer)

```
The message buffer is a classic channel, It can be consumed through a goroutine : 

```go
package main

import (
        "github.com/PierreKieffer/pubsubToolBox/consumer"
        "log"
)

var exit = make(chan bool)

func main() {

        // Init message buffer to receive pulled messages 
        var buffer = consumer.InitBuffer(10)

        // Launch the message buffer consumer to process messages 
        go ProcessBuffer(buffer)

        // Launch the pubsub consumer to pull messages 
	go consumer.Pull("PROJECT_ID", "SUBSCRIBER_NAME", "TOPIC_NAME", "private_key.json", buffer)

        <-exit

}

func ProcessBuffer(messageBuffer chan string) {
        for {
                // ... Process received messages
                log.Println("Message consumed : ", <-messageBuffer)
        }
}
```






