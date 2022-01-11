package router

import (
	"github/disorn-inc/go_mongo_sensor/controller"

	"github.com/gin-gonic/gin"
)

type MyContext struct {
	*gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}
func (c *MyContext) JSON(statuscode int, v interface{}) {
	c.Context.JSON(statuscode, v)
}

func NewGinHandler(handler func(controller.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default()

	return &MyRouter{r}
}

func (r *MyRouter) POST(path string, handler func(controller.Context)) {
	r.Engine.POST(path, NewGinHandler(handler))
}

func (r *MyRouter) GET(path string, handler func(controller.Context)) {
	r.Engine.GET(path, NewGinHandler(handler))
}