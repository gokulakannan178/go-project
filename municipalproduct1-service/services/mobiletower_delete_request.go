package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateMobileTowerDeleteRequest : ""
func (s *Service) BasicUpdateMobileTowerDeleteRequest(ctx *models.Context, request *models.MobileTowerDeleteRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldMobileTowerData, err := s.Daos.GetSingleMobileTower(ctx, request.MobileTowerID)
		if err != nil {
			return errors.New("error in geting old property" + err.Error())
		}
		if oldMobileTowerData == nil {
			return errors.New("trade license Not Found")
		}
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERDELETEREQUEST)
		request.Status = constants.MOBILETOWERDELETEREQUESTSTATUSINIT
		request.Requester.On = &t
		request.Requester.Scenario = "MobileTower Requested for Delete"
		err = s.Daos.SaveMobileTowerDeleteRequest(ctx, request)
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

// GetSingleMobileTowerDeleteRequest : ""
func (s *Service) GetSingleMobileTowerDeleteRequest(ctx *models.Context, UniqueID string) (*models.RefMobileTowerDeleteRequest, error) {
	pmr, err := s.Daos.GetSingleMobileTowerDeleteRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return pmr, nil
}

// AcceptMobileTowerDeleteRequestUpdate : ""
func (s *Service) AcceptMobileTowerDeleteRequestUpdate(ctx *models.Context, accept *models.AcceptMobileTowerDeleteRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.AcceptMobileTowerDeleteRequestUpdate(ctx, accept)
		if err != nil {
			return err
		}
		var resMobileTowerDelete *models.RefMobileTowerDeleteRequest
		resMobileTowerDelete, err = s.GetSingleMobileTowerDeleteRequest(ctx, accept.UniqueID)
		if err != nil {
			return errors.New("error in getting the property delete request" + err.Error())
		}
		err = s.Daos.UpdateMobileTowerStatusDeletedV2(ctx, resMobileTowerDelete.MobileTowerID)
		// err = s.Daos.UpdateMobileTowerEndDate(ctx, resMobileTowerDelete.MobileTowerID)
		if err != nil {
			return errors.New("error in updating property end date" + err.Error())
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

// RejectMobileTowerDeleteRequestUpdate : ""
func (s *Service) RejectMobileTowerDeleteRequestUpdate(ctx *models.Context, req *models.RejectMobileTowerDeleteRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectMobileTowerDeleteRequestUpdate(ctx, req)
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

// FilterMobileTowerDeleteRequest : ""
func (s *Service) FilterMobileTowerDeleteRequest(ctx *models.Context, filter *models.MobileTowerDeleteRequestFilter, pagination *models.Pagination) ([]models.RefMobileTowerDeleteRequest, error) {
	return s.Daos.FilterMobileTowerDeleteRequest(ctx, filter, pagination)
}
