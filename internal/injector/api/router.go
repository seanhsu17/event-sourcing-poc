package api

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/router"
)

var RouteSet = wire.NewSet(router.GinRouterSet)
