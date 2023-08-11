package serviceHttp

import (
	"github.com/atom-providers/app"
	httpFiber "github.com/atom-providers/http-fiber"
	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		app.DefaultProvider(),
		log.DefaultProvider(),
		httpFiber.DefaultProvider(),
	}, providers...)
}
