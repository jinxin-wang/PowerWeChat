package request

type RequestSearchOrder struct {
	Keyword  string `json:"keyword"`
	PageSize int    `json:"page_size,omitempty"`
	NextKey  string `json:"next_key,omitempty"`
}
