package models

type TradeLicenseCategoryType struct {
	UniqueID string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string    `json:"name" bson:"name,omitempty"`
	Desc     string    `json:"desc" bson:"desc,omitempty"`
	Created  CreatedV2 `json:"created" bson:"created,omitempty"`
	Status   string    `json:"status" bson:"status,omitempty"`
	TLBTID   string    `json:"tlbtId" bson:"tlbtId,omitempty"`
}

type TradeLicenseCategoryTypeFilter struct {
	TLBTID []string `json:"tlbtId" bson:"tlbtId,omitempty"`
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefTradeLicenseCategoryType struct {
	TradeLicenseCategoryType `bson:",inline"`
	Ref                      struct {
		TradeLicenseBusinessType TradeLicenseBusinessType `json:"tradeLicenseBusinessType" bson:"tradeLicenseBusinessType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
