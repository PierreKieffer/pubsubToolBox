package publisher

import (
        "context"
        "log"

        "cloud.google.com/go/pubsub"
)

type Publisher struct {
        Context      context.Context
        PubSubClient *pubsub.Client
        TopicID      string
}

func (p *Publisher) Publish(message string, attributes map[string]string) error {

        t := p.PubSubClient.Topic(p.TopicID)

        result := t.Publish(p.Context, &pubsub.Message{                                                                                
                Data:       []byte(message),
                Attributes: attributes,
        })  

        id, err := result.Get(p.Context)
        if err != nil {
                log.Println("ERROR : publisher.Publish : " + err.Error())
                return err 
        }   

        log.Println("INFO : publisher.Publish : Message published : " + id) 
        return nil 
}
