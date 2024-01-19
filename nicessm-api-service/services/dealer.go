package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDealer :""
func (s *Service) SaveDealer(ctx *models.Context, dealer *models.Dealer) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//Dealer.Code = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDealer)

	dealer.Status = constants.DEALERSTATUSACTIVE
	dealer.ActiveStatus = true
	dealer.Certification.Status = constants.DEALERCERTIFICATIONSTATUSNOTAPPLIED
	t := time.Now()
	dealer.Certification.ActionDate = t
	dealer.Certification.AppliedDate = t
	dealer.Certification.ExpiryDate = t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Dealer.created")
	dealer.Created = created
	log.Println("b4 Dealer.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDealer(ctx, dealer)
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

//UpdateDealer : ""
func (s *Service) UpdateDealer(ctx *models.Context, dealer *models.Dealer) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDealer(ctx, dealer)
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

//EnableDealer : ""
func (s *Service) EnableDealer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDealer(ctx, UniqueID)
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

//DisableDealer : ""
func (s *Service) DisableDealer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDealer(ctx, UniqueID)
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

//DeleteDealer : ""
func (s *Service) DeleteDealer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDealer(ctx, UniqueID)
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

//GetSingleDealer :""
func (s *Service) GetSingleDealer(ctx *models.Context, UniqueID string) (*models.RefDealer, error) {
	dealer, err := s.Daos.GetSingleDealer(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return dealer, nil
}

//FilterDealer :""
func (s *Service) FilterDealer(ctx *models.Context, filter *models.DealerFilter, pagination *models.Pagination) (Dealer []models.RefDealer, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDealer(ctx, filter, pagination)

}

func (s *Service) DealerUniquenessCheckRegistration(ctx *models.Context, Param string, Value string) (*models.DealerUniquinessChk, error) {
	dealer, err := s.Daos.DealerUniquenessCheckRegistration(ctx, Param, Value)
	if err != nil {
		return nil, err
	}
	return dealer, nil
}
func (s *Service) DealerNearBy(ctx *models.Context, dealernb *models.NearBy, pagination *models.Pagination) ([]models.RefDealer, error) {

	ulbs, err := s.Daos.DealerNearBy(ctx, dealernb, pagination)
	if err != nil {
		return nil, err
	}

	return ulbs, nil
}
func (s *Service) DealerCertificationApply(ctx *models.Context, dealerID string, certification *models.DealerCertification) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		certification.AppliedDate = t
		err := s.Daos.DealerCertificationApply(ctx, dealerID, certification)
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
func (s *Service) DealerCertificationApprove(ctx *models.Context, dealerID string, certification *models.DealerCertification) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		prod, er2 := s.Daos.GetactiveProductConfig(ctx, true)
		if er2 != nil {
			return er2
		}
		if prod == nil {
			return errors.New("product config is nil")

		}
		t := time.Now()
		fmt.Println(t)
		t2 := t.AddDate(0, int(prod.ExpiryMonth), 0)
		fmt.Println("t2  ", t2)
		certification.ActionDate = t
		certification.ExpiryDate = t2
		err := s.Daos.DealerCertificationApprove(ctx, dealerID, certification)
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
func (s *Service) DealerCertificationReject(ctx *models.Context, dealerID string, certification *models.DealerCertification) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		certification.AppliedDate = t
		err := s.Daos.DealerCertificationReject(ctx, dealerID, certification)
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
