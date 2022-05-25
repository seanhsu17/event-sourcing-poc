package gin

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/engine"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/framework/router"
)

type Engine struct {
	server *http.Server
}

func NewEngine() engine.HttpEngine {
	return &Engine{}
}

func (g *Engine) Init(r router.Router) {
	app := gin.Default()
	r.RegisterAPI(app)

	if configs.C.App.EnableProfile {
		pprof.Register(app, "debug/pprof")
	}

	g.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", configs.C.App.Port),
		Handler:      app,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (g *Engine) StartServer() {
	go func() {
		if err := g.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithField("err", err).Panic("listen and serve failed")
		}
	}()

	// graceful shout down
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	logrus.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := g.server.Shutdown(ctx); err != nil {
		logrus.WithField("err", err).Panic("server forced to shutdown")
	}
}

func (g *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	g.server.Handler.ServeHTTP(w, req)
}
