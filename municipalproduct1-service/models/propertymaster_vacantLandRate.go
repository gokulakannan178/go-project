package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//VacantLandRate : ""
type VacantLandRate struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name               string             `json:"name" bson:"name,omitempty"`
	Desc               string             `json:"desc" bson:"desc,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	Created            Created            `json:"created" bson:"created,omitempty"`
	Updated            []Updated          `json:"updated" bson:"updated,omitempty"`
	MunicipalityTypeID string             `json:"municipalityTypeId" bson:"municipalityTypeId,omitempty"`
	RoadTypeID         string             `json:"roadTypeId" bson:"roadTypeId,omitempty"`
	Rate               float64            `json:"rate" bson:"rate,omitempty"`
	RateType           string             `json:"rateType" bson:"rateType,omitempty"`
	DOE                *time.Time         `json:"doe" bson:"doe,omitempty"`
}

//RefVacantLandRate :""
type RefVacantLandRate struct {
	VacantLandRate `bson:",inline"`
	Ref            struct {
		MunicipalType *MunicipalType `json:"municipalType" bson:"municipalType,omitempty"`
		RoadType      *RoadType      `json:"roadType" bson:"roadType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//VacantLandRateFilter : ""
type VacantLandRateFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
