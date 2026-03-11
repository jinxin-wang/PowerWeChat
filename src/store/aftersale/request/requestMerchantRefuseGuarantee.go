package request

type RequestMerchantRefuseGuarantee struct {
	GuaranteeID string `json:"guarantee_id"`
	Reason      string `json:"reason"`
}
