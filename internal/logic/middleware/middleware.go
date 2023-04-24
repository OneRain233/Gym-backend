package middleware

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"Gym-backend/utility/response"

	"github.com/gogf/gf/v2/os/gctx"

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

// ResponseHandler used to handle all response
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	ctx := gctx.New()
	g.Log().Info(ctx, r.URL.Path, r.Response.Status)
	r.Middleware.Next()

	// if response has been written, skip it
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
	msg := ""
	// check if the code is 52 (database error)
	// there exist some sensitive information in the error message
	// filter it
	// TODO: check if there are more sensitive information in other error code
	if code.Code() == 52 {
		//msg = "Database error"
		msg = err.Error()
	} else {
		msg = err.Error()
	}
	response.Jsonify(r, code.Code(), msg)
}

// AuthHandler used to check if user have logged in
func (s *sMiddleware) AuthHandler(r *ghttp.Request) {
	user := service.Session().GetUser(r.Context())
	//fmt.Println("Current user:", user.Id)
	if user.Id == 0 {
		response.Jsonify(r, gcode.CodeNotAuthorized.Code(), "Unauthorized")
	}
	r.Middleware.Next()
}

// AdminAuthHandler used to handle admin user auth
func (s *sMiddleware) AdminAuthHandler(r *ghttp.Request) {
	user := service.Session().GetUser(r.Context())
	if user.Role != consts.RoleAdmin {
		response.Jsonify(r, gcode.CodeNotAuthorized.Code(), "Unauthorized")
	}
	r.Middleware.Next()
}

// ManagerAuthHandler used to handle manager user auth
func (s *sMiddleware) ManagerAuthHandler(r *ghttp.Request) {
	user := service.Session().GetUser(r.Context())
	if user.Role != consts.RoleManager && user.Role != consts.RoleAdmin {
		response.Jsonify(r, gcode.CodeNotAuthorized.Code(), "Unauthorized")
	}
	r.Middleware.Next()
}

// CorsHandler used to handle cors
func (s *sMiddleware) CorsHandler(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
