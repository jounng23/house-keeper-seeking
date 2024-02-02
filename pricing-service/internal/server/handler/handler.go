package handler

import (
	"pricing-svc/pkg/repositories/price"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetPrice(c *gin.Context)
}

type handler struct {
	priceRepo price.Repository
}

func NewHandler(priceRepo price.Repository) Handler {
	return &handler{priceRepo: priceRepo}
}
