package notification

type Notification struct {
	NotiID        string `json:"noti_id" bson:"noti_id"`
	JobID         string `json:"job_id" bson:"job_id"`
	ClientID      string `json:"client_id" bson:"client_id"`
	HouseKeeperID string `json:"house_keeper_id" bson:"house_keeper_id"`
	NotiMessage   string `json:"noti_message" bson:"noti_message"`
	CreatedAt     int64  `json:"created_at" bson:"created_at"`
}
