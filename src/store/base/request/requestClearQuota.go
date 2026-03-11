package request

type RequestClearQuota struct {
}

type RequestClearQuotaByAppSecret struct {
	Secret string `json:"secret"`
}
