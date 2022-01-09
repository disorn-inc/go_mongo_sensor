package main

import (
	"github/disorn-inc/go_mongo_sensor/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test_sensor", TestSensor)
	r.Run()
}

func TestSensor(c *gin.Context) {
	var sensor models.Sensor
	c.JSON(http.StatusOK, sensor)
}