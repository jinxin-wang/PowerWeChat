package reverse

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, error) {

	return NewClient(app)

}
