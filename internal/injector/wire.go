//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/event-sourcing-poc/internal/injector/api"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/injector/lib"
	"github.com/jerry-yt-chen/event-sourcing-poc/internal/injector/persistence"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// api
		api.InitGinEngine,
		api.RouteSet,
		api.ReceiverSet,
		api.ProvideReceiverList,

		// persistence
		persistence.InitMongo,

		// lib
		lib.InitEventPublisher,

		//Injector
		InjectorSet,
	)
	return new(Injector), nil, nil
}
