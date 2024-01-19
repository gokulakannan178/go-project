package models

import "time"

//EmployeeReport : ""
type EmployeeReport struct {
	UniqueID       string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string     `json:"name,omitempty" bson:"name,omitempty"`
	DOB            *time.Time `json:"dob" bson:"dob,omitempty"`
	Gender         string     `json:"gender" bson:"gender,omitempty"`
	Mobile         string     `json:"mobile" bson:"mobile,omitempty"`
	Email          string     `json:"email" bson:"email,omitempty"`
	Designation    string     `json:"designation,omitempty" bson:"designation,omitempty"`
	OrganisationID string     `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeStatus string     `json:"employeeStatus,omitempty" bson:"employeeStatus,omitempty"`
	ManagerID      string     `json:"managerId" bson:"managerId,omitempty"`
	Created        Created    `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string     `json:"status,omitempty" bson:"status,omitempty"`
}
