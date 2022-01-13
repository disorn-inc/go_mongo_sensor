package controller

import (
	"context"
	"github/disorn-inc/go_mongo_sensor/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type stringStorer interface {
	NewString(*models.TestString) error
	ReadString(*[]bson.M) (*mongo.Cursor ,error)
}

type StringHandler struct {
	stringStore stringStorer
}

func NewStringHandler(store stringStorer) *StringHandler {
	return &StringHandler{stringStore: store}
}

func (s *StringHandler) NewTestValue(c Context) {
	var test models.TestString
	if err := c.Bind(&test); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err := s.stringStore.NewString(&test)
	if err != nil {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,  map[string]interface{}{
		"massege": test.Massege,
	})
}

func (s *StringHandler) ListTestValue(c Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var tests []bson.M
	cur, err := s.stringStore.ReadString(&tests)
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