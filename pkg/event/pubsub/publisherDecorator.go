package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/mongo"
)

type PublisherDecorator struct {
	pub      Publisher
	mongoSvc mongo.Service
	topic    string
}

func NewPublisherDecorator(projectID, topic string, mongoSvc mongo.Service) Publisher {
	publisher, err := NewGcpPublisher(projectID, topic)
	if err != nil {
		panic(err)
	}
	return &PublisherDecorator{
		pub:      publisher,
		mongoSvc: mongoSvc,
		topic:    topic,
	}
}

func (d *PublisherDecorator) Send(payload interface{}, metadata event.Metadata) error {
	if err := d.pub.Send(payload, metadata); err != nil {
		return err
	}
	go d.saveRecord(payload, metadata)
	return nil
}

func (d *PublisherDecorator) saveRecord(payload interface{}, metadata event.Metadata) {
	p, _ := json.Marshal(payload)
	record := event.PublishRecord{
		TraceID:     metadata[event.TraceAttribute],
		Topic:       d.topic,
		EventID:     metadata["eventID"],
		Payload:     string(p),
		PublishTime: time.Now().Unix(),
		CreatedTime: time.Now().Unix(),
	}
	coll := d.mongoSvc.Collection(record)
	if result, err := coll.InsertOne(context.Background(), record); err != nil {
		logrus.WithField("err", err).Error("Insert record failed")
	} else {
		logrus.Info("Success insert: ", result)
	}
}
