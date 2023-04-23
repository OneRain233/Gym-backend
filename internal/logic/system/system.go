package system

import (
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gtime"
)

type sSystem struct {
}

func New() *sSystem {
	return &sSystem{}
}

func init() {
	service.RegisterSystem(New())
}

func (s *sSystem) GetLogs(ctx context.Context, date *gtime.Time) (res string, err error) {
	logPath := g.Cfg().MustGet(gctx.New(), "logger.path").String()
	// open file
	g.Log().Info(ctx, logPath+date.Format("Y-m-d")+".log")
	res = gfile.GetContents(logPath + date.Format("Y-m-d") + ".log")
	return res, nil
}

func (s *sSystem) GetSqlLogs(ctx context.Context, date *gtime.Time) (res string, err error) {
	logPath := g.Cfg().MustGet(gctx.New(), "database.logger.path").String()
	// open file
	g.Log().Info(ctx, logPath+date.Format("Y-m-d")+".log")
	res = gfile.GetContents(logPath + date.Format("Y-m-d") + ".log")
	return res, nil
}

func (s *sSystem) GetLogsList(ctx context.Context) (res []string, err error) {
	logPath := g.Cfg().MustGet(gctx.New(), "logger.path").String()
	// open file
	res, err = gfile.ScanDir(logPath, "*")
	if err != nil {
		return nil, err
	}
	return res, nil
}
