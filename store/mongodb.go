package store

import (
	"context"
	"github/disorn-inc/go_mongo_sensor/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore(col *mongo.Collection) *MongoDBStore {
	return &MongoDBStore{Collection: col}
}

func(s *MongoDBStore) New(sensor *models.Sensor) error {
	now := time.Now()
	sensor.CreatedAt = now
	sensor.UpdateAt = now
	_, err := s.InsertOne(context.Background(), sensor)
	return err
}

func(s *MongoDBStore) Read(testStrings *[]bson.M) (*mongo.Cursor, error) {
	// var tests []*models.TestString
	cur,err := s.Find(context.TODO(), bson.M{})
	return cur, err
}