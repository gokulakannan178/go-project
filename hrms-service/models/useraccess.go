package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//UserAcl : ""
type UserAcl struct {
	ID                 primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName           primitive.ObjectID `json:"userName" bson:"userName,omitempty"`
	UserType           string             `json:"userType" bson:"userType,omitempty"`
	Organisation       Access             `json:"organisation" bson:"organisation,omitempty"`
	Branch             Access             `json:"branch" bson:"branch,omitempty"`
	Department         Access             `json:"department" bson:"department,omitempty"`
	Designation        Access             `json:"designation" bson:"designation,omitempty"`
	Employee           Access             `json:"employee" bson:"employee,omitempty"`
	Attendance         Access             `json:"attendance" bson:"attendance,omitempty"`
	AdminAttendance    Access             `json:"adminAttendance" bson:"adminAttendance,omitempty"`
	Leave              Access             `json:"leave" bson:"leave,omitempty"`
	AdminLeave         Access             `json:"adminLeave" bson:"adminLeave,omitempty"`
	LeaveSettings      Access             `json:"leaveSettings" bson:"leaveSettings,omitempty"`
	Policy             Access             `json:"policy" bson:"policy,omitempty"`
	Reports            Access             `json:"reports" bson:"reports,omitempty"`
	AssetManagment     Access             `json:"assetManagment" bson:"assetManagment,omitempty"`
	BillClaim          Access             `json:"billClaim" bson:"billClaim,omitempty"`
	DocumentManagement Access             `json:"documentManagement" bson:"documentManagement,omitempty"`
	News               Access             `json:"news" bson:"news,omitempty"`
	PayRoll            Access             `json:"payRoll" bson:"payRoll,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	Created            *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	Updated            []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}
type Access struct {
	Read  string `json:"read" bson:"read,omitempty"`
	Write string `json:"write" bson:"write,omitempty"`
}
type ContentAndQueryAccess struct {
	CreateEdit      string `json:"createEdit" bson:"createEdit,omitempty"`
	BypassContent   string `json:"bypassContent" bson:"bypassContent,omitempty"`
	Manage          string `json:"manage" bson:"manage,omitempty"`
	TranslateReview string `json:"translateReview" bson:"translateReview,omitempty"`
	Upload          string `json:"upload" bson:"upload,omitempty"`
	QueryEdit       string `json:"queryEdit" bson:"queryEdit,omitempty"`
	Delete          string `json:"delete" bson:"delete,omitempty"`
	Review          string `json:"review" bson:"review,omitempty"`
	Translate       string `json:"translate" bson:"translate,omitempty"`
	Disseminate     string `json:"disseminate" bson:"disseminate,omitempty"`
	Search          string `json:"search" bson:"search,omitempty"`
}
type SpecialFeatures struct {
	WeatherData      string `json:"weatherData" bson:"weatherData,omitempty"`
	PickResolveQuery string `json:"pickResolveQuery" bson:"pickResolveQuery,omitempty"`
}

//RefUserAcl : ""
type RefUserAcl struct {
	UserAcl `bson:",inline"`
	Ref     struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserAclFilter : ""
type UserAclFilter struct {
	Status    []string             `json:"status,omitempty" bson:"status,omitempty"`
	UserType  []string             `json:"userType" bson:"userType,omitempty"`
	UserName  []primitive.ObjectID `json:"userName" bson:"userName,omitempty"`
	SortBy    string               `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int                  `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
