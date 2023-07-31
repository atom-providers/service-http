package serviceHttp

import (
	httpFiber "github.com/atom-providers/http-fiber"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"go.uber.org/dig"
)

type Http struct {
	dig.In

	Service  contracts.HttpService
	Initials []contracts.Initial `group:"initials"`
	Routes   []contracts.Route        `group:"routes"`
}

func Serve() error {
	return container.Container.Invoke(func(http Http) error {
		return http.Service.Serve()
	})
}
