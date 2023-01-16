package middleware

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct {
	LoginUrl string
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/login",
	}
}

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	service.BizCtx().Init(r, customCtx)
	if userEntity := service.Session().GetUser(r.Context()); userEntity.Id > 0 {
		customCtx.User = &model.ContextUser{
			Id:       uint(userEntity.Id),
			Username: userEntity.Username,
			Email:    userEntity.Email,
			Phone:    userEntity.Phone,
			Avatar:   userEntity.Avatar,
		}
	}
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	r.Middleware.Next()
}

func (s *sMiddleware) Auth(r *ghttp.Request) {

	r.Middleware.Next()
}
