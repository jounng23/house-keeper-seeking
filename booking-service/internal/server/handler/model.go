package handler

type BookingHouseKeeperClientInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type BookingHouseKeeperRequestBody struct {
	ClientInfo  BookingHouseKeeperClientInfo `json:"client"`
	BookingDate string                       `json:"booking_date"`
}

type BaseResponse struct {
	Data     interface{}      `json:"data"`
	Metadata MetadataResponse `json:"meta_data"`
}

type MetadataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
