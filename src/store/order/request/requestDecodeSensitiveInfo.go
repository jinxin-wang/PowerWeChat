package request

type RequestDecodeSensitiveInfo struct {
	OrderID       string `json:"order_id"`
	EncryptedData string `json:"encrypted_data"`
}
