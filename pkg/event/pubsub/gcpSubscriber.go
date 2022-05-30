package event

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
)

type GcpSubscriber struct {
	sub message.Subscriber
}

func NewGcpSubscriber(config configs.PubSubConfig) (Subscriber, error) {
	logger := watermill.NewStdLogger(true, false)
	subscriber, err := googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{
			GenerateSubscriptionName: func(topic string) string {
				return topic + "-sub"
			},
			ProjectID: config.ProjectID,
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	return &GcpSubscriber{
		sub: subscriber,
	}, nil
}

func (s *GcpSubscriber) Subscribe(topic string) (<-chan *message.Message, error) {
	return s.sub.Subscribe(context.Background(), topic)
}
