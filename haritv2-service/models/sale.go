package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	ID primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`

	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	CompanyID string     `json:"companyId" bson:"companyId,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	Company   struct {
		ID      string `json:"id" bson:"id,omitempty"`
		Type    string `json:"type" bson:"type,omitempty"`
		Name    string `json:"name" bson:"name,omitempty"`
		Logo    string `json:"logo" bson:"logo,omitempty"`
		Address string `json:"address" bson:"address,omitempty"`
		User    string `json:"user" bson:"user,omitempty"`
		Contact string `json:"contact" bson:"contact,omitempty"`
		Email   string `json:"email" bson:"email,omitempty"`
		State   struct {
			ID    string `json:"id" bson:"id,omitempty"`
			Label string `json:"label" bson:"label,omitempty"`
		} `json:"state" bson:"state,omitempty"`
	} `json:"company" bson:"company,omitempty"`
	Address struct {
		Billing  string `json:"billing" bson:"billing,omitempty"`
		Shipping string `json:"shipping" bson:"shipping,omitempty"`
	} `json:"address" bson:"address,omitempty"`
	Customer struct {
		ID          string  `json:"id" bson:"id,omitempty"`
		Type        string  `json:"type" bson:"type,omitempty"`
		Name        string  `json:"name" bson:"name,omitempty"`
		IsSelf      bool    `json:"isSelf" bson:"isSelf,omitempty"`
		Address     Address `json:"address" bson:"address,omitempty"`
		User        string  `json:"user" bson:"user,omitempty"`
		Contact     string  `json:"contact" bson:"contact,omitempty"`
		Gender      string  `json:"gender" bson:"gender,omitempty"`
		MaleCount   int     `json:"maleCount" bson:"maleCount,omitempty"`
		FemaleCount int     `json:"femaleCount" bson:"femaleCount,omitempty"`
		Email       string  `json:"email" bson:"email,omitempty"`
		Signature   string  `json:"signature" bson:"signature,omitempty"`
		Photo       string  `json:"photo" bson:"photo,omitempty"`
		State       struct {
			ID    string `json:"id" bson:"id,omitempty"`
			Label string `json:"label" bson:"label,omitempty"`
		} `json:"state" bson:"state,omitempty"`
	} `json:"customer" bson:"customer,omitempty"`
	Items            []SaleItems   `json:"items" bson:"items,omitempty"`
	Status           string        `json:"status" bson:"status,omitempty"`
	PaymentStatus    string        `json:"paymentStatus" bson:"paymentStatus,omitempty"`
	Transport        SaleTransport `json:"transport" bson:"transport,omitempty"`
	TotalAmount      float64       `json:"totalAmount" bson:"totalAmount,omitempty"`
	SubTotal         float64       `json:"subTotal" bson:"subTotal,omitempty"`
	RoundOff         float64       `json:"roundOff" bson:"roundOff,omitempty"`
	TotalTax         float64       `json:"totalTax" bson:"totalTax,omitempty"`
	Created          CreatedV2     `json:"createdOn" bson:"createdOn,omitempty"`
	Comments         []SaleComment `json:"comments" form:"comments," bson:"comments,omitempty"`
	DueDate          *time.Time    `json:"dueDate" bson:"dueDate,omitempty"`
	CustSaleCategory string        `json:"custSaleCategory" bson:"-"`
	IsSaleOnCredit   bool          `json:"isSaleOnCredit" bson:"-"`
	Payments         []Payment     `json:"payments" bson:"-"`
	Payment          Payment       `json:"payment" bson:"-"`
	ItemStatus       string        `json:"itemStatus" bson:"itemStatus,omitempty"`
}

//SaleTransport : ""
type SaleTransport struct {
	Name            string     `json:"name" bson:"name,omitempty"`
	Address         string     `json:"address" bson:"address,omitempty"`
	SPOCName        string     `json:"spocName" bson:"spocName,omitempty"`
	SPOCNumber      string     `json:"spocNumber" bson:"spocNumber,omitempty"`
	DriverName      string     `json:"driverName" bson:"driverName,omitempty"`
	DriverNumber    string     `json:"driverNumber" bson:"driverNumber,omitempty"`
	Charges         float64    `json:"charges" bson:"charges,omitempty"`
	IsPaymentDone   string     `json:"isPaymentDone" bson:"isPaymentDone,omitempty"`
	ReasonforReject string     `json:"reasonforreject" bson:"reasonforreject,omitempty"`
	CompanyID       string     `json:"companyId" bson:"companyId,omitempty"`
	VehicleNo       string     `json:"vehicleNo" bson:"vehicleNo,omitempty"`
	VehicleName     string     `json:"vehicleName" bson:"vehicleName,omitempty"`
	DriverID        string     `json:"driverId" bson:"driverId,omitempty"`
	ExepTransDate   *time.Time `json:"exepTransDate" bson:"exepTransDate,omitempty"`
	Status          string     `json:"status" bson:"status,omitempty"`
	DeliverRemarks  string     `json:"deliverRemarks" bson:"deliverRemarks,omitempty"`
	Type            string     `json:"type" bson:"type,omitempty"`
}

