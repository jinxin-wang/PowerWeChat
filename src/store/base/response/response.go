package base

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetStableAccessToken struct {
	response.ResponseStore
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type ResponseGetAPIQuota struct {
	response.ResponseStore
	Quota      string `json:"quota"`
	QuotaLimit string `json:"quota_limit"`
}

type ResponseCallbackCheck struct {
	response.ResponseStore
	Operator string `json:"operator"`
	Result   string `json:"result"`
}

type ResponseGetAPIDomainIP struct {
	response.ResponseStore
	IPList []string `json:"ip_list"`
}

type ResponseGetCallbackIP struct {
	response.ResponseStore
	IPList []string `json:"ip_list"`
}

type ResponseGetRidInfo struct {
	response.ResponseStore
	RequestInfo string `json:"request_info"`
	RequestMsg  string `json:"request_msg"`
}

type ResponseGetDataByMediaID struct {
	response.ResponseStore
	Data string `json:"data"`
}
