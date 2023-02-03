// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Gym-backend/internal/dao/internal"
)

// internalCommentResponseDao is internal type for wrapping internal DAO implements.
type internalCommentResponseDao = *internal.CommentResponseDao

// commentResponseDao is the data access object for table comment_response.
// You can define custom methods on it to extend its functionality as you wish.
type commentResponseDao struct {
	internalCommentResponseDao
}

var (
	// CommentResponse is globally public accessible object for table comment_response operations.
	CommentResponse = commentResponseDao{
		internal.NewCommentResponseDao(),
	}
)

// Fill with you ideas below.