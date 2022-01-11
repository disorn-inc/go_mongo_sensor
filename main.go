package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github/disorn-inc/go_mongo_sensor/controller"
	"github/disorn-inc/go_mongo_sensor/store"
	"github/disorn-inc/go_mongo_sensor/router"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	_, err := os.Create("/tmp/live")

	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove("/tmp/live")

	err = godotenv.Load(".env")
	if err != nil {
		log.Println("please consider environment variable: %s", err)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:123456@localhost:27017"))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database("example").Collection("todos")
	fmt.Println(collection)
	mongoStore := store.NewMongoDBStore(collection)
	handler := controller.NewSensorHandler(mongoStore)

	r := router.NewMyRouter()
	r.GET("/test_sensor", handler.TestSensor)
	r.POST("/sensor", handler.NewValue)
	r.Run()
}
