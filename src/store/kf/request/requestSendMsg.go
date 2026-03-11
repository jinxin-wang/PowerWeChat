package request

type RequestSendMsg struct {
	OpenID  string `json:"openid"`
	MsgType string `json:"msg_type"`
	Content string `json:"content"`
}
