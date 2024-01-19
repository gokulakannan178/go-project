package models

type SendWhatsAppText2 struct {
	MobileNo string `json:"mobileNo" bson:"mobileNo,omitempty"`

	Type []WhatsAppText `json:"type" bson:"type,omitempty"`
}
type WhatsAppText struct {
	Type string `json:"type" bson:"type,omitempty"`

	Text string `json:"text" bson:"text,omitempty"`
}
type SendWhatsAppText struct {
	To       string `json:"to" bson:"to,omitempty"`
	Type     string `json:"type" bson:"type,omitempty"`
	Template struct {
		NameSpace string `json:"namespace" bson:"namespace,omitempty"`
		Name      string `json:"name" bson:"name,omitempty"`
		Language  struct {
			Policy string `json:"policy" bson:"policy,omitempty"`
			Code   string `json:"code" bson:"code,omitempty"`
		} `json:"language" bson:"language,omitempty"`
		Components []WhatsAppComponent `json:"components" bson:"components,omitempty"`
	} `json:"template" bson:"template,omitempty"`
}
type WhatsAppComponent struct {
	Type       string         `json:"type" bson:"type,omitempty"`
	Parameters []WhatsAppText `json:"parameters" bson:"parameters,omitempty"`
}
