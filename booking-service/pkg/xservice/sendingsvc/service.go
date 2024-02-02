package sendingsvc

import (
	"booking-svc/pkg/utils"
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Service interface {
	PostNotification(c context.Context, req PostNotificationRequest) (BaseResponse, error)
}

type service struct {
	client  *http.Client
	baseUrl string
}

//go:generate mockgen -source=$GOFILE -package=sendingsvc_mocks -destination=$PWD/mocks/${GOFILE}
func NewService(client *http.Client, baseUrl string) Service {
	return &service{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (s *service) PostNotification(c context.Context, req PostNotificationRequest) (response BaseResponse, err error) {
	bodyByte, _ := json.Marshal(req)
	request, err := http.NewRequestWithContext(c, http.MethodPost, s.baseUrl+"/api/v1/notifications/", bytes.NewReader(bodyByte))
	if err != nil {
		log.Error().Msgf("Failed to create request due to: %s", err.Error())
		return
	}

	resp, err := s.client.Do(request)
	if err != nil {
		log.Error().Msgf("Failed to post notification to sending service due to: %s", err.Error())
		return
	}

	err = utils.ParseResponse(resp, &response)
	if err != nil {
		log.Error().Msgf("Failed to parse response from sending service due to: %s", err.Error())
	}
	return
}
