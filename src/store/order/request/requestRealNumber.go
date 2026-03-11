package request

type RequestApplyRealNumber struct {
	OrderID string `json:"order_id"`
	Reason  string `json:"reason"`
}

type RequestGetRealNumberStatus struct {
	ApplyID string `json:"apply_id"`
}

type RequestReapplyVirtualNumber struct {
	OrderID string `json:"order_id"`
}

type RequestDelayVirtualNumber struct {
	OrderID   string `json:"order_id"`
	DelayDays int    `json:"delay_days"`
}
