package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Applicant : ""
type Applicant struct {
	ID                     primitive.ObjectID  `json:"id"  bson:"_id,omitempty"`
	UserName               string              `json:"userName" bson:"userName,omitempty"`
	TypeID                 string              `json:"typeId" bson:"typeId,omitempty"`
	Name                   string              `json:"name" bson:"name,omitempty"`
	Description            string              `json:"description" bson:"description,omitempty"`
	Email                  string              `json:"email" bson:"email,omitempty"`
	MobileNumber           string              `json:"mobileNumber" bson:"mobileNumber,omitempty"`
	Status                 string              `json:"status" bson:"status,omitempty"`
	Certificates           []Certificate       `json:"certificates" bson:"certificates,omitempty"`
	EducationQualification []EducationDetail   `json:"educationQualification" bson:"educationQualification,omitempty"`
	Experiences            []ExperienceDetail  `json:"experiences" bson:"experiences,omitempty"`
	Address                Address             `json:"address" bson:"address,omitempty"`
	Created                Created             `json:"created" bson:"created,omitempty"`
	Updated                []Updated           `json:"updated" bson:"updated,omitempty"`
	ApplicantLog           []ApplicantTimeline `json:"applicantLog,omitempty"  bson:"applicantLog,omitempty"`
	Remarks                string              `json:"remarks" bson:"remarks,omitempty"`
}

// Certificate : ""
type Certificate struct {
	Name        string `json:"name" bson:"name,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	UploadURL   string `json:"uploadURL" bson:"uploadURL,omitempty"`
}

// EducationDetail : ""
type EducationDetail struct {
	Type        string  `json:"type" bson:"type,omitempty"`
	TypeName    string  `json:"typeName" bson:"typeName,omitempty"`
	YOP         int     `json:"yop" bson:"yop,omitempty"` //Year ofpassed out
	Percentage  float64 `json:"percentage" bson:"percentage,omitempty"`
	SchoolRuniv string  `json:"schoolRuniv" bson:"schoolRuniv,omitempty"`
}

// ExperienceDetail : ""
type ExperienceDetail struct {
	OrganizationName string `json:"organizationName" bson:"organizationName,omitempty"`
	Type             string `json:"type" bson:"type,omitempty"`
	TypeName         string `json:"typeName" bson:"typeName,omitempty"`
	From             string `json:"from" bson:"from,omitempty"`
	To               string `json:"to" bson:"to,omitempty"`
	CertURL          string `json:"certURL" bson:"certURL,omitempty"`
}

type RefApplicant struct {
	Applicant `bson:",inline"`
	Ref       struct {
		Organisation ULB        `json:"organisation,omitempty" bson:"organisation,omitempty"`
		Address      RefAddress `json:"address,omitempty" bson:"address,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ApplicantFilter : ""
type ApplicantFilter struct {
	Status        []string       `json:"status,omitempty" bson:"status,omitempty"`
	Address       *AddressSearch `json:"address"`
	ApplicantType []string       `json:"applicantType"`
	SortBy        string         `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder     int            `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//ApplicantStatusChange : ""
type ApplicantStatusChange struct {
	UserName string            `json:"userName,omitempty" bson:"userName,omitempty"`
	Info     ApplicantTimeline `json:"info,omitempty" bson:"info,omitempty"`
}

//ApplicantTimeline : ""
type ApplicantTimeline struct {
	On *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	By struct {
		ID   string `json:"id,omitempty" bson:"id,omitempty"`
		Type string `json:"type,omitempty" bson:"type,omitempty"`
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"by,omitempty" bson:"by,omitempty"`
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	TypeLabel string `json:"typeLabel,omitempty" bson:"typeLabel,omitempty"`
	Remarks   string `json:"remarks,omitempty" bson:"remarks,omitempty"`
}
