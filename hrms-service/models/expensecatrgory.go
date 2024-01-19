package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseCategory struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Desc           string             `json:"desc" bson:"desc,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	SubCategory    []string           `json:"subCategory,omitempty" bson:"-"`
	Created        *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	Updated        Updated            `json:"updated" bson:"updated,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefExpenseCategory struct {
	ExpenseCategory `bson:",inline"`
	Ref             struct {
		Expensecategorylist ExpenseCategoryList `json:"expensecategorylist" bson:"expensecategorylist,omitempty"`
		Expensesubcategory  ExpenseSubcategory  `json:"expensesubcategory" bson:"expensesubcategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterExpenseCategory struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
