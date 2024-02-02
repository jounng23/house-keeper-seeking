package pricefx

import (
	"net/http"

	"go.uber.org/fx"

	"booking-svc/pkg/config"
	"booking-svc/pkg/xservice/pricesvc"
)

var Module = fx.Provide(
	providePriceService,
)

func providePriceService(client *http.Client) pricesvc.Service {
	return pricesvc.NewService(client, config.ServerConfig().PricingSvcUrl)
}
