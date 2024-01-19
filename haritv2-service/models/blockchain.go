package models

import (
	"time"
)

type Inventory struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty"`
	From struct {
		ID              string  `json:"id,omitempty" bson:"id,omitempty"`
		Name            string  `json:"name,omitempty" bson:"name,omitempty"`
		Type            string  `json:"type,omitempty" bson:"type,omitempty"`
		BeforeInventory float64 `json:"beforeInventory,omitempty" bson:"beforeInventory,omitempty"`
		AfterInventory  float64 `json:"afterInventory,omitempty" bson:"afterInventory,omitempty"`
	} `json:"from,omitempty" bson:"from,omitempty"`

	To struct {
		ID              string  `json:"id,omitempty" bson:"id,omitempty"`
		Name            string  `json:"name,omitempty" bson:"name,omitempty"`
		Type            string  `json:"type,omitempty" bson:"type,omitempty"`
		BeforeInventory float64 `json:"beforeInventory,omitempty" bson:"beforeInventory,omitempty"`
		AfterInventory  float64 `json:"afterInventory,omitempty" bson:"afterInventory,omitempty"`
	} `json:"to,omitempty" bson:"to,omitempty"`
	Quantity   float64   `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price      float64   `json:"price,omitempty" bson:"price,omitempty"`
	TimeStramp time.Time `json:"timeStramp,omitempty" bson:"timeStramp,omitempty"`
}
