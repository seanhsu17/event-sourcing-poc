package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
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

func (d *PublisherDecorator) Send(traceID string, event Event) error {
	if err := d.pub.Send(traceID, event); err != nil {
		return err
	}
	go d.saveRecord(event)
	return nil
}

func (d *PublisherDecorator) saveRecord(event Event) {
	payload, _ := json.Marshal(event.Payload)
	record := Record{
		TraceID:       event.TraceID,
		EventType:     event.Type,
		Version:       event.Version,
		Payload:       string(payload),
		PublishedTime: event.Timestamp,
		CreatedTime:   time.Now().Unix(),
	}
	coll := d.mongoSvc.Collection(new(PublishedRecord))
	if result, err := coll.InsertOne(context.Background(), record); err != nil {
		logrus.WithField("err", err).Error("Insert record failed")
	} else {
		logrus.Info("Success insert: ", result)
	}
}
