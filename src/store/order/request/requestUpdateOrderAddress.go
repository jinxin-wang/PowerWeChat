package request

type RequestUpdateOrderAddress struct {
	OrderID      string `json:"order_id"`
	ReceiverName string `json:"receiver_name"`
	DetailInfo   string `json:"detail_info"`
	TelNumber    string `json:"tel_number"`
	PostalCode   string `json:"postal_code,omitempty"`
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	CountyName   string `json:"county_name"`
}
