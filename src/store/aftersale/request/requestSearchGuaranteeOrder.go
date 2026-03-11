package request

type RequestSearchGuaranteeOrder struct {
	PageSize int    `json:"page_size,omitempty"`
	NextKey  string `json:"next_key,omitempty"`
	Status   int    `json:"status,omitempty"`
}
