package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Announcement = cAnnouncement{}
var AnnouncementAdmin = cAnnouncementAdmin{}

type cAnnouncement struct{}
type cAnnouncementAdmin struct{}

func (c *cAnnouncement) GetAllAnnouncement(ctx context.Context, req *v1.GetAnnouncementsReq) (res *v1.GetAnnouncementsRes, err error) {
	res = &v1.GetAnnouncementsRes{}
	res.Announcement, err = service.Announcement().GetAnnouncements(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cAnnouncementAdmin) AddAnnouncement(ctx context.Context, req *v1.AddAnnouncementReq) (res *v1.AddAnnouncementRes, err error) {
	res = &v1.AddAnnouncementRes{}

	form := model.AddAnnouncement{
		Title:   req.Title,
		Content: req.Content,
		Images:  req.Images,
	}
	err = service.Announcement().AddAnnouncement(ctx, &form)
	if err != nil {
		return
	}
	return
}

func (c *cAnnouncementAdmin) DeleteAnnouncement(ctx context.Context, req *v1.DeleteAnnouncementReq) (res *v1.DeleteAnnouncementRes, err error) {
	res = &v1.DeleteAnnouncementRes{}
	err = service.Announcement().DeleteAnnouncement(ctx, req.Id)
	if err != nil {
		return
	}
	return
}

func (c *cAnnouncementAdmin) ModifyAnnouncement(ctx context.Context, req *v1.ModifyAnnouncementReq) (res *v1.ModifyAnnouncementRes, err error) {
	res = &v1.ModifyAnnouncementRes{}
	form := model.ModifyAnnouncement{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
		Images:  req.Images,
	}
	err = service.Announcement().ModifyAnnouncement(ctx, &form)
	if err != nil {
		return
	}
	return
}
