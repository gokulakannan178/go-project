package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//PropertyOwner : ""
type PropertyOwner struct {
	Title              string             `json:"title" bson:"title,omitempty"`
	ID                 primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID         string             `json:"propertyId" bson:"propertyId,omitempty"`
	Honoriffic         string             `json:"honoriffic" bson:"honoriffic,omitempty"`
	Name               string             `json:"name" bson:"name,omitempty"`
	Mobile             string             `json:"mobile" bson:"mobile,omitempty"`
	Email              string             `json:"email" bson:"email,omitempty"`
	Gender             string             `json:"gender" bson:"gender,omitempty"`
	FatherRpanRhusband string             `json:"fatherRpanRhusband" bson:"fatherRpanRhusband,omitempty"`
	GaurdianName       string             `json:"gaurdianName" bson:"gaurdianName,omitempty"`
	Relation           string             `json:"relation" bson:"relation,omitempty"`
	Address            string             `json:"address" bson:"address,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	Created            Created            `json:"created" bson:"created,omitempty"`
	Updated            []Updated          `json:"updated" bson:"updated,omitempty"`
	NewPropertyID      string             `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID      string             `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

//RefPropertyOwner :""
type RefPropertyOwner struct {
	PropertyOwner `bson:",inline"`
	Ref           struct {
		Honoriffic *Honoriffic `json:"honoriffic" bson:"honoriffic,omitempty"`
		Relation   *Relation   `json:"relation" bson:"relation,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

func (res *RefPropertyOwner) IncOwner(a int) int {
	return a + 1
}

//PropertyOwnerFilter : ""
type PropertyOwnerFilter struct {
	Status    []string `json:"status"`
	Mobile    []string `json:"mobile"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}

type PropertyIDsWithOwnerNames struct {
	PropertyIDs []string `json:"propertyIds" bson:"propertyIds,omitempty"`
}

type PropertyIDsWithMobileNos struct {
	PropertyIDs []string `json:"propertyIds" bson:"propertyIds,omitempty"`
}
