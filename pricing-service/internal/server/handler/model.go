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
