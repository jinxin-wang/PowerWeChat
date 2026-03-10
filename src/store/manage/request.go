package manage

type RequestGetShopQRCode struct {
	WeComCorpId string `json:"wecom_corp_id,omitempty"`
	WeComUserId string `json:"wecom_user_id,omitempty"`
}

type RequestGetShopTagLink struct {
	WeComCorpId string `json:"wecom_corp_id,omitempty"`
	WeComUserId string `json:"wecom_user_id,omitempty"`
}
