package routerfx

import (
	"go.uber.org/fx"

	"booking-svc/internal/server/handler"
	"booking-svc/internal/server/router"
)

var Module = fx.Provide(provideRouter)

func provideRouter(handler handler.Handler) router.Router {
	return router.NewRouter(handler)
}
