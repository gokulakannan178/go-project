package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceRequest struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty"`
	MinUser          *MinUser           `json:"minUser,omitempty" bson:"minUser,omitempty"`
	Location         *Location          `json:"location,omitempty" bson:"location,omitempty"`
	GCUser           *GCUser            `json:"gcUser,omitempty" bson:"gcUser,omitempty"`
	Citizen          *Citizen           `json:"citizen,omitempty" bson:"citizen,omitempty"`
	WardNo           string             `json:"wardNo,omitempty" bson:"wardNo,omitempty"`
	Description      string             `json:"description,omitempty" bson:"description,omitempty"`
	ClosingNotes     string             `json:"closingNotes,omitempty" bson:"closingNotes,omitempty"`
	OrganisationID   string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	AssignDate       *time.Time         `json:"assignDate" bson:"assignDate,omitempty"`
	CompletionDate   *time.Time         `json:"completionDate" bson:"completionDate,omitempty"`
	RequstedDate     *time.Time         `json:"requstedDate" bson:"requstedDate,omitempty"`
	IssueImageurl    []string           `json:"issueImageUrl" bson:"issueImageUrl,omitempty"`
	SolutionImageurl []string           `json:"solutionImageUrl" bson:"solutionImageUrl,omitempty"`
	EndDate          *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	Created          *CreatedV2         `json:"created" bson:"created,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
	WandAdmin        string             `json:"wardAdmin" bson:"wardAdmin,omitempty"`
}

type RefServiceRequest struct {
	ServiceRequest `bson:",inline"`
	Ref            struct {
		OrganisationID Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterServiceRequest struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	Name      []string `json:"name" bson:"name,omitempty"`
	UniqueID  []string `json:"uniqueID" bson:"uniqueID,omitempty"`
	GCID      []string `json:"gcId" bson:"gcId,omitempty"`
	ManagerID []string `json:"managerId" bson:"managerId,omitempty"`
	CitizenID []string `json:"citizenId" bson:"citizenId,omitempty"`

	Regex struct {
		Name        string `json:"name" bson:"name"`
		ManagerName string `json:"managerName" bson:"managerName"`
		GCName      string `json:"gcName" bson:"gcName"`
		CitizenName string `json:"citizenName" bson:"citizenName"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string  `json:"sortBy"`
	SortOrder float64 `json:"sortOrder"`
}
type CompletedServiceRequest struct {
	UniqueID     string `json:"uniqueID" bson:"uniqueID,omitempty"`
	ClosingNotes string `json:"closingNotes,omitempty" bson:"closingNotes,omitempty"`
	Imageurl     string `json:"imageUrl" bson:"imageUrl,omitempty"`
}
