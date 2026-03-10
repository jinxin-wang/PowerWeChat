package base

type RequestGetStableAccessToken struct {
	GrantType string `json:"grant_type"`
	AppID     string `json:"appid"`
	Secret    string `json:"secret"`
}

type RequestGetAPIQuota struct {
	CgiPath string `json:"cgi_path,omitempty"`
}

type RequestClearAPIQuota struct {
	CgiPath string `json:"cgi_path"`
}

type RequestClearQuotaByAppSecret struct {
	Secret string `json:"secret"`
}

type RequestCallbackCheck struct {
	Action   string `json:"action"`
	Operator string `json:"operator"`
}

type RequestGetRidInfo struct {
	Rid string `json:"rid"`
}

type RequestGetDataByMediaID struct {
	MediaID string `json:"media_id"`
}
