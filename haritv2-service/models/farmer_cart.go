package models

import "time"

type FarmerCart struct {
	CompanyID   string     `json:"companyId" bson:"companyId,omitempty"`
	CompanyType string     `json:"companyType" bson:"companyType,omitempty"`
	UniqueID    string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status      string     `json:"status" bson:"status,omitempty"`
	Created     *CreatedV2 `json:"created" bson:"created,omitempty"`
	DateTime    *time.Time `json:"dateTime,omitempty"  bson:"dateTime,omitempty"`
	Quantity    float64    `json:"quantity,omitempty"  bson:"quantity,omitempty"`
	ProductID   string     `json:"productId" bson:"productId,omitempty"`
}

//FarmerCartFilter : ""
type FarmerCartFilter struct {
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   []string `json:"status,omitempty" bson:"status,omitempty"`
}

//RefCart : ""
type RefFarmerCart struct {
	FarmerCart `bson:",inline"`
	Ref        struct {
		Ulb             RefULB  `json:"ulb" bson:"ulb,omitempty"`
		CalculatedPrice float64 `json:"calculatedPrice" bson:"calculatedPrice,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
