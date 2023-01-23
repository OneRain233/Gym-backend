package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session *ghttp.Session
	Data    g.Map
	User    *ContextUser
}

type ContextUser struct {
	Id       uint
	Username string
	Email    string
	Phone    string
	Avatar   string // avatar url
	Role     uint   // 0: Normal user, 1: Admin
}
