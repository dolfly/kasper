package main

import (
	"fmt"
	"log"
	"time"
	"github.com/movio/kasper"
)

type HelloWorldProcessorMultiplex struct{}

func (*HelloWorldProcessorMultiplex) Process(msg kasper.IncomingMessage, sender kasper.Sender, coordinator kasper.Coordinator) {
	key := msg.Key.(string)
	value := msg.Value.(string)
	offset := msg.Offset
	topic := msg.Topic
	partition := msg.Partition
	format := "Got message: key='%s', value='%s' at offset='%s' (topic='%s', partition='%d')\n"
	fmt.Printf(format, key, value, offset, topic, partition)
}

func main() {
	config := kasper.TopicProcessorConfig{
		BrokerList:  []string{"localhost:9092"},
		InputTopics: []string{"hello", "world"},
	}
	makeProcessor := func() kasper.MessageProcessor { return &HelloWorldProcessorMultiplex{} }
	topicProcessor := kasper.NewTopicProcessor(&config, makeProcessor, kasper.NewStringSerde(), kasper.NewStringSerde())
	topicProcessor.Run()
	log.Println("Running!")
	for {
		time.Sleep(2 * time.Second)
		log.Println("...")
	}
}
