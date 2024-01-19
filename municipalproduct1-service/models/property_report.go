package models

import "time"

// PropertyArrerAndCurrentReportFilter : ""
type PropertyArrearAndCurrentCollectionFilter struct {
	Status    []string `json:"status"`
	Ward      []string `json:"ward"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
}

// PropertyArrearAndCurrentCollectionReport : ""
type PropertyArrearAndCurrentCollectionReport struct {
	Ward     `bson:",inline"`
	Payments struct {
		TotalProperties    float64 `json:"totalProperties" bson:"totalProperties,omitempty"`
		FormFee            float64 `json:"formFee" bson:"formFee,omitempty"`
		ArrearTax          float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
		CurrentTax         float64 `json:"currentTax" bson:"currentTax,omitempty"`
		ArrearPenalty      float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
		CurrentPenalty     float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
		ArrearRebate       float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
		CurrentRebate      float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
		ArrearAlreadyPaid  float64 `json:"arrearAlreadyPaid" bson:"arrearAlreadyPaid,omitempty"`
		CurrentAlreadyPaid float64 `json:"currentAlreadyPaid" bson:"currentAlreadyPaid,omitempty"`
		OtherDemand        float64 `json:"otherDemand" bson:"otherDemand,omitempty"`
		Penalty            float64 `json:"penalty" bson:"penalty,omitempty"`

		TotalTax float64 `json:"totalTax" bson:"totalTax,omitempty"`
	} `json:"payments" bson:"payments,omitempty"`
}

// CounterReportV2Filter : ""
type CounterReportV2Filter struct {
	PropertyPaymentFilter `bson:",inline"`
}

type RefCounterReport struct {
	Property `bson:",inline"`
	Payments []struct {
		PropertyPayment `bson:",inline"`
		FYs             []struct {
			NoOfFys []struct {
				PropertyPaymentDemandFy `bson:",inline"`
			} `json:"noOfFys" bson:"noOfFys,omitempty"`
			TnxID          string  `json:"tnxId" bson:"tnxId,omitempty"`
			ArrearTax      float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
			CurrentTax     float64 `json:"currentTax" bson:"currentTax,omitempty"`
			ArrearPenalty  float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
			CurrentPenalty float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
			ArrearRebate   float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
			CurrentRebate  float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
			OtherDemand    float64 `json:"otherDemand" bson:"otherDemand,omitempty"`
			TotalTax       float64 `json:"totalTax" bson:"totalTax,omitempty"`
		} `json:"fys" bson:"fys,omitempty"`
		Ref struct {
			Collector *RefUser `json:"collector" bson:"collector,omitempty"`
		} `json:"ref" bson:"ref,omitempty"`
	} `json:"payments" bson:"payments,omitempty"`
	Ref struct {
		Ward      *RefWard          `json:"ward" bson:"ward,omitempty"`
		Owner     *RefPropertyOwner `json:"owner" bson:"owner,omitempty"`
		Activator *RefUser          `json:"activator" bson:"activator,omitempty"`
		Creator   *RefUser          `json:"creator" bson:"creator,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

// RefCounterReportV2 : ""
type RefCounterReportV2 struct {
	PropertyPayment `bson:",inline"`
	Basic           PropertyPaymentDemandBasic `json:"basic" bson:"basic,omitempty"`
	PaymentFYs      struct {
		ArrearTax          float64                   `json:"arrearTax" bson:"arrearTax,omitempty"`
		CurrentTax         float64                   `json:"currentTax" bson:"currentTax,omitempty"`
		ArrearPenalty      float64                   `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
		CurrentPenalty     float64                   `json:"currentPenalty" bson:"currentPenalty,omitempty"`
		ArrearRebate       float64                   `json:"arrearRebate" bson:"arrearRebate,omitempty"`
		CurrentRebate      float64                   `json:"currentRebate" bson:"currentRebate,omitempty"`
		ArrearAlreadyPaid  float64                   `json:"arrearAlreadyPaid" bson:"arrearAlreadyPaid,omitempty"`
		CurrentAlreadyPaid float64                   `json:"currentAlreadyPaid" bson:"currentAlreadyPaid,omitempty"`
		OtherDemand        float64                   `json:"otherDemand" bson:"otherDemand,omitempty"`
		TotalTax           float64                   `json:"totalTax" bson:"totalTax,omitempty"`
		FYs                []PropertyPaymentDemandFy `json:"fys" bson:"fys,omitempty"`
	} `json:"paymentFys,omitempty" bson:"paymentFys,omitempty"`
	Ref struct {
		Owner           PropertyOwner `json:"owner" bson:"owner,omitempty"`
		Ward            Ward          `json:"ward" bson:"ward,omitempty"`
		Collector       User          `json:"collector" bson:"collector,omitempty"`
		Creator         User          `json:"creator" bson:"creator,omitempty"`
		Activator       User          `json:"activator" bson:"activator,omitempty"`
		PropertyDetails Property      `json:"propertyDetails" bson:"propertyDetails,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

// FilterUserWisePropertyCollectionReport : ""
type UserWisePropertyCollectionFilter struct {
	UserFilter `bson:",inline"`
	DateFrom   *time.Time `json:"dateFrom" bson:"dateFrom,omitempty"`
	DateTo     *time.Time `json:"dateTo" bson:"dateTo,omitempty"`
}

//  RefUserWisePropertyCollection : ""
type RefUserWisePropertyCollection struct {
	User             `bson:",inline"`
	PropertyPayments []PropertyPayment      `json:"propertypayments" bson:"propertypayments,omitempty"`
	TLPayments       []TradeLicensePayments `json:"tlPayments" bson:"tlPayments,omitempty"`
	SRPayments       []ShopRentPayments     `json:"srPayments" bson:"srPayments,omitempty"`
	MTPayments       []MobileTowerPayments  `json:"mtPayments" bson:"mtPayments,omitempty"`
	PropertyCash     PaymentCollected       `json:"propertyCash" bson:"propertycash,omitempty"`
	PropertyCheque   PaymentCollected       `json:"propertyCheque" bson:"propertyCheque,omitempty"`
	PropertyDD       PaymentCollected       `json:"propertydd" bson:"propertydd,omitempty"`
	PropertyNB       PaymentCollected       `json:"propertynb" bson:"propertynb,omitempty"`
	TLCash           PaymentCollected       `json:"tlCash" bson:"tlCash,omitempty"`
	TLCheque         PaymentCollected       `json:"tlCheque" bson:"tlCheque,omitempty"`
	TLDD             PaymentCollected       `json:"tldd" bson:"tldd,omitempty"`
	TLNB             PaymentCollected       `json:"tlnb" bson:"tlnb,omitempty"`
	SRCash           PaymentCollected       `json:"srCash" bson:"srCash,omitempty"`
	SRCheque         PaymentCollected       `json:"srCheque" bson:"srCheque,omitempty"`
	SRDD             PaymentCollected       `json:"srdd" bson:"srdd,omitempty"`
	SRNB             PaymentCollected       `json:"srnb" bson:"srnb,omitempty"`
	MTCash           PaymentCollected       `json:"mtCash" bson:"mtCash,omitempty"`
	MTCheque         PaymentCollected       `json:"mtCheque" bson:"mtCheque,omitempty"`
	MTDD             PaymentCollected       `json:"mtdd" bson:"mtdd,omitempty"`
	MTNB             PaymentCollected       `json:"mtnb" bson:"mtnb,omitempty"`
}

type RefUserWisePropertyCollectionWithTotal struct {
	Collection []RefUserWisePropertyCollection `json:"collection" bson:"collection,omitempty"`

	Total struct {
		TotalAmount        float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		TotalCashAmount    float64 `json:"totalCashAmount" bson:"totalCashAmount,omitempty"`
		TotalChequeAmount  float64 `json:"totalChequeAmount" bson:"totalChequeAmount,omitempty"`
		TotalNBAmount      float64 `json:"totalNBAmount" bson:"totalNBAmount,omitempty"`
		TotalPayment       float64 `json:"totalPayment" bson:"totalPayment,omitempty"`
		TotalCashPayment   float64 `json:"totalCashPayment" bson:"totalCashPayment,omitempty"`
		TotalChequePayment float64 `json:"totalChequePayment" bson:"totalChequePayment,omitempty"`
		TotalNBPayment     float64 `json:"totalNBPayment" bson:"totalNBPayment,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	Property struct {
		TotalAmount        float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		TotalCashAmount    float64 `json:"totalCashAmount" bson:"totalCashAmount,omitempty"`
		TotalChequeAmount  float64 `json:"totalChequeAmount" bson:"totalChequeAmount,omitempty"`
		TotalNBAmount      float64 `json:"totalNBAmount" bson:"totalNBAmount,omitempty"`
		TotalPayment       float64 `json:"totalPayment" bson:"totalPayment,omitempty"`
		TotalCashPayment   float64 `json:"totalCashPayment" bson:"totalCashPayment,omitempty"`
		TotalChequePayment float64 `json:"totalChequePayment" bson:"totalChequePayment,omitempty"`
		TotalNBPayment     float64 `json:"totalNBPayment" bson:"totalNBPayment,omitempty"`
	} `json:"property" bson:"property,omitempty"`
	TradeLicense struct {
		TotalAmount        float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		TotalCashAmount    float64 `json:"totalCashAmount" bson:"totalCashAmount,omitempty"`
		TotalChequeAmount  float64 `json:"totalChequeAmount" bson:"totalChequeAmount,omitempty"`
		TotalNBAmount      float64 `json:"totalNBAmount" bson:"totalNBAmount,omitempty"`
		TotalPayment       float64 `json:"totalPayment" bson:"totalPayment,omitempty"`
		TotalCashPayment   float64 `json:"totalCashPayment" bson:"totalCashPayment,omitempty"`
		TotalChequePayment float64 `json:"totalChequePayment" bson:"totalChequePayment,omitempty"`
		TotalNBPayment     float64 `json:"totalNBPayment" bson:"totalNBPayment,omitempty"`
	} `json:"tradeLicense" bson:"tradeLicense,omitempty"`
	ShopRent struct {
		TotalAmount        float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		TotalCashAmount    float64 `json:"totalCashAmount" bson:"totalCashAmount,omitempty"`
		TotalChequeAmount  float64 `json:"totalChequeAmount" bson:"totalChequeAmount,omitempty"`
		TotalNBAmount      float64 `json:"totalNBAmount" bson:"totalNBAmount,omitempty"`
		TotalPayment       float64 `json:"totalPayment" bson:"totalPayment,omitempty"`
		TotalCashPayment   float64 `json:"totalCashPayment" bson:"totalCashPayment,omitempty"`
		TotalChequePayment float64 `json:"totalChequePayment" bson:"totalChequePayment,omitempty"`
		TotalNBPayment     float64 `json:"totalNBPayment" bson:"totalNBPayment,omitempty"`
	} `json:"shopRent" bson:"shopRent,omitempty"`
	MobileTower struct {
		TotalAmount        float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		TotalCashAmount    float64 `json:"totalCashAmount" bson:"totalCashAmount,omitempty"`
		TotalChequeAmount  float64 `json:"totalChequeAmount" bson:"totalChequeAmount,omitempty"`
		TotalNBAmount      float64 `json:"totalNBAmount" bson:"totalNBAmount,omitempty"`
		TotalPayment       float64 `json:"totalPayment" bson:"totalPayment,omitempty"`
		TotalCashPayment   float64 `json:"totalCashPayment" bson:"totalCashPayment,omitempty"`
		TotalChequePayment float64 `json:"totalChequePayment" bson:"totalChequePayment,omitempty"`
		TotalNBPayment     float64 `json:"totalNBPayment" bson:"totalNBPayment,omitempty"`
	} `json:"mobileTower" bson:"mobileTower,omitempty"`
}
type PaymentCollected struct {
	NoOfPayments     float64                `json:"noOfPayments" bson:"noOfPayments,omitempty"`
	TotalAmount      float64                `json:"totalAmount" bson:"totalAmount,omitempty"`
	PropertyPayments []PropertyPayment      `json:"propertypayments" bson:"propertypayments,omitempty"`
	TLPayments       []TradeLicensePayments `json:"tlPayments" bson:"tlPayments,omitempty"`
	SRPayments       []ShopRentPayments     `json:"srPayments" bson:"srPayments,omitempty"`
	MTPayments       []MobileTowerPayments  `json:"mtPayments" bson:"mtPayments,omitempty"`
}

// RefHoldingWiseCollectionReport : ""
type RefHoldingWiseCollectionReport struct {
	Basic PropertyPaymentDemandBasic `json:"basic" bson:"basic,omitempty"`
	// Payment struct {
	ArrearTax          float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
	CurrentTax         float64 `json:"currentTax" bson:"currentTax,omitempty"`
	ArrearPenalty      float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
	CurrentPenalty     float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
	ArrearRebate       float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
	CurrentRebate      float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
	ArrearAlreadyPaid  float64 `json:"arrearAlreadyPaid" bson:"arrearAlreadyPaid,omitempty"`
	CurrentAlreadyPaid float64 `json:"currentAlreadyPaid" bson:"currentAlreadyPaid,omitempty"`
	OtherDemand        float64 `json:"otherDemand" bson:"otherDemand,omitempty"`
	TotalTax           float64 `json:"totalTax" bson:"totalTax,omitempty"`
	TotalAmount        float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
	PropertyID         string  `json:"propertyId" bson:"propertyId,omitempty"`
	// } `json:"payment,omitempty" bson:"payment,omitempty"`
	Ref struct {
		Owner           PropertyOwner `json:"owner" bson:"owner,omitempty"`
		Ward            Ward          `json:"ward" bson:"ward,omitempty"`
		RoadType        RoadType      `json:"roadType" bson:"roadType,omitempty"`
		PropertyType    PropertyType  `json:"propertyType" bson:"propertyType,omitempty"`
		Creator         User          `json:"creator" bson:"creator,omitempty"`
		Activator       User          `json:"activator" bson:"activator,omitempty"`
		PropertyDetails Property      `json:"propertyDetails" bson:"propertyDetails,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

// ResPropertyWiseDemandandCollectionV2Report : ""
type ResPropertyWiseDemandandCollectionV2Report struct {
	Property `bson:",inline"`
	Ref      struct {
		Payments struct {
			Rebate     float64 `json:"rebate" bson:"rebate,omitempty"`
			FormFee    float64 `json:"formFee" bson:"formFee,omitempty"`
			PaymentFys struct {
				ArrearTax      float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
				ArrearPenalty  float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
				ArrearRebate   float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
				CurrentTax     float64 `json:"currentTax" bson:"currentTax,omitempty"`
				CurrentPenalty float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
				CurrentRebate  float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
				TotalTax       float64 `json:"totalTax" bson:"totalTax,omitempty"`
				OtherDemand    float64 `json:"otherDemand" bson:"otherDemand,omitempty"`
			} `json:"paymentFys" bson:"paymentFys,omitempty"`
		} `json:"payments,omitempty" bson:"payments,omitempty"`
		Address       RefAddress            `json:"address" bson:"address,omitempty"`
		PropertyType  *RefPropertyType      `json:"propertyType" bson:"propertyType,omitempty"`
		RoadType      *RefRoadType          `json:"roadType" bson:"roadType,omitempty"`
		PropertyOwner *RefPropertyOwner     `json:"propertyOwner" bson:"propertyOwner,omitempty"`
		Demand        OverallPropertyDemand `json:"demand" bson:"demand,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RefPropertyDemandAndCollectionReport struct {
	UniqueID    string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Collections struct {
		ArrearTax          float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
		ArrearPenalty      float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
		ArrearRebate       float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
		CurrentTax         float64 `json:"currentTax" bson:"currentTax,omitempty"`
		CurrentPenalty     float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
		CurrentRebate      float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
		ArrearOtherDemand  float64 `json:"arrearOtherDemand" bson:"arrearOtherDemand,omitempty"`
		CurrentOtherDemand float64 `json:"currentOtherDemand" bson:"currentOtherDemand,omitempty"`
		BoreCharge         float64 `json:"boreCharge" bson:"boreCharge,omitempty"`
		FormFee            float64 `json:"formFee" bson:"formFee,omitempty"`
	} `json:"collections" bson:"collections,omitempty"`
	Demand       OverallPropertyDemand `json:"demand" bson:"demand,omitempty"`
	PropertyType *RefPropertyType      `json:"propertyType" bson:"propertyType,omitempty"`
}

type RefPropertyCollectionReport struct {
	TnxID      string `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID string `json:"propertyId" bson:"propertyId,omitempty"`
	ReciptNo   string `json:"reciptNo" bson:"reciptNo,omitempty"`
	Owner      Owner  `json:"owners" bson:"owners,omitempty"`
	Status     string `json:"status" bson:"status,omitempty"`
	Collector  struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"collector" bson:"collector,omitempty"`
	Activator struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"activator" bson:"activator,omitempty"`
	Creator struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"creator" bson:"creator,omitempty"`
	CompletionDate *time.Time `json:"completionDate" bson:"completionDate,omitempty"`
	Details        *struct {
		Amount float64 `json:"amount" bson:"amount,omitempty"`
		MadeAt struct {
			At string `json:"at" bson:"at,omitempty"`
		} `json:"madeAt" bson:"madeAt,omitempty"`
		MOP struct {
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
		} `json:"mop" bson:"mop,omitempty"`
	} `json:"details" bson:"details,omitempty"`
	Demand *struct {
		FormFee            float64            `json:"formFee" bson:"formFee"`
		OtherDemand        float64            `json:"otherDemand" bson:"otherDemand"`
		BoreCharge         float64            `json:"boreCharge" bson:"boreCharge"`
		PreviousCollection PreviousCollection `json:"previousCollection" bson:"previousCollection"`
	} `json:"demand" bson:"demand,omitempty"`
	Property struct {
		Address struct {
			WardCode string `json:"wardCode" bson:"wardCode,omitempty"`
		} `json:"address,omitempty" bson:"address,omitempty"`
		ApplicationNo    string `json:"applicationNo" bson:"applicationNo,omitempty"`
		OldHoldingNumber string `json:"oldHoldingNumber" bson:"oldHoldingNumber,omitempty"`
	} `json:"property,omitempty" bson:"property,omitempty"`
	Address struct {
		AL1 string `json:"al1" bson:"al1,omitempty"`
		Al2 string `json:"al2" bson:"al2,omitempty"`
	} `json:"address,omitempty" bson:"address,omitempty"`
	PaymentFYs struct {
		ArrearTax          float64              `json:"arrearTax" bson:"arrearTax,omitempty"`
		CurrentTax         float64              `json:"currentTax" bson:"currentTax,omitempty"`
		ArrearPenalty      float64              `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
		CurrentPenalty     float64              `json:"currentPenalty" bson:"currentPenalty,omitempty"`
		ArrearRebate       float64              `json:"arrearRebate" bson:"arrearRebate,omitempty"`
		CurrentRebate      float64              `json:"currentRebate" bson:"currentRebate,omitempty"`
		ArrearAlreadyPaid  float64              `json:"arrearAlreadyPaid" bson:"arrearAlreadyPaid,omitempty"`
		CurrentAlreadyPaid float64              `json:"currentAlreadyPaid" bson:"currentAlreadyPaid,omitempty"`
		OtherDemand        float64              `json:"otherDemand" bson:"otherDemand,omitempty"`
		TotalTax           float64              `json:"totalTax" bson:"totalTax,omitempty"`
		FYs                []PropertyPaymentdfy `json:"fys" bson:"fys,omitempty"`
	} `json:"paymentFys,omitempty" bson:"paymentFys,omitempty"`
}
type Owner struct {
	Name               string `json:"name" bson:"name,omitempty"`
	Mobile             string `json:"mobile" bson:"mobile,omitempty"`
	FatherRpanRhusband string `json:"fatherRpanRhusband" bson:"fatherRpanRhusband,omitempty"`
}

type PropertyPaymentdfy struct {
	FY struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"fy,omitempty" bson:"fy,omitempty"`
}

type Summary struct {
	TnxID              string  `json:"tnxId" bson:"_id,omitempty"`
	ArrearTax          float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
	TotalTax           float64 `json:"totalTax" bson:"totalTax,omitempty"`
	ArrearRebate       float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
	ArrearPenalty      float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
	CurrentTax         float64 `json:"currentTax" bson:"currentTax,omitempty"`
	TotalCurrent       float64 `json:"totalCurrent" bson:"totalCurrent,omitempty"`
	TotalArrear        float64 `json:"totalArrear" bson:"totalArrear,omitempty"`
	CurrentRebate      float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
	CurrentPenalty     float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
	BoreCharge         float64 `json:"boreCharge" bson:"boreCharge,omitempty"`
	FormFee            float64 `json:"formFee" bson:"formFee,omitempty"`
	ArrearOtherDemand  float64 `json:"arrearOtherDemand" bson:"arrearOtherDemand,omitempty"`
	CurrentOtherDemand float64 `json:"currentOtherDemand" bson:"currentOtherDemand,omitempty"`
	FromFy             string  `json:"fromFy" bson:"fromFy,omitempty"`
	ToFy               string  `json:"toFy" bson:"toFy,omitempty"`
}

type RefHoldingWiseCollectionReportV2 struct {
	PropertyID string `json:"propertyId" bson:"propertyId,omitempty"`
	Basic      struct {
		ApplicationNo    string `json:"applicationNo" bson:"applicationNo,omitempty"`
		OldHoldingNumber string `json:"oldHoldingNumber" bson:"oldHoldingNumber,omitempty"`
		Address          struct {
			AL1 string `json:"al1" bson:"al1,omitempty"`
			Al2 string `json:"al2" bson:"al2,omitempty"`
		} `json:"address,omitempty" bson:"address,omitempty"`
	} `json:"basic" bson:"basic,omitempty"`
	Ward struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"ward" bson:"ward,omitempty"`
	Activator struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"activator" bson:"activator,omitempty"`
	PropertyType struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"propertyType" bson:"propertyType,omitempty"`
	RoadType struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"roadType" bson:"roadType,omitempty"`
	Creator struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"creator" bson:"creator,omitempty"`
	Owner   Owner `json:"owners" bson:"owners,omitempty"`
	Payment struct {
		ArrearTax          float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
		CurrentTax         float64 `json:"currentTax" bson:"currentTax,omitempty"`
		ArrearPenalty      float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
		CurrentPenalty     float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
		ArrearRebate       float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
		CurrentRebate      float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
		ArrearAlreadyPaid  float64 `json:"arrearAlreadyPaid" bson:"arrearAlreadyPaid,omitempty"`
		CurrentAlreadyPaid float64 `json:"currentAlreadyPaid" bson:"currentAlreadyPaid,omitempty"`
		OtherDemand        float64 `json:"otherDemand" bson:"otherDemand,omitempty"`
		TotalTax           float64 `json:"totalTax" bson:"totalTax,omitempty"`
		TotalAmount        float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		PropertyID         string  `json:"propertyId" bson:"propertyId,omitempty"`
	} `json:"payment,omitempty" bson:"payment,omitempty"`
}

type RefPropertyV2 struct {
	ApplicationNo string  `json:"applicationNo" bson:"applicationNo,omitempty"`
	UniqueID      string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	AreaOfPlot    float64 `json:"areaOfPlot" bson:"areaOfPlot,omitempty"`
	Advance       float64 `json:"advance" bson:"advance"`
	Address       struct {
		AL1 string `json:"al1" bson:"al1,omitempty"`
		Al2 string `json:"al2" bson:"al2,omitempty"`
	} `json:"address" bson:"address"`
	Ref struct {
		Address struct {
			Ward struct {
				Name string `json:"name" bson:"name,omitempty"`
			} `json:"ward" bson:"ward,omitempty"`
		} `json:"address" bson:"address"`
		Demand struct {
			Arrear struct {
				TotalTax float64 `json:"totalTax" bson:"totalTax,omitempty"`
			} `json:"arrear" bson:"arrear,omitempty"`
			Current struct {
				TotalTax float64 `json:"totalTax" bson:"totalTax,omitempty"`
			} `json:"current" bson:"current,omitempty"`
			Total struct {
				TotalTax float64 `json:"totalTax" bson:"totalTax,omitempty"`
				Ecess    float64 `json:"ecess" bson:"ecess,omitempty"`
			} `json:"total" bson:"total,omitempty"`
		} `json:"demand" bson:"demand"`
		PropertyType struct {
			Name string `json:"name" bson:"name,omitempty"`
		} `json:"propertyType" bson:"propertyType,omitempty"`
		RoadType struct {
			Name string `json:"name" bson:"name,omitempty"`
		} `json:"roadType" bson:"roadType,omitempty"`
		Collections struct {
			ArrearTax      float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
			ArrearPenalty  float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
			ArrearRebate   float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
			CurrentTax     float64 `json:"currentTax" bson:"currentTax,omitempty"`
			CurrentPenalty float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
			CurrentRebate  float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
			TotalTax       float64 `json:"totalTax" bson:"totalTax,omitempty"`
			OtherDemand    float64 `json:"otherDemand" bson:"otherDemand,omitempty"`
		} `json:"collections" bson:"collections,omitempty"`
		PropertyPayments struct {
			Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
			FormFee float64 `json:"formFee" bson:"formFee,omitempty"`
		} `json:"propertyPayments" bson:"propertyPayments,omitempty"`
		PropertyOwner []struct {
			Name               string `json:"name" bson:"name,omitempty"`
			Mobile             string `json:"mobile" bson:"mobile,omitempty"`
			FatherRpanRhusband string `json:"fatherRpanRhusband" bson:"fatherRpanRhusband,omitempty"`
		} `json:"propertyOwner" bson:"propertyOwner,omitempty"`
		// Inc           func(int) int
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
