package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Properties struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Location       Location           `json:"location" bson:"location"`
	Citizen        *Citizen           `json:"citizen" bson:"citizen"`
	GcUser         *GCUser            `json:"gcUser" bson:"gcUser"`
	MinUser        *MinUser           `json:"minUser" bson:"minUser"`
	RegisterBy     *RegisterBy        `json:"registerBy" bson:"registerBy"`
	Mobile         string             `json:"mobile" bson:"mobile,omitempty"`
	Email          string             `json:"email" bson:"email,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	ClosingNotes   string             `json:"closingNotes,omitempty" bson:"closingNotes,omitempty"`
	NfcID          string             `json:"nfcId,omitempty" bson:"nfcId,omitempty"`
	HoldingNumber  string             `json:"holdingNumber" bson:"holdingNumber,omitempty"`
	PropertyType   string             `json:"propertyType" bson:"propertyType,omitempty"`
	OwnerType      string             `json:"ownerType" bson:"ownerType,omitempty"`
	WandAdmin      string             `json:"wardAdmin" bson:"wardAdmin,omitempty"`
	CircleCode     string             `json:"circleCode" bson:"circleCode,omitempty"`
	SectorCode     string             `json:"sectorCode" bson:"sectorCode,omitempty"`
	WardCode       string             `json:"wardCode" bson:"wardCode,omitempty"`
	RequstedDate   *time.Time         `json:"requstedDate" bson:"requstedDate,omitempty"`
	Imageurl       string             `json:"imageUrl" bson:"imageUrl,omitempty"`
	RegisterDate   *time.Time         `json:"registerDate" bson:"registerDate,omitempty"`
	Created        *CreatedV2         `json:"created" bson:"created,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
	IdentityType   string             `json:"identityType" bson:"identityType,omitempty"`
	IdentityNo     string             `json:"identityNo" bson:"identityNo,omitempty"`
	Latitude       Location           `json:"latitude" bson:"latitude,omitempty"`
	Longitude      Location           `json:"longitude" bson:"longitude,omitempty"`
	PropertyNo     string             `json:"propertyNo,omitempty" bson:"propertyNo,omitempty"`
	PropertySource string             `json:"propertySource,omitempty" bson:"propertySource,omitempty"`
	Pincode        string             `json:"pincode,omitempty" bson:"pincode,omitempty"`
	HouseUID       string             `json:"houseUid,omitempty" bson:"houseUid,omitempty"`
	Roadtype       string             `json:"roadtype,omitempty" bson:"roadtype,omitempty"`
	Address        Address            `json:"address" bson:"address,omitempty"`
}

type RefProperties struct {
	Properties `bson:",inline"`
	Ref        struct {
		OrganisationID Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		State          *State       `json:"state,omitempty" bson:"state,omitempty"`
		District       *District    `json:"district,omitempty" bson:"district,omitempty"`
		Village        *Village     `json:"village,omitempty" bson:"village,omitempty"`
		Zone           *Zone        `json:"zone,omitempty" bson:"zone,omitempty"`
		Ward           *Ward        `json:"ward,omitempty" bson:"ward,omitempty"`
		Roadtype       RoadType     `json:"roadtype,omitempty" bson:"roadtype,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterProperties struct {
	Status       []string `json:"status" bson:"status,omitempty"`
	Name         []string `json:"name" bson:"name,omitempty"`
	UniqueID     []string `json:"uniqueID" bson:"uniqueID,omitempty"`
	GCID         []string `json:"gcId" bson:"gcId,omitempty"`
	ManagerID    []string `json:"managerId" bson:"managerId,omitempty"`
	CitizenID    []string `json:"citizenId" bson:"citizenId,omitempty"`
	RegisterID   []string `json:"registerId" bson:"registerId,omitempty"`
	Mobile       []string `json:"mobile" bson:"mobile,omitempty"`
	Pincode      []string `json:"pincode" bson:"pincode,omitempty"`
	OwnerType    []string `json:"ownerType" bson:"ownerType,omitempty"`
	PropertyType []string `json:"propertyType" bson:"propertyType"`
	VillageCode  []string `json:"villageCode" bson:"villageCode,omitempty"`
	StateCode    []string `json:"stateCode" bson:"stateCode,omitempty"`
	ZoneCode     []string `json:"zoneCode" bson:"zoneCode,omitempty"`
	WardCode     []string `json:"wardCode" bson:"wardCode,omitempty"`
	DistrictCode []string `json:"districtCode" bson:"districtCode,omitempty"`

	Regex struct {
		Name          string  `json:"name" bson:"name"`
		HouseUID      string  `json:"houseUid" bson:"houseUid"`
		Citizen       Citizen `json:"citizen,omitempty" bson:"citizen,omitempty"`
		HoldingNumber string  `json:"holdingNumber" bson:"holdingNumber,omitempty"`
		NfcID         string  `json:"nfcId,omitempty" bson:"nfcId,omitempty"`
		PropertyType  string  `json:"propertyType" bson:"propertyType,omitempty"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type CompletedPropertiees struct {
	UniqueID     string `json:"uniqueID" bson:"uniqueID,omitempty"`
	ClosingNotes string `json:"closingNotes,omitempty" bson:"closingNotes,omitempty"`
	Imageurl     string `json:"imageUrl" bson:"imageUrl,omitempty"`
}

type WardWisePropertyCount struct {
	Quantity int `json:"quantity" bson:"quantity,omitempty"`
}
