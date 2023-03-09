// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Announcement is the golang structure for table announcement.
type Announcement struct {
	Id         int         `json:"id"         ` // ID
	Title      string      `json:"title"      ` // Title
	Content    string      `json:"content"    ` // Content
	UpdateTime *gtime.Time `json:"updateTime" ` // Update time
	UserId     int         `json:"userId"     ` // User ID
	Delete     int         `json:"delete"     ` //
	Images     string      `json:"images"     ` //
}
