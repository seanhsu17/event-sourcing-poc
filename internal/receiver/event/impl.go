package event

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
	api.ResSuccess(c, http.StatusOK, struct{}{})
}
