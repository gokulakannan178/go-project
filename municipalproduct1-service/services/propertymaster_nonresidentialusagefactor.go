package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveNonResidentialUsageFactor :""
func (s *Service) SaveNonResidentialUsageFactor(ctx *models.Context, nonResidentialUsageFactor *models.NonResidentialUsageFactor) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	nonResidentialUsageFactor.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR)
	nonResidentialUsageFactor.Status = constants.NONRESIDENTIALUSAGEFACTORSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 nonResidentialUsageFactor.created")
	nonResidentialUsageFactor.Created = created
	log.Println("b4 nonResidentialUsageFactor.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveNonResidentialUsageFactor(ctx, nonResidentialUsageFactor)
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

//UpdateNonResidentialUsageFactor : ""
func (s *Service) UpdateNonResidentialUsageFactor(ctx *models.Context, nonResidentialUsageFactor *models.NonResidentialUsageFactor) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateNonResidentialUsageFactor(ctx, nonResidentialUsageFactor)
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

//EnableNonResidentialUsageFactor : ""
func (s *Service) EnableNonResidentialUsageFactor(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableNonResidentialUsageFactor(ctx, UniqueID)
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

//DisableNonResidentialUsageFactor : ""
func (s *Service) DisableNonResidentialUsageFactor(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableNonResidentialUsageFactor(ctx, UniqueID)
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

//DeleteNonResidentialUsageFactor : ""
func (s *Service) DeleteNonResidentialUsageFactor(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteNonResidentialUsageFactor(ctx, UniqueID)
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

//GetSingleNonResidentialUsageFactor :""
func (s *Service) GetSingleNonResidentialUsageFactor(ctx *models.Context, UniqueID string) (*models.RefNonResidentialUsageFactor, error) {
	nonResidentialUsageFactor, err := s.Daos.GetSingleNonResidentialUsageFactor(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return nonResidentialUsageFactor, nil
}

//FilterNonResidentialUsageFactor :""
func (s *Service) FilterNonResidentialUsageFactor(ctx *models.Context, nonResidentialUsageFactorfilter *models.NonResidentialUsageFactorFilter, pagination *models.Pagination) (nonResidentialUsageFactor []models.RefNonResidentialUsageFactor, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.FilterNonResidentialUsageFactor(ctx, nonResidentialUsageFactorfilter, pagination)
}
