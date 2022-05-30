package user

import "github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver"

type Receiver interface {
	receiver.Receiver
}
