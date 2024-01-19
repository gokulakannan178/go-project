package models

import "time"

//PropertyPayment : ""
type PropertyPayment struct {
	TnxID           string                  `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID      string                  `json:"propertyId" bson:"propertyId,omitempty"`
	ReciptNo        string                  `json:"reciptNo" bson:"reciptNo,omitempty"`
	FinancialYear   FinancialYear           `json:"financialYear" bson:"financialYear,omitempty"`
	Type            string                  `json:"type" bson:"type,omitempty"`
	Details         *PropertyPaymentDetails `json:"details" bson:"details,omitempty"`
	Demand          *PropertyPaymentDemand  `json:"demand" bson:"demand,omitempty"`
	RemainingAmount float64                 `json:"remainingAmount" bson:"remainingAmount,omitempty"`
	CompletionDate  *time.Time              `json:"completionDate" bson:"completionDate,omitempty"`
	Status          string                  `json:"status" bson:"status,omitempty"`
	PendingAmount   float64                 `json:"pendingAmount" bson:"pendingAmount,omitempty"`
	Address         Address                 `json:"address" bson:"address,omitempty"`
	ReciptURL       string                  `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark          string                  `json:"remark" bson:"remark,omitempty"`
	RejectedInfo    struct {
		RejectedActionDate *time.Time `json:"rejectedActionDate" bson:"rejectedActionDate,omitempty"`
		RejectedDate       *time.Time `json:"rejectedDate" bson:"rejectedDate,omitempty"`
		Remark             string     `json:"remark" bson:"remark,omitempty"`
		By                 string     `json:"by" bson:"by,omitempty"`
		ByType             string     `json:"byType" bson:"byType,omitempty"`
	} `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	BouncedInfo struct {
		BouncedActionDate *time.Time `json:"bouncedActionDate" bson:"bouncedActionDate,omitempty"`
		BouncedDate       *time.Time `json:"bouncedDate" bson:"bouncedDate,omitempty"`
		Remark            string     `json:"remark" bson:"remark,omitempty"`
		By                string     `json:"by" bson:"by,omitempty"`
		ByType            string     `json:"byType" bson:"byType,omitempty"`
	} `json:"bouncedInfo" bson:"bouncedInfo,omitempty"`
	VerifiedInfo struct {
		VerifiedActionDate *time.Time `json:"verifiedActionDate" bson:"verifiedActionDate,omitempty"`
		VerifiedDate       *time.Time `json:"verifiedDate" bson:"verifiedDate,omitempty"`
		Remark             string     `json:"remark" bson:"remark,omitempty"`
		By                 string     `json:"by" bson:"by,omitempty"`
		ByType             string     `json:"byType" bson:"byType,omitempty"`
	} `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	NotVerifiedInfo struct {
		NotVerifiedActionDate *time.Time `json:"notVerifiedActionDate" bson:"notVerifiedActionDate,omitempty"`
		NotVerifiedDate       *time.Time `json:"notVerifiedDate" bson:"notVerifiedDate,omitempty"`
		Remark                string     `json:"remark" bson:"remark,omitempty"`
		By                    string     `json:"by" bson:"by,omitempty"`
		ByType                string     `json:"byType" bson:"byType,omitempty"`
	} `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
	Summary            Summary            `json:"summary" bson:"summary,omitempty"`
	CollectionReceived CollectionReceived `json:"collectionReceived" bson:"collectionReceived,omitempty"`
	NewPropertyID      string             `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID      string             `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

//PropertyPaymentFilter : ""
type PropertyPaymentFilter struct {
	Status                      []string       `json:"status"`
	CollectionReceivedStatus    []string       `json:"collectionReceivedStatus"`
	CollectionReceivedBy        []string       `json:"collectionReceivedBy"`
	AvoidPartPayment            bool           `json:"avoidPartPayment"`
	Type                        []string       `json:"type"`
	PropertyIds                 []string       `json:"propertyIds"`
	MadeAt                      []string       `json:"madeAt"`
	MOP                         []string       `json:"mop"`
	ReceiptNo                   []string       `json:"receiptNo"`
	Collector                   []string       `json:"collector"`
	CollectorByType             []string       `json:"collectorByType"`
	DateRange                   *DateRange     `json:"dateRange"`
	BouncedDateRange            *DateRange     `json:"bouncedDateRange"`
	CollectionReceivedDateRange *DateRange     `json:"collectionReceivedDateRange"`
	Address                     *AddressSearch `json:"address"`
	SortBy                      string         `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder                   int            `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
	SearchText                  struct {
		HoldingNo   string `json:"holdingNo" bson:"holdingNo"`
		OwnerName   string `json:"ownerName" bson:"ownerName"`
		OwnerMobile string `json:"ownerMobile" bson:"ownerMobile"`
		ReceiptNo   string `json:"receiptNo" bson:"receiptNo"`
	} `json:"searchText" bson:"searchText"`
	StartDate *time.Time `json:"startDate" bson:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate" bson:"endDate,omitempty"`
}

