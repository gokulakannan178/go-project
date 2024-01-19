package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CommunicationCredit : ""
type CommunicationCredit struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueID,omitempty"  bson:"uniqueID,omitempty"`
	BalanceCredit    float64            `json:"balanceCredit,omitempty"  bson:"balanceCredit,omitempty"`
	ChartCountCredit int                `json:"chartCountCredit,omitempty"  bson:"chartCountCredit,omitempty"`
	Status           string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created          *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type CommunicationCreditFilter struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
}

type RefCommunicationCredit struct {
	CommunicationCredit `bson:",inline"`
	Ref                 struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
