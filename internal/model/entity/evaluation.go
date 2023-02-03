// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Evaluation is the golang structure for table evaluation.
type Evaluation struct {
	Id          int    `json:"id"          ` // ID
	UserId      int    `json:"userId"      ` // User id
	FacilityId  int    `json:"facilityId"  ` // Facility Id
	Score       string `json:"score"       ` // Score
	Description string `json:"description" ` // Description
	Anonymous   string `json:"anonymous"   ` // IS Anonymous
	Images      string `json:"images"      ` // Image urls
	Videos      string `json:"videos"      ` // Videos urls
}