//RefPropertyPayment : ""
type RefPropertyPayment struct {
	PropertyPayment `bson:",inline"`
	Basic           PropertyPaymentDemandBasic `json:"basic" bson:"basic,omitempty"`
	Fys             []PropertyPaymentDemandFy  `json:"fys" bson:"fys,omitempty"`
	Ref             struct {
		PropertyOwner            []RefPropertyOwner       `json:"propertyOwner" bson:"propertyOwner,omitempty"`
		Address                  RefAddress               `json:"address" bson:"address,omitempty"`
		Collector                User                     `json:"collector" bson:"collector,omitempty"`
		PartPayments             []RefPropertyPartPayment `json:"partPayments" bson:"partPayments,omitempty"`
		PartAmountCollected      float64                  `json:"partAmountCollected" bson:"partAmountCollected,omitempty"`
		CollectionReceived       User                     `json:"collectionReceivedBy" bson:"collectionReceivedBy,omitempty"`
		RejectedBy               User                     `json:"rejectedBy" bson:"rejectedBy,omitempty"`
		CollectionReceivedByType UserType                 `json:"collectionReceivedByType" bson:"collectionReceivedByType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type ArrerAndCurrentReport struct {
	Penalty            float64 `json:"penalty" bson:"penalty,omitempty"`
	ArrearTax          float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
	CurrentTax         float64 `json:"currentTax" bson:"currentTax,omitempty"`
	ArrearPenalty      float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
	CurrentPenalty     float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
	ArrearRebate       float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
	CurrentRebate      float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
	ArrearAlreadyPaid  float64 `json:"arrearAlreadyPaid" bson:"arrearAlreadyPaid,omitempty"`
	CurrentAlreadyPaid float64 `json:"currentAlreadyPaid" bson:"currentAlreadyPaid,omitempty"`
	Formfee            float64 `json:"formfee" bson:"formfee,omitempty"`
}

func (prop *RefPropertyPayment) Inc(a int) int {
	return a + 1
}

//PropertyPaymentDetails : ""
type PropertyPaymentDetails struct {
	PayeeName       string                        `json:"payeeName" bson:"payeeName,omitempty"`
	Amount          float64                       `json:"amount" bson:"amount,omitempty"`
	AmountInWords   string                        `json:"amountInWords" bson:"amountInWords,omitempty"`
	AmountPaid      float64                       `json:"amountPaid" bson:"amountPaid,omitempty"`
	IncomingPayment float64                       `json:"incomingPayment" bson:"incomingPayment,omitempty"`
	MadeAt          *PropertyPaymentDetailsMadeAt `json:"madeAt" bson:"madeAt,omitempty"`
	MOP             struct {
		Mode   string `json:"mode" bson:"mode,omitempty"`
		Cheque *struct {
			No         string     `json:"no" bson:"no,omitempty"`
			Date       *time.Time `json:"date" bson:"date,omitempty"`
			Bank       string     `json:"bank" bson:"bank,omitempty"`
			Branch     string     `json:"branch" bson:"branch,omitempty"`
			BounceDate *time.Time `json:"bounceDate" bson:"bounceDate,omitempty"`
		} `json:"cheque" bson:"cheque,omitempty"`
		DD *struct {
			No     string     `json:"no" bson:"no,omitempty"`
			Date   *time.Time `json:"date" bson:"date,omitempty"`
			Bank   string     `json:"bank" bson:"bank,omitempty"`
			Branch string     `json:"branch" bson:"branch,omitempty"`
		} `json:"dd" bson:"dd,omitempty"`
		CardRNet                *CardRNet `json:"cardRNet" bson:"cardRNet,omitempty"`
		PropertyPaymentCardRNet CardRNet  `json:"propertyPaymentCardRNet" bson:"propertyPaymentCardRNet,omitempty"`
		VendorInfo              struct {
			Paytm struct {
				BankName     string     `json:"bankName" bson:"bankName,omitempty"`
				BankTnxID    string     `json:"bankTnxId" bson:"bankTnxId,omitempty"`
				CheckSumHash string     `json:"checkSumHash" bson:"checkSumHash,omitempty"`
				Currency     string     `json:"currency" bson:"currency,omitempty"`
				GateWayName  string     `json:"gateWayName" bson:"gateWayName,omitempty"`
				MID          string     `json:"mId" bson:"mId,omitempty"`
				OrderId      string     `json:"orderId" bson:"orderId,omitempty"`
				PaymentMode  string     `json:"paymentMode" bson:"paymentMode,omitempty"`
				RespCode     string     `json:"respCode" bson:"respCode,omitempty"`
				RespMsg      string     `json:"respMsg" bson:"respMsg,omitempty"`
				Status       string     `json:"status" bson:"status,omitempty"`
				TxnAmount    string     `json:"txnAmount" bson:"txnAmount,omitempty"`
				TxnID        string     `json:"txnId" bson:"txnId,omitempty"`
				TxnDate      *time.Time `json:"TxnDate" bson:"TxnDate,omitempty"`
			} `json:"paytm" bson:"paytm,omitempty"`
			HDFC HDFCPaymentGatewayCheckPaymentStatusResponse
		} `json:"vendorInfo" bson:"vendorInfo,omitempty"`
	} `json:"mop" bson:"mop,omitempty"`
	Collector struct {
		ID       string `json:"id" bson:"id,omitempty"`
		Type     string `json:"type" bson:"type,omitempty"`
		Platform string `json:"platform" bson:"platform,omitempty"`
	} `json:"collector" bson:"collector,omitempty"`
}

type PropertyPaymentDetailsMadeAt struct {
	At   string `json:"at" bson:"at,omitempty"`
	Bank *struct {
		Name   string     `json:"name" bson:"name,omitempty"`
		Branch string     `json:"branch" bson:"branch,omitempty"`
		TxnID  string     `json:"txnId" bson:"txnId,omitempty"`
		DOP    *time.Time `json:"dop" bson:"dop,omitempty"`
	} `json:"bank" bson:"bank,omitempty"`
	Center *struct {
		ID       string     `json:"id" bson:"id,omitempty"`
		Location string     `json:"location" bson:"location,omitempty"`
		TxnID    string     `json:"txnId" bson:"txnId,omitempty"`
		DOP      *time.Time `json:"dop" bson:"dop,omitempty"`
	} `json:"center" bson:"center,omitempty"`
}
type CardRNet struct {
	Bank       string     `json:"bank" bson:"bank,omitempty"`
	Branch     string     `json:"branch" bson:"branch,omitempty"`
	TxnID      string     `json:"txnId" bson:"txnId,omitempty"`
	DOP        *time.Time `json:"dop" bson:"dop,omitempty"`
	VendorType string     `json:"vendorType" bson:"vendorType,omitempty"`
	TrackingID string     `json:"trackingId" bson:"trackingId,omitempty"`
	BankRefNo  string     `json:"bankRefNo" bson:"bankRefNo,omitempty"`
	CardName   string     `json:"cardname" bson:"cardName,omitempty"`
	CardType   string     `json:"cardType" bson:"cardType,omitempty"`
	Vendor     string     `json:"vendor" bson:"vendor,omitempty"`
	Proof      string     `json:"proof" bson:"proof,omitempty"`
}

//PropertyMakePayment : ""
type PropertyMakePayment struct {
	TnxID     string                 `json:"tnxId" bson:"tnxId,omitempty"`
	Details   PropertyPaymentDetails `json:"details" bson:"details,omitempty"`
	ReciptURL string                 `json:"reciptURL" bson:"reciptURL,omitempty"`
}

//PropertyPaymentDemand : ""
type PropertyPaymentDemand struct {
	PercentAreaBuildup        float64               `json:"percentAreaBuildup,omitempty" bson:"percentAreaBuildup,omitempty"`
	TaxableVacantLand         float64               `json:"taxableVacantLand,omitempty" bson:"taxableVacantLand,omitempty"`
	FYs                       []FinancialYearDemand `json:"fys,omitempty" bson:"-"`
	Property                  Property              `json:"property,omitempty" bson:"-"`
	ServiceCharge             float64               `json:"serviceCharge" bson:"serviceCharge"`
	IsServiceChargeApplicable bool                  `json:"isServiceChargeApplicable" bson:"isServiceChargeApplicable"`
	PropertyConfig            PropertyConfiguration `json:"propertyConfig" bson:"propertyConfig"`
	FYTax                     float64               `json:"fyTax" bson:"fyTax"`
	Tax                       float64               `json:"tax" bson:"tax"`
	FlTax                     float64               `json:"flTax" bson:"flTax"`
	VlTax                     float64               `json:"vlTax" bson:"vlTax"`
	CompositeTax              float64               `json:"compositeTax" bson:"compositeTax"`
	EducationChess            float64               `json:"educationChess" bson:"educationChess"`
	SWUC                      float64               `json:"swuc" bson:"swuc"`
	PenalCharge               float64               `json:"penalCharge" bson:"penalCharge"`
	OtherAmount               float64               `json:"otherAmount" bson:"otherAmount"`
	FormFee                   float64               `json:"formFee" bson:"formFee"`
	TotalTax                  float64               `json:"totalTax" bson:"totalTax"`
	BoreCharge                float64               `json:"boreCharge" bson:"boreCharge"`
	Arrear                    float64               `json:"arrear" bson:"arrear"`
	Current                   float64               `json:"current" bson:"current"`
	ArrearWP                  float64               `json:"arrearWp" bson:"arrearWp"`
	CurrentWP                 float64               `json:"currentWp" bson:"currentWp"`
	PanelCh                   float64               `json:"panelCh" bson:"panelCh"`
	Rebate                    float64               `json:"rebate" bson:"rebate"`
	OtherDemand               float64               `json:"otherDemand" bson:"otherDemand"`

	PreviousCollection PreviousCollection `json:"previousCollection" bson:"previousCollection"`
}

//PropertyPaymentDemandBasic : ""
type PropertyPaymentDemandBasic struct {
	TnxID         string             `json:"tnxId" bson:"tnxId,omitempty"`
	Property      Property           `json:"property,omitempty" bson:"property,omitempty"`
	Owners        []RefPropertyOwner `json:"owners,omitempty" bson:"owners,omitempty"`
	Floors        []RefPropertyFloor `json:"floors,omitempty" bson:"floors,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	NewPropertyID string             `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string             `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

//PropertyPaymentDemandFy : ""
type PropertyPaymentDemandFy struct {
	TnxID         string              `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID    string              `json:"propertyId" bson:"propertyId,omitempty"`
	FY            FinancialYearDemand `json:"fy,omitempty" bson:"fy,omitempty"`
	Status        string              `json:"status" bson:"status,omitempty"`
	NewPropertyID string              `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string              `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

func (prop *PropertyPaymentDemandFy) Inc(a int) int {
	return a + 1
}

//DashboardTotalCollectionFilter : ""
type DashboardTotalCollectionFilter struct {
	Range *struct {
		From *time.Time `json:"from" bson:"from,omitempty"`
		To   *time.Time `json:"to" bson:"to,omitempty"`
	} `json:"range" bson:"range,omitempty"`
}

//DashboardTotalCollection : ""
type DashboardTotalCollection struct {
	Current float64 `json:"current" bson:"current,omitempty"`
	Arriear float64 `json:"arriear" bson:"arriear,omitempty"`
	Total   float64 `json:"total" bson:"total,omitempty"`
}

//DashboardTotalCollectionRef : ""
type DashboardTotalCollectionRef struct {
	PropertyPayment `bson:",inline"`
	CurrentYear     PropertyPaymentDemandFy   `json:"currentYear" bson:"currentYear,omitempty"`
	ArriearYears    []PropertyPaymentDemandFy `json:"arriearYears" bson:"arriearYears,omitempty"`
}

type PropertyPaymentCalculate struct {
	PropertyPayment `bson:",inline"`
	FYs             []PropertyPaymentDemandFy `json:"fys,omitempty" bson:"fys,omitempty"`
}

type VerifyPayment struct {
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
	TnxID      string     `json:"tnxId" bson:"tnxId,omitempty"`
	Remarks    string     `json:"remarks" bson:"remarks,omitempty"`
}

type NotVerifiedPayment struct {
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	TnxID      string     `json:"tnxId" bson:"tnxId,omitempty"`
	Remarks    string     `json:"remarks" bson:"remarks,omitempty"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

type RejectPayment struct {
	Date       *time.Time `json:"date" bson:"date,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	UniqueID   string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	TnxID      string     `json:"tnxId,omitempty" bson:"tnxId,omitempty"`
	ReceiptNo  string     `json:"receiptNo,omitempty" bson:"receiptNo,omitempty"`
	Status     string     `json:"status" bson:"status,omitempty"`
	Remarks    string     `json:"remarks" bson:"remarks,omitempty"`
}

// BouncePayment : ""
type BouncePayment struct {
	Date       *time.Time `json:"date" bson:"date,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	UniqueID   string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	TnxID      string     `json:"tnxId,omitempty" bson:"tnxId,omitempty"`
	Status     string     `json:"status" bson:"status,omitempty"`
	Remarks    string     `json:"remarks" bson:"remarks,omitempty"`
}

// ZoneAndWardWiseFilter : ""
type ZoneAndWardWiseReportFilter struct {
	// MadeAt    []string `json:"madeAt"`
	// MOP       []string `json:"mop"`
	// Collector []string `json:"collector"`
	DateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateRange"`
	Address   *AddressSearch `json:"address"`
	SortBy    string         `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int            `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

type ZoneAndWardWiseReport struct {
	Zone  `bson:",inline"`
	Wards []struct {
		Ward     `bson:",inline"`
		Payments struct {
			TotalProperties  float64 `json:"totalProperties,omitempty" bson:"totalProperties,omitempty"`
			TotalCollections float64 `json:"totalCollections,omitempty" bson:"totalCollections,omitempty"`
		}
	} `json:"wards" bson:"wards"`
}

// DateWisePropertyPaymentReportFilter : ""
type DateWisePropertyPaymentReportFilter struct {
	Date *time.Time `json:"date" bson:"date,omitempty"`
}

// RefDateWisePropertyPaymentReport : ""
type RefDateWisePropertyPaymentReport struct {
	Report struct {
		Overall struct {
			ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
			CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
			TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
			TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
			PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
		} `json:"overall" bson:"overall,omitempty"`
		Year struct {
			ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
			CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
			TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
			TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
			PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			FyName            string  `json:"fyName" bson:"fyName,omitempty"`
		} `json:"year" bson:"year,omitempty"`
		Month struct {
			ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
			CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
			TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
			TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
			PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			FyMonth           string  `json:"fyMonth" bson:"fyMonth,omitempty"`
		} `json:"month" bson:"month,omitempty"`
		Week struct {
			ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
			CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
			TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
			TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
			PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			FyWeek            string  `json:"fyWeek" bson:"fyWeek,omitempty"`
		} `json:"week" bson:"week,omitempty"`
		Day struct {
			ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
			CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
			TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
			TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
			PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			FyDay             string  `json:"fyDay" bson:"fyDay,omitempty"`
		} `json:"day" bson:"day,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}
type DateWisePropertyPaymentReport struct {
	ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
	CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
	TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
	PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
	Date              string  `json:"date" bson:"date,omitempty"`
	ArrearPenalty     float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
	CurrentPenalty    float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
	RebateAmount      float64 `json:"rebateAmount" bson:"rebateAmount,omitempty"`
	AdvanceAmount     float64 `json:"advanceAmount" bson:"advanceAmount,omitempty"`
}

type OnlinePayment struct {
}

// UserWisePropertyCollectionReportFilter : ""
type UserWisePropertyCollectionReportFilter struct {
	UserType  []string   `json:"userType,omitempty" bson:"userType,omitempty"`
	Status    []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

// UserWisePropertyCollectionReport : ""
type UserWisePropertyCollectionReport struct {
	User     `bson:",inline"`
	Payments []struct {
		ID              int     `json:"day" bson:"_id,omitempty"`
		TotalCollection float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
	} `json:"payments" bson:"payments"`
}

type CollectionReceived struct {
	TnxID  string     `json:"tnxId" bson:"tnxId,omitempty"`
	Status string     `json:"status" bson:"status,omitempty"`
	Date   *time.Time `json:"date" bson:"date,omitempty"`
	Remark string     `json:"remark" bson:"remark,omitempty"`
	By     string     `json:"by" bson:"by,omitempty"`
	ByType string     `json:"byType" bson:"byType,omitempty"`
}

type CollectionReceivedRequest struct {
	TnxIDs []string `json:"tnxIds" bson:"tnxIds"`
	By     string   `json:"by" bson:"by,omitempty"`
	ByType string   `json:"byType" bson:"byType,omitempty"`
}
