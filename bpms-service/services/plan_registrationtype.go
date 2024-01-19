package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePlanRegistrationType :""
func (s *Service) SavePlanRegistrationType(ctx *models.Context, planRegistrationType *models.PlanRegistrationType) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	planRegistrationType.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPLANREGISTRATIONTYPE)
	planRegistrationType.Status = constants.PLANREGISTRATIONTYPESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 planRegistrationType.created")
	planRegistrationType.Created = created
	log.Println("b4 planRegistrationType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePlanRegistrationType(ctx, planRegistrationType)
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

//UpdatePlanRegistrationType : ""
func (s *Service) UpdatePlanRegistrationType(ctx *models.Context, planRegistrationType *models.PlanRegistrationType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePlanRegistrationType(ctx, planRegistrationType)
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

//EnablePlanRegistrationType : ""
func (s *Service) EnablePlanRegistrationType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePlanRegistrationType(ctx, UniqueID)
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

//DisablePlanRegistrationType : ""
func (s *Service) DisablePlanRegistrationType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePlanRegistrationType(ctx, UniqueID)
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

//DeletePlanRegistrationType : ""
func (s *Service) DeletePlanRegistrationType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePlanRegistrationType(ctx, UniqueID)
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

//GetSinglePlanRegistrationType :""
func (s *Service) GetSinglePlanRegistrationType(ctx *models.Context, UniqueID string) (*models.RefPlanRegistrationType, error) {
	planRegistrationType, err := s.Daos.GetSinglePlanRegistrationType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return planRegistrationType, nil
}

//FilterPlanRegistrationType :""
func (s *Service) FilterPlanRegistrationType(ctx *models.Context, planRegistrationTypefilter *models.PlanRegistrationTypeFilter, pagination *models.Pagination) (planRegistrationType []models.RefPlanRegistrationType, err error) {
	return s.Daos.FilterPlanRegistrationType(ctx, planRegistrationTypefilter, pagination)
}
