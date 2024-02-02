package handlerfx

import (
	"go.uber.org/fx"

	"booking-svc/internal/server/handler"
	"booking-svc/pkg/repositories/housekeeper"
	"booking-svc/pkg/repositories/job"
	"booking-svc/pkg/xservice/pricesvc"
	"booking-svc/pkg/xservice/sendingsvc"
)

var Module = fx.Provide(provideHandler)

func provideHandler(
	jobRepo job.Repository,
	housekeeperRepo housekeeper.Repository,
	priceSvc pricesvc.Service,
	sendingSvc sendingsvc.Service,
) handler.Handler {
	return handler.NewHandler(jobRepo, housekeeperRepo, priceSvc, sendingSvc)
}
