package kf

import (
	"testing"

	"github.com/jinxin-wang/PowerWeChat/v3/src/store/kf/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/kf/response"
	"github.com/go-playground/assert/v2"
)

func TestRequestCOSUpload(t *testing.T) {
	req := &request.RequestCOSUpload{
		MediaType: "image",
		MediaData: "base64_data_here",
	}

	assert.Equal(t, "image", req.MediaType)
	assert.Equal(t, "base64_data_here", req.MediaData)
}

func TestRequestSendMsg(t *testing.T) {
	req := &request.RequestSendMsg{
		OpenID:  "oXXXXX",
		MsgType: "text",
		Content: "Hello",
	}

	assert.Equal(t, "oXXXXX", req.OpenID)
	assert.Equal(t, "text", req.MsgType)
	assert.Equal(t, "Hello", req.Content)
}

func TestResponseCOSUpload(t *testing.T) {
	resp := &response.ResponseCOSUpload{
		MediaID: "media_id_123",
		URL:     "https://example.com/image.jpg",
	}

	assert.Equal(t, "media_id_123", resp.MediaID)
	assert.Equal(t, "https://example.com/image.jpg", resp.URL)
}

func TestResponseSendMsg(t *testing.T) {
	resp := &response.ResponseSendMsg{}

	assert.Equal(t, 0, resp.ErrCode)
}
