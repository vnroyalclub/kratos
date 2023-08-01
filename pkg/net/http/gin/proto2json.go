package gin

import (
	"github.com/vnroyalclub/kratos/pkg/ecode"
)

type JSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TTL     int         `json:"ttl"`
	Data    interface{} `json:"data,omitempty"`
}

func TOJSON(resp interface{}, err error) (data JSON) {
	bcode := ecode.Cause(err)

	return JSON{
		Code:    bcode.Code(),
		Message: bcode.Message(),
		TTL:     1,
		Data:    resp,
	}
}
