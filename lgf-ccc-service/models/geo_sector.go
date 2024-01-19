package models

//Sector : "Holds single state data"
type Sector struct {
	Name      string    `json:"name" bson:"name,omitempty"`
	Code      string    `json:"code"  bson:"code,omitempty"`
	WardCode  string    `json:"wardCode,omitempty"  bson:"wardCode,omitempty"`
	Status    string    `json:"status"  bson:"status,omitempty"`
	Created   Created   `json:"created"  bson:"created,omitempty"`
	Languages []string  `json:"languages"  bson:"languages,omitempty"`
	Updated   []Updated `json:"updated"  bson:"updated,omitempty"`
}

//RefSector : "Sector with refrence data such as language..."
type RefSector struct {
	Sector `bson:",inline"`
	Ref    struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
		Village  *District `json:"village,omitempty" bson:"village,omitempty"`
		Zone     *Zone     `json:"zone,omitempty" bson:"zone,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//SectorFilter : "Used for constructing filter query"
type SectorFilter struct {
	Codes     []string `json:"codes,omitempty" bson:"codes,omitempty"`
	ZoneCodes []string `json:"zoneCodes,omitempty" bson:"zoneCodes,omitempty"`
	WardCode  []string `json:"wardCode,omitempty"  bson:"wardCode,omitempty"`
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
