package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseDecodeSensitiveInfo struct {
	response.ResponseStore
	ReceiverName string `json:"receiver_name"`
	TelNumber    string `json:"tel_number"`
	DetailInfo   string `json:"detail_info"`
}

type SubOrderInfo struct {
	SubOrderID string `json:"sub_order_id"`
	Status     int    `json:"status"`
	OpenID     string `json:"openid"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type ResponseGetGiftOrderSubList struct {
	response.ResponseStore
	SubOrderList []SubOrderInfo `json:"sub_order_list"`
	NextKey      string         `json:"next_key"`
	HasMore      bool           `json:"has_more"`
}

type SKUChangeInfo struct {
	SKUChangeID string `json:"sku_change_id"`
	OrderID     string `json:"order_id"`
	ProductID   string `json:"product_id"`
	OldSkuID    string `json:"old_sku_id"`
	NewSkuID    string `json:"new_sku_id"`
	CreateTime  string `json:"create_time"`
	Status      int    `json:"status"`
}

type ResponseGetSKUChangeList struct {
	response.ResponseStore
	SKUChangeList []SKUChangeInfo `json:"sku_change_list"`
	NextKey       string          `json:"next_key"`
	HasMore       bool            `json:"has_more"`
}
