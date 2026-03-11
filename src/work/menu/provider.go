package menu

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	client, err := NewClient(app)

	return client, err

}
