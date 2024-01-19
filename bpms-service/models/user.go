package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//User : ""
type User struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UserName       string             `json:"userName" bson:"userName,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Mobile         string             `json:"mobile" bson:"mobile,omitempty"`
	Email          string             `json:"email" bson:"email,omitempty"`
	Password       string             `json:"-" bson:"password,omitempty"`
	Pass           string             `json:"password" bson:"-"`
	Type           string             `json:"type" bson:"type,omitempty"`
	Created        Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog      []Updated          `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	OrganisationID string             `json:"organisationId" bson:"organisationId,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
}

// RefUser : ""
type RefUser struct {
	User `bson:",inline"`
	Ref  struct {
		Applicant  *RefApplicant  `json:"applicant" bson:"applicant,omitempty"`
		ULB        *RefULB        `json:"ulb" bson:"ulb,omitempty"`
		Department *RefDepartment `json:"department" department:"userName,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserFilter : ""
type UserFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	UserType     []string `json:"userType,omitempty" bson:"userType,omitempty"`
	Organisation []string `json:"organisation,omitempty" bson:"organisation,omitempty"`
	SortBy       string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder    int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
