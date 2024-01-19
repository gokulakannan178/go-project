package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateSolidWasteReassessmentRequest : ""
func (s *Service) BasicUpdateSolidWasteReassessmentRequest(ctx *models.Context, request *models.SolidWasteReassessmentRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldSolidWasteData, err := s.Daos.GetSingleSolidWasteUserCharge(ctx, request.SolidWasteID)
		if err != nil {
			return errors.New("Error in geting old shoprent" + err.Error())
		}
		if oldSolidWasteData == nil {
			return errors.New("shoprent Not Found")
		}

		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSOLIDWASTEREASSESSMENTREQUEST)
		request.Previous.SolidWasteUserCharge = oldSolidWasteData.SolidWasteUserCharge
		request.Previous.Ref.Address = oldSolidWasteData.Ref.Address
		request.Previous.SolidWasteUserCharge.OwnerName = oldSolidWasteData.OwnerName
		request.Status = constants.SOLIDWASTEREASSESSMENTREQUESTSTATUSINIT
		request.Requester.On = &t

		err = s.Daos.SaveSolidWasteReassessmentRequestUpdate(ctx, request)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// GetSingleSolidWasteReassessmentRequest :""
func (s *Service) GetSingleSolidWasteReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefSolidWasteReassessmentRequest, error) {
	res, err := s.Daos.GetSingleSolidWasteReassessmentRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AcceptSolidWasteReassessmentRequestUpdate : ""
func (s *Service) AcceptSolidWasteReassessmentRequestUpdate(ctx *models.Context, req *models.AcceptSolidWasteReassessmentRequestUpdate) error {
	fmt.Println("shoprent accept service")

	res, err := s.Daos.GetSingleSolidWasteReassessmentRequest(ctx, req.UniqueID)
	if err != nil {
		return errors.New("Error in getting in solidWasteReassessment Request" + err.Error())
	}

	oldSolidWasteData, err := s.Daos.GetSingleSolidWasteUserCharge(ctx, res.SolidWasteID)
	if err != nil {
		return errors.New("Error in getting in old solidwaste" + err.Error())
	}

	if oldSolidWasteData == nil {
		return errors.New("solidwaste Not Found")
	}

	res.New.SolidWasteUserCharge.UniqueID = res.SolidWasteID
	err = s.Daos.UpdateSolidWasteUserCharge(ctx, &res.New.SolidWasteUserCharge)
	if err != nil {
		return errors.New("Error in updating in shoprent" + err.Error())
	}
	err = s.Daos.BasicSolidWasteReassessmentRequestUpdateToPayments(ctx, res)
	if err != nil {
		return errors.New("Error in upating in shoprentReassessment Request Payments" + err.Error())
	}
	err = s.Daos.AcceptSolidWasteReassessmentRequestUpdate(ctx, req)
	if err != nil {
		return errors.New("Error in upating in shoprentReassessment Request" + err.Error())
	}

	return nil
}

// RejectSolidWasteReassessmentRequestUpdate : ""
func (s *Service) RejectSolidWasteReassessmentRequestUpdate(ctx *models.Context, req *models.RejectSolidWasteReassessmentRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectSolidWasteReassessmentRequestUpdate(ctx, req)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// FilterSolidWasteReassessmentRequest : ""
func (s *Service) FilterSolidWasteReassessmentRequest(ctx *models.Context, filter *models.SolidWasteReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefSolidWasteReassessmentRequest, error) {
	return s.Daos.FilterSolidWasteReassessmentRequest(ctx, filter, pagination)

}
