package price

import (
	"context"
	"pricing-svc/pkg/utils"
	"time"
)

//go:generate mockgen -source=$GOFILE -package=price_mocks -destination=$PWD/mocks/${GOFILE}
type Repository interface {
	GetPriceByDatetime(c context.Context, datetime time.Time) (float64, error)
}

type repo struct {
	dbStorage Storage
}

func NewRepository(db Storage) Repository {
	return &repo{dbStorage: db}
}

func (repo *repo) GetPriceByDatetime(c context.Context, datetime time.Time) (float64, error) {
	if utils.IsSpecialDay(datetime) {
		return 200000, nil
	}
	return 100000, nil
}
