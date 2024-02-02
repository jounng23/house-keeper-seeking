package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const maxPoolSize = 500

func InitMongoDBClient(uri string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(uri).
		SetMaxPoolSize(maxPoolSize).
		SetReadPreference(readpref.Primary())
	return mongo.NewClient(opts)
}

//nolint:gocritic
func HandlePoolMonitor(evt *event.PoolEvent) {
	switch evt.Type {
	case event.PoolClosedEvent:
		fmt.Println("DB connection closed.")
	}
}
