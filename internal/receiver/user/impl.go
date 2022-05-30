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
)

type impl struct {
	pub event.Publisher
}

func ProvideReceiver(publisher event.Publisher) Receiver {
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

	event := event.Event{}
	event.TraceID = traceID
	event.Version = 1
	event.Type = "UserCreated"
	event.Payload = payload
	event.Timestamp = now

	im.pub.Send(traceID, event)

	api.ResJSON(c, http.StatusCreated, event)
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

	event := event.Event{}
	event.TraceID = traceID
	event.Version = 1
	event.Type = "UserUpdated"
	event.Payload = payload
	event.Timestamp = now

	im.pub.Send(traceID, event)

	api.ResJSON(c, http.StatusCreated, event)
}
