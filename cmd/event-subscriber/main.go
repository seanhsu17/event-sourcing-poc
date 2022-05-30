package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/mongo"
)

func main() {
	configs.InitConfigs()
	logger := watermill.NewStdLogger(false, false)
	subscriber, err := googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			// custom function to generate Subscription Name,
			// there are also predefined TopicSubscriptionName and TopicSubscriptionNameWithSuffix available.
			GenerateSubscriptionName: func(topic string) string {
				return topic + "-sub"
			},
			ProjectID: configs.C.Sub.ProjectID,
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	// Subscribe will create the subscription. Only messages that are sent after the subscription is created may be received.
	messages, err := subscriber.Subscribe(context.Background(), configs.C.Sub.Topic)
	if err != nil {
		panic(err)
	}

	// Mongo
	mongoSvc, _ := mongo.Init(configs.C.Mongo)
	process(messages, mongoSvc)
}

func process(messages <-chan *message.Message, mongoSvc mongo.Service) {
	for msg := range messages {
		log.Printf("received event: %s, event: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
		go saveRecord(msg, mongoSvc)
	}
}

func saveRecord(msg *message.Message, mongoSvc mongo.Service) {
	e := event.Event{}
	_ = json.Unmarshal(msg.Payload, &e)
	payload, _ := json.Marshal(e.Payload)
	record := event.Record{
		TraceID:      e.TraceID,
		EventType:    e.Type,
		Version:      e.Version,
		Payload:      string(payload),
		ReceivedTime: time.Now().Unix(),
		CreatedTime:  time.Now().Unix(),
	}
	mongoSvc.Collection(new(event.ReceivedRecord)).InsertOne(context.Background(), record)
}
