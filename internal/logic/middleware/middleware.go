package middleware

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"Gym-backend/utility/response"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/errors/gcode"

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

func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
	} else {
		response.Jsonify(r, code.Code(), "success", res)
	}

	response.Jsonify(r, code.Code(), err.Error())
}
