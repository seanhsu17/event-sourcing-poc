package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/mongo"
)

type PublisherDecorator struct {
	pub      Publisher
	mongoSvc mongo.Service
}

func NewPublisherDecorator(pubsubConfig configs.PubSubConfig, mongoSvc mongo.Service) Publisher {
	publisher, err := NewGcpPublisher(pubsubConfig)
	if err != nil {
		panic(err)
	}
	return &PublisherDecorator{
		pub:      publisher,
		mongoSvc: mongoSvc,
	}
}

func (d *PublisherDecorator) Send(msg event.Message) error {
	if err := d.pub.Send(msg); err != nil {
		return err
	}
	go d.saveRecord(msg)
	return nil
}

func (d *PublisherDecorator) saveRecord(msg event.Message) {
	payload, _ := json.Marshal(msg.Payload)
	record := event.Record{
		TraceID:       msg.TraceID,
		EventID:       msg.EventID,
		EventType:     msg.Type,
		Publisher:     msg.Source,
		Version:       msg.Version,
		Payload:       string(payload),
		PublishedTime: msg.Timestamp,
		CreatedTime:   time.Now().Unix(),
	}
	coll := d.mongoSvc.Collection(new(event.PublishedRecord))
	if result, err := coll.InsertOne(context.Background(), record); err != nil {
		logrus.WithField("err", err).Error("Insert record failed")
	} else {
		logrus.Info("Success insert: ", result)
	}
}
