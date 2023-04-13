package handler

import (
	"booking-service/internal/services/pricing"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	pricing pricing.IPricing
}

func NewHandler(pricing pricing.IPricing) *Handler {
	return &Handler{
		pricing: pricing,
	}
}

func (h *Handler) ConfigAPIRoute(router *gin.Engine) {
	routers := router.Group("v1")
	routers.GET("pricing", h.getPricing())
}
