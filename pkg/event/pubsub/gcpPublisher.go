package event

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
)

type ProjectID string

type GcpPublisher struct {
	pub message.Publisher
}

func NewGcpPublisher(projectID string) (Publisher, error) {
	logger := watermill.NewStdLogger(true, false)
	if publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: projectID,
	}, logger); err != nil {
		return nil, err
	} else {
		return &GcpPublisher{
			pub: publisher,
		}, nil
	}
}

func (p *GcpPublisher) Send(msg event.Message) error {
	content, _ := json.Marshal(msg)
	m := message.NewMessage(msg.EventID, content)
	if err := p.pub.Publish(msg.Topic, m); err != nil {
		return err
	}
	return nil
}
