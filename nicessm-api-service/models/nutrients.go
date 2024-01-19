package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Nutrients : ""
type Nutrients struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Name         string             `json:"name,omitempty"  bson:"name,omitempty"`
	Code         string             `json:"code,omitempty"  bson:"code,omitempty"`
	MaxVal       string             `json:"maxVal" bson:"maxVal,omitempty"`
	MinVal       string             `json:"minVal" bson:"minVal,omitempty"`
	Version      int                `json:"version,omitempty"  bson:"version,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Type         string             `json:"type,omitempty"  bson:"type,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}
type NutrientsFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Type         string   `json:"type,omitempty"  bson:"type,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
		Code string `json:"code,omitempty"  bson:"code,omitempty"`
		Type string `json:"type,omitempty"  bson:"type,omitempty"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefNutrients struct {
	Nutrients `bson:",inline"`
	Ref       struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
