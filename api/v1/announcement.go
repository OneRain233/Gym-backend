package v1

import (
	"Gym-backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetAnnouncementsReq struct {
	g.Meta `path:"/announcement/announcements" method:"get" tags:"Announcement" summary:"Get announcements"`
}

type GetAnnouncementsRes struct {
	g.Meta       `mime:"application/json" example:"string"`
	Announcement []*entity.Announcement `json:"announcement"`
}

type AddAnnouncementReq struct {
	g.Meta  `path:"/announcement/add" method:"post" mime:"application/json" tags:"Announcement" summary:"Add announcement"`
	Title   string   `json:"title" v:"required#Please input announcement title"`
	Content string   `json:"content" v:"required#Please input announcement content"`
	Images  []string `json:"images"`
}

type AddAnnouncementRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type DeleteAnnouncementReq struct {
	g.Meta `path:"/announcement/delete" method:"post" mime:"application/json" tags:"Announcement" summary:"Delete announcement"`
	Id     int `json:"id" v:"required#Please input announcement id"`
}

type DeleteAnnouncementRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type ModifyAnnouncementReq struct {
	g.Meta  `path:"/announcement/modify" method:"post" mime:"application/json" tags:"Announcement" summary:"Modify announcement"`
	Id      int      `json:"id" v:"required#Please input announcement id"`
	Title   string   `json:"title" v:"required#Please input announcement title"`
	Content string   `json:"content" v:"required#Please input announcement content"`
	Images  []string `json:"images"`
}

type ModifyAnnouncementRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
