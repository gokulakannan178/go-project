package models

//PlanRuleEngine : ""
type PlanRuleEngine struct {
	From     string `json:"from" bson:"from,omitempty"`
	To       string `json:"to" bson:"to,omitempty"`
	Scenario string `json:"scenario" bson:"scenario,omitempty"`
	Type     string `json:"type" bson:"type,omitempty"`
	Log      struct {
		Type  string `json:"type" bson:"type,omitempty"`
		Label string `json:"label" bson:"label,omitempty"`
	} `json:"log" bson:"log,omitempty"`
}
