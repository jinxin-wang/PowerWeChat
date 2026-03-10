package manage

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRequestGetShopQRCode(t *testing.T) {
	req := &RequestGetShopQRCode{
		WeComCorpId: "corp123",
		WeComUserId: "user456",
	}

	assert.Equal(t, "corp123", req.WeComCorpId)
	assert.Equal(t, "user456", req.WeComUserId)
}

func TestRequestGetShopQRCode_Empty(t *testing.T) {
	req := &RequestGetShopQRCode{}

	assert.Equal(t, "", req.WeComCorpId)
	assert.Equal(t, "", req.WeComUserId)
}

func TestRequestGetShopTagLink(t *testing.T) {
	req := &RequestGetShopTagLink{
		WeComCorpId: "corp789",
		WeComUserId: "user012",
	}

	assert.Equal(t, "corp789", req.WeComCorpId)
	assert.Equal(t, "user012", req.WeComUserId)
}

func TestRequestGetShopTagLink_Empty(t *testing.T) {
	req := &RequestGetShopTagLink{}

	assert.Equal(t, "", req.WeComCorpId)
	assert.Equal(t, "", req.WeComUserId)
}

func TestResponseGetShopBasicInfo(t *testing.T) {
	resp := &ResponseGetShopBasicInfo{
		Info: ShopBasicInfo{
			Nickname:    "Test Shop",
			HeadImgUrl:  "https://example.com/logo.png",
			SubjectType: "1",
			Status:      "2",
			Username:    "test_user",
		},
	}

	assert.Equal(t, "Test Shop", resp.Info.Nickname)
	assert.Equal(t, "https://example.com/logo.png", resp.Info.HeadImgUrl)
	assert.Equal(t, "1", resp.Info.SubjectType)
	assert.Equal(t, "2", resp.Info.Status)
	assert.Equal(t, "test_user", resp.Info.Username)
}

func TestResponseGetShopH5URL(t *testing.T) {
	resp := &ResponseGetShopH5URL{
		H5URL: "https://shop.example.com",
	}

	assert.Equal(t, "https://shop.example.com", resp.H5URL)
}

func TestResponseGetShopQRCode(t *testing.T) {
	resp := &ResponseGetShopQRCode{
		ShopQrcode: "https://example.com/qrcode.png",
	}

	assert.Equal(t, "https://example.com/qrcode.png", resp.ShopQrcode)
}

func TestResponseGetShopTagLink(t *testing.T) {
	resp := &ResponseGetShopTagLink{
		TagLink: "https://example.com/tag",
	}

	assert.Equal(t, "https://example.com/tag", resp.TagLink)
}

func TestShopBasicInfo(t *testing.T) {
	info := ShopBasicInfo{
		Nickname:    "My Shop",
		HeadImgUrl:  "https://cdn.example.com/logo.jpg",
		SubjectType: "2",
		Status:      "1",
		Username:    "admin",
	}

	assert.Equal(t, "My Shop", info.Nickname)
	assert.Equal(t, "https://cdn.example.com/logo.jpg", info.HeadImgUrl)
	assert.Equal(t, "2", info.SubjectType)
	assert.Equal(t, "1", info.Status)
	assert.Equal(t, "admin", info.Username)
}