//SaleFilter : ""
type SaleFilter struct {
	UniqueID        []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	CompanyID       []string `json:"companyId" bson:"companyId,omitempty"`
	CustomerID      []string `json:"customerId" bson:"customerId,omitempty"`
	CustomerType    []string `json:"customerType" bson:"customerType,omitempty"`
	CompanyType     []string `json:"companyType" bson:"companyType,omitempty"`
	TransportID     []string `json:"transportId" bson:"transportId,omitempty"`
	TransportStatus []string `json:"transportStatus" bson:"transportStatus,omitempty"`
	PaymentStatus   []string `json:"paymentStatus" bson:"paymentStatus,omitempty"`
	DriverID        []string `json:"driverId" bson:"driverId,omitempty"`
	Status          []string `json:"status" bson:"status,omitempty"`
	SortBy          string   `json:"sortBy"`
	SortOrder       int      `json:"sortOrder"`
	DateRange       *struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
}

//RefSale : ""
type RefSale struct {
	Sale `bson:",inline"`
	Ref  struct {
		PaymentsDone *PaymentsDone
	} `json:"ref" bson:"ref,omitempty"`
}

//SaleItems : ""
type SaleItems struct {
	SerialNos []string `json:"serialNos" bson:"serialNos,omitempty"`
	Quantity  float64  `json:"quantity" bson:"quantity,omitempty"`
	Product   struct {
		ID                 string  `json:"id" bson:"id,omitempty"`
		Name               string  `json:"name" bson:"name,omitempty"`
		CategoryID         string  `json:"categoryId" bson:"categoryId,omitempty"`
		Unit               string  `json:"unit" bson:"unit,omitempty"`
		HSN                string  `json:"hsn" bson:"hsn,omitempty"`
		IsTaxable          string  `json:"isTaxable" bson:"isTaxable,omitempty"`
		TaxExemptionReason string  `json:"taxExemptionReason" bson:"taxExemptionReason,omitempty"`
		IntraTaxRate       float64 `json:"intraTaxRate" bson:"intraTaxRate,omitempty"`
		InterTaxate        float64 `json:"interTaxRate" bson:"interTaxRate,omitempty"`
	} `json:"product" bson:"product,omitempty"`
	PkgType struct {
		ID    string `json:"id" bson:"id,omitempty"`
		Label string `json:"label" bson:"label,omitempty"`
	} `json:"pkgType" bson:"pkgType,omitempty"`
	BuyingUnitPrice  float64 `json:"buyingUnitPrice" bson:"buyingUnitPrice,omitempty"`
	SellingUnitPrice float64 `json:"sellingUnitPrice" bson:"sellingUnitPrice,omitempty"`
	Price            float64 `json:"price" bson:"price,omitempty"`
	Subcidy          Subcidy `json:"subcidy" bson:"subcidy,omitempty"`
	Tax              struct {
		GST struct {
			ID         string  `json:"id" bson:"id,omitempty"`
			Percentage float64 `json:"percentage" bson:"percentage,omitempty"`
			Label      string  `json:"label" bson:"label,omitempty"`
			CGST       float64 `json:"cgst" bson:"cgst,omitempty"`
			SGST       float64 `json:"sgst" bson:"sgst,omitempty"`
			IGST       float64 `json:"igst" bson:"igst,omitempty"`
			Total      float64 `json:"total" bson:"total,omitempty"`
		} `json:"gst" bson:"gst,omitempty"`
		Total float64 `json:"total" bson:"total,omitempty"`
	} `json:"tax" bson:"tax,omitempty"`
	Amount      float64 `json:"amount" bson:"amount,omitempty"`
	TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
}

type Subcidy struct {
	Type   string  `json:"type,omitempty" bson:"type,omitempty"`
	Unit   string  `json:"unit,omitempty" bson:"unit,omitempty"`
	Value  float64 `json:"value,omitempty" bson:"value,omitempty"`
	Amount float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	Desc   string  `json:"desc,omitempty" bson:"desc,omitempty"`
}

//SaleComment : ""
type SaleComment struct {
	Type string     `json:"type" bson:"type,omitempty"`
	Msg  string     `json:"msg" bson:"msg,omitempty"`
	On   *time.Time `json:"on" bson:"on,omitempty"`
	By   string     `json:"by" bson:"by,omitempty"`
}
