package job

import (
	"booking-svc/pkg/enum"
	"booking-svc/pkg/utils"
	"booking-svc/pkg/xservice/pricesvc"
	"context"
	"time"
)

//go:generate mockgen -source=$GOFILE -package=job_mocks -destination=$PWD/mocks/${GOFILE}
type Repository interface {
	InitJob(c context.Context, clientInfo ClientInfo, bookingDate time.Time) (Job, error)
	AssignHouseKeeperToJob(c context.Context, jobID string, bookingPrice float64, houseKeeperInfo HouseKeeperInfo) error
}

type repo struct {
	dbStorage  Storage
	pricingSvc pricesvc.Service
}

func NewRepository(db Storage, pricingSvc pricesvc.Service) Repository {
	return &repo{dbStorage: db, pricingSvc: pricingSvc}
}

func (repo *repo) InitJob(c context.Context, clientInfo ClientInfo, bookingDate time.Time) (newJob Job, err error) {
	newJob = Job{
		JobID:       utils.GenerateUUID(),
		ClientInfo:  clientInfo,
		BookingDate: bookingDate.Unix(),
		CreatedAt:   time.Now().Unix(),
		Status:      enum.JStatusNew,
	}
	err = repo.dbStorage.InitJob(c, newJob)
	return
}

func (repo *repo) AssignHouseKeeperToJob(c context.Context, jobID string, bookingPrice float64, houseKeeperInfo HouseKeeperInfo) (err error) {
	jobStatus := enum.JStatusAssigned
	updateParams := UpdateParams{
		BookingPrice:    &bookingPrice,
		HouseKeeperInfo: &houseKeeperInfo,
		JobStatus:       &jobStatus,
	}
	err = repo.dbStorage.UpdateJob(c, jobID, updateParams)
	return
}
