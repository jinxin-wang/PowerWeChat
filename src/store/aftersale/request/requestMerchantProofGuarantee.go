package request

type ProofInfo struct {
	Text     string   `json:"text,omitempty"`
	ImageIDs []string `json:"image_ids,omitempty"`
}

type RequestMerchantProofGuarantee struct {
	GuaranteeID string     `json:"guarantee_id"`
	ProofInfo   *ProofInfo `json:"proof_info,omitempty"`
}
