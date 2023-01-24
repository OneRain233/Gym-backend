package cmd

import (
	"Gym-backend/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"Gym-backend/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// register static files
			uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path").String()
			fmt.Println(uploadPath)
			if !gfile.Exists(uploadPath) {
				err := gfile.Mkdir(uploadPath)
				if err != nil {
					g.Log().Fatal(gctx.New(), err)
				}
			}
			s.AddStaticPath("/uploads", uploadPath)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					//ghttp.MiddlewareHandlerResponse,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					//controller.Hello,
					controller.Login,
					controller.Register,
					controller.File, // TODO: Move this to do auth
				)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().ResponseHandler, service.Middleware().AuthHandler)
					group.Bind(
						controller.Hello,
						controller.Facility,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
