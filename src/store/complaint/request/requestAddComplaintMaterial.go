package request

type RequestAddComplaintMaterial struct {
	ComplaintID string `json:"complaint_id"`
	Content     string `json:"content"`
}
