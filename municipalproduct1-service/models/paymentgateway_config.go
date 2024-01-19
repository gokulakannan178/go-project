package models

//PaymentGateway : ""
type PaymentGateway struct {
	UniqueID       string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	VendorName     string  `json:"vendorname" bson:"vendorname,omitempty"`
	Status         string  `json:"status" bson:"status,omitempty"`
	Created        Created `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated `json:"updated"  bson:"updated,omitempty"`
	MID            string  `json:"mid"  bson:"mid,omitempty"`
	MKey           string  `json:"mKey"  bson:"mKey,omitempty"`
	WebsiteName    string  `json:"websiteName"  bson:"websiteName,omitempty"`
	WebCallbackURL string  `json:"webCallbackURL"  bson:"webCallbackURL,omitempty"`
	Currency       string  `json:"currency"  bson:"currency,omitempty"`
	BaseURL        string  `json:"baseURL"  bson:"baseURL,omitempty"`
}

//PaymentGatewayFilter : ""
type PaymentGatewayFilter struct {
	Status []string `json:"status"`
	//UniqueID []string `json:"uniqueId"`
}
