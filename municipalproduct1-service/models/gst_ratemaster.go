package models


//GSTRateMaster : ""
type GSTRateMaster struct {
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Rate     float64            `json:"rate" bson:"rate,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}


//GSTRateMasterFilter : ""
type GSTRateMasterFilter struct {
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name       []string `json:"name,omitempty" bson:"name,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
}
