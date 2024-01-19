package models

import "time"

type FPO struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name      string     `json:"name" bson:"name,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	ChairMan  string     `json:"chairman" bson:"chairman,omitempty"`
	GSTNo     string     `json:"gstNo" bson:"gstNo,omitempty"`
	Email     string     `json:"email" bson:"email,omitempty"`
	Mobile    string     `json:"mobile" bson:"mobile,omitempty"`
	Created   *CreatedV2 `json:"created" bson:"created,omitempty"`
	Updated   Updated    `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog []Updated  `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	Address   Address    `json:"address" bson:"address,omitempty"`
	Logo      string     `json:"logo" bson:"logo,omitempty"`
}

type FPOFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Regex    struct {
		Name   string `json:"name" bson:"name"`
		Email  string `json:"email" bson:"email"`
		Mobile string `json:"mobile" bson:"mobile"`
	} `json:"regex" bson:"regex"`

	Address *AddressV4 `json:"address" bson:"address"`
}

type RefFPO struct {
	FPO `bson:",inline"`
	Ref struct {
		ChairMan RefUser    `json:"chairman" bson:"chairman,omitempty"`
		FPO      *FPO       `json:"fpo" bson:"fpo,omitempty"`
		Address  RefAddress `json:"address" bson:"address,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FPOReportFilter struct {
	FPOFilter `bson:",inline"`
	Date      *time.Time `json:"date,omitempty" bson:"date,omitempty"`
}

// FPOReport : ""
type FPOReport struct {
	FPO                      `bson:",inline"`
	CompostPurchasedTillDate struct {
		TotalPurchase float64 `json:"totalPurchase" bson:"totalPurchase"`
		Amount        float64 `json:"amount" bson:"amount"`
		Quantity      float64 `json:"quantity" bson:"quantity"`
		ULbs          int64   `json:"ulb" bson:"ulb"`
		FPO           int64   `json:"fpo" bson:"fpo"`
		Self          int64   `json:"self" bson:"self"`
		Customer      int64   `json:"customer" bson:"customer"`
	} `json:"compostPurchasedTillDate" bson:"compostPurchasedTillDate"`
	CompostPurchasedCurrMonth struct {
		TotalPurchase float64 `json:"totalPurchase" bson:"totalPurchase"`
		Amount        float64 `json:"amount" bson:"amount"`
		Quantity      float64 `json:"quantity" bson:"quantity"`
		ULbs          int64   `json:"ulb" bson:"ulb"`
		FPO           int64   `json:"fpo" bson:"fpo"`
		Self          int64   `json:"self" bson:"self"`
		Customer      int64   `json:"customer" bson:"customer"`
	} `json:"compostPurchasedCurrMonth" bson:"compostPurchasedCurrMonth"`
	PendingOrders struct {
		TotalPurchase float64 `json:"totalPurchase" bson:"totalPurchase"`
		Amount        float64 `json:"amount" bson:"amount"`
		Quantity      float64 `json:"quantity" bson:"quantity"`
		ULbs          int64   `json:"ulb" bson:"ulb"`
		FPO           int64   `json:"fpo" bson:"fpo"`
		Self          int64   `json:"self" bson:"self"`
		Customer      int64   `json:"customer" bson:"customer"`
	} `json:"pendingOrders" bson:"pendingOrders"`

	Ref struct {
		Address *RefAddress `json:"address" bson:"address"`
	} `json:"ref" bson:"ref,omitempty"`
}
type FPOMothWiseeport struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string `json:"name" bson:"name,omitempty"`
	ChairMan string `json:"chairman" bson:"chairman,omitempty"`
	Mobile   string `json:"mobile" bson:"mobile,omitempty"`
	Sale     struct {
		TotalCustomers  int `json:"totalCustomers" bson:"totalCustomers,omitempty"`
		TotalsaleAmount int `json:"totalsaleAmount" bson:"totalsaleAmount,omitempty"`
	} `json:"sales" bson:"sales,omitempty"`
}
type FPOMothWiseeportFilter struct {
	Month time.Month `json:"month" bson:"month,omitempty"`
	Year  int        `json:"year" bson:"year,omitempty"`
}
