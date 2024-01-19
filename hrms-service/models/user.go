package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User : ""
type User struct {
	UniqueID       string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UserName       string             `json:"userName,omitempty" bson:"userName,omitempty"`
	LoginId        string             `json:"loginId" bson:"loginId,omitempty"`
	Grade          string             `json:"grade,omitempty" bson:"grade,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	LastName       string             `json:"lastName" bson:"lastName,omitempty"`
	FatherName     string             `json:"fatherName" bson:"fatherName,omitempty"`
	SpouseName     string             `json:"spouseName" bson:"spouseName,omitempty"`
	OfficialEmail  string             `json:"officialEmail" bson:"officialEmail,omitempty"`
	JoiningDate    *time.Time         `json:"joiningDate" bson:"joiningDate,omitempty"`
	Profile        string             `json:"profile" bson:"profile,omitempty"`
	OrganisationID string             `json:"organisationId" bson:"organisationId,omitempty"`
	Password       string             `json:"-" bson:"password,omitempty"`
	Pass           string             `json:"password" bson:"-"`
	Role           string             `json:"role" bson:"role,omitempty"`
	Type           string             `json:"type" bson:"type,omitempty"`
	Created        Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog      []Updated          `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	Token          string             `json:"-" bson:"token,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	ManagerId      string             `json:"managerId" bson:"managerId,omitempty"`
	EmployeeId     string             `json:"employeeId" bson:"employeeId,omitempty"`
	BranchID       string             `json:"branchId,omitempty" bson:"branchId,omitempty"`
	DepartmentID   string             `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	DesignationID  string             `json:"designationId,omitempty" bson:"designationId,omitempty"`
	UpdateBioData  `bson:",inline"`
}

//RefUser :""
type RefUser struct {
	User `bson:",inline"`
	Ref  struct {
		Manager      User          `json:"manager" bson:"manager,omitempty"`
		Organisation *Organisation `json:"organisation" bson:"organisation,omitempty"`
		LastLocation *UserLocation `json:"lastLocation" bson:"lastLocation,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserFilter : ""
type UserFilter struct {
	Status            []string `json:"status"`
	UniqueID          []string `json:"uniqueId"`
	OmitID            []string `json:"omitId"`
	OrganisationID    []string `json:"organisationId" bson:"organisationId,omitempty"`
	Manager           []string `json:"manager" bson:"manager,omitempty"`
	Type              []string `json:"type" bson:"type"`
	GetRecentLocation bool     `json:"getRecentLocation"`
	Regex             struct {
		Name     string `json:"name" bson:"name"`
		Contact  string `json:"contact" bson:"contact"`
		UserName string `json:"userName" bson:"userName"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

//UserChangePassword : ""
type UserChangePassword struct {
	UserName    string `json:"userName" bson:"userName,omitempty"`
	OldPassword string `json:"oldPassword" bson:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword" bson:"newPassword,omitempty"`
}

type UserNewPassword struct {
	UserName string `json:"userName" bson:"userName,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
	Token    string `json:"token" bson:"token,omitempty"`
}

//UserLocation : ""
type UserLocation struct {
	ID        primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID  string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status    string             `json:"status" bson:"status,omitempty"`
	UserName  string             `json:"userName" bson:"userName,omitempty"`
	Time      *time.Time         `json:"time" bson:"time,omitempty"`
	Location  Location           `json:"location" bson:"location,omitempty"`
	UserType  string             `json:"userType" bson:"userType,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	EntryType string             `json:"entryType" bson:"entryType,omitempty"`
	ErrMsg    string             `json:"errMsg" bson:"errMsg,omitempty"`
}

//RefUserLocation :""
type RefUserLocation struct {
	UserLocation `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserLocationFilter : ""
type UserLocationFilter struct {
	Status   []string `json:"status"`
	UserType []string `json:"userType" bson:"userType,omitempty"`
}
type UserUniquinessChk struct {
	Success bool   `json:"success" bson:"success,omitempty"`
	Message string `json:"message" bson:"message,omitempty"`
}
