package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseImmediateDeliveryBindAccount struct {
	response.ResponseMiniProgram

	ShopList []*power.HashMap `json:"shop_list"`
}
