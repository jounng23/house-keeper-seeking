package housekeeper

import (
	"context"
	"math/rand"
	"time"
)

var HouseKeepers = []HouseKeeper{
	{
		HouseKeeperID: "1",
		Name:          "Nguyen Van A",
		PhoneNumber:   "0982214532",
	},
	{
		HouseKeeperID: "2",
		Name:          "Nguyen Thi B",
		PhoneNumber:   "0922877843",
	},
	{
		HouseKeeperID: "3",
		Name:          "Nguyen Van C",
		PhoneNumber:   "0344657758",
	},
}

//go:generate mockgen -source=$GOFILE -package=housekeeper_mocks -destination=$PWD/mocks/${GOFILE}
type Repository interface {
	PickAvailableHouseKeeper(c context.Context, bookingPrice float64, bookingDate time.Time) (HouseKeeper, error)
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (repo *repo) PickAvailableHouseKeeper(c context.Context, bookingPrice float64, bookingDate time.Time) (HouseKeeper, error) {
	// Note: Assuming returning random housekeeper
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(HouseKeepers))
	randomHouseKeeper := HouseKeepers[randomIndex]

	return randomHouseKeeper, nil
}
