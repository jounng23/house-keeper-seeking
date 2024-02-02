package price

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "prices"

type Storage interface{}

type storage struct {
	C *mongo.Collection
}

func NewStorage(db *mongo.Database) Storage {
	return &storage{C: db.Collection(collectionName)}
}
