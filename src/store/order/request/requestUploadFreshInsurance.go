package request

type RequestUploadFreshInsurance struct {
	OrderID       string `json:"order_id"`
	InsuranceInfo string `json:"insurance_info"`
}
