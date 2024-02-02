package notification

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "notifications"

type Storage interface {
	CreateNotification(c context.Context, noti Notification) error
}

type storage struct {
	C *mongo.Collection
}

func NewStorage(db *mongo.Database) Storage {
	return &storage{C: db.Collection(collectionName)}
}

func (s *storage) CreateNotification(c context.Context, noti Notification) (err error) {
	_, err = s.C.InsertOne(c, noti)
	return
}
