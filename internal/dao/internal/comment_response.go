// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommentResponseDao is the data access object for table comment_response.
type CommentResponseDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns CommentResponseColumns // columns contains all the column names of Table for convenient usage.
}

// CommentResponseColumns defines and stores column names for table comment_response.
type CommentResponseColumns struct {
	ID               string // ID
	PARENTRESPONSEID string // PARENT_RESPONSE_ID
	Title            string // title
	Content          string // content
	USERID           string // USER_ID
	CreatTime        string // creat_time
}

// commentResponseColumns holds the columns for table comment_response.
var commentResponseColumns = CommentResponseColumns{
	ID:               "ID",
	PARENTRESPONSEID: "PARENT_RESPONSE_ID",
	Title:            "title",
	Content:          "content",
	USERID:           "USER_ID",
	CreatTime:        "creat_time",
}

// NewCommentResponseDao creates and returns a new DAO object for table data access.
func NewCommentResponseDao() *CommentResponseDao {
	return &CommentResponseDao{
		group:   "default",
		table:   "comment_response",
		columns: commentResponseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CommentResponseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CommentResponseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CommentResponseDao) Columns() CommentResponseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CommentResponseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CommentResponseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CommentResponseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
