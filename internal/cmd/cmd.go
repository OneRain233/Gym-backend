package cmd

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"Gym-backend/utility/cron"
	"context"
	"encoding/base64"
	"fmt"
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
			//oai := s.GetOpenApi()
			//oai.Info.Title = `API Reference`
			//oai.Config.CommonResponse = response.JsonRes{}
			//oai.Config.CommonResponseDataField = `Data`

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
				for k, _ := range customConfigs.Map() {
					value := g.Cfg().MustGet(gctx.New(), "customConfig."+k+".Value").String()
					valType := g.Cfg().MustGet(gctx.New(), "customConfig."+k+".Type").String()
					fmt.Println(k, value, valType)
					config := &model.Config{
						Key:   k,
						Value: strings.TrimSpace(gconv.String(value)),
						Type:  valType,
					}
					err := service.Config().UpdateConfig(gctx.New(), config)
					if err != nil {
						g.Log().Fatal(gctx.New(), err)
					}
				}
			}

			// check if default avatar file exists
			avatarPath := g.Cfg().MustGet(gctx.New(), "upload.path").String() + "avatar/" + g.Cfg().MustGet(gctx.New(), "upload.defaultAvatar").String()
			if !gfile.Exists(avatarPath) {
				g.Log().Info(gctx.New(), "Default avatar file not found, generating new one...")
				err := gfile.Mkdir(g.Cfg().MustGet(gctx.New(), "upload.path").String() + "avatar/")
				if err != nil {
					g.Log().Fatal(gctx.New(), err)
				}
				defaultAvatarBase64 := g.Cfg().MustGet(gctx.New(), "upload.defaultAvatarBase64").String()
				// decode base64 string to image
				rawData, err := base64.StdEncoding.DecodeString(defaultAvatarBase64)
				if err != nil {
					g.Log().Fatal(gctx.New(), err)
				}
				g.Log().Info(gctx.New(), "Default avatar file generated successfully")
				err = gfile.PutBytes(avatarPath, gconv.Bytes(rawData))
				if err != nil {
					g.Log().Fatal(gctx.New(), err)
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

			// cron job
			err = cron.OrderExpired()
			if err != nil {
				g.Log().Fatal(gctx.New(), err)
			}
			err = cron.SubscriptionExpired()
			if err != nil {
				g.Log().Fatal(gctx.New(), err)
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
					controller.File,
					controller.User,
					controller.Facility,
					controller.Bank,
					controller.Announcement,
					controller.SubscriptionUnauthorized,
					controller.Tag,
					controller.EvaluationUnauth,
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
						controller.Wallet,
						controller.Credit,

						controller.Subscription,
						controller.Evaluation,
						controller.Course,
					)
				})

				// Manager user
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middleware().ResponseHandler,
						service.Middleware().AuthHandler,
						service.Middleware().ManagerAuthHandler,
					)
					group.Bind(
						controller.OrderManager,
						controller.UserManger,
						controller.Visualize,
						controller.SubscriptionAdmin,
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
						controller.BankAdmin, //TODO: Move this to do auth
						controller.FacilityAdmin,
						controller.PaymentAdmin,
						controller.OrderAdmin,
						controller.AnnouncementAdmin,
						controller.Config,
						controller.UserAdmin,
						controller.CreditAdmin,
						controller.EvaluationAdmin,
						controller.TagAdmin,
						controller.System,
						controller.CourseAdmin,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
