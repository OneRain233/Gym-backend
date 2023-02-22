package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// FileUploadReq File Upload api
type FileUploadReq struct {
	g.Meta `path:"/upload" method:"post" mime:"multipart/form-data" tags:"File" summary:"Upload file"`
	File   *ghttp.UploadFile `json:"file" v:"required#Please select file to upload"`
	Name   string            `json:"name" v:"required#Please input file name"`
	Type   string            `json:"type" v:"required#Please input file type"`
}

// FileUploadRes File Upload response
type FileUploadRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Url    string `json:"url"`
	Name   string `json:"name"`
}
