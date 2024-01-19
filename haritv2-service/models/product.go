package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Product : ""
type Product struct {
	ID primitive.ObjectID `json:"id"  bson:"_id,omitempty"`

	// ID                 bson.ObjectId `json:"id,omitempty" form:"id" bson:"_id,omitempty"`
	UniqueID           string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name               string     `json:"name" bson:"name,omitempty"`
	CategoryID         string     `json:"categoryId" bson:"categoryId,omitempty"`
	Unit               string     `json:"unit" bson:"unit,omitempty"`
	HSN                string     `json:"hsn" bson:"hsn,omitempty"`
	IsDefault          bool       `json:"isDefault" bson:"isDefault,omitempty"`
	IsTaxable          string     `json:"isTaxable" bson:"isTaxable,omitempty"`
	TaxExemptionReason string     `json:"taxExemptionReason" bson:"taxExemptionReason,omitempty"`
	IntraTaxRate       float64    `json:"intraTaxRate" bson:"intraTaxRate"`
	InterTaxRate       float64    `json:"interTaxRate" bson:"interTaxRate"`
	Created            *CreatedV2 `json:"createdOn" bson:"createdOn,omitempty"`
	Status             string     `json:"status" bson:"status,omitempty"`
}

//ProductFilter : ""
type ProductFilter struct {
	Name       []string `json:"name" bson:"name,omitempty"`
	Status     []string `json:"status" bson:"status,omitempty"`
	CategoryID []string `json:"categoryId" bson:"categoryId,omitempty"`
	// ProjectionFields []string `json:"projectionFields" bson:"projectionFields,omitempty"`
	SortField string `json:"sortField" bson:"sortField,omitempty"`
	SortOrder int    `json:"sortOrder" bson:"sortOrder,omitempty"`
}

//RefProduct : ""
type RefProduct struct {
	Product `bson:",inline"`
	Ref     struct {
		Category *ProductCategory `json:"category" bson:"category,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}
