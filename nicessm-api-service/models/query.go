package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Query : "Holds single Query data"
type Query struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Status          string             `json:"status" bson:"status,omitempty"`
	Query           string             `json:"query" bson:"query,omitempty"`
	UniqueID        string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	QueryType       string             `json:"queryType" bson:"queryType,omitempty"`
	ActiveStatus    bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	IsSolutionSent  bool               `json:"isSolutionSent" bson:"isSolutionSent,omitempty"`
	Created         Created            `json:"createdOn" bson:"createdOn,omitempty"`
	CreatedBy       primitive.ObjectID `json:"createdBy" bson:"createdBy,omitempty"`
	CreatedType     string             `json:"createdType" bson:"createdType,omitempty"`
	Solution        string             `json:"solution" bson:"solution,omitempty"`
	Title           string             `json:"title" bson:"title,omitempty"`
	Date            *time.Time         `json:"date" bson:"date,omitempty"`
	ResolvedDate    *time.Time         `json:"resolvedDate" bson:"resolvedDate,omitempty"`
	ContentID       primitive.ObjectID `json:"contentId"  bson:"contentId,omitempty"`
	KnowledgeDomain primitive.ObjectID `json:"knowledgeDomain"  bson:"knowledgeDomain,omitempty"`
	SubDomain       primitive.ObjectID `json:"subDomain"  bson:"subDomain,omitempty"`
	GramPanchayat   primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village         primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	Version         int64              `json:"version"  bson:"version,omitempty"`
	AssignedTo      primitive.ObjectID `json:"assignedTo"  bson:"assignedTo,omitempty"`
	AssignedDate    *time.Time         `json:"assignedDate"  bson:"assignedDate,omitempty"`
	ResolvedBy      primitive.ObjectID `json:"resolvedBy"  bson:"resolvedBy,omitempty"`
	Rating          string             `json:"rating"  bson:"rating,omitempty"`
	State           primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block           primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District        primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Images          []string           `json:"images"  bson:"images,omitempty"`
	Farmer          primitive.ObjectID `json:"farmer"  bson:"farmer,omitempty"`
}

//RefQuery : "Query with refrence data such as language..."
type RefQuery struct {
	Query `bson:",inline"`
	Ref   struct {
		KnowledgeDomain KnowledgeDomain `json:"knowledgeDomain"  bson:"knowledgeDomain,omitempty"`
		SubDomain       SubDomain       `json:"subDomain"  bson:"subDomain,omitempty"`
		State           State           `json:"state"  bson:"state,omitempty"`
		Block           Block           `json:"block"  bson:"block,omitempty"`
		District        District        `json:"district"  bson:"district,omitempty"`
		GramPanchayat   GramPanchayat   `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
		Village         Village         `json:"village"  bson:"village,omitempty"`
		AssignedTo      User            `json:"assignedTo"  bson:"assignedTo,omitempty"`
		ResolvedBy      User            `json:"resolvedBy"  bson:"resolvedBy,omitempty"`
		CreatedByFarmer Farmer          `json:"createdByFarmer"  bson:"createdByFarmer,omitempty"`
		CreatedByUser   User            `json:"createdByUser"  bson:"createdByUser,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//QueryFilter : "Used for constructing filter query"
type QueryFilter struct {
	ActiveStatus    []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	CreatedBy       []primitive.ObjectID `json:"createdOn" bson:"createdOn,omitempty"`
	Query           []primitive.ObjectID `json:"query"  bson:"query,omitempty"`
	ContentID       []primitive.ObjectID `json:"contentId"  bson:"contentId,omitempty"`
	KnowledgeDomain []primitive.ObjectID `json:"knowledgeDomain"  bson:"knowledgeDomain,omitempty"`
	SubDomain       []primitive.ObjectID `json:"subDomain"  bson:"subDomain,omitempty"`
	QueryType       []string             `json:"queryType" bson:"queryType,omitempty"`
	AssignedTo      []primitive.ObjectID `json:"assignedTo"  bson:"assignedTo,omitempty"`
	ResolvedBy      []primitive.ObjectID `json:"resolvedBy"  bson:"resolvedBy,omitempty"`
	GramPanchayat   []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village         []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	State           []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block           []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District        []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Farmer          []primitive.ObjectID `json:"farmer"  bson:"farmer,omitempty"`
	Status          []string             `json:"status" bson:"status,omitempty"`
	CreatedFrom     struct {
		StartDate *time.Time `json:"startDate"`
		EndDate   *time.Time `json:"endDate"`
	} `json:"createdFrom"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
	Regex      struct {
		Query string `json:"query" bson:"query,omitempty"`
	} `json:"regex" bson:"regex"`
}
type AssignUserToQuery struct {
	QueryId      primitive.ObjectID `json:"queryId"  bson:"queryId,omitempty"`
	UserId       primitive.ObjectID `json:"userId"  bson:"userId,omitempty"`
	AssignerId   primitive.ObjectID `json:"assignerId"  bson:"assignerId,omitempty"`
	AssignedDate *time.Time         `json:"assignedDate"  bson:"assignedDate,omitempty"`
}
type ResolveQuery struct {
	QueryId  primitive.ObjectID `json:"queryId"  bson:"queryId,omitempty"`
	UserId   primitive.ObjectID `json:"userId"  bson:"userId,omitempty"`
	Solution string             `json:"solution" bson:"solution,omitempty"`
}
