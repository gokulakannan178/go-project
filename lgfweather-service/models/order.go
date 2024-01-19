package models

type Order struct {
	Sale         `bson:",inline"`
	OrderPayment OrderPayment `json:"orderPayment" bson:"orderPayment,omitempty"`
	Requested    struct {
		Is bool   `json:"is" bson:"is,omitempty"`
		By string `json:"by" bson:"by,omitempty"`
	} `json:"requested" bson:"requested,omitempty"`
	Rejected struct {
		Type string `json:"type" bson:"type,omitempty"`
		By   string `json:"by" bson:"by,omitempty"`
	} `json:"rejected" bson:"rejected,omitempty"`
}

type OrderPayment struct {
	Type  string `json:"type" bson:"type,omitempty"`
	TnxID string `json:"tnxId" bson:"tnxId,omitempty"`
}

//OrderFilter : ""
type OrderFilter struct {
	UniqueID        []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	CompanyID       []string `json:"companyId" bson:"companyId,omitempty"`
	CustomerID      []string `json:"customerId" bson:"customerId,omitempty"`
	CustomerType    []string `json:"customerType" bson:"customerType,omitempty"`
	CompanyType     []string `json:"companyType" bson:"companyType,omitempty"`
	TransportID     []string `json:"transportId" bson:"transportId,omitempty"`
	TransportStatus []string `json:"transportStatus" bson:"transportStatus,omitempty"`
	TransportType   []string `json:"transportType" bson:"transportType,omitempty"`
	PaymentStatus   []string `json:"paymentStatus" bson:"paymentStatus,omitempty"`
	DriverID        []string `json:"driverId" bson:"driverId,omitempty"`
	Status          []string `json:"status" bson:"status,omitempty"`
	RequestedIs     []bool   `json:"requestedIs" bson:"requestedIs,omitempty"`
	RequestedBy     []string `json:"requestedBy" bson:"requestedBy,omitempty"`
	SortBy          string   `json:"sortBy"`
	SortOrder       int      `json:"sortOrder"`
}
type OrderCancelFilter struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
}
type OrderNotification struct {
	CustomerMobileNo string  `json:"customerMobileNo" bson:"customerMobileNo,omitempty"`
	CustomerEmailID  string  `json:"customerEmailID" bson:"customerEmailID,omitempty"`
	CustomerAppToken string  `json:"customerAppToken" bson:"customerAppToken,omitempty"`
	CustomerName     string  `json:"customerName" bson:"customerName,omitempty"`
	CustomerFirmName string  `json:"customerFirmName" bson:"customerFirmName,omitempty"`
	CompanyAppToken  string  `json:"companyAppToken" bson:"companyAppToken,omitempty"`
	CompanyMobileNo  string  `json:"companyMobileNo" bson:"companyMobileNo,omitempty"`
	CompanyEmailID   string  `json:"companyEmailID" bson:"companyEmailID,omitempty"`
	CompanyName      string  `json:"companyName" bson:"companyName,omitempty"`
	CompanySpocName  string  `json:"companySpocName" bson:"companySpocName,omitempty"`
	Quantity         float64 `json:"quantity" bson:"quantity,omitempty"`
}

//RefOrder : ""
type RefOrder struct {
	Order `bson:",inline"`
	Ref   struct {
	} `json:"ref" bson:"ref,omitempty"`
}

type CreateOrder struct {
	Customer struct {
		ID          string `json:"id" bson:"id,omitempty"`
		Type        string `json:"type" bson:"type,omitempty"`
		Name        string `json:"name" bson:"name,omitempty"`
		Mobile      string `json:"mobile" bson:"mobile"`
		MaleCount   int    `json:"maleCount" bson:"maleCount,omitempty"`
		FemaleCount int    `json:"femaleCount" bson:"femaleCount,omitempty"`
		Gender      string `json:"gender" bson:"gender"`
		PinCode     string `json:"pinCode" bson:"pinCode"`
	} `json:"customer" bson:"customer,omitempty"`
	Company struct {
		ID   string `json:"id" bson:"id,omitempty"`
		Type string `json:"type" bson:"type,omitempty"`
	} `json:"company" bson:"company,omitempty"`
	Product struct {
		ID       string  `json:"id" bson:"id,omitempty"`
		Quantity float64 `json:"quantity" bson:"quantity,omitempty"`
	} `json:"product" bson:"product,omitempty"`
	OrderPayment OrderPayment `json:"orderPayment" bson:"orderPayment,omitempty"`
	TotalAmount  float64      `json:"totalAmount" bson:"totalAmount,omitempty"`
	SubTotal     float64      `json:"subTotal" bson:"subTotal,omitempty"`
	RoundOff     float64      `json:"roundOff" bson:"roundOff,omitempty"`
	TotalTax     float64      `json:"totalTax" bson:"totalTax,omitempty"`
}
