package handler

import (
	housekeeper_mocks "booking-svc/pkg/repositories/housekeeper/mocks"
	job_mocks "booking-svc/pkg/repositories/job/mocks"
	pricesvc_mocks "booking-svc/pkg/xservice/pricesvc/mocks"
	sendingsvc_mocks "booking-svc/pkg/xservice/sendingsvc/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	// . "github.com/smartystreets/goconvey/convey"
)

var defaultJobID = "23eaa454-13ab-4bcf-a7fc-28f163495b38"
var defaulClient = BookingHouseKeeperClientInfo{
	ID:          "23eaa454-13ab-4bcf-a7fc-28f163499b27",
	Name:        "Nguyen Van X",
	PhoneNumber: "0987654321",
}

type jobHandlerSuite struct {
	goMockCtrl      *gomock.Controller
	jobRepo         *job_mocks.MockRepository
	houseKeeperRepo *housekeeper_mocks.MockRepository
	priceSvc        *pricesvc_mocks.MockService
	sendingSvc      *sendingsvc_mocks.MockService

	handler Handler
}

func TestIsReturnCorrectPriceOnBookingDate(t *testing.T) {
	// var jobHandlerSuite jobHandlerSuite
	// jobHandlerSuite.SetupTestJobHandler(t)

	// defer jobHandlerSuite.goMockCtrl.Finish()

	// testcases := []struct {
	// 	name       string
	// 	genInputFn func() (c *gin.Context)
	// 	expectFn   func(bookingPrice float64)
	// }{
	// 	{
	// 		name: "Behavior_BookingOnSpecialDate_Return_SuccessOnCorrectPrice",
	// 		genInputFn: func() (c *gin.Context) {
	// 			jsonBody, _ := json.Marshal(BookingHouseKeeperRequestBody{
	// 				ClientInfo:  defaulClient,
	// 				BookingDate: "2024-01-01 08:30:00",
	// 			})
	// 			w := httptest.NewRecorder()
	// 			w.Body = bytes.NewBuffer(jsonBody)
	// 			c, _ = gin.CreateTestContext(w)
	// 			return
	// 		},
	// 		expectFn: func(bookingPrice float64) {
	// 			Convey("The booking price should be equal to 200000", func() {
	// 				So(bookingPrice, ShouldEqual, 200000)
	// 			})
	// 		},
	// 	},
	// }

	// for _, tc := range testcases {
	// c := tc.genInputFn()
	// tc.expectFn()
	// }
}

func (s jobHandlerSuite) SetupTestJobHandler(t *testing.T) {
	s.goMockCtrl = gomock.NewController(t)
	s.jobRepo = job_mocks.NewMockRepository(s.goMockCtrl)
	s.houseKeeperRepo = housekeeper_mocks.NewMockRepository(s.goMockCtrl)
	s.priceSvc = pricesvc_mocks.NewMockService(s.goMockCtrl)
	s.sendingSvc = sendingsvc_mocks.NewMockService(s.goMockCtrl)

	s.handler = NewHandler(s.jobRepo, s.houseKeeperRepo, s.priceSvc, s.sendingSvc)
}
