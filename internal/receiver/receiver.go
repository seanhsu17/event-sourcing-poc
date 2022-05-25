package receiver

import (
	"github.com/gin-gonic/gin"
)

type Receiver interface {
	GetRouteInfos() []RouteInfo
}

type RouteInfo struct {
	Method      string
	Path        string
	Middlewares []gin.HandlerFunc
	Handler     gin.HandlerFunc
}

func (r *RouteInfo) GetFlow() []gin.HandlerFunc {
	var flow []gin.HandlerFunc
	// Append specific middleware here
	flow = append(flow, r.Middlewares...)

	// Append handler
	flow = append(flow, r.Handler)
	return flow
}
