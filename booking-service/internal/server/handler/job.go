package handler

import (
	"booking-svc/pkg/repositories/housekeeper"
	"booking-svc/pkg/repositories/job"
	"booking-svc/pkg/xservice/pricesvc"
	"booking-svc/pkg/xservice/sendingsvc"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func (h *handler) BookHouseKeeper(c *gin.Context) {
	var response BaseResponse
	ctx := c.Request.Context()

	var requestBody BookingHouseKeeperRequestBody
	_ = c.BindJSON(&requestBody)

	clientInfo := buildClientInfoFromRequest(requestBody.ClientInfo)

	bookingDate, err := time.Parse(time.DateTime, requestBody.BookingDate)
	if err != nil {
		log.Error().Msgf("Failed to parse booking datetime due to: %v", err.Error())
		response.Metadata = MetadataResponse{Message: "failed to parse booking datetime"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var getPriceResp pricesvc.GetPriceReponse
	var newJob job.Job

	var wg errgroup.Group

	wg.Go(func() error {
		getPriceResp, err = h.priceSvc.GetPrice(ctx, bookingDate.Unix())
		if err != nil {
			log.Error().Msgf("Failed to get price due to: %v", err.Error())
			return fmt.Errorf("failed to get booking price")
		}
		return nil
	})

	wg.Go(func() error {
		newJob, err = h.jobRepo.InitJob(ctx, clientInfo, bookingDate)
		if err != nil {
			log.Error().Msgf("Failed to init job due to: %v", err.Error())
			return errors.New("failed to init job")
		}
		return nil
	})

	err = wg.Wait()
	if err != nil {
		response.Metadata = MetadataResponse{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	newJob.BookingPrice = getPriceResp.Data.Price
	housekeeper, err := h.housekeeperRepo.PickAvailableHouseKeeper(ctx, newJob.BookingPrice, bookingDate)
	if err != nil {
		log.Error().Msgf("Failed to pick house keeper due to: %v", err.Error())
		response.Metadata = MetadataResponse{Message: "failed to pick available house keeper"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = h.jobRepo.AssignHouseKeeperToJob(ctx, newJob.JobID, newJob.BookingPrice, buildHouseKeeperInfo(housekeeper))
	if err != nil {
		log.Error().Msgf("Failed to assign house keeper to job due to: %v", err.Error())
		response.Metadata = MetadataResponse{Message: "failed to assign house keeper to job"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	_, err = h.sendingSvc.PostNotification(c, sendingsvc.PostNotificationRequest{
		JobID:         newJob.JobID,
		ClientID:      clientInfo.ID,
		HouseKeeperID: housekeeper.HouseKeeperID,
	})
	if err != nil {
		log.Error().Msgf("Failed to post notification due to: %v", err.Error())
		response.Metadata = MetadataResponse{Message: "failed to post notification"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data = newJob
	response.Metadata.Message = "Job is successfully created"
	c.JSON(http.StatusOK, response)
}

func buildHouseKeeperInfo(housekeeper housekeeper.HouseKeeper) job.HouseKeeperInfo {
	return job.HouseKeeperInfo{
		ID:          housekeeper.HouseKeeperID,
		Name:        housekeeper.Name,
		PhoneNumber: housekeeper.PhoneNumber,
	}
}

func buildClientInfoFromRequest(reqClientInfo BookingHouseKeeperClientInfo) job.ClientInfo {
	return job.ClientInfo{
		ID:          reqClientInfo.ID,
		Name:        reqClientInfo.Name,
		PhoneNumber: reqClientInfo.PhoneNumber,
	}
}
