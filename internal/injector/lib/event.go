package lib

import (
	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	pubsub "github.com/jerry-yt-chen/event-sourcing-poc/pkg/event/pubsub"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/mongo"
)

func InitEventPublisher(mongoSvc mongo.Service) pubsub.Publisher {
	return pubsub.NewPublisherDecorator(configs.C.Pub.ProjectID, mongoSvc)
}
