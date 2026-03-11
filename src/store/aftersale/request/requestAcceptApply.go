package request

type RequestAcceptApply struct {
	AftersaleID string `json:"aftersale_id"`
	AddressID   string `json:"address_id,omitempty"`
}
