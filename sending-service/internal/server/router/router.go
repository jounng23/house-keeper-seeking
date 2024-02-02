package router

import (
	"sending-svc/internal/server/handler"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(g gin.IRouter)
}

type router struct {
	handler handler.Handler
}

func NewRouter(handler handler.Handler) Router {
	return &router{handler: handler}
}

func (r *router) Register(g gin.IRouter) {
	priceGroup := g.Group("/notifications")
	{
		priceGroup.POST("/", r.handler.PostNotification)
	}
}
