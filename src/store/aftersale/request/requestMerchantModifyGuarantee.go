package request

type RequestMerchantModifyGuarantee struct {
	GuaranteeID string `json:"guarantee_id"`
	Amount      int    `json:"amount"`
}
