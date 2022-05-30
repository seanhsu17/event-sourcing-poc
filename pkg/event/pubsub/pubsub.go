package event

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
)

type Publisher interface {
	Send(traceID string, msg event.Message) error
}

type Subscriber interface {
	Subscribe(topic string) (<-chan *message.Message, error)
}
