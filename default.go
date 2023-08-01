package serviceHttp

import (
	httpFiber "github.com/atom-providers/http-fiber"
	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		httpFiber.DefaultProvider(),
	}, providers...)
}
