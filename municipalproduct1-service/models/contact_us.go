package models

import (
	"fmt"
	"time"
)

type ContactUs struct {
	Name     string     `json:"name" bson:"name,omitempty"`
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Email    string     `json:"email" bson:"email,omitempty"`
	Subject  string     `json:"subject" bson:"subject,omitempty"`
	Message  string     `json:"message" bson:"message,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	Created  *CreatedV2 `json:"created" bson:"created,omitempty"`
}

type RefContactUs struct {
	ContactUs `bson:",inline"`
	Ref       struct {
	} `json:"ref" bson:"ref,omitempty"`
}

func (cu *ContactUs) ConvertEmailMsg() string {
	return fmt.Sprintf("Mr/Ms %v with email id %v has contacted you on %v on behalf of %v - %v", cu.Name, cu.Email, time.Now(), cu.Subject, cu.Message)
}
func (cu *ContactUs) AutoResponse() string {
	return fmt.Sprintf("Hi %v, \n Thanks for contacting us. We will respond to you soon.", cu.Name)
}

type FilterContactUs struct {
	Status []string `json:"status" bson:"status,omitempty"`
}
