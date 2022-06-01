package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
	pubsub "github.com/jerry-yt-chen/event-sourcing-poc/pkg/event/pubsub"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/mongo"
)

func main() {
	configs.InitConfigs()
	// Mongo
	mongoSvc, _ := mongo.Init(configs.C.Mongo)
	subscriber, _ := pubsub.NewGcpSubscriber(configs.C.Sub)

	// Subscribe will create the subscription. Only messages that are sent after the subscription is created may be received.
	messages, err := subscriber.Subscribe(configs.C.Sub.Topic)
	if err != nil {
		panic(err)
	}

	process(messages, mongoSvc)
}

func process(messages <-chan *message.Message, mongoSvc mongo.Service) {
	for msg := range messages {
		log.Printf("received event: %s, event: %s", msg.UUID, string(msg.Payload))
		receivedTime := time.Now()

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
		go saveRecord(msg, mongoSvc, "SomeSubscriber", receivedTime)
	}
}

func saveRecord(m *message.Message, mongoSvc mongo.Service, subscriberName string, receivedTime time.Time) {
	msg := event.Message{}
	_ = json.Unmarshal(m.Payload, &msg)
	record := event.ReceivedRecord{
		Topic:        msg.Topic,
		TraceID:      msg.TraceID,
		EventID:      msg.EventID,
		Subscriber:   subscriberName,
		ReceivedTime: receivedTime.Unix(),
		CreatedTime:  time.Now().Unix(),
	}
	mongoSvc.Collection(new(event.ReceivedRecord)).InsertOne(context.Background(), record)
}
