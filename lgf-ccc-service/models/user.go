package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User : ""
type User struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	//ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UserName       string     `json:"userName,omitempty" bson:"userName,omitempty"`
	Name           string     `json:"name" bson:"name,omitempty"`
	LastName       string     `json:"lastName" bson:"lastName,omitempty"`
	FatherName     string     `json:"fatherName" bson:"fatherName,omitempty"`
	SpouseName     string     `json:"spouseName" bson:"spouseName,omitempty"`
	OfficialEmail  string     `json:"officialEmail" bson:"officialEmail,omitempty"`
	JoiningDate    *time.Time `json:"joiningDate" bson:"joiningDate,omitempty"`
	Profile        string     `json:"profile" bson:"profile,omitempty"`
	OrganisationID string     `json:"organisationId" bson:"organisationId,omitempty"`
	Password       string     `json:"password" bson:"password,omitempty"`
	Pass           string     `json:"pass" bson:"-"`
	Role           string     `json:"role" bson:"role,omitempty"`
	Type           string     `json:"type" bson:"type,omitempty"`
	Created        Created    `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated    `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog      []Updated  `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	Token          string     `json:"-" bson:"token,omitempty"`
	Status         string     `json:"status" bson:"status,omitempty"`
	Manager        *MinUser   `json:"manager" bson:"manager,omitempty"`
	EmployeeId     string     `json:"employeeId" bson:"employeeId,omitempty"`
	BranchID       string     `json:"branchId,omitempty" bson:"branchId,omitempty"`
	VehicleID      string     `json:"vehicleId,omitempty" bson:"vehicleId,omitempty"`
	DepartmentID   string     `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	DumbSiteID     string     `json:"dumbSiteId,omitempty" bson:"dumbSiteId,omitempty"`
	DesignationID  string     `json:"designationId,omitempty" bson:"designationId,omitempty"`
	IsAreaAssign   string     `json:"isAreaAssign" bson:"isAreaAssign,omitempty"`
	Area           Address    `json:"area" bson:"area,omitempty"`
	UpdateBioData  `bson:",inline"`
	LicenseNo      string     `json:"licenseNo,omitempty" bson:"licenseNo,omitempty"`
	LicenseType    string     `json:"licenseType" bson:"licenseType,omitempty"`
	LicenseExpiry  *time.Time `json:"licenseExpiry" bson:"licenseExpiry,omitempty"`
	AadharNo       string     `json:"AadharNo" bson:"AadharNo,omitempty"`
}

type MinUser struct {
	Name string `json:"name" bson:"name,omitempty"`
	Id   string `json:"id" bson:"id,omitempty"`
}
type CreatedBy struct {
	Name string `json:"name" bson:"name,omitempty"`
	Id   string `json:"id" bson:"id,omitempty"`
}
type GCUser struct {
	Name string `json:"name" bson:"name,omitempty"`
	Id   string `json:"id" bson:"id,omitempty"`
}
type Citizen struct {
	Name string `json:"name" bson:"name,omitempty"`
	Id   string `json:"id" bson:"id,omitempty"`
}
type PropertyDetail struct {
	Name          string `json:"name" bson:"name,omitempty"`
	HoldingNumber string `json:"holdingNumber" bson:"holdingNumber,omitempty"`
	Id            string `json:"id" bson:"id,omitempty"`
}
type RegisterBy struct {
	Name string `json:"name" bson:"name,omitempty"`
	Id   string `json:"id" bson:"id,omitempty"`
}

// RefUser :""
type RefUser struct {
	User `bson:",inline"`
	Ref  struct {
		Manager      User          `json:"manager" bson:"manager,omitempty"`
		Organisation *Organisation `json:"organisation" bson:"organisation,omitempty"`
		LastLocation *UserLocation `json:"lastLocation" bson:"lastLocation,omitempty"`
		DumbSite     DumpSite      `json:"dumbSite" bson:"dumbSite,omitempty"`
		VehicleID    Vechile       `json:"vehicle,omitempty" bson:"vehicle,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RegistrationUser struct {
	User `bson:",inline"`
	OTP  string `json:"otp"`
}

// UserFilter : ""
type UserFilter struct {
	Status         []string `json:"status"`
	UniqueID       []string `json:"uniqueId"`
	Name           []string `json:"name" bson:"name"`
	OmitID         []string `json:"omitId"`
	OrganisationID []string `json:"organisationId" bson:"organisationId,omitempty"`
	Manager        []string `json:"manager" bson:"manager,omitempty"`
	Type           []string `json:"type" bson:"type"`
	DumbSiteID     []string `json:"dumbSiteId" bson:"dumbSiteId,omitempty"`
	Mobile         []string `json:"mobile" bson:"mobile,omitempty"`

	GetRecentLocation bool `json:"getRecentLocation"`
	Regex             struct {
		Name     string  `json:"name" bson:"name"`
		Mobile   string  `json:"mobile" bson:"mobile"`
		UserName string  `json:"userName" bson:"userName"`
		Manager  MinUser `json:"manager" bson:"manager"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// UserChangePassword : ""
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
type RegistrationCitizen struct {
	Name     string `json:"userName" bson:"userName,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
	Token    string `json:"token" bson:"token,omitempty"`
}

// UserLocation : ""
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

// RefUserLocation :""
type RefUserLocation struct {
	UserLocation `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

// UserLocationFilter : ""
type UserLocationFilter struct {
	Status   []string `json:"status"`
	UserType []string `json:"userType" bson:"userType,omitempty"`
}
type UserUniquinessChk struct {
	Success bool   `json:"success" bson:"success,omitempty"`
	Message string `json:"message" bson:"message,omitempty"`
}

type UpdateBioData struct {
	Mobile        string     `json:"mobile" bson:"mobile,omitempty"`
	Email         string     `json:"email" bson:"email,omitempty"`
	DOB           *time.Time `json:"dob" bson:"dob,omitempty"`
	Address       Address    `json:"address" bson:"address,omitempty"`
	Name          string     `json:"name" bson:"name,omitempty"`
	Gender        string     `json:"gender" bson:"gender,omitempty"`
	LineManager   string     `json:"lineManager" bson:"lineManager,omitempty"`
	ProfileImg    string     `json:"profileImg" bson:"profileImg,omitempty"`
	DepartmentID  string     `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	DesignationID string     `json:"designationId,omitempty" bson:"designationId,omitempty"`
}
