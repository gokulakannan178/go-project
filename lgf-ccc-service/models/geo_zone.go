package models

//Zone : "Holds single state data"
type Zone struct {
	Name        string    `json:"name" bson:"name,omitempty"`
	Code        string    `json:"code"  bson:"code,omitempty"`
	WardCode    string    `json:"wardCode,omitempty"  bson:"wardCode,omitempty"`
	VillageCode string    `json:"villageCode,omitempty"  bson:"villageCode,omitempty"`
	Status      string    `json:"status"  bson:"status,omitempty"`
	Created     Created   `json:"created"  bson:"created,omitempty"`
	Languages   []string  `json:"languages"  bson:"languages,omitempty"`
	Updated     []Updated `json:"updated"  bson:"updated,omitempty"`
}

//RefZone : "Zone with refrence data such as language..."
type RefZone struct {
	Zone `bson:",inline"`
	Ref  struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
		Village  *Village  `json:"village,omitempty" bson:"village,omitempty"`
		Ward     *Ward     `json:"ward,omitempty" bson:"ward,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ZoneFilter : "Used for constructing filter query"
type ZoneFilter struct {
	Codes        []string `json:"codes,omitempty" bson:"codes,omitempty"`
	VillageCodes []string `json:"villageCodes,omitempty" bson:"villageCodes,omitempty"`
	WardCode     []string `json:"wardCode,omitempty"  bson:"wardCode,omitempty"`
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder    int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
