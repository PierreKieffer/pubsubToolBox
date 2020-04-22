package main

import (
	"github.com/PierreKieffer/pubsubToolBox/producer"
)

func main() {

	message := `{"Message" : "Hello world"}`
	producer.Publish("PROJECT_NAME", "TOPIC_NAME", "private_key.json", message)

}
