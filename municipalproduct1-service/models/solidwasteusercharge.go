package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SolidWasteUserCharge : ""
type SolidWasteUserCharge struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	UniqueID      string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID    string             `json:"propertyId" bson:"propertyId,omitempty"`
	CategoryID    string             `json:"categoryId" bson:"categoryId,omitempty"`
	SubCategoryID string             `json:"subCategoryId" bson:"subCategoryId,omitempty"`
	OwnerName     string             `json:"ownerName" bson:"ownerName,omitempty"`
	Desc          string             `json:"description" bson:"description,omitempty"`
	MobileNo      string             `json:"mobileNo" bson:"mobileNo,omitempty"`
	Address       Address            `json:"address" bson:"address,omitempty"`
	DateFrom      *time.Time         `json:"dateFrom" bson:"dateFrom,omitempty"`
	DateTo        *time.Time         `json:"dateTo" bson:"dateTo,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
}

//SolidWasteUserChargeFilter : ""
type SolidWasteUserChargeFilter struct {
	UniqueIDs     []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	CategoryID    []string `json:"categoryId" bson:"categoryId,omitempty"`
	SubCategoryID []string `json:"subCategoryId" bson:"subCategoryId,omitempty"`
	Status        []string `json:"status" bson:"status,omitempty"`
	StateCode     []string `json:"stateCode" bson:"stateCode,omitempty"`
	DistrictCode  []string `json:"districtCode" bson:"districtCode,omitempty"`
	VillageCode   []string `json:"villageCode" bson:"villageCode,omitempty"`
	ZoneCode      []string `json:"zoneCode" bson:"zoneCode,omitempty"`
	WardCode      []string `json:"wardCode" bson:"wardCode,omitempty"`

	Regex struct {
		UniqueID  string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Name      string `json:"name" bson:"name"`
		OwnerName string `json:"ownerName" bson:"ownerName,omitempty"`
		MobileNo  string `json:"mobileNo" bson:"mobileNo,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefSolidWasteUserCharge : ""
type RefSolidWasteUserCharge struct {
	SolidWasteUserCharge `bson:",inline"`
	Ref                  struct {
		Address     RefAddress                      `json:"address" bson:"address,omitempty"`
		Category    SolidWasteUserChargeCategory    `json:"category" bson:"category,omitempty"`
		SubCategory SolidWasteUserChargeSubCategory `json:"subCategory" bson:"subCategory,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}
