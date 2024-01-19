package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User : ""
type User struct {
	ID                           primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	UserName                     string               `json:"userName" bson:"userName,omitempty"`
	Name                         string               `json:"name" bson:"name,omitempty"`
	FirstName                    string               `json:"firstName" bson:"firstName,omitempty"`
	Lastname                     string               `json:"lastname" bson:"lastname,omitempty"`
	Project                      []primitive.ObjectID `json:"project" bson:"_,omitempty"`
	AlternateNumber              string               `json:"alternateNumber" bson:"alternateNumber,omitempty"`
	FatherName                   string               `json:"fatherName" bson:"fatherName,omitempty"`
	Father_HusbandName           string               `json:"father_husbandName" bson:"father_husbandName,omitempty"`
	State                        string               `json:"state" bson:"state,omitempty"`
	StateCode                    primitive.ObjectID   `json:"stateCode" bson:"stateCode,omitempty"`
	District                     string               `json:"district" bson:"district,omitempty"`
	DistrictCode                 primitive.ObjectID   `json:"districtCode" bson:"districtCode,omitempty"`
	City                         string               `json:"city" bson:"city,omitempty"`
	PinCode                      int64                `json:"pinCode" bson:"pinCode,omitempty"`
	SpouseName                   string               `json:"spouseName" bson:"spouseName,omitempty"`
	Gender                       string               `json:"gender" bson:"gender,omitempty"`
	Mobile                       string               `json:"mobileNumber" bson:"mobileNumber,omitempty"`
	Email                        string               `json:"email" bson:"email,omitempty"`
	DateOfBirth                  *time.Time           `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	CreatedDate                  *time.Time           `json:"createdDate" bson:"createdDate,omitempty"`
	Address                      string               `json:"address" bson:"address,omitempty"`
	Profile                      string               `json:"profile" bson:"profile,omitempty"`
	OrganisationID               string               `json:"organisationId" bson:"organisationId,omitempty"`
	Password                     string               `json:"-" bson:"password,omitempty"`
	ConfirmPassword              string               `json:"confirmpassword" bson:"confirmpassword,omitempty"`
	Pass                         string               `json:"password" bson:"-"`
	Role                         string               `json:"role" bson:"role,omitempty"`
	Designation                  string               `json:"designation" bson:"designation,omitempty"`
	Type                         string               `json:"type" bson:"type,omitempty"`
	Created                      Created              `json:"createdOn" bson:"createdOn,omitempty"`
	Updated                      Updated              `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog                    []Updated            `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	Token                        string               `json:"token" bson:"-"`
	Status                       string               `json:"status" bson:"status,omitempty"`
	ManagerID                    string               `json:"managerId" bson:"managerId,omitempty"`
	BloodGroup                   string               `json:"bloodGroup" bson:"bloodGroup,omitempty"`
	CollectionLimit              CollectionLimit      `json:"collectionLimit" bson:"collectionLimit,omitempty"`
	AccessPrivilege              AccessPrivilege      `json:"accessPrivilege" bson:"accessPrivilege,omitempty"`
	Experience                   float64              `json:"experience" bson:"experience,omitempty"`
	EducationalQualification     string               `json:"educationalQualification" bson:"educationalQualification,omitempty"`
	KnowledgeDomains             []primitive.ObjectID `json:"knowledgeDomains" bson:"knowledgeDomains,omitempty"`
	SubDomains                   []primitive.ObjectID `json:"subDomains" bson:"subDomains,omitempty"`
	Organisation                 string               `json:"organisation" bson:"organisation,omitempty"`
	Officenumber                 string               `json:"officenumber" bson:"officenumber,omitempty"`
	Officeaddress                string               `json:"officeaddress" bson:"officeaddress,omitempty"`
	OrganisationNatureofBusiness string               `json:"organisationNatureofBusiness" bson:"organisationNatureofBusiness,omitempty"`
	KdExpertise                  string               `json:"kdExpertise" bson:"kdExpertise,omitempty"`
	MailNotify                   bool                 `json:"mailNotify" bson:"mailNotify,omitempty"`
	SmsNotify                    bool                 `json:"smsNotify" bson:"smsNotify,omitempty"`
	IsSelfRegistration           bool                 `json:"isSelfRegistration" bson:"isSelfRegistration,omitempty"`
	LanguageExpertise            primitive.ObjectID   `json:"languageExpertise" bson:"languageExpertise"`
	LanguagesKnown               []primitive.ObjectID `json:"languagesKnown" bson:"languagesKnown"`
	ViewLanguage                 primitive.ObjectID   `json:"viewLanguage" bson:"viewLanguage"`
	SubjectExpertise             string               `json:"subjectExpertise" bson:"subjectExpertise,omitempty"`
	Occupation                   string               `json:"occupation" bson:"occupation,omitempty"`
	UserType                     string               `json:"userType" bson:"userType,omitempty"`
	ProofType                    string               `json:"proofType" bson:"proofType,omitempty"`
	ProofNo                      string               `json:"proofNo" bson:"proofNo,omitempty"`
	UserOrg                      primitive.ObjectID   `json:"userOrg,omitempty" bson:"userOrg,omitempty"`
	Village                      string               `json:"village" bson:"village,omitempty"`
	VillageCode                  primitive.ObjectID   `json:"villageCode" bson:"villageCode,omitempty"`
	SmsCount                     int64                `json:"smsCount" bson:"smsCount,omitempty"`
	Block                        string               `json:"block" bson:"block,omitempty"`
	BlockCode                    primitive.ObjectID   `json:"blockCode" bson:"blockCode,omitempty"`
	Grampanchayat                string               `json:"grampanchayat" bson:"grampanchayat,omitempty"`
	GrampanchayatCode            primitive.ObjectID   `json:"grampanchayatCode" bson:"grampanchayatCode,omitempty"`
}
type AccessPrivilege struct {
	AccessLevel    string               `json:"accessLevel" bson:"accessLevel,omitempty"`
	Districts      []primitive.ObjectID `json:"districts" bson:"districts,omitempty"`
	States         []primitive.ObjectID `json:"states" bson:"states,omitempty"`
	Villages       []primitive.ObjectID `json:"villages" bson:"villages,omitempty"`
	Blocks         []primitive.ObjectID `json:"blocks" bson:"blocks,omitempty"`
	Grampanchayats []primitive.ObjectID `json:"grampanchayats" bson:"grampanchayats,omitempty"`
}
type UserAccess struct {
	Is       bool   `json:"is" bson:"is,omitempty"`
	UserName string `json:"userName" bson:"userName,omitempty"`
}

//RefUser :""
type RefUser struct {
	User `bson:",inline"`
	Ref  struct {
		Block                Block             `json:"block" bson:"block,omitempty"`
		Grampanchayat        GramPanchayat     `json:"grampanchayat" bson:"grampanchayat,omitempty"`
		State                State             `json:"state" bson:"state,omitempty"`
		District             District          `json:"district" bson:"district,omitempty"`
		Manager              User              `json:"manager" bson:"manager,omitempty"`
		Village              Village           `json:"village" bson:"village,omitempty"`
		Organisation         *Organisation     `json:"organisation" bson:"organisation,omitempty"`
		AccessDistricts      []District        `json:"accessDistricts" bson:"accessDistricts,omitempty"`
		AccessStates         []State           `json:"accessStates" bson:"accessStates,omitempty"`
		AccessVillages       []Village         `json:"accessVillages" bson:"accessVillages,omitempty"`
		AccessBlocks         []Block           `json:"accessBlocks" bson:"accessBlocks,omitempty"`
		AccessGrampanchayats []GramPanchayat   `json:"accessGrampanchayats" bson:"accessGrampanchayats,omitempty"`
		Projects             []RefProjectUser  `json:"projects" bson:"projects,omitempty"`
		KnowledgeDomains     []KnowledgeDomain `json:"knowledgeDomains" bson:"knowledgeDomains,omitempty"`
		Type                 UserType          `json:"type" bson:"type,omitempty"`
		SubDomains           []SubDomain       `json:"subDomains" bson:"subDomains,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserFilter : ""
type UserFilter struct {
	Status             []string             `json:"status" bson:"status,omitempty"`
	UniqueID           []string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	OmitID             []string             `json:"omitId" bson:"omitId,omitempty"`
	OrganisationID     []primitive.ObjectID `json:"organisationId" bson:"organisationId,omitempty"`
	States             []primitive.ObjectID `json:"states" bson:"states,omitempty"`
	Districts          []primitive.ObjectID `json:"districts" bson:"districts,omitempty"`
	Villages           []primitive.ObjectID `json:"villages" bson:"villages,omitempty"`
	Blocks             []primitive.ObjectID `json:"blocks" bson:"blocks,omitempty"`
	Grampanchayats     []primitive.ObjectID `json:"grampanchayats" bson:"grampanchayats,omitempty"`
	IsSelfRegistration []bool               `json:"isSelfRegistration" bson:"isSelfRegistration,omitempty"`
	Project            []primitive.ObjectID `json:"project" bson:"project,omitempty"`
	AccessLevel        []string             `json:"accessLevel" bson:"accessLevel,omitempty"`
	Manager            []string             `json:"manager" bson:"manager,omitempty"`
	Type               []string             `json:"type" bson:"type"`
	GetRecentLocation  bool                 `json:"getRecentLocation"`
	CreatedFrom        struct {
		StartDate *time.Time `json:"startDate"`
		EndDate   *time.Time `json:"endDate"`
	} `json:"createdFrom"`
	OmitProjectUser struct {
		Is      bool               `json:"is"`
		Project primitive.ObjectID `json:"project"`
	} `json:"omitProjectUser"`
	Regex struct {
		Name      string `json:"name" bson:"name"`
		FirstName string `json:"firstName" bson:"firstName"`
		Lastname  string `json:"lastname" bson:"lastname"`
		Contact   string `json:"contact" bson:"contact"`
		UserName  string `json:"userName" bson:"userName"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}

//duplicate user filter
type DuplicateUserFilter struct {
	UserFilter `bson:",inline"`
	By         string `json:"by" bson:"by,omitempty"`
}
type DuplicateUserReport struct {
	ID    string    `json:"id" bson:"_id,omitempty"`
	Users []RefUser `json:"users" bson:"users,omitempty"`
}

//UserChangePassword : ""
type UserChangePassword struct {
	UserName    string `json:"userName" bson:"userName,omitempty"`
	OldPassword string `json:"oldPassword" bson:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword" bson:"newPassword,omitempty"`
}

//UserLocation : ""
type UserLocation struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	UserName  string     `json:"userName" bson:"userName,omitempty"`
	Time      *time.Time `json:"time" bson:"time,omitempty"`
	Location  Location   `json:"location" bson:"location,omitempty"`
	UserType  string     `json:"userType" bson:"userType,omitempty"`
	Name      string     `json:"name" bson:"name,omitempty"`
	EntryType string     `json:"entryType" bson:"entryType,omitempty"`
	ErrMsg    string     `json:"errMsg" bson:"errMsg,omitempty"`
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
type RefPassword struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Token    string `json:"token" bson:"token,omitempty"`
	OTP      int64  `json:"otp" bson:"otp,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}
type CollectionLimit struct {
	Cash float64 `json:"cash,omitempty" bson:"cash,omitempty"`
}
type UserStory struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName string     `json:"userName" bson:"userName,omitempty"`
	On       *time.Time `json:"on" bson:"on,omitempty"`
	By       string     `json:"by" bson:"by,omitempty"`
	ByType   string     `json:"byType" bson:"byType,omitempty"`
	Msg      string     `json:"msg" bson:"msg,omitempty"`
}

type UserUniquinessChk struct {
	Success bool   `json:"success" bson:"success,omitempty"`
	Message string `json:"message" bson:"message,omitempty"`
}
type DissiminateUser struct {
	ID                   primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name                 string             `json:"name" bson:"name,omitempty"`
	MobileNumber         string             `json:"mobileNumber"  bson:"mobileNumber,omitempty"`
	Email                string             `json:"email" bson:"email,omitempty"`
	UserID               string             `json:"userID"  bson:"userID,omitempty"`
	AppRegistrationToken string             `json:"appRegistrationToken" bson:"appRegistrationToken,omitempty"`
}
type UserOTPLogin struct {
	User     `bson:",inline"`
	OTP      string `json:"otp"`
	Scenario string `json:"scenario"`
}
