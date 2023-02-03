// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Gym-backend/internal/dao/internal"
)

// internalCreditDao is internal type for wrapping internal DAO implements.
type internalCreditDao = *internal.CreditDao

// creditDao is the data access object for table credit.
// You can define custom methods on it to extend its functionality as you wish.
type creditDao struct {
	internalCreditDao
}

var (
	// Credit is globally public accessible object for table credit operations.
	Credit = creditDao{
		internal.NewCreditDao(),
	}
)

// Fill with you ideas below.