package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HouseVisited struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	//Citizen        *Citizen           `json:"citizen,omitempty" bson:"citizen,omitempty"`
	Property       PropertyDetail `json:"property" bson:"property,omitempty"`
	MinUser        MinUser        `json:"minUser,omitempty" bson:"minUser,omitempty"`
	GCUser         GCUser         `json:"gcUser,omitempty" bson:"gcUser,omitempty"`
	Location       Location       `json:"location,omitempty" bson:"location,omitempty"`
	Description    string         `json:"description,omitempty" bson:"description,omitempty"`
	OrganisationID string         `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           *time.Time     `json:"date" bson:"date,omitempty"`
	Created        *CreatedV2     `json:"created" bson:"created,omitempty"`
	IsStatus       string         `json:"isStatus,omitempty" bson:"isStatus,omitempty"`
	Status         string         `json:"status,omitempty" bson:"status,omitempty"`
	IsAvailable    string         `json:"isAvailable" bson:"isAvailable"`
	HoldingNumber  string         `json:"holdingNumber,omitempty" bson:"holdingNumber,omitempty"`
	Remark         string         `json:"remark,omitempty" bson:"remark,omitempty"`
	WardNo         string         `json:"wardNo" bson:"wardNo,omitempty"`
	CircleNo       string         `json:"circleNo" bson:"circleNo,omitempty"`
	WandAdmin      string         `json:"wardAdmin" bson:"wardAdmin,omitempty"`
	SectorCode     string         `json:"sectorCode,omitempty" bson:"sectorCode,omitempty"`
	OwnerType      string         `json:"ownerType,omitempty" bson:"ownerType,omitempty"`
	PropertyType   string         `json:"propertyType" bson:"propertyType,omitempty"`
	Distance       string         `json:"distance" bson:"distance,omitempty"`
	HouseUID       string         `json:"houseUid,omitempty" bson:"houseUid,omitempty"`
}

type RefHouseVisited struct {
	HouseVisited `bson:",inline"`
	Ref          struct {
		OrganisationID Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterHouseVisited struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	IsStatus   []string `json:"isStatus,omitempty" bson:"isStatus,omitempty"`
	Name       []string `json:"name" bson:"name,omitempty"`
	UniqueID   []string `json:"uniqueID" bson:"uniqueID,omitempty"`
	ManagerID  []string `json:"managerId" bson:"managerId,omitempty"`
	DumbsiteID []string `json:"dumbsiteId" bson:"dumbsiteId,omitempty"`
	GCID       []string `json:"gcId" bson:"gcId,omitempty"`
	CitizenID  []string `json:"citizenId" bson:"citizenId,omitempty"`
	WardNo     []string `json:"wardNo" bson:"wardNo,omitempty"`
	CircleNo   []string `json:"circleNo" bson:"circleNo,omitempty"`

	Regex struct {
		Name        string `json:"name" bson:"name"`
		ManagerName string `json:"managerName" bson:"managerName"`
		GCName      string `json:"gcName" bson:"gcName"`
		CitizenName string `json:"citizenName" bson:"citizenName"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
