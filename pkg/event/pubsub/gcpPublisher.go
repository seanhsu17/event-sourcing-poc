package event

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
)

type ProjectID string

type GcpPublisher struct {
	pub    message.Publisher
	config configs.PubSubConfig
}

func NewGcpPublisher(config configs.PubSubConfig) (Publisher, error) {
	logger := watermill.NewStdLogger(true, false)
	if publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: config.ProjectID,
	}, logger); err != nil {
		return nil, err
	} else {
		return &GcpPublisher{
			pub:    publisher,
			config: config,
		}, nil
	}
}

func (p *GcpPublisher) Send(traceID string, msg event.Message) error {
	p.pub.Close()
	content, _ := json.Marshal(msg)
	m := message.NewMessage(traceID, content)
	if err := p.pub.Publish(p.config.Topic, m); err != nil {
		return err
	}
	return nil
}
