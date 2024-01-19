package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePlanReqDocument :""
func (s *Service) SavePlanReqDocument(ctx *models.Context, planReqDocument *models.PlanReqDocument) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	planReqDocument.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPLANREQDOCUMENT)
	planReqDocument.Status = constants.PLANREQDOCUMENTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 planReqDocument.created")
	planReqDocument.Created = created
	log.Println("b4 planReqDocument.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePlanReqDocument(ctx, planReqDocument)
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

//UpdatePlanReqDocument : ""
func (s *Service) UpdatePlanReqDocument(ctx *models.Context, planReqDocument *models.PlanReqDocument) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePlanReqDocument(ctx, planReqDocument)
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

//EnablePlanReqDocument : ""
func (s *Service) EnablePlanReqDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePlanReqDocument(ctx, UniqueID)
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

//DisablePlanReqDocument : ""
func (s *Service) DisablePlanReqDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePlanReqDocument(ctx, UniqueID)
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

//DeletePlanReqDocument : ""
func (s *Service) DeletePlanReqDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePlanReqDocument(ctx, UniqueID)
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

//GetSinglePlanReqDocument :""
func (s *Service) GetSinglePlanReqDocument(ctx *models.Context, UniqueID string) (*models.RefPlanReqDocument, error) {
	planReqDocument, err := s.Daos.GetSinglePlanReqDocument(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return planReqDocument, nil
}

//FilterPlanReqDocument :""
func (s *Service) FilterPlanReqDocument(ctx *models.Context, planReqDocumentfilter *models.PlanReqDocumentFilter, pagination *models.Pagination) (planReqDocument []models.RefPlanReqDocument, err error) {
	return s.Daos.FilterPlanReqDocument(ctx, planReqDocumentfilter, pagination)
}
