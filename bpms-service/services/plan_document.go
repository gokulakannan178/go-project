package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePlanDocument :""
func (s *Service) SavePlanDocument(ctx *models.Context, planDocument *models.PlanDocument) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	planDocument.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPLANDOCUMENT)
	planDocument.Status = constants.PLANDOCUMENTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 planDocument.created")
	planDocument.Created = created
	planDocument.IssuedDate = &t
	log.Println("b4 planDocument.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.UpsertPlanDocument(ctx, planDocument)
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

//UpdatePlanDocument : ""
func (s *Service) UpdatePlanDocument(ctx *models.Context, planDocument *models.PlanDocument) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePlanDocument(ctx, planDocument)
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

//EnablePlanDocument : ""
func (s *Service) EnablePlanDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePlanDocument(ctx, UniqueID)
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

//DisablePlanDocument : ""
func (s *Service) DisablePlanDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePlanDocument(ctx, UniqueID)
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

//DeletePlanDocument : ""
func (s *Service) DeletePlanDocument(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePlanDocument(ctx, UniqueID)
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

//GetSinglePlanDocument :""
func (s *Service) GetSinglePlanDocument(ctx *models.Context, UniqueID string) (*models.RefPlanDocument, error) {
	planDocument, err := s.Daos.GetSinglePlanDocument(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return planDocument, nil
}

//FilterPlanDocument :""
func (s *Service) FilterPlanDocument(ctx *models.Context, planDocumentfilter *models.PlanDocumentFilter, pagination *models.Pagination) (planDocument []models.RefPlanDocument, err error) {
	return s.Daos.FilterPlanDocument(ctx, planDocumentfilter, pagination)
}

//GetPendingDocuments : ""
func (s *Service) GetPendingDocuments(ctx *models.Context, planDocumentfilter *models.GetPendingPlanDocumentFilter, pagination *models.Pagination) ([]models.RefPlanDocument, error) {
	return s.Daos.GetPendingDocuments(ctx, planDocumentfilter, pagination)
}
