package job

import "booking-svc/pkg/enum"

type Job struct {
	JobID           string          `json:"job_id" bson:"job_id"`
	ClientInfo      ClientInfo      `json:"client_info" bson:"client_info"`
	HouseKeeperInfo HouseKeeperInfo `json:"house_keeper_info" bson:"house_keeper_info"`
	BookingDate     int64           `json:"booking_date" bson:"booking_date"`
	BookingPrice    float64         `json:"booking_price" bson:"booking_price"`
	Status          enum.JobStatus  `json:"status" bson:"status"`
	CreatedAt       int64           `json:"created_at" bson:"created_at"`
}

type ClientInfo struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type HouseKeeperInfo struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type UpdateParams struct {
	JobStatus       *enum.JobStatus
	BookingPrice    *float64
	HouseKeeperInfo *HouseKeeperInfo
}
