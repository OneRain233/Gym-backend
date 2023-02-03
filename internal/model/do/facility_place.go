// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FacilityPlace is the golang structure of table facility_place for DAO operations like Where/Data.
type FacilityPlace struct {
	g.Meta      `orm:"table:facility_place, do:true"`
	Id          interface{} // ID
	FacilityId  interface{} // Facility ID
	Name        interface{} // Place Name
	Description interface{} // Place Description
	Cost        interface{} // Place order code
}
