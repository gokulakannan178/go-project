package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveUser :""
func (s *Service) PaytmPaymentInitiateTransaction(ctx *models.Context, upp *models.UserPaytmPaymentInit) (interface{}, error) {
	log.Println("transaction start")
	var responseData interface{}
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return responseData, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		pg, err := s.Daos.GetDefaultPaymentGateway(ctx)
		if err != nil {
			return errors.New("Err in geting pg config" + err.Error())
		}
		if pg == nil {
			return errors.New("Err in geting pg config - nil")
		}
		var pit models.PaytmtInitTranscation
		pit.Body.RequestType = "Payment"
		pit.Body.MID = pg.MID
		pit.Body.WebsiteName = pg.WebsiteName
		pit.Body.OrderID = upp.OrderId
		pit.Body.TxnAmount.Value = fmt.Sprintf("%v", upp.TxnAmount)
		pit.Body.TxnAmount.Currency = pg.Currency
		pit.Body.UserInfo.CustId = upp.CustomerID
		pit.Body.CallBackUrl = pg.WebCallbackURL
		signature, err := s.Shared.PatymCheckSum.GenerateSignature(pit.Body, pg.MKey)
		if err != nil {
			return errors.New("Error in generating signature - " + err.Error())
		}
		fmt.Println("signature generated suddccessfully")
		pit.Head.Signature = signature
		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"
		resp, err := s.Shared.Post(pit.GetInitiateTransactionAPIURL(pg), headers, pit)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		//Read the response body
		var pitr *models.PaytmtInitTranscationResponse

		err = json.NewDecoder(resp.Body).Decode(&pitr)
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	return err
		// }
		// err = json.Unmarshal(body, &pit)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		responseData = pitr
		log.Println(responseData)
		return nil

	}); err != nil {
		return responseData, err
	}

	return responseData, nil
}

func (s *Service) CreateChecksum(ctx *models.Context, body interface{}) (string, error) {
	log.Println("transaction start")
	pg, err := s.Daos.GetSinglePaymentGateway(ctx, "1")
	if err != nil {
		return "", errors.New("Err in geting pg config" + err.Error())
	}
	signature, err := s.Shared.PatymCheckSum.GenerateSignature(body, pg.MKey)
	if err != nil {
		return "", errors.New("Error in generating signature - " + err.Error())
	}
	fmt.Println("signature generated suddccessfully")

	return signature, nil
}

// GetDefaultPaymentGateway : ""
func (s *Service) GetDefaultPaymentGateway(ctx *models.Context) (*models.PaymentGateway, error) {
	return s.Daos.GetDefaultPaymentGateway(ctx)
}
