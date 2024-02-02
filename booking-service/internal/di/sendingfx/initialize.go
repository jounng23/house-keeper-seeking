package sendingfx

import (
	"net/http"

	"go.uber.org/fx"

	"booking-svc/pkg/config"
	"booking-svc/pkg/xservice/sendingsvc"
)

var Module = fx.Provide(
	provideSendingService,
)

func provideSendingService(client *http.Client) sendingsvc.Service {
	return sendingsvc.NewService(client, config.ServerConfig().SendingSvcUrl)
}
