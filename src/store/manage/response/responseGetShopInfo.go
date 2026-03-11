package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ShopBasicInfo struct {
	Nickname    string `json:"nickname"`
	HeadImgUrl  string `json:"headimg_url"`
	SubjectType string `json:"subject_type"`
	Status      string `json:"status"`
	Username    string `json:"username"`
}

type ResponseGetShopBasicInfo struct {
	response.ResponseStore
	Info ShopBasicInfo `json:"info"`
}

type ResponseGetShopH5URL struct {
	response.ResponseStore
	H5URL string `json:"h5_url"`
}

type ResponseGetShopQRCode struct {
	response.ResponseStore
	ShopQrcode string `json:"shop_qrcode"`
}

type ResponseGetShopTagLink struct {
	response.ResponseStore
	TagLink string `json:"tag_link"`
}
