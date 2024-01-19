package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdatePropertyDeleteRequest : ""
func (s *Service) BasicUpdatePropertyDeleteRequest(ctx *models.Context, request *models.PropertyDeleteRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldPropertyData, err := s.Daos.GetSingleProperty(ctx, request.PropertyID)
		if err != nil {
			return errors.New("error in geting old property" + err.Error())
		}
		if oldPropertyData == nil {
			return errors.New("property Not Found")
		}
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYDELETEREQUEST)
		request.Status = constants.PROPERTYDELETEREQUESTSTATUSINIT
		request.Requester.On = &t
		request.Requester.Scenario = "Property Requested for Delete"
		err = s.Daos.SavePropertyDeleteRequest(ctx, request)
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

// GetSinglePropertyDeleteRequest : ""
func (s *Service) GetSinglePropertyDeleteRequest(ctx *models.Context, UniqueID string) (*models.RefPropertyDeleteRequest, error) {
	pmr, err := s.Daos.GetSinglePropertyDeleteRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return pmr, nil
}

// AcceptPropertyDeleteRequestUpdate : ""
func (s *Service) AcceptPropertyDeleteRequestUpdate(ctx *models.Context, accept *models.AcceptPropertyDeleteRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.AcceptPropertyDeleteRequestUpdate(ctx, accept)
		if err != nil {
			return err
		}
		var resPropertyDelete *models.RefPropertyDeleteRequest
		resPropertyDelete, err = s.GetSinglePropertyDeleteRequest(ctx, accept.UniqueID)
		if err != nil {
			return errors.New("error in getting the property delete request" + err.Error())
		}
		err = s.Daos.UpdatePropertyStatusDeletedV2(ctx, resPropertyDelete.PropertyID)
		// err = s.Daos.UpdatePropertyEndDate(ctx, resPropertyDelete.PropertyID)
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

// RejectPropertyDeleteRequestUpdate : ""
func (s *Service) RejectPropertyDeleteRequestUpdate(ctx *models.Context, req *models.RejectPropertyDeleteRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectPropertyDeleteRequestUpdate(ctx, req)
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

// FilterPropertyDeleteRequest : ""
func (s *Service) FilterPropertyDeleteRequest(ctx *models.Context, filter *models.PropertyDeleteRequestFilter, pagination *models.Pagination) ([]models.RefPropertyDeleteRequest, error) {
	return s.Daos.FilterPropertyDeleteRequest(ctx, filter, pagination)
}
