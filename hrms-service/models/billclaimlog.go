package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillClaimLog struct {
	ID        primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID  string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	BillClaim string             `json:"billClaim,omitempty" bson:"billClaim,omitempty"`
	Log       Updated            `json:"log" bson:"log,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
	New       BillClaim          `json:"new,omitempty" bson:"new,omitempty"`
	Previous  RefBillClaim       `json:"previous,omitempty" bson:"previous,omitempty"`
}

type RefBillClaimLog struct {
	BillClaimLog `bson:",inline"`
	Ref          struct {
		//	BillClaim RefBillClaim `json:"billClaim,omitempty" bson:"billClaim,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterBillClaimLog struct {
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeId []string `json:"employeeId" bson:"employeeId,omitempty"`
	Regex      struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
