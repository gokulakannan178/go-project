package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BirthCertificate : ""
type BirthCertificate struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	FatherName       string             `json:"fatherName" bson:"fatherName,omitempty"`
	MotherName       string             `json:"motherName" bson:"motherName,omitempty"`
	Gender           string             `json:"gender" bson:"gender,omitempty"`
	Mobile           string             `json:"mobile" bson:"mobile,omitempty"`
	Email            string             `json:"email" bson:"email,omitempty"`
	DOB              *time.Time         `json:"dob" bson:"dob,omitempty"`
	Address          Address            `json:"address" bson:"address,omitempty"`
	PlaceOfBirth     string             `json:"placeOfBirth" bson:"placeOfBirth,omitempty"`
	HospitalID       string             `json:"hospitalId" bson:"hospitalId,omitempty"`
	PermanentAddress Address            `json:"permanentAddress" bson:"permanentAddress,omitempty"`
	Remarks          string             `json:"remarks" bson:"remarks,omitempty"`
	Action           Action             `json:"action" bson:"action,omitempty"`
	Status           string             `json:"status" bson:"status,omitempty"`
	Created          CreatedV2          `json:"created" bson:"created,omitempty"`
}

// BirthCertificateFilter : ""
type BirthCertificateFilter struct {
	UniqueIDs    []string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status       []string   `json:"status" bson:"status,omitempty"`
	HospitalID   []string   `json:"hospitalId" bson:"hospitalId,omitempty"`
	Name         []string   `json:"name" bson:"name,omitempty"`
	FatherName   []string   `json:"fatherName" bson:"fatherName,omitempty"`
	Gender       []string   `json:"gender" bson:"gender,omitempty"`
	DOB          *DateRange `json:"dob" bson:"dob,omitempty"`
	PlaceOfBirth []string   `json:"placeOfBirth" bson:"placeOfBirth,omitempty"`

	Regex struct {
		Name       string `json:"name" bson:"name,omitempty"`
		FatherName string `json:"fatherName" bson:"fatherName,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefBirthCertificate : ""
type RefBirthCertificate struct {
	BirthCertificate `bson:",inline"`
	Ref              struct {
		Hospital Hospital `json:"hospital" bson:"hospital,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

type Approved struct {
	On      *time.Time `json:"on" bson:"on,omitempty"`
	By      string     `json:"by" bson:"by,omitempty"`
	ByType  string     `json:"byType" form:"byType" bson:"byType,omitempty"`
	Remarks string     `json:"remarks" bson:"remarks,omitempty"`
}
type Rejeceted struct {
	On      *time.Time `json:"on" bson:"on,omitempty"`
	By      string     `json:"by" bson:"by,omitempty"`
	ByType  string     `json:"byType" form:"byType" bson:"byType,omitempty"`
	Remarks string     `json:"remarks" bson:"remarks,omitempty"`
}
