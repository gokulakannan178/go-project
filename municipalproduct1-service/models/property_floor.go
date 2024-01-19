package models

import "time"

//PropertyFloor : ""
type PropertyFloor struct {
	Title            string     `json:"title" bson:"title,omitempty"`
	PropertyID       string     `json:"propertyId" bson:"propertyId,omitempty"`
	UniqueID         string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	No               string     `json:"no" bson:"no,omitempty"`
	UsageType        string     `json:"usageType" bson:"usageType,omitempty"`
	ConstructionType string     `json:"constructionType" bson:"constructionType,omitempty"`
	BuildUpArea      float64    `json:"buildUpArea" bson:"buildUpArea,omitempty"`
	OccupancyType    string     `json:"occupancyType" bson:"occupancyType,omitempty"`
	NonResUsageType  string     `json:"nonResUsageType" bson:"nonResUsageType,omitempty"`
	RatableAreaType  string     `json:"ratableAreaType" bson:"ratableAreaType,omitempty"`
	DateFrom         *time.Time `json:"dateFrom" bson:"dateFrom,omitempty"`
	DateTo           *time.Time `json:"dateTo" bson:"dateTo,omitempty"`
	Status           string     `json:"status" bson:"status,omitempty"`
	Created          Created    `json:"created" bson:"created,omitempty"`
	Updated          []Updated  `json:"updated" bson:"updated,omitempty"`
	SortOrder        int64      `json:"sortOrder"  bson:"sortOrder,omitempty"`
	NewPropertyID    string     `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID    string     `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

//RefPropertyFloor :""
type RefPropertyFloor struct {
	PropertyFloor `bson:",inline"`
	Ref           struct {
		UsageType        *UsageType                 `json:"usageType" bson:"usageType,omitempty"`
		ConstructionType *ConstructionType          `json:"constructionType" bson:"constructionType,omitempty"`
		OccupancyType    *OccupancyType             `json:"occupancyType" bson:"occupancyType,omitempty"`
		NonResUsageType  *NonResidentialUsageFactor `json:"nonResUsageType" bson:"nonResUsageType,omitempty"`
		FloorRatableArea *FloorRatableArea          `json:"floorRatableArea" bson:"floorRatableArea,omitempty"`
		FloorNo          *RefFloorType              `json:"floorNo" bson:"floorNo,omitempty"`
		AVR              *AVR                       `json:"avr" bson:"avr,omitempty"`
		VLR              VacantLandRate             `json:"vlr" bson:"vlr,omitempty"`
		CompositeTax     CompositeTaxRateMaster     `json:"compositeTax" bson:"compositeTax,omitempty"`
		PropertyTax      *PropertyTax               `json:"propertyTax" bson:"propertyTax,omitempty"`
		RatableArea      float64                    `json:"ratableArea" bson:"ratableArea,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

func (res *RefPropertyFloor) IncFloor(a int) int {
	return a + 1
}

//PropertyFloorFilter : ""
type PropertyFloorFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
