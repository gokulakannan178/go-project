package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Project : ""
type Project struct {
	ID                primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	Name              string               `json:"name,omitempty"  bson:"name,omitempty"`
	ActiveStatus      bool                 `json:"activeStatus,omitempty"  bson:"activeStatus,omitempty"`
	UniqueID          string               `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status            string               `json:"status,omitempty"  bson:"status,omitempty"`
	Budget            float64              `json:"budget,omitempty"  bson:"budget,omitempty"`
	NationalLevel     bool                 `json:"nationalLevel"  bson:"nationalLevel"`
	Mail              string               `json:"mail,omitempty"  bson:"mail,omitempty"`
	Organisation      primitive.ObjectID   `json:"organisation,omitempty" bson:"organisation,omitempty"`
	KnowledgeDomainID []primitive.ObjectID `json:"knowledgeDomainId,omitempty" bson:"-"`
	StateID           []primitive.ObjectID `json:"stateId,omitempty" bson:"-"`
	PartnerID         []primitive.ObjectID `json:"partnerId,omitempty" bson:"-"`
	StartDate         *time.Time           `json:"startDate,omitempty"  bson:"startDate,omitempty"`
	EndDate           *time.Time           `json:"endDate,omitempty"  bson:"endDate,omitempty"`
	Version           float64              `json:"version,omitempty"  bson:"version,omitempty"`
	Created           *CreatedV2           `json:"created,omitempty"  bson:"created,omitempty"`
	Remarks           string               `json:"remarks"  bson:"remarks,omitempty"`
}

type ProjectFilter struct {
	Status        []string             `json:"status,omitempty" bson:"status,omitempty"`
	OmitID        []string             `json:"omitId" bson:"omitId,omitempty"`
	NationalLevel []bool               `json:"nationalLevel,omitempty"  bson:"nationalLevel,omitempty"`
	Organisation  []primitive.ObjectID `json:"organisation,omitempty"  bson:"organisation,omitempty"`
	BudgetRange   *struct {
		From float64 `json:"from"`
		To   float64 `json:"to"`
	} `json:"budgetRange"`
	StartDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"startDateRange"`
	EndDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"endDateRange"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
	Regex     struct {
		Name string `json:"name" bson:"name"`
		Mail string `json:"mail,omitempty"  bson:"mail,omitempty"`
	} `json:"regex" bson:"regex"`
	UserAccess         UserAccess `json:"userAccess"`
	OmitProjectPartner struct {
		Is      bool               `json:"is"`
		Project primitive.ObjectID `json:"project"`
	} `json:"omitProjectPartner"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}

type RefProject struct {
	Project `bson:",inline"`
	Ref     struct {
		Organisation Organisation `json:"organisation,omitempty" bson:"organisation,omitempty"`
		//States          []RefProjectState    `json:"states,omitempty" bson:"states,omitempty"`
		StateIDs          []primitive.ObjectID `json:"stateIDs,omitempty" bson:"stateIDs,omitempty"`
		KnowledgeDomaiIDs []primitive.ObjectID `json:"knowledgeDomaiIDs,omitempty" bson:"knowledgeDomaiIDs,omitempty"`
		ProjectPartnerIDs []primitive.ObjectID `json:"projectPartnerIDs,omitempty" bson:"projectPartnerIDs,omitempty"`
		//KnowledgeDomain []RefKnowledgeDomain `json:"knowledgeDomain,omitempty" bson:"knowledgeDomain,omitempty"`
		//ProjectPartner  []RefProjectPartner  `json:"partner,omitempty" bson:"partner,omitempty"`
		ProjectUsers []RefProjectUser    `json:"projectUsers,omitempty" bson:"projectUsers,omitempty"`
		Partners     []RefProjectPartner `json:"partners" bson:"partners,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
