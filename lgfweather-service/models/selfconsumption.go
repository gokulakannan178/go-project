package models

import "time"

//SelfConsumption : ""
type SelfConsumption struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Type      string     `json:"type" bson:"type,omitempty"`
	CompanyID string     `json:"companyId" bson:"companyId"`
	Quantity  float64    `json:"quantity" bson:"quantity"`
	Date      *time.Time `json:"date" bson:"date"`
	By        string     `json:"by" bson:"by"`
	ByTYpe    string     `json:"byType" bson:"byType"`
	Created   CreatedV2  `json:"createdOn" bson:"createdOn,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	ULBID     string     `json:"ulbId" bson:"ulbId,omitempty"`
}

//SelfConsumptionFilter : ""
type SelfConsumptionFilter struct {
	UniqueID             []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	CompanyID            []string `json:"companyId" bson:"companyId"`
	ULBID                []string `json:"ulbId" bson:"ulbId,omitempty"`
	Status               []string `json:"status" bson:"status,omitempty"`
	SortBy               string   `json:"sortBy"`
	SortOrder            int      `json:"sortOrder"`
	SelfConsumptionRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"selfConsumptionRange"`
	Regex struct {
		Name    string `json:"name" bson:"name,omitempty"`
		UlbName string `json:"ulbName" bson:"ulbName,omitempty"`
	} `json:"regex"`
}

//RefSelfConsumption : ""
type RefSelfConsumption struct {
	SelfConsumption `bson:",inline"`
	Ref             struct {
		ULBID    RefULB      `json:"ulbId" bson:"ulbId,omitempty"`
		User     User        `json:"user" bson:"user,omitempty"`
		UserType UserType    `json:"userType" bson:"userType,omitempty"`
		Company  interface{} `json:"company" bson:"company,omitempty"`
		Address  *RefAddress
	} `json:"ref" bson:"ref,omitempty"`
}
