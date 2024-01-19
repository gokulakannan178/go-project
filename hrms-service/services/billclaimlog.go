package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveBillClaimLog :""
func (s *Service) SaveBillClaimLog(ctx *models.Context, billClaimLog *models.BillClaimLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	billClaimLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIMLOG)
	billClaimLog.Status = constants.BILLCLAIMLOGSTATUSPENDING
	//t := time.Now()

	log.Println("b4 BillClaimLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveBillClaimLog(ctx, billClaimLog)
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

//UpdateBillClaimLog : ""
func (s *Service) UpdateBillClaimLog(ctx *models.Context, billClaimLog *models.BillClaimLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateBillClaimLog(ctx, billClaimLog)
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

//EnableBillClaimLog : ""
func (s *Service) EnableBillClaimLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableBillClaimLog(ctx, UniqueID)
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

//DisableBillClaimLog : ""
func (s *Service) DisableBillClaimLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableBillClaimLog(ctx, UniqueID)
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

//DeleteBillClaimLog : ""
func (s *Service) DeleteBillClaimLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteBillClaimLog(ctx, UniqueID)
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

//GetSingleBillClaimLog :""
func (s *Service) GetSingleBillClaimLog(ctx *models.Context, UniqueID string) (*models.RefBillClaimLog, error) {
	billClaimLog, err := s.Daos.GetSingleBillClaimLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return billClaimLog, nil
}

//FilterBillClaimLog :""
func (s *Service) FilterBillClaimLog(ctx *models.Context, filter *models.FilterBillClaimLog, pagination *models.Pagination) ([]models.RefBillClaimLog, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterBillClaimLog(ctx, filter, pagination)

}
