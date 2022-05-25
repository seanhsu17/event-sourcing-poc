package main

import (
	"github.com/sirupsen/logrus"

	"github.com/jerry-yt-chen/event-sourcing-poc/configs"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/injector"
)

func main() {
	configs.InitConfigs()

	injector, cleanup, err := injector.BuildInjector()
	if err != nil {
		logrus.WithField("err", err).Panic("Fail to build injector")
	}
	defer cleanup()

	// start http server
	injector.HttpEngine.StartServer()
}
