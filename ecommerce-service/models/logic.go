package models

type VarientInputLogic struct {
	VarientID   string `json:"varientId" bson:"varientId,omitempty"`
	VarientName string `json:"varientName" bson:"varientName,omitempty"`
	Varients    string `json:"varients" bson:"varients,omitempty"`
}

type VarientOutputLogic struct {
	VarientID   string   `json:"varientId" bson:"varientId,omitempty"`
	VarientName string   `json:"varientName" bson:"varientName,omitempty"`
	Varients    []string `json:"varients" bson:"varients,omitempty"`
}
