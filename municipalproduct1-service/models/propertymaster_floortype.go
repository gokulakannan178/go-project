package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//FloorType : ""
type FloorType struct {
	ID                  primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID            string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name                string             `json:"name" bson:"name,omitempty"`
	Desc                string             `json:"desc" bson:"desc,omitempty"`
	Status              string             `json:"status" bson:"status,omitempty"`
	Discount            float64            `json:"discount" bson:"discount,omitempty"`
	MaintenanceDiscount float64            `json:"maintenanceDiscount" bson:"maintenanceDiscount,omitempty"`
	Created             Created            `json:"created" bson:"created,omitempty"`
	Updated             []Updated          `json:"updated" bson:"updated,omitempty"`
	SortOrder           int64              `json:"sortOrder"  bson:"sortOrder,omitempty"`
}

//RefFloorType :""
type RefFloorType struct {
	FloorType `bson:",inline"`
	Ref       struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//FloorTypeFilter : ""
type FloorTypeFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
