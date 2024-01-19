package models

import "time"

//PropertyVistLog : ""
type PropertyVisitLog struct {
	UniqueID      string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status        string     `json:"status" bson:"status,omitempty"`
	Created       CreatedV2  `json:"created,omitempty" bson:"created,omitempty"`
	Updated       []Updated  `json:"updated" bson:"updated,omitempty"`
	PropertyID    string     `json:"propertyId" bson:"propertyId,omitempty"`
	UserID        string     `json:"userId" bson:"userId,omitempty"`
	UserType      string     `json:"usertype,omitempty" bson:"usertype,omitempty"`
	OwnerName     string     `json:"ownerName" bson:"ownerName,omitempty"`
	Address       Address    `json:"address" bson:"address,omitempty"`
	MobileNo      string     `json:"mobileNo" bson:"mobileNo,omitempty"`
	Location      Location   `json:"location" bson:"location,omitempty"`
	RemarkID      string     `json:"remarkId" bson:"remarkId,omitempty"`
	Remarks       string     `json:"remarks" bson:"remarks,omitempty"`
	NextVisitDate *time.Time `json:"nextVisitDate" bson:"nextVisitDate,omitempty"`
	NewPropertyID string     `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string     `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

//RefPropertyVistLog : ""
type RefPropertyVisitLog struct {
	PropertyVisitLog `bson:",inline"`

	Ref struct {
		Property RefProperty                `json:"property" bson:"property,omitempty"`
		Remark   PropertyVisitLogRemarkType `json:"remark" bson:"remark,omitempty"`
		Address  RefAddress                 `json:"address" bson:"address,omitempty"`
		Amount   float64                    `json:"amount" bson:"amount,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type PropertyVisitLogFilter struct {
	Status        []string       `json:"status" bson:"status,omitempty"`
	UniqueID      []string       `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID    []string       `json:"propertyId" bson:"propertyId,omitempty"`
	UserId        []string       `json:"userId" bson:"userId,omitempty"`
	UserType      []string       `json:"userType" bson:"userType,omitempty"`
	Address       *AddressSearch `json:"address" bson:"address,omitempty"`
	NextDateRange *DateRange     `json:"dateRange,omitempty"  bson:"dateRange,omitempty"` //nextVisitDate Key in visit Log

	SearchBox struct {
		Ownername  string `json:"ownername" bson:"ownername,omitempty"`
		MobileNo   string `json:"mobileNo" bson:"mobileNo,omitempty"`
		PropertyID string `json:"propertyId" bson:"propertyId,omitempty"`
	} `json:"searchBox,omitempty" bson:"searchBox,omitempty"`
}
