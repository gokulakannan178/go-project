package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Department : ""
type Department struct {
	ID               primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	Desc             string             `json:"desc" bson:"desc,omitempty"`
	Address          Address            `json:"address" bson:"address,omitempty"`
	OrganisationID   string             `json:"organisationId" bson:"organisationId,omitempty"`
	DistrictCode     string             `json:"districtCode" bson:"districtCode,omitempty"`
	DepartmentTypeID string             `json:"departmentTypeId" bson:"departmentTypeId,omitempty"`
	Level            string             `json:"level" bson:"level,omitempty"`
	Status           string             `json:"status" bson:"status,omitempty"`
	Created          Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated          []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefDepartment : ""
type RefDepartment struct {
	Department `bson:",inline"`
	Ref        struct {
		Organisation     *ULB            `json:"organisation,omitempty" bson:"organisation,omitempty"`
		Address          *RefAddress     `json:"address,omitempty" bson:"address,omitempty"`
		DepartmentTypeID *DepartmentType `json:"departmentType" bson:"departmentType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DepartmentFilter : ""
type DepartmentFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	Organisation []string `json:"organisation,omitempty" organisation:"status,omitempty"`
	DistrictCode []string `json:"districtCode" bson:"districtCode,omitempty"`
	SortBy       string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder    int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
