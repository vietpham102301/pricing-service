package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getPricing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("start getPricing...")
		q, data := ctx.Request.URL.Query(), make(map[string]int)
		for k, v := range q {
			if len(v) > 0 {
				val, err := strconv.Atoi(v[0])
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				data[k] = val
			}
		}
		record, err := h.pricing.GetPricing(data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if record == nil {
			ctx.JSON(http.StatusBadRequest, "No matched jobType")
			return
		}

		ctx.JSON(http.StatusOK, record)
	}
}
