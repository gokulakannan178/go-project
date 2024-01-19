package models

//State : "Holds single state data"
type State struct {
	Name      string    `json:"name" bson:"name,omitempty"`
	Code      string    `json:"code"  bson:"code,omitempty"`
	Desc      string    `json:"desc"  bson:"desc,omitempty"`
	Status    string    `json:"status"  bson:"status,omitempty"`
	Created   Created   `json:"created"  bson:"created,omitempty"`
	Languages []string  `json:"languages"  bson:"languages,omitempty"`
	Updated   []Updated `json:"updated"  bson:"updated,omitempty"`
}

//RefState : "State with refrence data such as language..."
type RefState struct {
	State `bson:",inline"`
	Ref   struct {
		Languages []Language `json:"languages,omitempty" bson:"languages,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//StateFilter : "Used for constructing filter query"
type StateFilter struct {
	Codes     []string `json:"codes,omitempty" bson:"codes,omitempty"`
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
