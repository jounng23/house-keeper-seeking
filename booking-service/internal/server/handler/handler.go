package handler

import (
	"booking-svc/pkg/repositories/housekeeper"
	"booking-svc/pkg/repositories/job"
	"booking-svc/pkg/xservice/pricesvc"
	"booking-svc/pkg/xservice/sendingsvc"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	BookHouseKeeper(c *gin.Context)
}

type handler struct {
	jobRepo         job.Repository
	housekeeperRepo housekeeper.Repository
	priceSvc        pricesvc.Service
	sendingSvc      sendingsvc.Service
}

func NewHandler(
	jobRepo job.Repository,
	housekeeperRepo housekeeper.Repository,
	priceSvc pricesvc.Service,
	sendingSvc sendingsvc.Service,
) Handler {
	return &handler{
		jobRepo:         jobRepo,
		housekeeperRepo: housekeeperRepo,
		priceSvc:        priceSvc,
		sendingSvc:      sendingSvc,
	}
}
