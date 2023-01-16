package v1

import "github.com/gogf/gf/v2/frame/g"

type ProfileReq struct {
	g.Meta `path:"/profile" tags:"Profile" method:"get" summary:"Profile"`
}
