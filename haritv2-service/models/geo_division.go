package models

//Division : "Holds single state data"
type Division struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name      string     `json:"name" bson:"name,omitempty"`
	Code      string     `json:"code"  bson:"code,omitempty"`
	StateCode string     `json:"stateCode,omitempty"  bson:"stateCode,omitempty"`
	Status    string     `json:"status"  bson:"status,omitempty"`
	Created   *CreatedV2 `json:"created"  bson:"created,omitempty"`
	Languages []string   `json:"languages"  bson:"languages,omitempty"`
	Updated   []Updated  `json:"updated"  bson:"updated,omitempty"`
}

//RefDivision : "Division with refrence data such as language..."
type RefDivision struct {
	Division `bson:",inline"`
	Ref      struct {
		State *State `json:"state,omitempty" bson:"state,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DivisionFilter : "Used for constructing filter query"
type DivisionFilter struct {
	Codes     []string `json:"codes,omitempty" bson:"codes,omitempty"`
	UniqueID  []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	StateCode []string `json:"stateCode,omitempty"  bson:"stateCode,omitempty"`
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
