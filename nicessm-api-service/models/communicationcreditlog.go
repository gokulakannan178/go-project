package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CommunicationCreditLog : ""
type CommunicationCreditLog struct {
	ID                primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID          string             `json:"uniqueID,omitempty"  bson:"uniqueID,omitempty"`
	PreCredit         float64            `json:"preCredit"  bson:"preCredit,omitempty"`
	PostCredit        float64            `json:"postCredit"  bson:"postCredit,omitempty"`
	Credit            float64            `json:"credit"  bson:"credit,omitempty"`
	CommunicationMode string             `json:"communicationMode,omitempty"  bson:"communicationMode,omitempty"`
	Status            string             `json:"status,omitempty"  bson:"status,omitempty"`
	CreateData        *time.Time         `json:"createData,omitempty"  bson:"createData,omitempty"`
	Created           *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type CommunicationCreditLogFilter struct {
	Status            []string `json:"status,omitempty" bson:"status,omitempty"`
	CommunicationMode []string `json:"communicationMode,omitempty"  bson:"communicationMode,omitempty"`
}

type RefCommunicationCreditLog struct {
	CommunicationCreditLog `bson:",inline"`
	Ref                    struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
