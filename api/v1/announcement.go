package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAnnouncementsReq struct {
	g.Meta `path:"/announcement/announcements" method:"get" tags:"Announcement" summary:"Get announcements"`
}

type GetAnnouncementsRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   []*entity.Announcement `json:"data"`
}

type AddAnnouncementReq struct {
	g.Meta  `path:"/announcement/add" method:"post" mime:"application/json" tags:"Announcement" summary:"Add announcement"`
	Title   string `json:"title" v:"required#Please input announcement title"`
	Content string `json:"content" v:"required#Please input announcement content"`
}

type AddAnnouncementRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
