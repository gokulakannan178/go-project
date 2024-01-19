package models

//District : "Holds single district data"
type District struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`

	Name         string     `json:"name,omitempty"  bson:"name,omitempty"`
	Code         string     `json:"code,omitempty"  bson:"code,omitempty"`
	StateCode    string     `json:"stateCode,omitempty"  bson:"stateCode,omitempty"`
	DivisionCode string     `json:"divisionCode,omitempty"  bson:"divisionCode,omitempty"`
	Status       string     `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *CreatedV2 `json:"created,omitempty"  bson:"created,omitempty"`
	Languages    []string   `json:"languages,omitempty" bson:"languages,omitempty"`
	Updated      []Updated  `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefDistrict : ""
type RefDistrict struct {
	District `bson:",inline"`
	Ref      struct {
		Division *Division `json:"division,omitempty" bson:"division,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DistrictFilter : "Used for constructing filter query"
type DistrictFilter struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`

	Codes        []string `json:"codes,omitempty" bson:"codes,omitempty"`
	DivisionCode []string `json:"divisionCode,omitempty"  bson:"divisionCode,omitempty"`
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder    int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
