package serviceHttp

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

type Http struct {
	dig.In

	Service  contracts.HttpService
	Initials []contracts.Initial   `group:"initials"`
	Routes   []contracts.HttpRoute `group:"routes"`
}

func Serve() error {
	return container.Container.Invoke(func(http Http) error {
		return http.Service.Serve()
	})
}

func ServeE(cmd *cobra.Command, args []string) error {
	return Serve()
}
