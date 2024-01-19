package models

//Village : "Holds single state data"
type Village struct {
	Name         string    `json:"name" bson:"name,omitempty"`
	Code         string    `json:"code"  bson:"code,omitempty"`
	DistrictCode string    `json:"districtCode,omitempty"  bson:"districtCode,omitempty"`
	Status       string    `json:"status"  bson:"status,omitempty"`
	Created      Created   `json:"created"  bson:"created,omitempty"`
	Languages    []string  `json:"languages"  bson:"languages,omitempty"`
	Updated      []Updated `json:"updated"  bson:"updated,omitempty"`
}

//RefVillage : "Village with refrence data such as language..."
type RefVillage struct {
	Village `bson:",inline"`
	Ref     struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//VillageFilter : "Used for constructing filter query"
type VillageFilter struct {
	Codes         []string `json:"codes,omitempty" bson:"codes,omitempty"`
	DistrictCodes []string `json:"districtCodes,omitempty" bson:"districtCodes,omitempty"`
	Status        []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy        string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder     int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
