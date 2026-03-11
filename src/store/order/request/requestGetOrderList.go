package request

type RequestGetOrderList struct {
	PageSize  int    `json:"page_size,omitempty"`
	NextKey   string `json:"next_key,omitempty"`
	Status    int    `json:"status,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}
