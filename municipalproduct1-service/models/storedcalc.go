package models

type StoredCalc struct {
	UniqueID      string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	Fys           string  `json:"fys" bson:"fys,omitempty"`
	Property      string  `json:"property" bson:"property,omitempty"`
	Status        string  `json:"status"  bson:"status,omitempty"`
	VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
	SumFloorTax   float64 `json:"sumFloorTax" bson:"sumFloorTax"`
	Created       Created `json:"created"  bson:"created,omitempty"`
}
type RefStoredCalc struct {
	StoredCalc `bson:",inline"`
	Ref        struct {
	} `json:"ref" bson:"ref,omitempty"`
}
type StoredCalcFilter struct {
	UniqueID  []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
