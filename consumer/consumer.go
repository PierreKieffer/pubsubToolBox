package consumer                                                                                                                                                                                                                                                                

import (
        "context"
        "log"
        "os"

        "cloud.google.com/go/pubsub"
)

type Message struct {
        Data       string
        Attributes map[string]string
}

type Consumer struct {
        Context        context.Context
        PubSubClient   *pubsub.Client
        SubscriberName string
        TopicID        string
        Buffer         chan Message
}

func InitBuffer(bufferSize int) chan Message {
        var buffer = make(chan Message, bufferSize)
        return buffer
}

func (c *Consumer) Pull() {

        topic := c.PubSubClient.Topic(c.TopicID)

        // Create topic subscription if it does not yet exist.
        sub := c.PubSubClient.Subscription(c.SubscriberName)
        exists, err := sub.Exists(c.Context)
        if err != nil {
                log.Println("ERROR : consumer.Pull : " + err.Error())
                os.Exit(1)
        }   
        if !exists {
                if _, err = c.PubSubClient.CreateSubscription(c.Context, c.SubscriberName, pubsub.SubscriptionConfig{Topic: topic}); err != nil {
                        log.Println("ERROR : consumer.Pull : " + err.Error())
                        os.Exit(1)
                }   
        }   

        err = sub.Receive(c.Context, func(ctx context.Context, msg *pubsub.Message) {
                message := Message{
                        Data:       string(msg.Data),
                        Attributes: msg.Attributes,
                }   

                c.Buffer <- message
                log.Println("INFO : consumer.Pull : Message received and pushed to buffer")
                msg.Ack()
        })  
        if err != nil {
                log.Println("ERROR : consumer.Pull : " + err.Error())
                os.Exit(1)
        }   
}
