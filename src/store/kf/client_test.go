package kf

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/kf/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/kf/response"
)

func TestRequestCOSUpload(t *testing.T) {
	req := &request.RequestCOSUpload{
		MediaType: "image",
		MediaData: "base64_data_here",
	}

	assert.Equal(t, "image", req.MediaType)
	assert.Equal(t, "base64_data_here", req.MediaData)
}

func TestRequestCOSUpload_Empty(t *testing.T) {
	req := &request.RequestCOSUpload{}

	assert.Equal(t, "", req.MediaType)
	assert.Equal(t, "", req.MediaData)
}

func TestRequestCOSUpload_Image(t *testing.T) {
	req := &request.RequestCOSUpload{
		MediaType: "image/jpeg",
		MediaData: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==",
	}

	assert.Equal(t, "image/jpeg", req.MediaType)
}

func TestRequestCOSUpload_Video(t *testing.T) {
	req := &request.RequestCOSUpload{
		MediaType: "video/mp4",
		MediaData: "video_base64_data",
	}

	assert.Equal(t, "video/mp4", req.MediaType)
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

func TestRequestSendMsg_Empty(t *testing.T) {
	req := &request.RequestSendMsg{}

	assert.Equal(t, "", req.OpenID)
	assert.Equal(t, "", req.MsgType)
	assert.Equal(t, "", req.Content)
}

func TestRequestSendMsg_Image(t *testing.T) {
	req := &request.RequestSendMsg{
		OpenID:  "oYYYYY",
		MsgType: "image",
		Content: "media_id_123",
	}

	assert.Equal(t, "oYYYYY", req.OpenID)
	assert.Equal(t, "image", req.MsgType)
	assert.Equal(t, "media_id_123", req.Content)
}

func TestRequestSendMsg_LongContent(t *testing.T) {
	longContent := "这是一条很长的消息内容，用于测试消息长度限制的情况。在实际业务场景中，客服消息可能会有较长的文本内容，需要确保结构体能够正常处理。"
	req := &request.RequestSendMsg{
		OpenID:  "oZZZZZ",
		MsgType: "text",
		Content: longContent,
	}

	assert.Equal(t, "oZZZZZ", req.OpenID)
	assert.Equal(t, longContent, req.Content)
}

func TestResponseCOSUpload(t *testing.T) {
	resp := &response.ResponseCOSUpload{
		MediaID: "media_id_123",
		URL:     "https://example.com/image.jpg",
	}

	assert.Equal(t, "media_id_123", resp.MediaID)
	assert.Equal(t, "https://example.com/image.jpg", resp.URL)
}

func TestResponseCOSUpload_Empty(t *testing.T) {
	resp := &response.ResponseCOSUpload{}

	assert.Equal(t, "", resp.MediaID)
	assert.Equal(t, "", resp.URL)
	assert.Equal(t, 0, resp.ErrCode)
}

func TestResponseCOSUpload_WithError(t *testing.T) {
	resp := &response.ResponseCOSUpload{}
	resp.ErrCode = 40001
	resp.ErrMsg = "invalid media type"

	assert.Equal(t, 40001, resp.ErrCode)
	assert.Equal(t, "invalid media type", resp.ErrMsg)
}

func TestResponseCOSUpload_HTTPSURL(t *testing.T) {
	resp := &response.ResponseCOSUpload{
		MediaID: "media_video_456",
		URL:     "https://mmecimage.cn/video.mp4",
	}

	assert.Equal(t, "media_video_456", resp.MediaID)
	assert.Equal(t, "https://mmecimage.cn/video.mp4", resp.URL)
}

func TestResponseSendMsg(t *testing.T) {
	resp := &response.ResponseSendMsg{}

	assert.Equal(t, 0, resp.ErrCode)
}

func TestResponseSendMsg_Success(t *testing.T) {
	resp := &response.ResponseSendMsg{}
	resp.ErrCode = 0
	resp.ErrMsg = "ok"

	assert.Equal(t, 0, resp.ErrCode)
	assert.Equal(t, "ok", resp.ErrMsg)
}

func TestResponseSendMsg_Fail(t *testing.T) {
	resp := &response.ResponseSendMsg{}
	resp.ErrCode = 45015
	resp.ErrMsg = "response out of time limit or subscription is canceled"

	assert.Equal(t, 45015, resp.ErrCode)
}
