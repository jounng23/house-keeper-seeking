package notificationfx

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"

	"sending-svc/pkg/repositories/notification"
)

var Module = fx.Provide(
	providePriceStorage,
	providePriceRepository,
)

func providePriceStorage(db *mongo.Database) notification.Storage {
	return notification.NewStorage(db)
}

func providePriceRepository(dbStorage notification.Storage) notification.Repository {
	return notification.NewRepository(dbStorage)
}
