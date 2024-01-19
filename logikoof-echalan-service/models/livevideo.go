package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//LiveVideo : ""
type LiveVideo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChannelName string             `json:"channelName" bson:"channelName,omitempty"`
	URL         string             `json:"url" bson:"url,omitempty"`
	UniqueID    string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
	Created     Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Updated     Updated            `json:"updated"  bson:"updated,omitempty"`
	UpdateLog   []Updated          `json:"updatedLog" bson:"updatedLog,omitempty"`
}

//RefLiveVideo :""
type RefLiveVideo struct {
	LiveVideo `bson:",inline"`
	Ref       struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//LiveVideoFilter : ""
type LiveVideoFilter struct {
	Status    []string `json:"status"`
	RegNo     []string `json:"regNo"`
	OmitID    []string `json:"omitId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
