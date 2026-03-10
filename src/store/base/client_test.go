package base

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRequestGetStableAccessToken(t *testing.T) {
	req := &RequestGetStableAccessToken{
		GrantType: "client_credential",
		AppID:     "test_appid",
		Secret:    "test_secret",
	}

	assert.Equal(t, "client_credential", req.GrantType)
	assert.Equal(t, "test_appid", req.AppID)
	assert.Equal(t, "test_secret", req.Secret)
}

func TestRequestGetAPIQuota(t *testing.T) {
	req := &RequestGetAPIQuota{
		CgiPath: "/wxa/getunlimited",
	}

	assert.Equal(t, "/wxa/getunlimited", req.CgiPath)
}

func TestRequestGetAPIQuota_Empty(t *testing.T) {
	req := &RequestGetAPIQuota{}

	assert.Equal(t, "", req.CgiPath)
}

func TestRequestClearAPIQuota(t *testing.T) {
	req := &RequestClearAPIQuota{
		CgiPath: "/wxa/getunlimited",
	}

	assert.Equal(t, "/wxa/getunlimited", req.CgiPath)
}

func TestRequestClearQuotaByAppSecret(t *testing.T) {
	req := &RequestClearQuotaByAppSecret{
		Secret: "test_secret",
	}

	assert.Equal(t, "test_secret", req.Secret)
}

func TestRequestCallbackCheck(t *testing.T) {
	req := &RequestCallbackCheck{
		Action:   "callback",
		Operator: "DEFAULT",
	}

	assert.Equal(t, "callback", req.Action)
	assert.Equal(t, "DEFAULT", req.Operator)
}

func TestRequestGetRidInfo(t *testing.T) {
	req := &RequestGetRidInfo{
		Rid: "test_rid_123456",
	}

	assert.Equal(t, "test_rid_123456", req.Rid)
}

func TestRequestGetDataByMediaID(t *testing.T) {
	req := &RequestGetDataByMediaID{
		MediaID: "test_media_id",
	}

	assert.Equal(t, "test_media_id", req.MediaID)
}

func TestResponseGetStableAccessToken(t *testing.T) {
	resp := &ResponseGetStableAccessToken{
		AccessToken: "test_access_token",
		ExpiresIn:   7200,
	}

	assert.Equal(t, "test_access_token", resp.AccessToken)
	assert.Equal(t, 7200, resp.ExpiresIn)
}

func TestResponseGetAPIQuota(t *testing.T) {
	resp := &ResponseGetAPIQuota{
		Quota:      "1000",
		QuotaLimit: "2000",
	}

	assert.Equal(t, "1000", resp.Quota)
	assert.Equal(t, "2000", resp.QuotaLimit)
}

func TestResponseCallbackCheck(t *testing.T) {
	resp := &ResponseCallbackCheck{
		Operator: "DEFAULT",
		Result:   "OK",
	}

	assert.Equal(t, "DEFAULT", resp.Operator)
	assert.Equal(t, "OK", resp.Result)
}

func TestResponseGetAPIDomainIP(t *testing.T) {
	resp := &ResponseGetAPIDomainIP{
		IPList: []string{"127.0.0.1", "127.0.0.2"},
	}

	assert.Equal(t, 2, len(resp.IPList))
	assert.Equal(t, "127.0.0.1", resp.IPList[0])
}

func TestResponseGetCallbackIP(t *testing.T) {
	resp := &ResponseGetCallbackIP{
		IPList: []string{"127.0.0.1", "127.0.0.2", "127.0.0.3"},
	}

	assert.Equal(t, 3, len(resp.IPList))
}

func TestResponseGetRidInfo(t *testing.T) {
	resp := &ResponseGetRidInfo{
		RequestInfo: "test_request_info",
		RequestMsg:  "test_request_msg",
	}

	assert.Equal(t, "test_request_info", resp.RequestInfo)
	assert.Equal(t, "test_request_msg", resp.RequestMsg)
}

func TestResponseGetDataByMediaID(t *testing.T) {
	resp := &ResponseGetDataByMediaID{
		Data: "test_data_content",
	}

	assert.Equal(t, "test_data_content", resp.Data)
}
