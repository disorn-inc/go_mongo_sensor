package controller

import (
	"github/disorn-inc/go_mongo_sensor/models"
	"net/http"
)

type storer interface {
	New(*models.Sensor) error
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