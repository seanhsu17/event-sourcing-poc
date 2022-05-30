package event

import "github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"

type Publisher interface {
	Send(traceID string, msg event.Message) error
}
