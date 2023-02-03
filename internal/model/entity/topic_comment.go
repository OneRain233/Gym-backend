// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TopicComment is the golang structure for table topic_comment.
type TopicComment struct {
	Id              int         `json:"id"              ` // ID
	UserId          int         `json:"userId"          ` // USER_ID
	TopicId         int         `json:"topicId"         ` // TOPIC_ID
	ParentCommentId int         `json:"parentCommentId" ` // Parent comment ID
	Title           string      `json:"title"           ` // Comment_title
	Content         string      `json:"content"         ` // Comment_conrtent
	Approve         string      `json:"approve"         ` // comment_approve
	Disapprove      string      `json:"disapprove"      ` // comment_disapprove
	UpdateTime      *gtime.Time `json:"updateTime"      ` // creat_time
	Images          string      `json:"images"          ` // Image urls
}