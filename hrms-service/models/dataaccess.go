package models

type DataAccessRequest struct {
	IsAccess bool   `json:"isAccess,omitempty" bson:"isAccess,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
}
type DataAccess struct {
	Organisation []RefOrganisation `json:"organisation,omitempty" bson:"organisation,omitempty"`
	User         *RefUser          `json:"user,omitempty" bson:"user,omitempty"`
	SuperAdmin   bool              `json:"superAdmin,omitempty" bson:"superAdmin,omitempty"`
	Branch       string            `json:"branch,omitempty" bson:"branch,omitempty"`
	Department   string            `json:"department,omitempty" bson:"department,omitempty"`
	Designation  string            `json:"designation,omitempty" bson:"designation,omitempty"`

	// AccessDistricts      []District        `json:"accessDistricts" bson:"accessDistricts,omitempty"`
	// AccessStates         []State           `json:"accessStates" bson:"accessStates,omitempty"`
	// AccessVillages       []Village         `json:"accessVillages" bson:"accessVillages,omitempty"`
	// AccessBlocks         []Block           `json:"accessBlocks" bson:"accessBlocks,omitempty"`
	// AccessGrampanchayats []GramPanchayat   `json:"accessGrampanchayats" bson:"accessGrampanchayats,omitempty"`
	// Projects             []RefProjectUser  `json:"projects" bson:"projects,omitempty"`
}
