package announcement

import (
	_ "Gym-backend/api/v1"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
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

	var images string
	images = strings.Join(input.Images, ",")

	announcement := entity.Announcement{
		Title:      input.Title,
		Content:    input.Content,
		UpdateTime: gtime.Now(),
		UserId:     service.Session().GetUser(ctx).Id,
		//Delete:     0,
		Images: images,
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

func (s *sAnnouncement) ModifyAnnouncement(ctx context.Context, input *model.ModifyAnnouncement) error {
	if input.Id == 0 {
		err := gerror.New("id is empty")
		return err
	}
	var images string
	images = strings.Join(input.Images, ",")
	announcement := entity.Announcement{
		Id:         input.Id,
		Title:      input.Title,
		Content:    input.Content,
		UpdateTime: gtime.Now(),
		UserId:     service.Session().GetUser(ctx).Id,
		//Delete:     0,
		Images: images,
	}
	_, err := dao.Announcement.Ctx(ctx).Where("id", input.Id).Update(announcement)
	if err != nil {
		return err
	}
	return nil

}
