package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//CreatedBankDeposit :""
func (s *Service) CreatedBankDeposit(ctx *models.Context, bankDeposit *models.BankDeposit) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	bankDeposit.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBANKDEPOSIT)
	bankDeposit.Status = constants.BANKDEPOSITSTATUSPENDING
	t := time.Now()
	bankDeposit.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveBankDeposit(ctx, bankDeposit)
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

//VerifierBankDeposit
func (s *Service) VerifierBankDeposit(ctx *models.Context, uniqueID string, Verifier *models.BankDepositVerifier) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	t := time.Now()
	Verifier.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.VerifyBankDeposit(ctx, uniqueID, constants.BANKDEPOSITSTATUSVERIFIED, Verifier)
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

//GetSingleBankDeposit
func (s *Service) GetSingleBankDeposit(ctx *models.Context, uniqueID string) (*models.BankDeposit, error) {
	return s.Daos.GetSingleBankDeposit(ctx, uniqueID)

}

//NotVerifierBankDeposit
func (s *Service) NotVerifierBankDeposit(ctx *models.Context, uniqueID string, Verifier *models.BankDepositVerifier) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	t := time.Now()
	Verifier.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.NotVerifyBankDeposit(ctx, uniqueID, constants.BANKDEPOSITSTATUSNOTVERIFIED, Verifier)
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

//BankDepositFilter
func (s *Service) BankDepositFilter(ctx *models.Context, bdf *models.BankDepositFilter, pagination *models.Pagination) ([]models.BankDeposit, error) {
	return s.Daos.BankDepositFilter(ctx, bdf, pagination)
}
