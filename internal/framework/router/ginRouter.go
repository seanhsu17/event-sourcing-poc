package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	api "github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/engine/gin/render"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/receiver"
)

var GinRouterSet = wire.NewSet(ProvideRouteV1, wire.Bind(new(Router), new(*GinRouter)))

type GinRouter struct {
	routesInfo []receiver.RouteInfo
}

func (r *GinRouter) RegisterAPI(engine *gin.Engine) {
	defaultRoutInfo := receiver.RouteInfo{
		Method: http.MethodGet,
		Path:   "",
		Handler: func(context *gin.Context) {
			api.ResSuccess(context, http.StatusOK, struct{}{})
		},
	}
	engine.Handle(defaultRoutInfo.Method, defaultRoutInfo.Path, defaultRoutInfo.GetFlow()...)
	for _, routeInfo := range r.routesInfo {
		engine.Handle(routeInfo.Method, r.generatePath(routeInfo), routeInfo.GetFlow()...)
	}
}

func (r *GinRouter) GetRoutesInfo() []receiver.RouteInfo {
	return r.routesInfo
}

func ProvideRouteV1(receivers ...receiver.Receiver) *GinRouter {
	return &GinRouter{
		routesInfo: extractRouteInfo(receivers...),
	}
}

func (r *GinRouter) PrefixPath() string {
	prefixes := []string{
		"v1",
		configs.C.App.Name,
	}
	return strings.Join(prefixes, "/")
}

func (r *GinRouter) generatePath(routeInfo receiver.RouteInfo) string {
	paths := []string{configs.C.App.BaseRoute, r.PrefixPath(), routeInfo.Path}
	return strings.Join(paths, "/")
}

func extractRouteInfo(receivers ...receiver.Receiver) []receiver.RouteInfo {
	var routeInfos []receiver.RouteInfo

	for _, receiver := range receivers {
		routeInfos = append(routeInfos, receiver.GetRouteInfos()...)
	}

	return routeInfos
}
