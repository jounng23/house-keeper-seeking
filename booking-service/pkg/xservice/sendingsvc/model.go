package sendingsvc

type PostNotificationRequest struct {
	JobID         string `json:"job_id"`
	ClientID      string `json:"client_id"`
	HouseKeeperID string `json:"house_keeper_id"`
}

type BaseResponse struct {
	Metadata interface{} `json:"meta_data"`
}

type MetadataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
