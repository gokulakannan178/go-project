package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// PaytmtQrCodeInitTranscation :""
func (s *Service) PaytmtQrCodeInitTranscation(ctx *models.Context, upp *models.QrCodePaytmPaymentInit) (interface{}, error) {
	log.Println("transaction start")
	var responseData interface{}
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return responseData, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		pg, err := s.Daos.GetSinglePaymentGateway(ctx, "1")
		if err != nil {
			return errors.New("Err in geting pg config" + err.Error())
		}
		if pg == nil {
			return errors.New("Err in geting pg config - nil")
		}
		var pit models.PaytmtQrCodeInitTranscation
		pit.Body.BusinessType = upp.BusinessType
		pit.Body.MID = pg.MID
		pit.Body.PosID = upp.PosID
		pit.Body.OrderID = upp.OrderId
		pit.Body.Amount = fmt.Sprintf("%v", upp.TxnAmount)

		signature, err := s.Shared.PatymCheckSum.GenerateSignature(pit.Body, pg.MKey)
		if err != nil {
			return errors.New("Error in generating signature - " + err.Error())
		}
		fmt.Println("signature generated suddccessfully")
		pit.Head.Signature = signature
		pit.Head.ClientId = upp.ClientID
		pit.Head.Version = upp.PosID
		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"
		resp, err := s.Shared.Post(pit.GetInitiateQRCodeTransactionAPIURL(pg), headers, pit)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		//Read the response body
		var pitr *models.PaytmtQrCodeInitTranscationResponse

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
