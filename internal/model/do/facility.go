// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Facility is the golang structure of table facility for DAO operations like Where/Data.
type Facility struct {
	g.Meta      `orm:"table:facility, do:true"`
	ID          interface{} // Facility ID
	NAME        interface{} // Facility Name
	DESCRIPTION interface{} // Facility Description
	LOCATION    interface{} // Facility Location
	IMAGES      interface{} // Image urls
}
