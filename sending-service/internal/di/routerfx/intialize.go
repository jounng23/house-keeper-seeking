package routerfx

import (
	"go.uber.org/fx"

	"sending-svc/internal/server/handler"
	"sending-svc/internal/server/router"
)

var Module = fx.Provide(provideRouter)

func provideRouter(handler handler.Handler) router.Router {
	return router.NewRouter(handler)
}
