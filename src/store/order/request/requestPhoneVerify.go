package request

type RequestAddPhoneVerifyCode struct {
	OrderID string `json:"order_id"`
	Phone   string `json:"phone"`
}

type RequestSendPhoneVerifyCode struct {
	OrderID string `json:"order_id"`
}

type RequestGetPhoneStatus struct {
	OrderID string `json:"order_id"`
}
