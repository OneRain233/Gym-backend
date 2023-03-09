package announcement

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

type sAnnouncement struct {
}

func init() {
	service.RegisterAnnouncement(New())
}

func New() *sAnnouncement {
	return &sAnnouncement{}
}

func (s *sAnnouncement) GetAnnouncements(ctx context.Context) (announcements []*entity.Announcement, err error) {
	err = dao.Announcement.Ctx(ctx).Scan(&announcements)
	if err != nil {
		return
	}
	return
}

func (s *sAnnouncement) AddAnnouncement(ctx context.Context, input *model.AddAnnouncement) error {
	announcement := entity.Announcement{
		Title:      input.Title,
		Content:    input.Content,
		UpdateTime: gtime.Now(),
		UserId:     service.Session().GetUser(ctx).Id,
	}
	_, err := dao.Announcement.Ctx(ctx).Insert(announcement)
	if err != nil {
		return err
	}

	return nil
}

func (s *sAnnouncement) DeleteAnnouncement(ctx context.Context, id int) error {
	_, err := dao.Announcement.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
