// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TopicApprovesDao is the data access object for table topic_approves.
type TopicApprovesDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns TopicApprovesColumns // columns contains all the column names of Table for convenient usage.
}

// TopicApprovesColumns defines and stores column names for table topic_approves.
type TopicApprovesColumns struct {
	ID         string // ID
	USERID     string // USER_ID
	TOPICID    string // TOPIC_ID
	TYPE       string // TYPE
	UPDATETIME string // create_time
}

// topicApprovesColumns holds the columns for table topic_approves.
var topicApprovesColumns = TopicApprovesColumns{
	ID:         "ID",
	USERID:     "USER_ID",
	TOPICID:    "TOPIC_ID",
	TYPE:       "TYPE",
	UPDATETIME: "UPDATE_TIME",
}

// NewTopicApprovesDao creates and returns a new DAO object for table data access.
func NewTopicApprovesDao() *TopicApprovesDao {
	return &TopicApprovesDao{
		group:   "default",
		table:   "topic_approves",
		columns: topicApprovesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TopicApprovesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TopicApprovesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TopicApprovesDao) Columns() TopicApprovesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TopicApprovesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TopicApprovesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TopicApprovesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
