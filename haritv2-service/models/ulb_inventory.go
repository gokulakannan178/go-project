package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ULBInventory struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ItemID      string             `json:"itemId" bson:"itemId,omitempty"`
	CompanyID   string             `json:"companyId" bson:"companyId,omitempty"`
	ProductID   string             `json:"productId" bson:"productId,omitempty"`
	PkgType     string             `json:"pkgType" bson:"pkgType,omitempty"`
	Price       float64            `json:"price" bson:"price,omitempty"`
	BuyingPrice float64            `json:"buyingPrice" bson:"buyingPrice,omitempty"`
	CompanyType string             `json:"companyType" bson:"companyType,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
	Quantity    float64            `json:"quantity" bson:"quantity,omitempty"`
}

type RefULBInventory struct {
	ULBInventory `bson:",inline"`
}
