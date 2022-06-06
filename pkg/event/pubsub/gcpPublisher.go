package event

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
)

type ProjectID string

type GcpPublisher struct {
	pub   message.Publisher
	topic string
}

func NewGcpPublisher(projectID, topic string) (Publisher, error) {
	logger := watermill.NewStdLogger(true, true)
	if publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: projectID,
	}, logger); err != nil {
		return nil, err
	} else {
		return &GcpPublisher{
			pub:   publisher,
			topic: topic,
		}, nil
	}
}

func (p *GcpPublisher) Send(payload interface{}, metadata event.Metadata) error {
	content, _ := json.Marshal(payload)
	// TODO Separate traceID and spanID from TraceAttribute
	m := message.NewMessage(metadata[event.TraceAttribute], content)
	m.Metadata = message.Metadata(metadata)
	fmt.Println("send: ", m.Metadata)
	if err := p.pub.Publish(p.topic, m); err != nil {
		return err
	}
	return nil
}
