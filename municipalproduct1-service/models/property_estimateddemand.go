package models

import (
	"time"
)

// EstimatedPropertyDemand : ""
type EstimatedPropertyDemand struct {
	EstimatedDemand `bson:",inline"`
	Created         Created `json:"created,omitempty" bson:"created,omitempty"`
	Status          string  `json:"status" bson:"status,omitempty"`
	UniqueID        string  `json:"uniqueId" bson:"uniqueId,omitempty"`
}

// EstimatedDemand : ""
type EstimatedDemand struct {
	MunicipalityID string          `json:"municipalityId" bson:"municipalityId,omitempty"`
	YOA            string          `json:"yoa" bson:"yoa,omitempty"`
	IsMatched      string          `json:"isMatched" bson:"isMatched,omitempty"`
	AreaOfPlot     float64         `json:"areaOfPlot" bson:"areaOfPlot,omitempty"`
	BuiltUpArea    float64         `json:"builtUpArea" bson:"builtUpArea,omitempty"`
	PropertyTypeID string          `json:"propertyTypeId" bson:"propertyTypeId,omitempty"`
	RoadTypeID     string          `json:"roadTypeId" bson:"roadTypeId,omitempty"`
	DOA            *time.Time      `json:"doa" bson:"doa,omitempty"`
	IsGovtProperty string          `json:"isGovtProperty" bson:"isGovtProperty,omitempty"`
	Owner          []PropertyOwner `json:"owner" bson:"-"`
	OwnerID        []string        `json:"ownerId" bson:"ownerIds,omitempty"` // multiple owners
	Floors         []PropertyFloor `json:"floors" bson:"-"`
	Legacy         struct {
		IsLegacy       bool               `json:"isLegacy" bson:"isLegacy,omitempty"`
		LegacyProperty *RegLegacyProperty `json:"legacyProperty" bson:"legacyProperty,omitempty"`
	} `json:"legacy" bson:"-"`
	Demand             UpdateDemand                `json:"demand" bson:"demand"`
	Collection         UpdateCollection            `json:"collection" bson:"collection"`
	Penalty            UpdatePenalty               `json:"penalty" bson:"penalty"`
	Rebate             UpdateRebate                `json:"rebate" bson:"rebate"`
	Advance            float64                     `json:"advance" bson:"advance"`
	PreviousCollection PreviousCollection          `json:"previousCollection" bson:"previousCollection"`
	NDemand            PropertyTaxTotalDemand      `json:"ndemand" bson:"ndemand,omitempty"`
	NCollection        PropertyTaxTotalCollection  `json:"ncollection" bson:"ncollection,omitempty"`
	NOutstanding       PropertyTaxTotalOutStanding `json:"noutstanding" bson:"noutstanding,omitempty"`
	NPending           PropertyTaxTotalPending     `json:"npending" bson:"npending,omitempty"`
	ParkPenalty        bool                        `json:"parkPenalty" bson:"parkPenalty,omitempty"`
}

// RefEstimatedPropertyDemand :""
type RefEstimatedPropertyDemand struct {
	EstimatedPropertyDemand `bson:",inline"`
	Ref                     struct {
		PropertyOwner []RefPropertyOwner    `json:"propertyOwner" bson:"propertyOwner,omitempty"`
		Floors        []RefPropertyFloor    `json:"floors" bson:"floors,omitempty"`
		PropertyType  *RefPropertyType      `json:"propertyType" bson:"propertyType,omitempty"`
		RoadType      *RefRoadType          `json:"roadType" bson:"roadType,omitempty"`
		YOA           *RefFinancialYear     `json:"yoa" bson:"yoa,omitempty"`
		MunicipalType *RefMunicipalType     `json:"municipalType" bson:"municipalType,omitempty"`
		Demand        OverallPropertyDemand `json:"demand" bson:"demand,omitempty"`
		Payments      struct {
			Payments float64 `json:"payments" bson:"payments,omitempty"`
			Amount   float64 `json:"amount" bson:"amount,omitempty"`
		} `json:"payments" bson:"payments,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
