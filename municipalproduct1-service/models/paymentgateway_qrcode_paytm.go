package models

import "fmt"

//PaytmtQrCodeInitTranscation : ""
type PaytmtQrCodeInitTranscation struct {
	Body struct {
		BusinessType string `json:"businessType" bson:"businessType,omitempty"`
		MID          string `json:"mid"  bson:"mid,omitempty"`
		PosID        string `json:"posId"  bson:"posId,omitempty"`
		OrderID      string `json:"orderId"  bson:"orderId,omitempty"`
		Amount       string `json:"amount"  bson:"amount,omitempty"`
	} `json:"body" bson:"body,omitempty"`
	Head struct {
		ClientId  string `json:"clientId" bson:"clientId,omitempty"`
		Version   string `json:"version" bson:"version,omitempty"`
		Signature string `json:"signature" bson:"signature,omitempty"`
	} `json:"head" bson:"head,omitempty"`
}

func (p *PaytmtQrCodeInitTranscation) GetInitiateQRCodeTransactionAPIURL(pg *PaymentGateway) string {
	url := fmt.Sprintf("%v/paymentservices/qr/create", pg.BaseURL)
	return url
}

// PaytmtQrCodeInitTranscationResponse : "Used to get the response from the PaytmQrCode Init API"
type PaytmtQrCodeInitTranscationResponse struct {
	Head struct {
		ResponseTimestamp string `json:"responseTimestamp"`
		Version           string `json:"version"`
		ClientID          string `json:"clientId"`
		Signature         string `json:"signature"`
	} `json:"head"`
	Body struct {
		ResultInfo struct {
			ResultStatus string `json:"resultStatus"`
			ResultCode   string `json:"resultCode"`
			ResultMsg    string `json:"resultMsg"`
		} `json:"resultInfo"`
		QrData   string `json:"qrData"`
		QrCodeId string `json:"qrCodeId"`
		Image    string `json:"image"`
	} `json:"body"`
}

//QrCodePaytmPaymentInit : ""
type QrCodePaytmPaymentInit struct {
	TxnAmount    float64 `json:"txnAmount"  bson:"txnAmount,omitempty"`
	PosID        string  `json:"posId"  bson:"posId,omitempty"`
	OrderId      string  `json:"orderId"  bson:"orderId,omitempty"`
	BusinessType string  `json:"businessType"  bson:"businessType,omitempty"`
	ClientID     string  `json:"clientId" bson:"clientId,omitempty"`
}
