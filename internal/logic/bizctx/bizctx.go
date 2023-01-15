package bizctx

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

// SBizCtx Context is the context for business logic.
type SBizCtx struct{}

// New creates and returns a new context.
func init() {
	service.RegisterBizCtx(New())
}

func New() *SBizCtx {
	return &SBizCtx{}
}

func (s *SBizCtx) Init(r ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)

}
func (s *SBizCtx) Get(ctx context.Context) *model.Context {
	return ctx.Value(consts.ContextKey).(*model.Context)
}

func (s *SBizCtx) GetSession(ctx context.Context) *ghttp.Session {
	return s.Get(ctx).Session
}

func (s *SBizCtx) GetUser(ctx context.Context) *model.ContextUser {
	return s.Get(ctx).User
}

func (s *SBizCtx) SetUser(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}
