package models

//PropertyVisitLogRemarkType : "Used for Providing DropDown In While Adding Remarks"
type PropertyVisitLogRemarkType struct {
	UniqueID string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   string    `json:"status" bson:"status,omitempty"`
	Created  CreatedV2 `json:"created,omitempty" bson:"created,omitempty"`
	Updated  []Updated `json:"updated" bson:"updated,omitempty"`
	Name     string    `json:"name" bson:"name,omitempty"`
	Desc     string    `json:"desc" bson:"desc,omitempty"`
	IsEntry  bool      `json:"isEntry" bson:"isEntry,omitempty"`
}

//RefPropertyVisitLogRemarkType : "Used for Providing DropDown In While Adding Remarks"
type RefPropertyVisitLogRemarkType struct {
	PropertyVisitLogRemarkType `bson:",inline"`
}

type PropertyVisitLogRemarkTypeFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}
