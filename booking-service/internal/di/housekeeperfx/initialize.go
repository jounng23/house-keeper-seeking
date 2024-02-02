package housekeeperfx

import (
	"go.uber.org/fx"

	"booking-svc/pkg/repositories/housekeeper"
)

var Module = fx.Provide(
	provideHouseKeeperRepository,
)

func provideHouseKeeperRepository() housekeeper.Repository {
	return housekeeper.NewRepository()
}
