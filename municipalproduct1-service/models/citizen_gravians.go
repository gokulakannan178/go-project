package models

import "time"

// CitizenGrievance : ""
type CitizenGrievance struct {
	UniqueID      string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	MobileNo      string     `json:"mobileNo" bson:"mobileNo,omitempty"`
	Requestor     Action     `json:"requestor" bson:"requestor,omitempty"`
	Activator     Action     `json:"activator" bson:"activator,omitempty"`
	Completed     Action     `json:"completed" bson:"completed,omitempty"`
	Name          string     `json:"name" bson:"name,omitempty"`
	Type          string     `json:"type" bson:"type,omitempty"`
	Created       *CreatedV2 `json:"created" bson:"created,omitempty"`
	Problem       string     `json:"problem" bson:"problem,omitempty"`
	Solution      string     `json:"solution" bson:"solution,omitempty"`
	SolutionImage []string   `json:"solutionImage" bson:"solutionImage,omitempty"`
	SolutionBy    string     `json:"solutionBy" bson:"solutionBy,omitempty"`
	SolutionDate  *time.Time `json:"solutionDate" bson:"solutionDate,omitempty"`
	Description   string     `json:"description" bson:"description,omitempty"`
	Image         []string   `json:"image" bson:"image,omitempty"`
	Status        string     `json:"status" bson:"status,omitempty"`
	By            string     `json:"by" bson:"by,omitempty"`
	ByID          string     `json:"byId" bson:"byId,omitempty"`
	ByType        string     `json:"byType" bson:"byType,omitempty"`
	RaisedFrom    string     `json:"raisedFrom" bson:"raisedFrom,omitempty"`
}

// CitizenGrievanceFilter : ""
type CitizenGrievanceFilter struct {
	Status        []string  `json:"status" bson:"status,omitempty"`
	Requestor     []string  `json:"requestor" bson:"requestor,omitempty"`
	Activator     []string  `json:"activator" bson:"activator,omitempty"`
	Completed     []string  `json:"completed" bson:"completed,omitempty"`
	MobileNo      []string  `json:"mobileNo" bson:"mobileNo,omitempty"`
	RequestedDate DateRange `json:"requestedDate" bson:"requestedDate,omitempty"`
	ActivatedDate DateRange `json:"activatedDate" bson:"activatedDate,omitempty"`
	CompletedDate DateRange `json:"completedDate" bson:"completedDate,omitempty"`
	SortBy        string    `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder     int       `json:"sortOrder" bson:"sortOrder,omitempty"`
}

type RefCitizenGrievance struct {
	CitizenGrievance `bson:",inline"`
	Ref              struct {
	} `json:"ref" bson:"ref,omitempty"`
}
type RejectedCitizengravians struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Remarks  string `json:"remarks" bson:"remarks,omitempty"`
	By       string `json:"by" bson:"by,omitempty"`
	ByID     string `json:"byId" bson:"byId,omitempty"`
	ByType   string `json:"byType" bson:"byType,omitempty"`
}

// CitizenGrievanceSolution : ""
type CitizenGrievanceSolution struct {
	CitizenGrievanceID string     `json:"citizenGrievanceId" bson:"citizenGrievanceId,omitempty"`
	Solution           string     `json:"solution" bson:"solution,omitempty"`
	SolutionImage      []string   `json:"solutionImage" bson:"solutionImage,omitempty"`
	Status             string     `json:"status" bson:"status,omitempty"`
	SolutionDate       *time.Time `json:"solutionDate" bson:"solutionDate,omitempty"`
	By                 string     `json:"by" bson:"by,omitempty"`
	ByID               string     `json:"byId" bson:"byId,omitempty"`
	ByType             string     `json:"byType" bson:"byType,omitempty"`
}
