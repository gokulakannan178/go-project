package models

import (
	"fmt"
)

//PaymentGateway : ""
type PaytmtInitTranscation struct {
	Body struct {
		RequestType string `json:"requestType" bson:"requestType,omitempty"`
		MID         string `json:"mid"  bson:"mid,omitempty"`
		WebsiteName string `json:"websiteName"  bson:"websiteName,omitempty"`
		OrderID     string `json:"orderId"  bson:"orderId,omitempty"`
		TxnAmount   struct {
			Value    string `json:"value"  bson:"value,omitempty"`
			Currency string `json:"currency"  bson:"currency,omitempty"`
		} `json:"txnAmount"  bson:"txnAmount,omitempty"`
		UserInfo struct {
			CustId string `json:"custId"  bson:"custId,omitempty"`
		} `json:"userInfo"  bson:"userInfo,omitempty"`

		CallBackUrl string `json:"callbackUrl"  bson:"callbackUrl,omitempty"`
	} `json:"body" bson:"body,omitempty"`
	Head struct {
		Signature string `json:"signature" bson:"signature,omitempty"`
	} `json:"head" bson:"head,omitempty"`
}

func (pit *PaytmtInitTranscation) GetInitiateTransactionAPIURL(pg *PaymentGateway) string {
	url := fmt.Sprintf("%v/theia/api/v1/initiateTransaction?mid=%v&orderId=%v", pg.BaseURL, pit.Body.MID, pit.Body.OrderID)
	return url
}

type PaytmtInitTranscationResponse struct {
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
		TxnToken         string `json:"txnToken"`
		IsPromoCodeValid bool   `json:"isPromoCodeValid"`
		Authenticated    bool   `json:"authenticated"`
	} `json:"body"`
}

//PaymentGateway : ""
type UserPaytmPaymentInit struct {
	TxnAmount  float64 `json:"txnAmount"  bson:"txnAmount,omitempty"`
	CustomerID string  `json:"customerId"  bson:"customerId,omitempty"`
	OrderId    string  `json:"orderId"  bson:"orderId,omitempty"`
}

//UserPaytmPaymentInitResponse : ""
type UserPaytmPaymentInitResponse struct {
	TxnAmount  float64 `json:"txnAmount"  bson:"txnAmount,omitempty"`
	CustomerID string  `json:"customerId"  bson:"customerId,omitempty"`
	OrderId    string  `json:"orderId"  bson:"orderId,omitempty"`
}
