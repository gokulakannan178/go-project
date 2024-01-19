package models

//Dashboard : ""
type Dashboard struct {
	Total    int64 `json:"total" bson:"total,omitempty"`
	Active   int64 `json:"active" bson:"active,omitempty"`
	Disabled int64 `json:"disabled" bson:"disabled,omitempty"`
}
