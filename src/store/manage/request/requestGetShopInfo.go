package request

type RequestGetShopBasicInfo struct {
}

type RequestGetShopH5URL struct {
}

type RequestGetShopQRCode struct {
	WeComCorpId string `json:"wecom_corp_id,omitempty"`
	WeComUserId string `json:"wecom_user_id,omitempty"`
}

type RequestGetShopTagLink struct {
	WeComCorpId string `json:"wecom_corp_id,omitempty"`
	WeComUserId string `json:"wecom_user_id,omitempty"`
}
