package event

import (
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	api "github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/engine/gin/render"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver"
)

type impl struct {
}

func ProvideReceiver() Receiver {
	im := &impl{}
	return im
}

func (im *impl) GetRouteInfos() []receiver.RouteInfo {
	return []receiver.RouteInfo{
		{
			Method:      http.MethodPost,
			Path:        "/events",
			Middlewares: nil,
			Handler:     im.sendEvent,
		},
	}
}

func (im *impl) sendEvent(c *gin.Context) {
	logger := watermill.NewStdLogger(true, false)
	publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: configs.C.Pub.Project,
	}, logger)
	if err != nil {
		panic(err)
	}
	msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
	publishMessages(publisher, msg)
	api.ResSuccess(c, http.StatusOK, publisher)
}

func publishMessages(publisher message.Publisher, msg *message.Message) {
	if err := publisher.Publish(configs.C.Pub.Topic, msg); err != nil {
		panic(err)
	}
}
