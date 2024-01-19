package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FarmerLiveStock : ""
type FarmerLiveStock struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	LiveStock    primitive.ObjectID `json:"liveStock" bson:"liveStock,omitempty"`
	Farmer       primitive.ObjectID `json:"farmer,omitempty"  bson:"farmer,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	Stage        primitive.ObjectID `json:"stage,omitempty"  bson:"stage,omitempty"`
	Version      int                `json:"version,omitempty"  bson:"version,omitempty"`
	Quantity     int                `json:"quantity,omitempty"  bson:"quantity,omitempty"`
	Veriety      primitive.ObjectID `json:"veriety,omitempty"  bson:"veriety,omitempty"`
}

type FarmerLiveStockFilter struct {
	Status    []string             `json:"status,omitempty"  bson:"status,omitempty"`
	Stage     []primitive.ObjectID `json:"stage,omitempty"  bson:"stage,omitempty"`
	LiveStock []primitive.ObjectID `json:"LiveStock" bson:"LiveStock,omitempty"`
	Category  []primitive.ObjectID `json:"category" bson:"category,omitempty"`
	Farmer    []primitive.ObjectID `json:"farmer,omitempty"  bson:"farmer,omitempty"`
	Veriety   []primitive.ObjectID `json:"veriety,omitempty"  bson:"veriety,omitempty"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefFarmerLiveStock struct {
	FarmerLiveStock `bson:",inline"`
	Ref             struct {
		LiveStock Commodity         `json:"liveStock" bson:"liveStock,omitempty"`
		Farmer    Farmer            `json:"farmer,omitempty"  bson:"farmer,omitempty"`
		Stage     CommodityStage    `json:"stage,omitempty"  bson:"stage,omitempty"`
		Category  CommodityCategory `json:"category" bson:"category,omitempty"`
		Veriety   CommodityVariety  `json:"veriety,omitempty"  bson:"veriety,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type FarmerLiveStockCount struct {
	Count float64 `json:"count" bson:"count,omitempty"`
}
