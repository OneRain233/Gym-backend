package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadForm struct {
	File *ghttp.UploadFile
	//Name string
	Type string // file type: avatar or facility
}

type FileUploadRes struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}
