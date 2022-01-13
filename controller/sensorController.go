package controller

import (
	"context"
	"github/disorn-inc/go_mongo_sensor/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type storer interface {
	New(*models.Sensor) error
	Read(*[]bson.M) (*mongo.Cursor ,error)
}

type SensorHandler struct {
	store storer
}

func NewSensorHandler(store storer) *SensorHandler {
	return &SensorHandler{store: store}
}

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
}

func (s *SensorHandler) NewValue(c Context) {
	var sensor models.Sensor
	if err := c.Bind(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err := s.store.New(&sensor)
	if err != nil {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,  map[string]interface{}{
		"ID": sensor.ID,
	})
}

func (s *SensorHandler) TestSensor(c Context) {
	var sensor models.Sensor
	c.JSON(http.StatusOK, sensor)
}

func (s *SensorHandler) ListSensorValue(c Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var tests []bson.M
	cur, err := s.store.Read(&tests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	if err = cur.All(ctx, &tests); err != nil {
        c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        return
    }
	defer cancel()
	c.JSON(http.StatusOK, tests)
}