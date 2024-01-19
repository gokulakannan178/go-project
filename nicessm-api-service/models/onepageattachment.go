package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//OnePageAttachment : ""
type OnePageAttachment struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Image       string             `json:"image,omitempty"  bson:"image,omitempty"`
	ImageHeight string             `json:"imageHeight,omitempty"  bson:"imageHeight,omitempty"`
	ImageWidth  string             `json:"imageWidth,omitempty"  bson:"imageWidth,omitempty"`
	Status      string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created     *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type OnePageAttachmentFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefOnePageAttachment struct {
	OnePageAttachment `bson:",inline"`
	Ref               struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
