package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

var System = cSystem{}

type cSystem struct{}

func (c *cSystem) GetLogs(ctx context.Context, req *v1.GetLogReq) (res *v1.GetLogRes, err error) {
	res = &v1.GetLogRes{}
	dateString := req.Date
	date, err := gtime.StrToTime(dateString)
	if err != nil {
		return
	}
	res.SysLogs, err = service.System().GetLogs(ctx, date)
	if err != nil {
		return
	}
	res.SqlLogs, err = service.System().GetSqlLogs(ctx, date)
	if err != nil {
		return
	}
	return
}
