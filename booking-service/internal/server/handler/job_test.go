package handler

import (
	"booking-svc/pkg/repositories/housekeeper"
	housekeeper_mocks "booking-svc/pkg/repositories/housekeeper/mocks"
	"booking-svc/pkg/repositories/job"
	job_mocks "booking-svc/pkg/repositories/job/mocks"
	"booking-svc/pkg/xservice/pricesvc"
	pricesvc_mocks "booking-svc/pkg/xservice/pricesvc/mocks"
	"booking-svc/pkg/xservice/sendingsvc"
	sendingsvc_mocks "booking-svc/pkg/xservice/sendingsvc/mocks"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

var DefaultTestClient = BookingHouseKeeperClientInfo{
	ID:          "402b9df9-c5b0-44ba-ac04-0e8c33cdf233",
	Name:        "Nguyen Van X",
	PhoneNumber: "0987654321",
}

var DefaultTestHouseKeeper = housekeeper.HouseKeeper{
	HouseKeeperID: "402b9df9-c5b0-44ba-ac04-0e8c33cdf233",
	Name:          "Nguyen Van Y",
	PhoneNumber:   "0912345678",
}

var DefaultTestJob = job.Job{
	JobID: "402b9df9-c5b0-44ba-ac04-0e8c33cdf223",
}

type handlerTestSuite struct {
	jobRepo         *job_mocks.MockRepository
	housekeeperRepo *housekeeper_mocks.MockRepository
	priceSvc        *pricesvc_mocks.MockService
	sendingSvc      *sendingsvc_mocks.MockService
}

func initHandlerTestSuite(ctrl *gomock.Controller) handlerTestSuite {
	return handlerTestSuite{
		jobRepo:         job_mocks.NewMockRepository(ctrl),
		housekeeperRepo: housekeeper_mocks.NewMockRepository(ctrl),
		priceSvc:        pricesvc_mocks.NewMockService(ctrl),
		sendingSvc:      sendingsvc_mocks.NewMockService(ctrl),
	}
}

func newTestHandler(ts handlerTestSuite) *handler {
	return &handler{
		jobRepo:         ts.jobRepo,
		housekeeperRepo: ts.housekeeperRepo,
		priceSvc:        ts.priceSvc,
		sendingSvc:      ts.sendingSvc,
	}
}

func Test_bookHouseKeeperHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testSuite := initHandlerTestSuite(ctrl)

	testcases := []struct {
		name                                string
		genRequestBodyFn                    func() BookingHouseKeeperRequestBody
		genGetPriceResultFn                 func() (pricesvc.GetPriceReponse, error)
		genInitJobResultFn                  func() (job.Job, error)
		genPickAvailableHousekeeperResultFn func() (housekeeper.HouseKeeper, error)
		genAssignHousekeeperToJobResultFn   func() error
		genPostNotifyJobResultFn            func() (sendingsvc.BaseResponse, error)
		expectFn                            func(tcName string, newJob job.Job, err error)
	}{
		{
			name: "Checking Data Response Correct",
			genGetPriceResultFn: func() (pricesvc.GetPriceReponse, error) {
				return pricesvc.GetPriceReponse{
					Data: pricesvc.GetPriceResponseData{
						Price: 200000,
					},
				}, nil
			},
			genRequestBodyFn: func() BookingHouseKeeperRequestBody {
				return BookingHouseKeeperRequestBody{
					ClientInfo:  DefaultTestClient,
					BookingDate: "2024-01-01 08:30:00",
				}
			},
			genInitJobResultFn: func() (job.Job, error) {
				var j = DefaultTestJob
				j.ClientInfo = job.ClientInfo{
					ID:          DefaultTestClient.ID,
					Name:        DefaultTestClient.Name,
					PhoneNumber: DefaultTestClient.PhoneNumber,
				}
				d, _ := time.Parse(time.DateTime, "2024-01-01 08:30:00")
				j.BookingDate = d.Unix()
				return j, nil
			},
			genPickAvailableHousekeeperResultFn: func() (housekeeper.HouseKeeper, error) {
				return DefaultTestHouseKeeper, nil
			},
			genAssignHousekeeperToJobResultFn: func() error { return nil },
			genPostNotifyJobResultFn: func() (sendingsvc.BaseResponse, error) {
				return sendingsvc.BaseResponse{}, nil
			},
			expectFn: func(tcName string, newJob job.Job, err error) {
				Convey(tcName, t, func() {
					Convey("Check is client info correct", func() {
						Convey("Check is client ID correct", func() {
							So(newJob.ClientInfo.ID, ShouldEqual, DefaultTestClient.ID)
						})

						Convey("Check is client name correct", func() {
							So(newJob.ClientInfo.Name, ShouldEqual, DefaultTestClient.Name)
						})

						Convey("Check is client phone number correct", func() {
							So(newJob.ClientInfo.PhoneNumber, ShouldEqual, DefaultTestClient.PhoneNumber)
						})
					})

					Convey("Check is housekeeper info correct", func() {
						Convey("Check is housekeeper ID correct", func() {
							So(newJob.HouseKeeperInfo.ID, ShouldEqual, DefaultTestHouseKeeper.HouseKeeperID)
						})

						Convey("Check is housekeeper name correct", func() {
							So(newJob.HouseKeeperInfo.Name, ShouldEqual, DefaultTestHouseKeeper.Name)
						})

						Convey("Check is housekeeper phone number correct", func() {
							So(newJob.HouseKeeperInfo.PhoneNumber, ShouldEqual, DefaultTestHouseKeeper.PhoneNumber)
						})
					})

					Convey("Check is booking date correct", func() {
						expectedRs, _ := time.Parse(time.DateTime, "2024-01-01 08:30:00")
						So(newJob.BookingDate, ShouldEqual, expectedRs.Unix())
					})

					Convey("Check is booking price correct", func() {
						So(newJob.BookingPrice, ShouldEqual, 200000)
					})
				})
			},
		},
	}

	for _, tc := range testcases {
		if tc.genInitJobResultFn != nil {
			testSuite.jobRepo.
				EXPECT().
				InitJob(context.Background(), gomock.Any(), gomock.Any()).
				Return(tc.genInitJobResultFn()).
				AnyTimes()
		}

		if tc.genGetPriceResultFn != nil {
			testSuite.priceSvc.
				EXPECT().
				GetPrice(context.Background(), gomock.Any()).
				Return(tc.genGetPriceResultFn()).
				AnyTimes()
		}

		if tc.genPickAvailableHousekeeperResultFn != nil {
			testSuite.housekeeperRepo.
				EXPECT().
				PickAvailableHouseKeeper(context.Background(), gomock.Any(), gomock.Any()).
				Return(tc.genPickAvailableHousekeeperResultFn()).
				AnyTimes()
		}

		if tc.genAssignHousekeeperToJobResultFn != nil {
			testSuite.jobRepo.
				EXPECT().
				AssignHouseKeeperToJob(context.Background(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tc.genAssignHousekeeperToJobResultFn()).
				AnyTimes()
		}

		if tc.genPostNotifyJobResultFn != nil {
			testSuite.sendingSvc.
				EXPECT().
				PostNotification(context.Background(), gomock.Any()).
				Return(tc.genPostNotifyJobResultFn()).
				AnyTimes()
		}

		h := newTestHandler(testSuite)
		newJob, err := h.bookHouseKeeperHandler(context.Background(), tc.genRequestBodyFn())
		tc.expectFn(tc.name, newJob, err)
	}
}
