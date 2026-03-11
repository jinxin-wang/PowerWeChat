package request

type RequestAddComplaintProof struct {
	ComplaintID string      `json:"complaint_id"`
	ProofInfo   interface{} `json:"proof_info"`
}
