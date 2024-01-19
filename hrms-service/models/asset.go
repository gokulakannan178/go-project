package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Asset struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	Description      string             `json:"description,omitempty" bson:"description,omitempty"`
	AssetTypeId      string             `json:"assetTypeId,omitempty" bson:"assetTypeId,omitempty"`
	AssetPropertysId []AssetPropertys   `json:"assetPropertysId,omitempty" bson:"-"`
	AssignId         string             `json:"assignId,omitempty" bson:"assignId,omitempty"`
	EmployeeId       string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationID   string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Make             string             `json:"make,omitempty" bson:"make,omitempty"`
	ModelNo          string             `json:"modelNo,omitempty" bson:"modelNo,omitempty"`
	Created          Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
	Remark           string             `json:"remark,omitempty" bson:"remark,omitempty"`
}

type RefAsset struct {
	Asset `bson:",inline"`
	Ref   struct {
		AssetTypeId          AssetType        `json:"assetTypeId,omitempty" bson:"assetTypeId,omitempty"`
		AssetTypePepropertys []AssetPropertys `json:"assetTypePepropertys,omitempty" bson:"assetTypePepropertys,omitempty"`
		AssetId              AssetAssign      `json:"assetId,omitempty" bson:"assetId,omitempty"`
		Assetlog             AssetLog         `json:"assetlog,omitempty" bson:"assetlog,omitempty"`
		Employee             Employee         `json:"employee,omitempty" bson:"employee,omitempty"`
		OrganisationID       Organisation     `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterAsset struct {
	EmployeeId     []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
type AssetAssign struct {
	UniqueID   string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	AssetId    string `json:"assetId,omitempty" bson:"assetId,omitempty"`
	AssignId   string `json:"assignId,omitempty" bson:"assignId,omitempty"`
	EmployeeId string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Status     string `json:"status,omitempty" bson:"status,omitempty"`
	IsLog      bool   `json:"isLog,omitempty" bson:"isLog,omitempty"`
	Remark     string `json:"remark,omitempty" bson:"remark,omitempty"`
}
