package jobfx

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"

	"booking-svc/pkg/repositories/job"
	"booking-svc/pkg/xservice/pricesvc"
)

var Module = fx.Provide(
	provideJobStorage,
	provideJobRepository,
)

func provideJobStorage(db *mongo.Database) job.Storage {
	return job.NewStorage(db)
}

func provideJobRepository(dbStorage job.Storage, priceSvc pricesvc.Service) job.Repository {
	return job.NewRepository(dbStorage, priceSvc)
}
