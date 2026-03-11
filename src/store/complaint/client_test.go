package complaint

import (
	"testing"

	"github.com/jinxin-wang/PowerWeChat/v3/src/store/complaint/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/complaint/response"
	"github.com/go-playground/assert/v2"
)

func TestRequestAddComplaintMaterial(t *testing.T) {
	req := &request.RequestAddComplaintMaterial{
		ComplaintID: "complaint_123",
		Content:     "补充说明内容",
	}

	assert.Equal(t, "complaint_123", req.ComplaintID)
	assert.Equal(t, "补充说明内容", req.Content)
}

func TestRequestAddComplaintProof(t *testing.T) {
	req := &request.RequestAddComplaintProof{
		ComplaintID: "complaint_456",
		ProofInfo:   map[string]interface{}{"type": "image"},
	}

	assert.Equal(t, "complaint_456", req.ComplaintID)
}

func TestRequestGetComplaintOrder(t *testing.T) {
	req := &request.RequestGetComplaintOrder{
		ComplaintID: "complaint_789",
	}

	assert.Equal(t, "complaint_789", req.ComplaintID)
}

func TestResponseAddComplaintMaterial(t *testing.T) {
	resp := &response.ResponseAddComplaintMaterial{}

	assert.Equal(t, 0, resp.ErrCode)
}

func TestResponseAddComplaintProof(t *testing.T) {
	resp := &response.ResponseAddComplaintProof{}

	assert.Equal(t, 0, resp.ErrCode)
}

func TestResponseGetComplaintOrder(t *testing.T) {
	resp := &response.ResponseGetComplaintOrder{
		ComplaintInfo: &response.ComplaintInfo{
			ComplaintID:     "complaint_001",
			OrderID:         "order_001",
			OpenID:          "oXXXXX",
			ComplaintType:   1,
			ComplaintDetail: "商品质量问题",
			ComplaintTime:   "2024-01-01 10:00:00",
			Status:          0,
		},
	}

	assert.Equal(t, "complaint_001", resp.ComplaintInfo.ComplaintID)
	assert.Equal(t, "order_001", resp.ComplaintInfo.OrderID)
	assert.Equal(t, 1, resp.ComplaintInfo.ComplaintType)
}
