package engine

import (
	"net/http"

	"github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/router"
)

type HttpEngine interface {
	Init(r router.Router)
	StartServer()
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
