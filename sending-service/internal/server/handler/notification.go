package handler

import (
	"net/http"
	"sending-svc/pkg/repositories/notification"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) PostNotification(c *gin.Context) {
	ctx := c.Request.Context()
	var response BaseResponse

	var requestBody PostNotificationRequestBody
	_ = c.BindJSON(&requestBody)

	notification := buildNotificationFromRequestBody(requestBody)
	err := h.notificationRepo.CreateNotification(ctx, notification)
	if err != nil {
		log.Error().Msgf("Failed to create notification due to: %v", err.Error())
		response.Metadata = MetadataResponse{Message: "Failed to parse datetime"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func buildNotificationFromRequestBody(reqBody PostNotificationRequestBody) notification.Notification {
	return notification.Notification{
		JobID:         reqBody.JobID,
		ClientID:      reqBody.ClientID,
		HouseKeeperID: reqBody.HouseKeeperID,
	}
}
