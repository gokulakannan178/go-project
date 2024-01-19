package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Dissemination : ""
type Dissemination struct {
	ID                  primitive.ObjectID    `json:"id"  bson:"id,omitempty"`
	Content             primitive.ObjectID    `json:"content"  bson:"content,omitempty"`
	ContentTranslation  primitive.ObjectID    `json:"contentTranslation"  bson:"contentTranslation,omitempty"`
	CreatedBy           primitive.ObjectID    `json:"createdBy"  bson:"createdBy,omitempty"`
	DateCreated         *time.Time            `json:"dateCreated"  bson:"dateCreated,omitempty"`
	DateOfDissemination *time.Time            `json:"dateOfDissemination"  bson:"dateOfDissemination,omitempty"`
	Farmers             []FarmerDissemination `json:"farmers"  bson:"farmers,omitempty"`
	Users               []UserDissemination   `json:"users"  bson:"users,omitempty"`
	FarmersCount        int                   `json:"farmersCount"  bson:"farmersCount,omitempty"`
	UsersCount          int                   `json:"usersCount"  bson:"usersCount,omitempty"`
	IsSent              bool                  `json:"isSent" bson:"isSent,omitempty"`
	Status              string                `json:"status" bson:"status,omitempty"`
	Message             string                `json:"message" bson:"message,omitempty"`
	Created             *Created              `json:"created"  bson:"created,omitempty"`
	Version             string                `json:"version" bson:"version,omitempty"`
	Type                string                `json:"type" bson:"type,omitempty"`
	Mode                string                `json:"mode" bson:"mode,omitempty"`
}
type DisseminationFilter struct {
	IsSent       []bool               `json:"isSent,omitempty"  bson:"isSent,omitempty"`
	Content      []primitive.ObjectID `json:"content"  bson:"content,omitempty"`
	CreatedBy    []primitive.ObjectID `json:"createdBy"  bson:"createdBy,omitempty"`
	Status       []string             `json:"status" form:"status" bson:"status,omitempty"`
	Organisation []primitive.ObjectID `json:"organisation"  bson:"organisation,omitempty"`
	Project      []primitive.ObjectID `json:"project"  bson:"project,omitempty"`
	State        []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	District     []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Block        []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	Mode         []string             `json:"mode" bson:"mode,omitempty"`

	DateDisseminationRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateDisseminationRange"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
	Regex     struct {
		Message string `json:"message" bson:"message,omitempty"`
	} `json:"regex" bson:"regex"`
}
type DisseminationReportFilter struct {
	State                  []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	District               []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	DateDisseminationRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateDisseminationRange"`
}
type RefDisseminationReport struct {
	//ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Ref struct {
		State State `json:"state"  bson:"state,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
	Districts []struct {
		// ID struct {
		// 	State    primitive.ObjectID `json:"state"  bson:"state,omitempty"`
		// 	District primitive.ObjectID `json:"district"  bson:"district,omitempty"`
		// } `json:"id"  bson:"_id,omitempty"`
		Farmers float64 `json:"farmer"  bson:"farmer,omitempty"`
		//	State          primitive.ObjectID `json:"state"  bson:"state,omitempty"`
		Dessiminations float64  `json:"dessiminations"  bson:"dessiminations,omitempty"`
		Sms            float64  `json:"sms"  bson:"sms,omitempty"`
		Voice          float64  `json:"voice"  bson:"voice,omitempty"`
		Poster         float64  `json:"poster"  bson:"poster,omitempty"`
		Document       float64  `json:"document"  bson:"document,omitempty"`
		Video          float64  `json:"video"  bson:"video,omitempty"`
		Count          float64  `json:"count"  bson:"count,omitempty"`
		District       District `json:"district"  bson:"district,omitempty"`
	} `json:"districts"  bson:"districts,omitempty"`
}
type RefDissemination struct {
	Dissemination `bson:",inline"`
	Ref           struct {
		Content   Content `json:"content"  bson:"content,omitempty"`
		CreatedBy User    `json:"createdBy"  bson:"createdBy,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type FarmerDissemination struct {
	ID                   primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	Name                 string             `json:"name" form:"name" bson:"name,omitempty"`
	MobileNumber         string             `json:"mobileNumber" form:"status" bson:"mobileNumber,omitempty"`
	AppRegistrationToken string             `json:"appRegistrationToken" bson:"appRegistrationToken,omitempty"`
}

type UserDissemination struct {
	ID                   primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	Name                 string             `json:"name" form:"name" bson:"name,omitempty"`
	AppRegistrationToken string             `json:"appRegistrationToken" bson:"appRegistrationToken,omitempty"`
	MobileNumber         string             `json:"mobileNumber" form:"status" bson:"mobileNumber,omitempty"`
}
