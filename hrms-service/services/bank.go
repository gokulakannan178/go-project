package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveBankInformation :""
func (s *Service) SaveBankInformation(ctx *models.Context, bankInformation *models.BankInformation) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	bankInformation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBANKINFORMATION)
	bankInformation.Status = constants.BANKINFORMATIONSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 BankInformation.created")
	bankInformation.Created = &created
	log.Println("b4 BankInformation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveBankInformation(ctx, bankInformation)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//UpdateBankInformation : ""
func (s *Service) UpdateBankInformation(ctx *models.Context, bankInformation *models.BankInformation) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateBankInformation(ctx, bankInformation)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//EnableBankInformation : ""
func (s *Service) EnableBankInformation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableBankInformation(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DisableBankInformation : ""
func (s *Service) DisableBankInformation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableBankInformation(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DeleteBankInformation : ""
func (s *Service) DeleteBankInformation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteBankInformation(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//GetSingleBankInformation :""
func (s *Service) GetSingleBankInformation(ctx *models.Context, UniqueID string) (*models.RefBankInformation, error) {
	bankInformation, err := s.Daos.GetSingleBankInformation(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return bankInformation, nil
}

//FilterBankInformation :""
func (s *Service) FilterBankInformation(ctx *models.Context, filter *models.BankInformationFilter, pagination *models.Pagination) (BankInformation []models.RefBankInformation, err error) {
	return s.Daos.FilterBankInformation(ctx, filter, pagination)
}
func (s *Service) UpdateEmployeeBankInformation(ctx *models.Context, bankInformation *models.BankInformation) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	bankInformation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBANKINFORMATION)
	bankInformation.Status = constants.BANKINFORMATIONSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 BankInformation.created")
	bankInformation.Created = &created
	log.Println("b4 BankInformation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeBankInformation(ctx, bankInformation)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}
func (s *Service) GetSingleBankInformationWithEmployee(ctx *models.Context, UniqueID string) (*models.RefBankInformation, error) {
	bankInformation, err := s.Daos.GetSingleBankInformationWithEmployee(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return bankInformation, nil
}
