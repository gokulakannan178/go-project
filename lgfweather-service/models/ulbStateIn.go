package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Authentication : "using for auth"
type ULBStateIn struct {
	ID                primitive.ObjectID `json:"id bson:"id,omitempty"`
	StateID           string             `json:"stateID" bson:"stateID,omitempty"`
	CertificateStatus []string           `json:"certificateStatus" bson:"certificateStatus,omitempty"`
}
