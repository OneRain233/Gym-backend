package bizctx

import (
	"Gym-backend/internal/service"
	"context"

	"Gym-backend/internal/consts"
	"Gym-backend/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

// sBizCtx Context is the context for business logic.
type sBizCtx struct{}

// New creates and returns a new context.
func init() {
	service.RegisterBizCtx(New())
}

func New() *sBizCtx {
	return &sBizCtx{}
}

func (s *sBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)

}
func (s *sBizCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

func (s *sBizCtx) GetSession(ctx context.Context) *ghttp.Session {
	return s.Get(ctx).Session
}

func (s *sBizCtx) GetUser(ctx context.Context) *model.ContextUser {
	return s.Get(ctx).User
}

func (s *sBizCtx) SetUser(ctx context.Context, user *model.ContextUser) {
	s.Get(ctx).User = user
}
