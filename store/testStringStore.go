package store

import (
	"context"
	"github/disorn-inc/go_mongo_sensor/models"
	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestStringStore struct {
	*mongo.Collection
}

func NewTestStringStore(col *mongo.Collection) *MongoDBStore {
	return &MongoDBStore{Collection: col}
}

func(s *MongoDBStore) NewString(testStrings *models.TestString) error {
	_, err := s.InsertOne(context.Background(), testStrings)
	return err
}

func(s *MongoDBStore) ReadString(testStrings *[]bson.M) (*mongo.Cursor, error) {
	// var tests []*models.TestString
	cur,err := s.Find(context.TODO(), bson.M{})
	return cur, err
}

