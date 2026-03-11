package complaint

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/complaint/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/complaint/response"
)

func TestRequestAddComplaintMaterial(t *testing.T) {
	req := &request.RequestAddComplaintMaterial{
		ComplaintID: "complaint_123",
		Content:     "补充说明内容",
	}

	assert.Equal(t, "complaint_123", req.ComplaintID)
	assert.Equal(t, "补充说明内容", req.Content)
}

func TestRequestAddComplaintMaterial_Empty(t *testing.T) {
	req := &request.RequestAddComplaintMaterial{}

	assert.Equal(t, "", req.ComplaintID)
	assert.Equal(t, "", req.Content)
}

func TestRequestAddComplaintMaterial_LongContent(t *testing.T) {
	longContent := "这是一个很长的补充说明内容，用于描述纠纷的详细情况。商家需要向平台补充说明关于此纠纷的具体信息，包括订单情况、沟通记录、商品状态等。"
	req := &request.RequestAddComplaintMaterial{
		ComplaintID: "complaint_long_123",
		Content:     longContent,
	}

	assert.Equal(t, "complaint_long_123", req.ComplaintID)
	assert.Equal(t, longContent, req.Content)
}

func TestRequestAddComplaintProof(t *testing.T) {
	req := &request.RequestAddComplaintProof{
		ComplaintID: "complaint_456",
		ProofInfo:   map[string]interface{}{"type": "image"},
	}

	assert.Equal(t, "complaint_456", req.ComplaintID)
}

func TestRequestAddComplaintProof_Empty(t *testing.T) {
	req := &request.RequestAddComplaintProof{}

	assert.Equal(t, "", req.ComplaintID)
	assert.Equal(t, nil, req.ProofInfo)
}

func TestRequestAddComplaintProof_WithSlice(t *testing.T) {
	req := &request.RequestAddComplaintProof{
		ComplaintID: "complaint_789",
		ProofInfo: []map[string]string{
			{"url": "https://example.com/proof1.jpg", "desc": "物流凭证"},
			{"url": "https://example.com/proof2.jpg", "desc": "聊天记录"},
		},
	}

	assert.Equal(t, "complaint_789", req.ComplaintID)
	proofSlice, ok := req.ProofInfo.([]map[string]string)
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, len(proofSlice))
}

func TestRequestAddComplaintProof_WithString(t *testing.T) {
	req := &request.RequestAddComplaintProof{
		ComplaintID: "complaint_string",
		ProofInfo:   "简单的举证说明文本",
	}

	assert.Equal(t, "complaint_string", req.ComplaintID)
	proofStr, ok := req.ProofInfo.(string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "简单的举证说明文本", proofStr)
}

func TestRequestGetComplaintOrder(t *testing.T) {
	req := &request.RequestGetComplaintOrder{
		ComplaintID: "complaint_789",
	}

	assert.Equal(t, "complaint_789", req.ComplaintID)
}

func TestRequestGetComplaintOrder_Empty(t *testing.T) {
	req := &request.RequestGetComplaintOrder{}

	assert.Equal(t, "", req.ComplaintID)
}

func TestRequestGetComplaintOrder_InvalidID(t *testing.T) {
	req := &request.RequestGetComplaintOrder{
		ComplaintID: "invalid_id_format",
	}

	assert.Equal(t, "invalid_id_format", req.ComplaintID)
}

func TestResponseAddComplaintMaterial(t *testing.T) {
	resp := &response.ResponseAddComplaintMaterial{}

	assert.Equal(t, 0, resp.ErrCode)
}

func TestResponseAddComplaintMaterial_Success(t *testing.T) {
	resp := &response.ResponseAddComplaintMaterial{}
	resp.ErrCode = 0
	resp.ErrMsg = "ok"

	assert.Equal(t, 0, resp.ErrCode)
	assert.Equal(t, "ok", resp.ErrMsg)
}

func TestResponseAddComplaintMaterial_Fail(t *testing.T) {
	resp := &response.ResponseAddComplaintMaterial{}
	resp.ErrCode = 40001
	resp.ErrMsg = "complaint not found"

	assert.Equal(t, 40001, resp.ErrCode)
}

func TestResponseAddComplaintProof(t *testing.T) {
	resp := &response.ResponseAddComplaintProof{}

	assert.Equal(t, 0, resp.ErrCode)
}

func TestResponseAddComplaintProof_Success(t *testing.T) {
	resp := &response.ResponseAddComplaintProof{}
	resp.ErrCode = 0
	resp.ErrMsg = "ok"

	assert.Equal(t, 0, resp.ErrCode)
}

func TestResponseAddComplaintProof_InvalidProof(t *testing.T) {
	resp := &response.ResponseAddComplaintProof{}
	resp.ErrCode = 40002
	resp.ErrMsg = "invalid proof format"

	assert.Equal(t, 40002, resp.ErrCode)
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

func TestResponseGetComplaintOrder_Empty(t *testing.T) {
	resp := &response.ResponseGetComplaintOrder{}

	assert.Equal(t, (*response.ComplaintInfo)(nil), resp.ComplaintInfo)
	assert.Equal(t, 0, resp.ErrCode)
}

func TestResponseGetComplaintOrder_NotFound(t *testing.T) {
	resp := &response.ResponseGetComplaintOrder{}
	resp.ErrCode = 40003
	resp.ErrMsg = "complaint order not found"

	assert.Equal(t, 40003, resp.ErrCode)
	assert.Equal(t, nil, resp.ComplaintInfo)
}

func TestComplaintInfo(t *testing.T) {
	info := &response.ComplaintInfo{
		ComplaintID:     "complaint_test",
		OrderID:         "order_test",
		OpenID:          "oABCDEFG",
		ComplaintType:   2,
		ComplaintDetail: "未收到货",
		ComplaintTime:   "2024-06-15 14:30:00",
		Status:          1,
	}

	assert.Equal(t, "complaint_test", info.ComplaintID)
	assert.Equal(t, "order_test", info.OrderID)
	assert.Equal(t, "oABCDEFG", info.OpenID)
	assert.Equal(t, 2, info.ComplaintType)
	assert.Equal(t, "未收到货", info.ComplaintDetail)
	assert.Equal(t, "2024-06-15 14:30:00", info.ComplaintTime)
	assert.Equal(t, 1, info.Status)
}

func TestComplaintInfo_Refund(t *testing.T) {
	info := &response.ComplaintInfo{
		ComplaintID:     "complaint_refund",
		OrderID:         "order_refund",
		OpenID:          "o123456",
		ComplaintType:   3,
		ComplaintDetail: "申请退款",
		ComplaintTime:   "2024-07-20 09:00:00",
		Status:          2,
	}

	assert.Equal(t, 3, info.ComplaintType)
	assert.Equal(t, 2, info.Status)
}

func TestComplaintInfo_Exchange(t *testing.T) {
	info := &response.ComplaintInfo{
		ComplaintID:     "complaint_exchange",
		OrderID:         "order_exchange",
		OpenID:          "o789012",
		ComplaintType:   4,
		ComplaintDetail: "换货申请",
		ComplaintTime:   "2024-08-01 16:45:00",
		Status:          0,
	}

	assert.Equal(t, 4, info.ComplaintType)
	assert.Equal(t, 0, info.Status)
}
