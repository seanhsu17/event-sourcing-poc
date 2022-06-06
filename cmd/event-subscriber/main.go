package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sirupsen/logrus"

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
		logrus.Printf("received event: %s, event: %s\n", msg.UUID, string(msg.Payload))
		receivedTime := time.Now()
		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
		go saveRecord(msg, mongoSvc, "SomeSubscriber", receivedTime)
	}
}

func saveRecord(m *message.Message, mongoSvc mongo.Service, subscriberName string, receivedTime time.Time) {
	options := event.PublishOption{}
	_ = json.Unmarshal([]byte(m.Metadata.Get("Options")), &options)
	traceID := m.Metadata.Get("Cloud-Trace-Context")
	record := event.ReceivedRecord{
		Topic:        configs.C.Sub.Topic,
		TraceID:      traceID,
		EventID:      options.Key,
		Subscriber:   subscriberName,
		ReceivedTime: receivedTime.Unix(),
		CreatedTime:  time.Now().Unix(),
	}
	mongoSvc.Collection(new(event.ReceivedRecord)).InsertOne(context.Background(), record)
}
