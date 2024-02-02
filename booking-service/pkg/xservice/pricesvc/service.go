package pricesvc

import (
	"booking-svc/pkg/utils"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

type Service interface {
	GetPrice(c context.Context, datetime int64) (GetPriceReponse, error)
}

type service struct {
	client  *http.Client
	baseUrl string
}

//go:generate mockgen -source=$GOFILE -package=pricesvc_mocks -destination=$PWD/mocks/${GOFILE}
func NewService(client *http.Client, baseUrl string) Service {
	return &service{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (s *service) GetPrice(c context.Context, datetime int64) (response GetPriceReponse, err error) {
	parsedURL, _ := url.Parse(s.baseUrl + "/api/v1/price")

	query := url.Values{}
	query.Set("datetime", fmt.Sprint(datetime))

	parsedURL.RawQuery = query.Encode()

	request := &http.Request{
		Method: http.MethodGet,
		URL:    parsedURL,
	}

	resp, err := s.client.Do(request)
	if err != nil {
		log.Error().Msgf("Failed to get price from pricesvc due to: %s", err.Error())
		return
	}

	err = utils.ParseResponse(resp, &response)
	if err != nil {
		log.Error().Msgf("Failed to get price from pricesvc due to: %s", err.Error())
	}
	return
}
