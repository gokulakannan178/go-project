package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveHDFCPaymentGateway : ""
func (s *Service) SaveHDFCPaymentGateway(ctx *models.Context, hdfc *models.HDFCPaymentGateway) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	hdfc.UniqueID = "1"
	hdfc.Status = constants.HDFCPAYMENTGATEWAYACTIVE
	t := time.Now()
	hdfc.Created.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.UpsertHDFCPaymentGateway(ctx, hdfc)
		if dberr != nil {
			return dberr
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

//GetSingleProductConfiguration : ""
func (s *Service) GetSingleDefaultHDFCPaymentGateway(ctx *models.Context) (*models.RefHDFCPaymentGateway, error) {
	return s.Daos.GetSingleDefaultHDFCPaymentGateway(ctx)
}

//GetSingleMerchantIDHDFCPaymentGateway : ""
func (s *Service) GetSingleMerchantIDHDFCPaymentGateway(ctx *models.Context) (string, error) {
	return s.Daos.GetSingleMerchantIDHDFCPaymentGateway(ctx)
}

//check Payment Status
func (s *Service) CheckPaymentStatus(ctx *models.Context, cps *models.HDFCPaymentGatewayCheckPaymentStatus) (*models.HDFCPaymentGatewayCheckPaymentStatusResponse, error) {
	jsonData, err := json.Marshal(cps)
	if err != nil {
		return nil, errors.New("Err in marshal data - " + err.Error())
	}
	hdfcPGConfig, err := s.GetSingleDefaultHDFCPaymentGateway(ctx)
	if err != nil {
		return nil, errors.New("Err in getting pg config  - " + err.Error())
	}
	if hdfcPGConfig == nil {
		return nil, errors.New("pg config  is nil")
	}
	encRequest := s.Shared.HDFCencrypt(jsonData, hdfcPGConfig.WorkingKey)
	data := url.Values{}
	data.Set("enc_request", encRequest)
	data.Set("access_code", hdfcPGConfig.AccessCode)
	data.Set("request_type", "JSON")
	data.Set("response_type", "JSON")
	data.Set("command", "orderStatusTracker")
	data.Set("order_no", cps.OrderNo)
	fmt.Println(data)

	req, err := http.NewRequest("POST", "https://apitest.ccavenue.com/apis/servlet/DoWebTrans",
		bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Err in getting status api response  - " + err.Error())
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Err in unmarshaling status api response data  - " + err.Error())
	}
	fmt.Println("Response body:", string(bodyBytes))
	response, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return nil, errors.New("Error decoding response body  - " + err.Error())
	}
	enc_response := response.Get("enc_response")
	respStatus := response.Get("status")
	if respStatus == "1" {
		return nil, errors.New("status api failed  - " + enc_response)
	}
	decryptedRespData := s.Shared.HDFCdecrypt(enc_response, hdfcPGConfig.WorkingKey)
	// if err := json.NewDecoder(resp.Body).Decode(&respJSON); err != nil {
	// 	return errors.New("Err in unmarshaling status api response data  - " + err.Error())
	// }
	// pairs := strings.Split(decryptedRespData, "&")
	// // Create a new map
	// respData := make(map[string]string)

	// for _, pair := range pairs {
	// 	parts := strings.Split(pair, "=")
	// 	if len(parts) >= 2 {
	// 		key := parts[0]
	// 		value := parts[1]
	// 		respData[key] = value
	// 	}
	// }
	fmt.Println("decryptedRespData==========================")
	fmt.Println(decryptedRespData)
	var hpgcpsr models.HDFCPaymentGatewayCheckPaymentStatusResponse
	// var hpgcpsr2 interface{}
	// d := []byte(decryptedRespData)
	// if err := json.Unmarshal(d, &hpgcpsr); err != nil {
	// 	return nil, errors.New("Error in unmarshaling dectypted data  - " + err.Error())
	// }
	if err := json.NewDecoder(strings.NewReader(decryptedRespData)).Decode(&hpgcpsr); err != nil {
		return nil, errors.New("Error in unmarshaling dectypted data  - " + err.Error())
	}

	fmt.Println("hpgcpsr=========================")
	fmt.Println(hpgcpsr)

	return &hpgcpsr, nil
}
