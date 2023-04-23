package file

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

type sFile struct{}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

func (s *sFile) UploadFile(ctx context.Context, input model.FileUploadForm) (output model.FileUploadRes, err error) {
	uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path").String()

	file := input.File
	uploadType := input.Type
	switch uploadType {
	case "avatar":
		uploadPath += "avatar/"
	case "facility":
		uploadPath += "facility/"
	case "course_image":
		uploadPath += "course/image/"
	case "course_video":
		uploadPath += "course/video/"
	case "common":
		uploadPath += "common/"
	default:
		uploadPath += "common/"
	}
	if !gfile.Exists(uploadPath) {
		err = gfile.Mkdir(uploadPath)
		if err != nil {
			return
		}
	}
	filename, err := file.Save(uploadPath, true)
	if err != nil {
		return
	}
	output.URL = uploadPath + filename
	output.Name = filename
	return
}
