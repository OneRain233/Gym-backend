package mail

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, msg string) error {
	m := gomail.NewMessage()
	server := g.Cfg().MustGet(gctx.New(), "mail.host").String()
	port := g.Cfg().MustGet(gctx.New(), "mail.port")
	username := g.Cfg().MustGet(gctx.New(), "mail.username").String()
	password := g.Cfg().MustGet(gctx.New(), "mail.password").String()
	if server == "" || port.String() == "" || username == "" || password == "" {
		return gerror.New("mail config error")
	}
	m.SetHeader("From", username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Gym Booking")
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(server, port.Int(), username, password)
	if err := d.DialAndSend(m); err != nil {
		g.Log().Error(gctx.New(), err)
	} else {
		g.Log().Info(gctx.New(), "email sent")
	}
	return nil
}

func SendEmailWithAttachment(to string, msg string, attachment string) error {
	m := gomail.NewMessage()
	server := g.Cfg().MustGet(gctx.New(), "mail.host").String()
	port := g.Cfg().MustGet(gctx.New(), "mail.port")
	username := g.Cfg().MustGet(gctx.New(), "mail.username").String()
	password := g.Cfg().MustGet(gctx.New(), "mail.password").String()
	if server == "" || port.String() == "" || username == "" || password == "" {
		return gerror.New("mail config error")
	}
	m.SetHeader("From", username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Gym Booking")
	m.SetBody("text/html", msg)
	m.Attach(attachment)

	d := gomail.NewDialer(server, port.Int(), username, password)
	if err := d.DialAndSend(m); err != nil {
		g.Log().Error(gctx.New(), err)
	} else {
		g.Log().Info(gctx.New(), "email sent")
	}
	return nil
}
