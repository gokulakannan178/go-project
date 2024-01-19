package models

type Advertisement struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Title    string   `json:"title" bson:"title,omitempty"`
	Status   string   `json:"status" bson:"status,omitempty"`
	Created  *Created `json:"createdOn" bson:"createdOn,omitempty"`
	Desc     string   `json:"description" bson:"description,omitempty"`
	Type     string   `json:"type" bson:"type,omitempty"`
	ImgUrl   string   `json:"imgurl" bson:"imgurl,omitempty"`
}
type AdvertisementFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}
type RefAdvertisement struct {
	Advertisement `bson:",inline"`
}
