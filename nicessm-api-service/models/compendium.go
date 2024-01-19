package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Compendium : ""
type Compendium struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
	File         string             `json:"file,omitempty" bson:"file,omitempty"`
	DateCreated  *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
	//	Version      int                `json:"version" form:"version" bson:"version,omitempty"`
}
type CompendiumFilter struct {
	Status    []string             `json:"status" form:"status" bson:"status,omitempty"`
	Content   []primitive.ObjectID `json:"content,omitempty" bson:"content,omitempty"`
	SortBy    string               `json:"sortBy"`
	SortOrder int                  `json:"sortOrder"`
	SearchBox struct {
		Comment string `json:"Comment,omitempty" bson:"Comment,omitempty"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefCompendium struct {
	Compendium `bson:",inline"`
	Ref        struct {
		Content Content `json:"content,omitempty" bson:"content,omitempty"`
		User    User    `json:"user" bson:"user,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type CompendiumFileUpload struct {
	File    string  `json:"file,omitempty" bson:"file,omitempty"`
	Content Content `json:"content,omitempty" bson:"content,omitempty"`
	DoSave  bool    `json:"doSave,omitempty" bson:"doSave,omitempty"`
}

type CompendiumData struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Comment        string             `json:"Comment,omitempty" bson:"Comment,omitempty"`
	ActiveStatus   bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
	Content        primitive.ObjectID `json:"content,omitempty" bson:"content,omitempty"`
	DateCreated    *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
	User           primitive.ObjectID `json:"user" bson:"user,omitempty"`
	Version        int                `json:"version" form:"version" bson:"version,omitempty"`
	Tag            string             `json:"tag,omitempty" bson:"tag,omitempty"`
	Date           string             `json:"date,omitempty" bson:"date,omitempty"`
	Title          string             `json:"title,omitempty" bson:"title,omitempty"`
	ContentStrings []string           `json:"contentStrings,omitempty" bson:"contentStrings,omitempty"`
}
