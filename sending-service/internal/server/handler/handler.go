package handler

import (
	"sending-svc/pkg/repositories/notification"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	PostNotification(c *gin.Context)
}

type handler struct {
	notificationRepo notification.Repository
}

func NewHandler(notificationRepo notification.Repository) Handler {
	return &handler{notificationRepo: notificationRepo}
}
