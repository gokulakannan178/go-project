package models

type PaymentMode struct {
	Name     string  `json:"name" bson:"name,omitempty"`
	UniqueID string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	Desc     string  `json:"desc" bson:"desc,omitempty"`
	Created  Created `json:"created" bson:"created,omitempty"`
	Status   string  `json:"status" bson:"status,omitempty"`
}

//PaymentMode : ""
type PaymentModeFilter struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	SearchText struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"searchText" bson:"searchText"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

//RefPaymentMode : ""
type RefPaymentMode struct {
	PaymentMode `bson:",inline"`
	// Ref     struct {
	// } `json:"ref" bson:"ref,omitempty"`
}
