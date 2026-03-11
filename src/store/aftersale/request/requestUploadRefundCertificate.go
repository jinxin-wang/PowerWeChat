package request

type CertificateInfo struct {
	ImageIDs []string `json:"image_ids,omitempty"`
}

type RequestUploadRefundCertificate struct {
	AftersaleID     string           `json:"aftersale_id"`
	CertificateInfo *CertificateInfo `json:"certificate_info,omitempty"`
}
