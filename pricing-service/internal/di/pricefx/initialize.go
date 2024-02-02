package pricefx

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"

	"pricing-svc/pkg/repositories/price"
)

var Module = fx.Provide(
	providePriceStorage,
	providePriceRepository,
)

func providePriceStorage(db *mongo.Database) price.Storage {
	return price.NewStorage(db)
}

func providePriceRepository(dbStorage price.Storage) price.Repository {
	return price.NewRepository(dbStorage)
}
