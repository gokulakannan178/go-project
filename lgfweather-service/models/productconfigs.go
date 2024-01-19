package models

type ProductConfig struct {
	UniqueID                string  `json:"Uniqueid" form:"UniqueId," bson:"UniqueId,omitempty"`
	Status                  string  `json:"status" form:"status," bson:"status,omitempty"`
	ContactUsMobile         string  `json:"contactusmobile" form:"contactusmobile," bson:"contactusmobile,omitempty"`
	Created                 Created `json:"createdOn" bson:"createdOn,omitempty"`
	CeritificateExpiryMonth int     `json:"ceritificateExpiryMonth" bson:"ceritificateExpiryMonth,omitempty"`
}
