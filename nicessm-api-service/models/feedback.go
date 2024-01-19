package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FeedBack : "Holds single FeedBack data"
type FeedBack struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Status          string             `json:"status" bson:"status,omitempty"`
	Feedback        string             `json:"feedback" bson:"feedback,omitempty"`
	FeedbackType    string             `json:"feedbackType" bson:"feedbackType,omitempty"`
	ActiveStatus    bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created         Created            `json:"createdOn" bson:"createdOn,omitempty"`
	CreatedBy       primitive.ObjectID `json:"createdBy" bson:"createdBy,omitempty"`
	CreatedType     string             `json:"createdType" bson:"createdType,omitempty"`
	Date            *time.Time         `json:"date" bson:"date,omitempty"`
	Query           primitive.ObjectID `json:"query"  bson:"query,omitempty"`
	Content         primitive.ObjectID `json:"content"  bson:"content,omitempty"`
	KnowledgeDomain primitive.ObjectID `json:"knowledgeDomain"  bson:"knowledgeDomain,omitempty"`
	Farmer          primitive.ObjectID `json:"farmer"  bson:"farmer,omitempty"`
	SubDomain       primitive.ObjectID `json:"subDomain"  bson:"subDomain,omitempty"`
	GramPanchayat   primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village         primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	//Version         string             `json:"version"  bson:"version,omitempty"`
	Completeness   float64            `json:"completeness"  bson:"completeness,omitempty"`
	Relevance      float64            `json:"relevance"  bson:"relevance,omitempty"`
	Timeliness     float64            `json:"timeliness"  bson:"timeliness,omitempty"`
	Understandable float64            `json:"understandable"  bson:"understandable,omitempty"`
	Rating         float64            `json:"rating"  bson:"rating,omitempty"`
	State          primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block          primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District       primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Type           string             `json:"type"  bson:"type,omitempty"`
}

//RefFeedBack : "FeedBack with refrence data such as language..."
type RefFeedBack struct {
	FeedBack `bson:",inline"`
	Ref      struct {
		KnowledgeDomain KnowledgeDomain `json:"knowledgeDomain"  bson:"knowledgeDomain,omitempty"`
		SubDomain       SubDomain       `json:"subDomain"  bson:"subDomain,omitempty"`
		Query           Query           `json:"query"  bson:"query,omitempty"`
		Content         Content         `json:"content"  bson:"content,omitempty"`
		State           State           `json:"state"  bson:"state,omitempty"`
		Block           Block           `json:"block"  bson:"block,omitempty"`
		District        District        `json:"district"  bson:"district,omitempty"`
		GramPanchayat   GramPanchayat   `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
		Village         Village         `json:"village"  bson:"village,omitempty"`
		Farmer          Farmer          `json:"farmer"  bson:"farmer,omitempty"`
		CreatedBy       User            `json:"createdBy" bson:"createdBy,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//FeedBackFilter : "Used for constructing filter query"
type FeedBackFilter struct {
	ActiveStatus        []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	IsFeedbackRating    bool                 `json:"isFeedbackRating" bson:"isFeedbackRating,omitempty"`
	CreatedBy           []primitive.ObjectID `json:"createdOn" bson:"createdOn,omitempty"`
	Query               []primitive.ObjectID `json:"query"  bson:"query,omitempty"`
	KnowledgeDomain     []primitive.ObjectID `json:"knowledgeDomain"  bson:"knowledgeDomain,omitempty"`
	SubDomain           []primitive.ObjectID `json:"subDomain"  bson:"subDomain,omitempty"`
	FeedbackType        []string             `json:"feedbackType" bson:"feedbackType,omitempty"`
	Content             []primitive.ObjectID `json:"content"  bson:"content,omitempty"`
	GramPanchayat       []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village             []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	State               []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block               []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District            []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	ContentOrganisation []primitive.ObjectID `json:"contentOrganisation" bson:"contentOrganisation,omitempty"`
	ContentProject      []primitive.ObjectID `json:"contentProject" bson:"contentProject,omitempty"`
	Type                []string             `json:"type"  bson:"type,omitempty"`
	Status              []string             `json:"status" bson:"status,omitempty"`
	SortBy              string               `json:"sortBy"`
	SortOrder           int                  `json:"sortOrder"`
	DateRange           *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateRange"`
	Regex struct {
		Type     string `json:"type"  bson:"type"`
		Feedback string `json:"feedback" bson:"feedback,omitempty"`
	} `json:"regex" bson:"regex"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}

type FeedBackRating struct {
	A1         int     `json:"rating1"  bson:"1,omitempty"`
	A2         int     `json:"rating2"  bson:"2,omitempty"`
	A3         int     `json:"rating3"  bson:"3,omitempty"`
	A4         int     `json:"rating4"  bson:"4,omitempty"`
	A5         int     `json:"rating5"  bson:"5,omitempty"`
	A6         int     `json:"rating6"  bson:"6,omitempty"`
	A7         int     `json:"rating7"  bson:"7,omitempty"`
	A8         int     `json:"rating8"  bson:"8,omitempty"`
	A9         int     `json:"rating9"  bson:"9,omitempty"`
	A10        int     `json:"rating10"  bson:"10,omitempty"`
	Average    float64 `json:"average"  bson:"average,omitempty"`
	AverageStr string  `json:"averageStr"  bson:"averageStr,omitempty"`
}
