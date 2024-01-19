package models

//Block : "Holds single state data"
type Block struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`

	Name         string     `json:"name" bson:"name,omitempty"`
	Code         string     `json:"code"  bson:"code,omitempty"`
	DistrictCode string     `json:"districtCode,omitempty"  bson:"districtCode,omitempty"`
	Status       string     `json:"status"  bson:"status,omitempty"`
	Created      *CreatedV2 `json:"created"  bson:"created,omitempty"`
	Languages    []string   `json:"languages"  bson:"languages,omitempty"`
	Updated      []Updated  `json:"updated"  bson:"updated,omitempty"`
}

//RefBlock : "Village with refrence data such as language..."
type RefBlock struct {
	Village `bson:",inline"`
	Ref     struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//BlockFilter : "Used for constructing filter query"
type BlockFilter struct {
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`

	Codes         []string `json:"codes,omitempty" bson:"codes,omitempty"`
	DistrictCodes []string `json:"districtCodes,omitempty" bson:"districtCodes,omitempty"`
	Status        []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy        string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder     int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
