package models

//GramPanjayat : "Holds single state data"
type GramPanchayat struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name      string     `json:"name" bson:"name,omitempty"`
	Code      string     `json:"code"  bson:"code,omitempty"`
	BlockCode string     `json:"blockCode,omitempty"  bson:"blockCode,omitempty"`
	Status    string     `json:"status"  bson:"status,omitempty"`
	Created   *CreatedV2 `json:"created"  bson:"created,omitempty"`
	Languages []string   `json:"languages"  bson:"languages,omitempty"`
	Updated   []Updated  `json:"updated"  bson:"updated,omitempty"`
}

//RefGramPanjayat : "Village with refrence data such as language..."
type RefGramPanchayat struct {
	Village `bson:",inline"`
	Ref     struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//GramPanjayatFilter : "Used for constructing filter query"
type GramPanchayatFilter struct {
	UniqueID   []string `json:"name" bson:"name,omitempty"`
	Codes      []string `json:"codes,omitempty" bson:"codes,omitempty"`
	BlockCodes []string `json:"blockCodes,omitempty" bson:"blockCodes,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy     string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder  int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
