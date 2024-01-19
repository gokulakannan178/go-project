package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetLog struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	AssetTypeId    string             `json:"assetTypeId,omitempty" bson:"assetTypeId,omitempty"`
	AssetId        string             `json:"assetId,omitempty" bson:"assetId,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Make           string             `json:"make,omitempty" bson:"make,omitempty"`
	ModelNo        string             `json:"modelNo,omitempty" bson:"modelNo,omitempty"`
	Action         struct {
		UserName string     `json:"userName,omitempty" bson:"username,omitempty"`
		UserType string     `json:"userType,omitempty" bson:"usertype,omitempty"`
		UserID   string     `json:"userId,omitempty" bson:"userId,omitempty"`
		Date     *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	} `json:"action,omitempty" bson:"action,omitempty"`
	Remark    string     `json:"remark,omitempty" bson:"remark,omitempty"`
	Created   *Created   `json:"createdOn" bson:"createdOn,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
	//IsLog     string     `json:"isLog,omitempty" bson:"isLog,omitempty"`
}

type RefAssetLog struct {
	AssetLog `bson:",inline"`
	Ref      struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterAssetLog struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex," bson:"regex"`
}
