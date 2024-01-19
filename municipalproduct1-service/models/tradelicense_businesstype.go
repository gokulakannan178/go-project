package models

type TradeLicenseBusinessType struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string     `json:"name" bson:"name,omitempty"`
	Desc     string     `json:"desc" bson:"desc,omitempty"`
	Created  *CreatedV2 `json:"created,omitempty" bson:"created,omitempty"`

	Status string `json:"status" bson:"status,omitempty"`
}

type TradeLicenseBusinessTypeFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefTradeLicenseBusinessType struct {
	TradeLicenseBusinessType `bson:",inline"`
}
