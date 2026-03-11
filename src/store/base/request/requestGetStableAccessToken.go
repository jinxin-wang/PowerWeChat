package request

type RequestGetStableAccessToken struct {
	GrantType string `json:"grant_type"`
	AppID     string `json:"appid"`
	Secret    string `json:"secret"`
}
