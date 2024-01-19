package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Product : ""
type Product struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus  bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Name          string             `json:"name"  bson:"name,omitempty"`
	CategoryId    primitive.ObjectID `json:"categoryId" " bson:"categoryId,omitempty"`
	SubCategoryID primitive.ObjectID `json:"subCategoryId" " bson:"subCategoryId,omitempty"`
	DealerId      primitive.ObjectID `json:"dealerId" " bson:"dealerId,omitempty"`
	Description   string             `json:"description" bson:"description,omitempty"`
	Version       int                `json:"version"  bson:"version,omitempty"`
	Status        string             `json:"status"  bson:"status,omitempty"`
	BuyingPrice   float64            `json:"buyingPrice"  bson:"buyingPrice,omitempty"`
	SellingPrice  float64            `json:"sellingPrice"  bson:"sellingPrice,omitempty"`
	Make          string             `json:"make"  bson:"make,omitempty"`
	Model         string             `json:"model"  bson:"model,omitempty"`
	Code          string             `json:"code"  bson:"code,omitempty"`
	Quantity      float64            `json:"quantity"  bson:"quantity,omitempty"`
	Created       *Created           `json:"created"  bson:"created,omitempty"`
}

type ProductFilter struct {
	OmitIDs        []primitive.ObjectID `json:"omitIds,omitempty"`
	Status         []string             `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus   []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Classification []string             `json:"classification" bson:"classification,omitempty"`
	SortBy         string               `json:"sortBy"`
	SortOrder      int                  `json:"sortOrder"`
	SearchBox      struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefProduct struct {
	Product `bson:",inline"`
	Ref     struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
