package app

import (
	"github.com/gin-gonic/gin"
)

type Context interface {
	Params(param string) string
}

type MyContext struct {
	*gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

func (g *MyContext) Params(param string) string {
	return g.Context.Param(param)
}

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}
