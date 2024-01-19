package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrganisationConfig struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Logo         string             `json:"logo" bson:"logo,omitempty"`
	WaterMark    string             `json:"waterMark" bson:"waterMark,omitempty"`
	LogoWithName string             `json:"logoWithName" bson:"logoWithName,omitempty"`
	Email        struct {
		ContactUs string `json:"contactUs" bson:"contactUs,omitempty"`
		SendEmail string `json:"sendEmail" bson:"sendEmail,omitempty"`
	} `json:"email" bson:"email,omitempty"`
	Mobile     string   `json:"mobile" bson:"mobile,omitempty"`
	Phone      string   `json:"phone" bson:"phone,omitempty"`
	Address    string   `json:"address" bson:"address,omitempty"`
	Copyrights string   `json:"copyrights" bson:"copyrights,omitempty"`
	PoweredBy  string   `json:"poweredBy" bson:"poweredBy,omitempty"`
	Rights     string   `json:"rights" bson:"rights,omitempty"`
	Status     string   `json:"status" bson:"status,omitempty"`
	UniqueID   string   `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Created    *Created `json:"created" bson:"created,omitempty"`
	IsDefault  bool     `json:"isdefault" bson:"isdefault,omitempty"`
	IsSms      bool     `json:"issms" bson:"issms,omitempty"`
	APIURL     string   `json:"-" bson:"apiUrl,omitempty"`
	UIURL      string   `json:"-" bson:"uiUrl,omitempty"`
	LossOfPay  float64  `json:"lossOfPay,omitempty" bson:"lossOfPay,omitempty"`
	PartialPay float64  `json:"partialPay,omitempty" bson:"partialPay,omitempty"`
}
type OrganisationConfigFilter struct {
	IsDefault []bool   `json:"isdefault,omitempty"`
	Status    []string `json:"status" form:"status" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Searchbox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}
type RefOrganisationConfig struct {
	OrganisationConfig `bson:",inline"`
	Ref                struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
