package handler

type BaseResponse struct {
	Data     interface{}      `json:"data"`
	Metadata MetadataResponse `json:"meta_data"`
}

type MetadataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetPriceResponseData struct {
	Price float64 `json:"price"`
}

type PostNotificationRequestBody struct {
	JobID         string `json:"job_id"`
	ClientID      string `json:"client_id"`
	HouseKeeperID string `json:"house_keeper_id"`
}
