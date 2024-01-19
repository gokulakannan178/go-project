package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//GetSingleCRF :""
func (s *Service) GetSingleCRF(ctx *models.Context, UniqueID string) (*models.RefCRF, error) {
	crf, err := s.Daos.GetSingleCRF(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return crf, nil
}

//FilterCRF :""
func (s *Service) FilterCRF(ctx *models.Context, crffilter *models.CRFFilter, pagination *models.Pagination) (crf []models.RefCRF, err error) {
	return s.Daos.FilterCRF(ctx, crffilter, pagination)
}

//CRF Inspection APIS

//StartPlanCRFInspection : ""
func (s *Service) StartPlanCRFInspection(ctx *models.Context, crf *models.PlanCRFStartInspectionReqPayload) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		checklist, err := s.Daos.GetChecklistsOfDept(ctx, crf.DeptID)
		if err != nil {
			return errors.New("Error in geting checklist - " + err.Error())
		}
		if checklist == nil {
			return errors.New("Checklist Nil- ")
		}
		var crfis []models.CRFInspection
		for _, v := range checklist {
			var crfi models.CRFInspection
			crfi.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCRFINSPECTION)
			crfi.Status = constants.PLANCRFSTATUSPENDING
			crfi.PlanID = crf.PlanID
			crfi.PlanRegTypeID = crf.PlanRegTypeID
			crfi.DeptID = crf.DeptID
			crfi.CheckListID = v.UniqueID
			crfis = append(crfis, crfi)
		}
		if len(crfis) > 0 {
			if err := s.Daos.SaveCRFInspection(ctx, crfis); err != nil {
				return errors.New("Error in creating inspection - " + err.Error())
			}
		}
		if err := s.CommonCRFFlowUpdate(ctx, crf.Flow.CRFID, crf.Flow.Scenario, &crf.Flow.PlanTimeline); err != nil {
			return err
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

//GetCRFInspectionOfPlan : ""
func (s *Service) GetCRFInspectionOfPlan(ctx *models.Context, planID, deptID string) ([]models.RefCRFInspection, error) {
	return s.Daos.GetCRFInspectionOfPlan(ctx, planID, deptID)
}

//SubmitCRFInspection : ""
func (s *Service) SubmitCRFInspection(ctx *models.Context, crfsInspection *models.CRFInspection) error {
	t := time.Now()
	crfsInspection.Time = &t
	crfsInspection.Status = constants.PLANCRFSTATUSCOMPLETED
	return s.Daos.SubmitCRFInspection(ctx, crfsInspection)
}

//EndPlanCRFInspection : ""
func (s *Service) EndPlanCRFInspection(ctx *models.Context, crf *models.PlanCRFEndInspectionReqPayload) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if err := s.CommonCRFFlowUpdate(ctx, crf.Flow.CRFID, crf.Flow.Scenario, &crf.Flow.PlanTimeline); err != nil {
			return err
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
