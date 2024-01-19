package models

type Order struct {
	Sale          `bson:",inline"`
	InitateRemark []string `json:"initateRemark" bson:"-"`
}

//SaleFilter : ""
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
	SortBy          string   `json:"sortBy"`
	SortOrder       int      `json:"sortOrder"`
}

//RefOrder : ""
type RefOrder struct {
	Order `bson:",inline"`
	// Ref     struct {
	// } `json:"ref" bson:"ref,omitempty"`
}

type CreateOrder struct {
	Customer struct {
		ID      string  `json:"id" bson:"id,omitempty"`
		Type    string  `json:"type" bson:"type,omitempty"`
		Email   string  `json:"email" bson:"email,omitempty"`
		Name    string  `json:"name" bson:"name,omitempty"`
		Address Address `json:"address" bson:"address,omitempty"`
	} `json:"customer" bson:"customer,omitempty"`
	Company struct {
		ID   string `json:"id" bson:"id,omitempty"`
		Type string `json:"type" bson:"type,omitempty"`
	} `json:"company" bson:"company,omitempty"`
	Products    []CreateOrderProduct `json:"products" bson:"products,omitempty"`
	TotalAmount float64              `json:"totalAmount" bson:"totalAmount,omitempty"`
	SubTotal    float64              `json:"subTotal" bson:"subTotal,omitempty"`
	RoundOff    float64              `json:"roundOff" bson:"roundOff,omitempty"`
	TotalTax    float64              `json:"totalTax" bson:"totalTax,omitempty"`
}
type CreateOrderProduct struct {
	InventoryID string  `json:"inventoryid" bson:"inventoryid,omitempty"`
	Quantity    float64 `json:"quantity" bson:"quantity,omitempty"`
	Amount      float64 `json:"amount" bson:"amount,omitempty"`
	Price       float64 `json:"price" bson:"price,omitempty"`
}
