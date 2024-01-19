package models

// TradeLicensePaymentsPart2 : ""
type TradeLicensePaymentsPart2 struct {
	TradeLicensePayments `bson:",inline"`
}

// TradeLicensePaymentsBasicsPart2 : ""
type TradeLicensePaymentsBasicsPart2 struct {
	TradeLicensePaymentsBasics `bson:",inline"`
}

// MakeTradeLicensePaymentReqPart2 : ""
type MakeTradeLicensePaymentReqPart2 struct {
	MakeTradeLicensePaymentReq `bson:",inline"`
}

// TradeLicensePaymentsFilterPart2 : ""
type TradeLicensePaymentsFilterPart2 struct {
	TradeLicensePaymentsFilter `bson:",inline"`
}

// RefTradeLicensePaymentsPart2 : ""

type RefTradeLicensePaymentsPart2 struct {
	RefTradeLicensePayments `bson:",inline"`
}
type MakeTradeLicensePaymentsActionPart2 struct {
	MakeTradeLicensePaymentsAction `bson:",inline"`
}

type RefBasicTradeLicenseUpdateLogV2Part2 struct {
	RefBasicTradeLicenseUpdateLogV2 `bson:",inline"`
}
