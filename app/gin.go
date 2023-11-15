package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context interface {
	Params(param string) string
	OK(res any)
}

type MyContext struct {
	*gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

func (c *MyContext) Params(param string) string {
	return c.Context.Param(param)
}

func (c *MyContext) OK(res any) {
	c.Context.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"response": res,
	})
}

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}
