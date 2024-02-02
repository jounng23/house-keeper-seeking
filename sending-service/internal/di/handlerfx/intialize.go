package handlerfx

import (
	"go.uber.org/fx"

	"sending-svc/internal/server/handler"
	"sending-svc/pkg/repositories/notification"
)

var Module = fx.Provide(provideHandler)

func provideHandler(notificationRepo notification.Repository) handler.Handler {
	return handler.NewHandler(notificationRepo)
}
