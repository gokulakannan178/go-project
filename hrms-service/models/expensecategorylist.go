package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExpenseCategoryList struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID       string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	SubcategoryId  string             `json:"subcategoryId,omitempty" bson:"subcategoryId,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	CategoryId     string             `json:"categoryId,omitempty" bson:"categoryId,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefExpenseCategoryList struct {
	ExpenseCategoryList `bson:",inline"`
	Ref                 struct {
		ExpenseSubCategory ExpenseSubcategory `json:"expenseSubcategory" bson:"expenseSubcategory,omitempty"`
		OrganisationID     Organisation       `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		ExpenseCategory    ExpenseCategory    `json:"expensecategory" bson:"expensecategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterExpenseCategoryList struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
