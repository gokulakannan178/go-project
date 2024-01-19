package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Dealer : "Holds single KnowlegdeDomain data"
type Dealer struct {
	ID            primitive.ObjectID  `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID      string              `json:"uniqueId"  bson:"uniqueId,omitempty"`
	Name          string              `json:"name" bson:"name,omitempty"`
	Mobile        string              `json:"mobile"  bson:"mobile,omitempty"`
	Email         string              `json:"email"  bson:"email,omitempty"`
	Status        string              `json:"status" bson:"status,omitempty"`
	ActiveStatus  bool                `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created       Created             `json:"createdOn" bson:"createdOn,omitempty"`
	GramPanchayat primitive.ObjectID  `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village       primitive.ObjectID  `json:"village"  bson:"village,omitempty"`
	State         primitive.ObjectID  `json:"state"  bson:"state,omitempty"`
	Block         primitive.ObjectID  `json:"block"  bson:"block,omitempty"`
	District      primitive.ObjectID  `json:"district"  bson:"district,omitempty"`
	PinCode       float64             `json:"pinCode"  bson:"pinCode,omitempty"`
	Location      Location            `json:"location"  bson:"location,omitempty"`
	Description   string              `json:"description"  bson:"description,omitempty"`
	Certification DealerCertification `json:"certification"  bson:"certification,omitempty"`
}

//RefDealer : "RefDealer with refrence data such as language..."
type RefDealer struct {
	Dealer `bson:",inline"`
	Ref    struct {
		GramPanchayat GramPanchayat `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
		Village       Village       `json:"village"  bson:"village,omitempty"`
		State         State         `json:"state"  bson:"state,omitempty"`
		Block         Block         `json:"block"  bson:"block,omitempty"`
		District      District      `json:"district"  bson:"district,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DealerFilter : "Used for constructing filter query"
type DealerFilter struct {
	GramPanchayat                []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	CertificationStatus          []string             `json:"certificationStatus"  bson:"certificationStatus,omitempty"`
	Village                      []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	ActiveStatus                 []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status                       []string             `json:"status" bson:"status,omitempty"`
	SortBy                       string               `json:"sortBy"`
	SortOrder                    int                  `json:"sortOrder"`
	CertificationExpiryDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"certificationExpiryDateRange"`
	CertificationAppliedDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"certificationAppliedDateRange"`
	Regex struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}

type DealerUniquinessChk struct {
	Success bool   `json:"success" bson:"success,omitempty"`
	Message string `json:"message" bson:"message,omitempty"`
}
type DealerCertification struct {
	Status      string    `json:"status" bson:"status,omitempty"`
	AppliedDate time.Time `json:"appliedDate"  bson:"appliedDate,omitempty"`
	ExpiryDate  time.Time `json:"expiryDate"  bson:"expiryDate,omitempty"`
	ActionDate  time.Time `json:"actionDate"  bson:"actionDate,omitempty"`
	URL         string    `json:"url"  bson:"url,omitempty"`
}
