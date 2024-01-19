package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//BillclaimConfig : "Holds single billclaimconfig data"
type BillclaimConfig struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID     string             `json:"uniqueID" bson:"uniqueID,omitempty"`
	Grade        string             `json:"grade"  bson:"grade,omitempty"`
	Level        int64              `json:"level"  bson:"level,omitempty"`
	MinAmount    float64            `json:"minAmount"  bson:"minAmount,omitempty"`
	MaxAmount    float64            `json:"maxAmount"  bson:"maxAmount,omitempty"`
	Organisation string             `json:"organisation"  bson:"organisation,omitempty"`
	Desc         string             `json:"desc"  bson:"desc,omitempty"`
	Status       string             `json:"status"  bson:"status,omitempty"`
	Created      Created            `json:"created"  bson:"created,omitempty"`
	Updated      []Updated          `json:"updated"  bson:"updated,omitempty"`
}

//RefBillclaimConfig : "BillclaimConfig with refrence data such as language..."
type RefBillclaimConfig struct {
	BillclaimConfig `bson:",inline"`
	Ref             struct {
		OrganisationID Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//BillclaimConfigFilter : "Used for constructing filter query"
type BillclaimConfigFilter struct {
	Codes          []string `json:"codes,omitempty" bson:"codes,omitempty"`
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	SortBy         string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder      int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
