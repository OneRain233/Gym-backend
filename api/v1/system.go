package v1

import "github.com/gogf/gf/v2/frame/g"

type GetLogReq struct {
	g.Meta `path:"/logs" method:"post" tags:"System" summary:"Get logs"`
	Date   string `json:"date"`
}

type GetLogRes struct {
	g.Meta  `mime:"application/json" example:"string"`
	SysLogs string `json:"sysLogs"`
	SqlLogs string `json:"sqlLogs"`
}
