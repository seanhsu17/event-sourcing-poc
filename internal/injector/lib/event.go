package lib

import (
	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/mongo"
)

func InitEventPublisher(mongoSvc mongo.Service) event.Publisher {
	return event.NewPublisherDecorator(configs.C.Pub, mongoSvc)
}
