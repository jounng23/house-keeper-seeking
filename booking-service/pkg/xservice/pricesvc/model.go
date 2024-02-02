package pricesvc

type BaseResponse struct {
	Metadata interface{} `json:"meta_data"`
}

type GetPriceReponse struct {
	BaseResponse
	Data GetPriceResponseData `json:"data"`
}

type MetadataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetPriceResponseData struct {
	Price float64 `json:"price"`
}
