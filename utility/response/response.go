package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type JsonRes struct {
	Code     int         `json:"code"`     // error code(0:success, 1:fail, >1:error code)
	Message  string      `json:"message"`  // msg
	Data     interface{} `json:"data"`     // data result
	Redirect string      `json:"redirect"` // redirect url
}

func Jsonify(r *ghttp.Request, code int, message string, data ...interface{}) {
	var res interface{}
	if len(data) > 0 {
		res = data[0]
	} else {
		res = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    res,
	})

	r.Exit()
}
