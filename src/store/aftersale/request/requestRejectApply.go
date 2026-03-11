package request

type RequestRejectApply struct {
	AftersaleID string `json:"aftersale_id"`
	Reason      string `json:"reason"`
}
