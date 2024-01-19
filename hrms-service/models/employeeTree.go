package models

type EmployeeTree struct {
	Employee *LineManagerEmployee  `json:"employee" bson:"employee,omitempty"`
	Child    []LineManagerEmployee `json:"childran" bson:"childran,omitempty"`
}
type Name struct {
	UniqueID    string `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName    string `json:"userName,omitempty" bson:"userName,omitempty"`
	Name        string `json:"name" bson:"name,omitempty"`
	LineManager string `json:"lineManager" bson:"lineManager,omitempty"`
}
type EmployeeTreev2 struct {
	Employee *LineManagerEmployee `json:"employee" bson:"employee,omitempty"`
	Child    []EmployeeTreev2     `json:"childran" bson:"childran,omitempty"`
}
type Orgchart struct {
	Employee `bson:",inline"`
	Child    []Orgchart `json:"children" bson:"children,omitempty"`
}
type Partent struct {
	Parent map[*Employee][]Employee `json:"parent" bson:"parent,omitempty"`
}
type EmployeeLinemanagerCheck struct {
	LineManager bool `json:"lineManager" bson:"lineManager,omitempty"`
}
