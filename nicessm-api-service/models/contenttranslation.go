package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ContentTranslation : ""
type ContentTranslation struct {
	ID                primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Content           primitive.ObjectID `json:"content,omitempty" bson:"content,omitempty"`
	DateCreated       *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
	Language          primitive.ObjectID `json:"language,omitempty" bson:"language,omitempty"`
	ActiveStatus      bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	Status            string             `json:"status,omitempty" bson:"status,omitempty"`
	TranslatedTitle   string             `json:"translatedTitle,omitempty" bson:"translatedTitle,omitempty"`
	Linktype          string             `json:"linktype,omitempty" bson:"linktype,omitempty"`
	Translator        primitive.ObjectID `json:"translator" bson:"translator,omitempty"`
	TranslatedContent string             `json:"translatedContent,omitempty" bson:"translatedContent,omitempty"`
	Version           int                `json:"version" form:"version" bson:"version,omitempty"`
	DateReviewed      *time.Time         `json:"dateReviewed" form:"dateReviewed" bson:"dateReviewed,omitempty"`
	ReviewedBy        primitive.ObjectID `json:"reviewedBy" bson:"reviewedBy,omitempty"`
}
type ContentTranslationFilter struct {
	Status              []string             `json:"status" form:"status" bson:"status,omitempty"`
	ContentType         []string             `json:"contentType" form:"contentType" bson:"contentType,omitempty"`
	ContentOrganisation []primitive.ObjectID `json:"contentOrganisation" form:"contentOrganisation" bson:"contentOrganisation,omitempty"`
	ContentProject      []primitive.ObjectID `json:"contentProject" form:"contentProject" bson:"contentProject,omitempty"`
	ContentStatus       []string             `json:"contentStatus" form:"contentStatus" bson:"contentStatus,omitempty"`
	Content             []primitive.ObjectID `json:"content,omitempty" bson:"content,omitempty"`
	SortBy              string               `json:"sortBy"`
	SortOrder           int                  `json:"sortOrder"`
	SearchBox           struct {
		TranslatedContent string `json:"translatedContent,omitempty" bson:"translatedContent,omitempty"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefContentTranslation struct {
	ContentTranslation `bson:",inline"`
	Ref                struct {
		Content      Content      `json:"content,omitempty" bson:"content,omitempty"`
		Translator   User         `json:"translator" bson:"translator,omitempty"`
		Language     Language     `json:"language,omitempty" bson:"language,omitempty"`
		ReviewedBy   User         `json:"reviewedBy" bson:"reviewedBy,omitempty"`
		Organisation Organisation `json:"organisation" bson:"organisation,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type ApprovedContentTranslation struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	ReviewedBy  primitive.ObjectID `json:"reviewedBy" bson:"reviewedBy,omitempty"`
	DateCreated *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
}
type RejectedContentTranslation struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	ReviewedBy  primitive.ObjectID `json:"reviewedBy" bson:"reviewedBy,omitempty"`
	DateCreated *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
}
