package api

import (
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/engine"
	ginEngine "github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/engine/gin"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/router"
)

func InitGinEngine(r router.Router) engine.HttpEngine {
	engine := ginEngine.NewEngine()
	engine.Init(r)
	return engine
}
