package user

import (
	"net/http"
	"time"

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
	Payload interface{}         `json:"payload"`
	Options event.PublishOption `json:"options"`
	TraceID string              `json:"traceID"`
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
	now := time.Now().Unix()
	traceID := uuid.NewString()

	user := userModel.User{
		UserId: uuid.NewString(),
		Age:    18,
		Name:   "Zakk Wylde",
		Gender: "male",
	}

	options := event.PublishOption{
		Key:       uuid.NewString(),
		EventType: "createUser",
		Source:    "UserService",
		Timestamp: now,
	}
	metadata := event.Metadata{
		event.TraceAttribute:   traceID,
		event.OptionsAttribute: options.ToString(),
	}
	im.pub.Send(user, metadata)

	api.ResJSON(c, http.StatusCreated, Resp{
		Payload: user,
		TraceID: traceID,
		Options: options,
	})
}
