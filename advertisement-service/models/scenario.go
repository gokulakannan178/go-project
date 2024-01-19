package models

type Scenario struct {
	Name     string  `json:"name" bson:"name,omitempty"`
	UniqueID string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	Desc     string  `json:"desc" bson:"desc,omitempty"`
	Created  Created `json:"created" bson:"created,omitempty"`
	Status   string  `json:"status" bson:"status,omitempty"`
	Message  string  `json:"message" bson:"message,omitempty"`
}

//ScenarioFilter : ""
type ScenarioFilter struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	SearchText struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"searchText" bson:"searchText"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

//RefScenario : ""
type RefScenario struct {
	Scenario `bson:",inline"`
	// Ref     struct {
	// } `json:"ref" bson:"ref,omitempty"`
}
