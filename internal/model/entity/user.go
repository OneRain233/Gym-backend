// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	ID         int         `json:"iD"         ` // USER ID
	USERNAME   string      `json:"uSERNAME"   ` // USERNAME
	PASSWORD   string      `json:"pASSWORD"   ` // PASSWORD
	AVATAR     string      `json:"aVATAR"     ` // AVATAR URL
	EMAIL      string      `json:"eMAIL"      ` // EMAIL ADDRESS
	PHONE      string      `json:"pHONE"      ` // PHONE NUMBER
	GENDER     int         `json:"gENDER"     ` // GENDER
	CREATETIME *gtime.Time `json:"cREATETIME" ` // CREATE AT
	ROLE       int         `json:"rOLE"       ` // ROLE
	LOGINTIME  *gtime.Time `json:"lOGINTIME"  ` // LAST LOGIN TIME
}
