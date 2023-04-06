// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Gym-backend/internal/dao/internal"
)

// internalSubscriptionPermissionDao is internal type for wrapping internal DAO implements.
type internalSubscriptionPermissionDao = *internal.SubscriptionPermissionDao

// subscriptionPermissionDao is the data access object for table subscription_permission.
// You can define custom methods on it to extend its functionality as you wish.
type subscriptionPermissionDao struct {
	internalSubscriptionPermissionDao
}

var (
	// SubscriptionPermission is globally public accessible object for table subscription_permission operations.
	SubscriptionPermission = subscriptionPermissionDao{
		internal.NewSubscriptionPermissionDao(),
	}
)

// Fill with you ideas below.