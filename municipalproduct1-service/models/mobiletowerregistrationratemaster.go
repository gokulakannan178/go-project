package models

import (
	"time"
)

type MobileTowerRegistrationRateMaster struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string     `json:"name" form:"name," bson:"name,omitempty"`
	Desc     string     `json:"desc" form:"desc," bson:"desc,omitempty"`
	Rate     float32    `json:"rate" form:"rate," bson:"rate,omitempty"`
	Doe      *time.Time `json:"doe" form:"doe," bson:"doe,omitempty"`
	Status   string     `json:"status" form:"status," bson:"status,omitempty"`
	Created  Created    `json:"createdOn" bson:"createdOn,omitempty"`
}
type MobileTowerRegistrationRateMasterFilter struct {
	Status []string `json:"status" form:"status," bson:"status,omitempty"`
}
type RefMobileTowerRegistrationRateMaster struct {
	MobileTowerRegistrationRateMaster `bson:",inline"`
}
