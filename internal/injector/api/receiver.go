package api

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver/event"
)

var ReceiverSet = wire.NewSet(
	event.ProvideReceiver,
)

func ProvideReceiverList(a event.Receiver) []receiver.Receiver {
	return []receiver.Receiver{a}
}
