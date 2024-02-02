package routerfx

import (
	"go.uber.org/fx"

	"pricing-svc/internal/server/handler"
	"pricing-svc/internal/server/router"
)

var Module = fx.Provide(provideRouter)

func provideRouter(handler handler.Handler) router.Router {
	return router.NewRouter(handler)
}
