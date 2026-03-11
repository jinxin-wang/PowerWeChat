package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetStableAccessToken struct {
	response.ResponseStore
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
