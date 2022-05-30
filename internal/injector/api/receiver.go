package api

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver/user"
)

var ReceiverSet = wire.NewSet(
	user.ProvideReceiver,
)

func ProvideReceiverList(u user.Receiver) []receiver.Receiver {
	return []receiver.Receiver{u}
}
