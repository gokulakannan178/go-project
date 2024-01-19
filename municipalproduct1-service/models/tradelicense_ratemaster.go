package models

import "time"

type TradeLicenseRateMaster struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string     `json:"name" bson:"name,omitempty"`
	Desc     string     `json:"desc" bson:"desc,omitempty"`
	Created  CreatedV2  `json:"created" bson:"created,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	Type     string     `json:"type" bson:"type,omitempty"`
	TLCTID   string     `json:"tlctId" bson:"tlctId,omitempty"`
	TLBTID   string     `json:"tlbtId" bson:"tlbtId,omitempty"`
	Rate     float64    `json:"rate" bson:"rate,omitempty"`
	DOE      *time.Time `json:"doe" bson:"doe,omitempty"`
}

type TradeLicenseRateMasterFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
	TLCTID []string `json:"tlctId" bson:"tlctId,omitempty"`
	TLBTID []string `json:"tlbtId" bson:"tlbtId,omitempty"`
}

type RefTradeLicenseRateMaster struct {
	TradeLicenseRateMaster `bson:",inline"`
	Ref                    struct {
		TradeLicenseBusinessType TradeLicenseBusinessType `json:"tradeLicenseBusinessType" bson:"tradeLicenseBusinessType,omitempty"`
		TradeLicenseCategoryType TradeLicenseCategoryType `json:"tradeLicenseCategoryType" bson:"tradeLicenseCategoryType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
