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

// ResponseHandler used to handle all response
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// if response has been written, skip it
	if r.Response.BufferLength() > 0 {
		// replace the message if the response is error
		code := gerror.Code(r.GetError())
		if code != gcode.CodeOK {
			response.Jsonify(r, code.Code(), "An error occurred, please check the log")
		}
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

	response.Jsonify(r, code.Code(), "An error occurred, please check the log")
}

// AuthHandler used to check if user have logged in
func (s *sMiddleware) AuthHandler(r *ghttp.Request) {
	r.Middleware.Next()
	user := service.Session().GetUser(r.Context())
	//fmt.Println("Current user:", user.Id)
	if user.Id == 0 {
		response.Jsonify(r, gcode.CodeNotAuthorized.Code(), "Unauthorized")
	}
}

// AdminAuthHandler used to handle admin user auth
func (s *sMiddleware) AdminAuthHandler(r *ghttp.Request) {
	r.Middleware.Next()
	user := service.Session().GetUser(r.Context())
	if user.Role != 1 {
		response.Jsonify(r, gcode.CodeNotAuthorized.Code(), "Unauthorized")
	}
}
