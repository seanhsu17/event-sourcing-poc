package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	userModel "github.com/jerry-yt-chen/event-sourcing-poc/internal/domain/user/model"
	api "github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/engine/gin/render"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver"
	"github.com/jerry-yt-chen/event-sourcing-poc/pkg/event"
	pubsub "github.com/jerry-yt-chen/event-sourcing-poc/pkg/event/pubsub"
)

type impl struct {
	pub pubsub.Publisher
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
		{
			Method:      http.MethodPut,
			Path:        "/users",
			Middlewares: nil,
			Handler:     im.update,
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

	payload := userModel.EventPayload{
		Data: user,
	}

	msg := event.Message{
		TraceID:   traceID,
		EventID:   uuid.NewString(),
		Topic:     configs.C.Pub.Topic,
		Source:    "UserService",
		Version:   1,
		Type:      "UserCreated",
		Payload:   payload,
		Timestamp: now,
	}

	im.pub.Send(msg)

	api.ResJSON(c, http.StatusCreated, msg)
}

func (im *impl) update(c *gin.Context) {
	now := time.Now().Unix()
	traceID := uuid.NewString()

	user := userModel.User{
		UserId: uuid.NewString(),
		Age:    20,
		Name:   "Zakk Wylde",
		Gender: "male",
	}

	payload := userModel.EventPayload{
		Data: user,
		Metadata: userModel.Metadata{
			ModifiedFields: []string{
				"Age",
				"Gender",
			},
		},
	}

	msg := event.Message{
		TraceID:   traceID,
		EventID:   uuid.NewString(),
		Topic:     configs.C.Pub.Topic,
		Source:    "UserService",
		Version:   1,
		Type:      "UserUpdated",
		Payload:   payload,
		Timestamp: now,
	}

	im.pub.Send(msg)

	api.ResJSON(c, http.StatusOK, msg)
}
