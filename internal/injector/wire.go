//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/event-sourcing-poc/internal/injector/api"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// api
		api.InitGinEngine,
		api.RouteSet,
		api.ReceiverSet,
		api.ProvideReceiverList,

		//Injector
		InjectorSet,
	)
	return new(Injector), nil, nil
}
