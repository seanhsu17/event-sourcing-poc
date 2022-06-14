package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	userModel "github.com/jerry-yt-chen/event-sourcing-poc/internal/domain/user/model"
	api "github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/engine/gin/render"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
	pubsub "github.com/jerry-yt-chen/event-sourcing-poc/pkg/event/pubsub"
)

type impl struct {
	pub pubsub.Publisher
}

type Resp struct {
	Payload  interface{}       `json:"payload"`
	TraceID  string            `json:"traceID"`
	Metadata map[string]string `json:"metadata"`
}

func ProvideReceiver(publisher pubsub.Publisher) Receiver {
	return &impl{
		pub: publisher,
	}
}

func (im *impl) GetRouteInfos() []receiver.RouteInfo {
	return []receiver.RouteInfo{
		{
			Method:      http.MethodPost,
			Path:        "/users",
			Middlewares: nil,
			Handler:     im.create,
		},
	}
}

func (im *impl) create(c *gin.Context) {
	traceID := uuid.NewString()

	user := userModel.User{
		UserId: "ef0df7a5-4c94-49df-926d-ac22380f8f91",
		Age:    18,
		Name:   "Zakk Wylde",
		Gender: "male",
	}

	metadata := event.Metadata{
		event.TraceAttribute: traceID,
		"eventID":            uuid.NewString(),
	}
	im.pub.Send(user, metadata)

	api.ResJSON(c, http.StatusCreated, Resp{
		Payload:  user,
		TraceID:  traceID,
		Metadata: metadata,
	})
}
