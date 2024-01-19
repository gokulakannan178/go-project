package models

type HDFCPaymentGateway struct {
	UniqueID   string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	WorkingKey string  `json:"workingKey" bson:"workingKey,omitempty"`
	AccessCode string  `json:"accessCode" bson:"accessCode,omitempty"`
	IsDefault  bool    `json:"isDefault" bson:"isDefault,omitempty"`
	Status     string  `json:"status" bson:"status,omitempty"`
	BaseURL    string  `json:"baseUrl" bson:"baseUrl,omitempty"`
	MerchantID string  `json:"merchantId" bson:"merchantId,omitempty"`
	Created    Created `json:"created" bson:"created,omitempty"`
}

type HDFCPaymentGatewayFilter struct {
	SortBy    string `json:"sortBy,omitempty"`
	SortOrder int    `json:"sortOrder,omitempty"`
}

type RefHDFCPaymentGateway struct {
	HDFCPaymentGateway `bson:",inline"`
	Ref                struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type HDFCPaymentGatewayCheckPaymentStatus struct {
	ReferenceNO string `json:"reference_no" bson:"reference_no,omitempty"`
	OrderNo     string `json:"order_no" bson:"order_no,omitempty"`
}

type HDFCPaymentGatewayCheckPaymentStatusResponse struct {
	OrderStatusResult HDFCOrderStatusResult `json:"Order_Status_Result" bson:"Order_Status_Result,omitempty"`
}
type HDFCOrderStatusResult struct {
	OrderGtwID          string `json:"order_gtw_id"`
	OrderNo             string `json:"order_no"`
	OrderShipZip        string `json:"order_ship_zip"`
	OrderShipAddress    string `json:"order_ship_address"`
	OrderBillEmail      string `json:"order_bill_email"`
	OrderCaptAmt        int    `json:"order_capt_amt"`
	OrderShipTel        string `json:"order_ship_tel"`
	OrderShipName       string `json:"order_ship_name"`
	OrderBillCountry    string `json:"order_bill_country"`
	OrderCardName       string `json:"order_card_name"`
	OrderStatus         string `json:"order_status"`
	OrderBillState      string `json:"order_bill_state"`
	OrderTax            int    `json:"order_tax"`
	OrderBillCity       string `json:"order_bill_city"`
	OrderShipState      string `json:"order_ship_state"`
	OrderDiscount       int    `json:"order_discount"`
	OrderTDS            int    `json:"order_TDS"`
	OrderDateTime       string `json:"order_date_time"`
	OrderShipCountry    string `json:"order_ship_country"`
	OrderBillAddress    string `json:"order_bill_address"`
	OrderFeePercValue   int    `json:"order_fee_perc_value"`
	OrderIP             string `json:"order_ip"`
	OrderOptionType     string `json:"order_option_type"`
	OrderBankRefNo      int64  `json:"order_bank_ref_no"`
	OrderCurrncy        string `json:"order_currncy"`
	OrderFeeFlat        int    `json:"order_fee_flat"`
	OrderShipCity       string `json:"order_ship_city"`
	OrderBillTel        string `json:"order_bill_tel"`
	OrderDeviceType     string `json:"order_device_type"`
	OrderGrossAmt       int    `json:"order_gross_amt"`
	OrderAmt            int    `json:"order_amt"`
	OrderFraudStatus    string `json:"order_fraud_status"`
	OrderBillZip        string `json:"order_bill_zip"`
	OrderBillName       string `json:"order_bill_name"`
	ReferenceNo         int64  `json:"reference_no"`
	OrderBankResponse   string `json:"order_bank_response"`
	OrderStatusDateTime string `json:"order_status_date_time"`
	Status              int    `json:"status"`
}
