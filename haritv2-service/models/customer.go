package models

//Prefixs : ""
const (
	PREFIXCUSTOMER     string = "CUSTOMER"
	COLLECTIONCUSTOMER string = "customer"
)

//Customer : ""
type Customer struct {
	UniqueID       string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name           string    `json:"name" bson:"name,omitempty"`
	Profile        string    `json:"profile" bson:"profile,omitempty"`
	Email          string    `json:"email" bson:"email"`
	Mobile         string    `json:"mobile" bson:"mobile"`
	Gender         string    `json:"gender" bson:"gender"`
	PinCode        string    `json:"pinCode" bson:"pinCode"`
	Created        CreatedV2 `json:"createdOn" bson:"createdOn,omitempty"`
	Address        Address   `json:"address" bson:"address,omitempty"`
	Status         string    `json:"status" bson:"status,omitempty"`
	PrimaryContact Contact   `json:"primaryContact" bson:"primaryContact,omitempty"`
}

//CustomerFilter : ""
type CustomerFilter struct {
	Name     string   `json:"name" bson:"name,omitempty"`
	UniqueID []string `json:"uniqueID" bson:"uniqueID,omitempty"`
	Status   []string `json:"status" bson:"status,omitempty"`
	Regex    struct {
		Name   string `json:"name" bson:"name"`
		Email  string `json:"email" bson:"email"`
		Mobile string `json:"mobile" bson:"mobile"`
	} `json:"regex" bson:"regex"`

	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

//RefCustomer : ""
type RefCustomer struct {
	Customer `bson:",inline"`
	Ref      struct {
		Company interface{} `json:"company" bson:"company,omitempty"`
		Address *RefAddress
	} `json:"ref" bson:"ref,omitempty"`
}

//Tax : ""
type Tax struct {
	GSTType string `json:"gsttType" bson:"gsttType,omitempty"`
	GSTIN   string `json:"gstIn" bson:"gstIn,omitempty"`
	UIN     string `json:"uin" bson:"uin,omitempty"`
	//Source Of Supply
	SOS      string  `json:"sos" bson:"sos,omitempty"`
	Currency string  `json:"currency" bson:"currency,omitempty"`
	TDS      float64 `json:"tds" bson:"tds,omitempty"`
}
type Bank struct {
	Name      string `json:"name" bson:"name,omitempty"`
	AccountNo string `json:"accountNo" bson:"accountNo,omitempty"`
	IFSC      string `json:"ifsc" bson:"ifsc,omitempty"`
}
type RegistrationCustomer struct {
	Customer `bson:",inline"`
	OTP      string `json:"otp"`
}
