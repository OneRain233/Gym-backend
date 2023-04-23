package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// FileUploadReq File Upload api
type FileUploadReq struct {
	g.Meta `path:"/upload" method:"post" mime:"multipart/form-data" tags:"File" summary:"Upload file"`
	File   *ghttp.UploadFile `json:"file" v:"required#Please select file to upload"`
	//Name   string            `json:"name" v:"required#Please input file name"`
	//Type   string            `json:"type" v:"required#Please input file type"`
}

// FileUploadRes File Upload response
type FileUploadRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Url    string `json:"url"`
	Name   string `json:"name"`
}

type AvatarUploadReq struct {
	g.Meta `path:"/upload/avatar" method:"post" mime:"multipart/form-data" tags:"File" summary:"Upload avatar"`
	File   *ghttp.UploadFile `json:"file" v:"required#Please select file to upload"`
}

type AvatarUploadRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Url    string `json:"url"`
	Name   string `json:"name"`
}

type FacilityImageUploadReq struct {
	g.Meta `path:"/upload/facility" method:"post" mime:"multipart/form-data" tags:"File" summary:"Upload facility image"`
	File   *ghttp.UploadFile `json:"file" v:"required#Please select file to upload"`
}

type FacilityImageUploadRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Url    string `json:"url"`
	Name   string `json:"name"`
}

type CourseImageUploadReq struct {
	g.Meta `path:"/upload/course/image" method:"post" mime:"multipart/form-data" tags:"File" summary:"Upload course image"`
	File   *ghttp.UploadFile `json:"file" v:"required#Please select file to upload"`
}

type CourseImageUploadRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Url    string `json:"url"`
	Name   string `json:"name"`
}

type CourseVideoUploadReq struct {
	g.Meta `path:"/upload/course/video" method:"post" mime:"multipart/form-data" tags:"File" summary:"Upload course video"`
	File   *ghttp.UploadFile `json:"file" v:"required#Please select file to upload"`
}

type CourseVideoUploadRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Url    string `json:"url"`
	Name   string `json:"name"`
}
