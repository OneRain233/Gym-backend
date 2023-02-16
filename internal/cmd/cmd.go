package cmd

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"Gym-backend/utility/response"
	"context"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"

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
			oai := s.GetOpenApi()
			oai.Info.Title = `API Reference`
			oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

			// set time zone
			err = gtime.SetTimeZone(g.Cfg().MustGet(gctx.New(), "timeZone").String())
			if err != nil {
				g.Log().Fatal(gctx.New(), err)
			}
			// update config in database
			// to update the config database from file, just edit the config.yaml file and run the command
			// customConfigUpdate: 1
			// Note: this operation will overwrite the config in database
			status := g.Cfg().MustGet(gctx.New(), "customConfigUpdate").Uint()
			if status == 1 {
				customConfigs, _ := g.Cfg().Get(gctx.New(), "customConfig")
				for k, v := range customConfigs.Map() {
					config := &model.Config{
						Key:   k,
						Value: strings.TrimSpace(gconv.String(v)),
					}
					err := service.Config().UpdateConfig(gctx.New(), config)
					if err != nil {
						g.Log().Fatal(gctx.New(), err)
					}
				}
			}

			// register static files
			uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path").String()
			if !gfile.Exists(uploadPath) {
				err := gfile.Mkdir(uploadPath)
				if err != nil {
					g.Log().Fatal(gctx.New(), err)
				}
			}
			s.AddStaticPath("/uploads", uploadPath)

			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(
					//ghttp.MiddlewareHandlerResponse,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
					service.Middleware().CorsHandler,
				)
				// Unauthorized user
				group.Bind(
					controller.User,
					controller.File, // TODO: Move this to do auth
					controller.Facility,
					controller.Bank,
					controller.BankAdmin, //TODO: Move this to do auth
				)

				// Normal user
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middleware().ResponseHandler,
						service.Middleware().AuthHandler,
					)
					group.Bind(
						controller.Profile,
						controller.Order,
						controller.Payment,
						controller.Card,
						controller.Announcement,
						controller.Wallet,
					)
				})

				// Admin user
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middleware().ResponseHandler,
						service.Middleware().AuthHandler,
						service.Middleware().AdminAuthHandler,
					)
					group.Bind(
						controller.FacilityAdmin,
						controller.PaymentAdmin,
						controller.OrderAdmin,
						controller.AnnouncementAdmin,
						controller.Config,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
