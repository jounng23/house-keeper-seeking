package notification

import (
	"context"
	"sending-svc/pkg/utils"
	"time"
)

//go:generate mockgen -source=$GOFILE -package=notification_mocks -destination=$PWD/mocks/${GOFILE}
type Repository interface {
	CreateNotification(c context.Context, noti Notification) error
}

type repo struct {
	dbStorage Storage
}

func NewRepository(db Storage) Repository {
	return &repo{dbStorage: db}
}

func (repo *repo) CreateNotification(c context.Context, noti Notification) error {
	newNoti := noti
	newNoti.NotiID = utils.GenerateUUID()
	newNoti.CreatedAt = time.Now().Unix()
	return repo.dbStorage.CreateNotification(c, newNoti)
}
