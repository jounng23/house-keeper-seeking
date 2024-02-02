package job

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "jobs"

type Storage interface {
	InitJob(c context.Context, newJob Job) error
	UpdateJob(c context.Context, jobID string, updateParams UpdateParams) error
}

type storage struct {
	C *mongo.Collection
}

func NewStorage(db *mongo.Database) Storage {
	return &storage{C: db.Collection(collectionName)}
}

func (s *storage) InitJob(c context.Context, newJob Job) (err error) {
	_, err = s.C.InsertOne(c, newJob)
	return
}

func (s *storage) UpdateJob(c context.Context, jobID string, updateParams UpdateParams) error {
	filters := bson.M{"job_id": jobID}

	updates := bson.M{}

	if updateParams.BookingPrice != nil {
		updates["booking_price"] = updateParams.BookingPrice
	}

	if updateParams.HouseKeeperInfo != nil {
		updates["house_keeper_info"] = updateParams.HouseKeeperInfo
	}

	_, err := s.C.UpdateOne(c, filters, bson.M{"$set": updates})
	return err
}
