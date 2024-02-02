package handlerfx

import (
	"go.uber.org/fx"

	"pricing-svc/internal/server/handler"
	"pricing-svc/pkg/repositories/price"
)

var Module = fx.Provide(provideHandler)

func provideHandler(priceRepo price.Repository) handler.Handler {
	return handler.NewHandler(priceRepo)
}
