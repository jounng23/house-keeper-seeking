package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) GetPrice(c *gin.Context) {
	ctx := c.Request.Context()
	var response BaseResponse
	datetimeUnix, err := strconv.Atoi(c.Query("datetime"))
	if err != nil {
		log.Error().Msgf("Failed to parse datetime due to: %v", err.Error())
		response.Metadata = MetadataResponse{Message: "Failed to parse datetime"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	datetime := time.Unix(int64(datetimeUnix), 0)
	price, _ := h.priceRepo.GetPriceByDatetime(ctx, datetime)
	response.Data = GetPriceResponseData{
		Price: price,
	}
	c.JSON(http.StatusOK, response)
}
