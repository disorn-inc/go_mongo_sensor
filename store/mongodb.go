package store

import (
	"context"
	"github/disorn-inc/go_mongo_sensor/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore(col *mongo.Collection) *MongoDBStore {
	return &MongoDBStore{Collection: col}
}

func(s *MongoDBStore) New(sensor *models.Sensor) error {
	_, err := s.InsertOne(context.Background(), sensor)
	return err
}