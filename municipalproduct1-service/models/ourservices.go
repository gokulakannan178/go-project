package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OurService struct {
	ID                             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID                       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OurServiceName                 string             `json:"OurServiceName" bson:"OurServiceName,omitempty"`
	CollectionName                 string             `json:"collectionName" bson:"collectionName,omitempty"`
	Title                          string             `json:"title" bson:"title,omitempty"`
	Name                           string             `json:"name" bson:"name,omitempty"`
	Desc                           string             `json:"description" bson:"description,omitempty"`
	Mobile                         string             `json:"mobile" bson:"mobile,omitempty"`
	Email                          string             `json:"email" bson:"email,omitempty"`
	Designation                    string             `json:"designation" bson:"designation,omitempty"`
	TotalRegister                  string             `json:"totalRegister" bson:"totalRegister,omitempty"`
	Highlights                     string             `json:"highlights" bson:"highlights,omitempty"`
	FileUrl                        string             `json:"fileUrl" bson:"fileUrl,omitempty"`
	ImageUrl                       string             `json:"imageUrl" bson:"imageUrl,omitempty"`
	ProfilePic                     string             `json:"profilePic" bson:"profilePic,omitempty"`
	Status                         string             `json:"status" bson:"status,omitempty"`
	ShowOfficeDirectoryInDashboard string             `json:"showOfficeDirectoryInDashboard" bson:"showOfficeDirectoryInDashboard,omitempty"`
	Date                           *time.Time         `json:"date"  bson:"date,omitempty"`
	Created                        *Created           `json:"created"  bson:"created,omitempty"`
}

type RefOurService struct {
	OurService `bson:",inline"`
	Ref        struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterOurService struct {
	Status                         []string `json:"status,omitempty" bson:"status,omitempty"`
	ShowOfficeDirectoryInDashboard []string `json:"showOfficeDirectoryInDashboard,omitempty" bson:"showOfficeDirectoryInDashboard,omitempty"`
	SortBy                         string   `json:"sortBy"`
	SortOrder                      int      `json:"sortOrder"`
	Regex                          struct {
		Name  string `json:"name" bson:"name"`
		Title string `json:"title" bson:"title,omitempty"`
	} `json:"regex" bson:"regex"`
}
