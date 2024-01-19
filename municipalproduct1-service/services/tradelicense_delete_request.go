package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateTradeLicenseDeleteRequest : ""
func (s *Service) BasicUpdateTradeLicenseDeleteRequest(ctx *models.Context, request *models.TradeLicenseDeleteRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldTradeLicenseData, err := s.Daos.GetSingleTradeLicense(ctx, request.TradeLicenseID)
		if err != nil {
			return errors.New("error in geting old property" + err.Error())
		}
		if oldTradeLicenseData == nil {
			return errors.New("trade license Not Found")
		}
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSEDELETEREQUEST)
		request.Status = constants.TRADELICENSEDELETEREQUESTSTATUSINIT
		request.Requester.On = &t
		request.Requester.Scenario = "TradeLicense Requested for Delete"
		err = s.Daos.SaveTradeLicenseDeleteRequest(ctx, request)
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

// GetSingleTradeLicenseDeleteRequest : ""
func (s *Service) GetSingleTradeLicenseDeleteRequest(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseDeleteRequest, error) {
	pmr, err := s.Daos.GetSingleTradeLicenseDeleteRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return pmr, nil
}

// AcceptTradeLicenseDeleteRequestUpdate : ""
func (s *Service) AcceptTradeLicenseDeleteRequestUpdate(ctx *models.Context, accept *models.AcceptTradeLicenseDeleteRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.AcceptTradeLicenseDeleteRequestUpdate(ctx, accept)
		if err != nil {
			return err
		}
		var resTradeLicenseDelete *models.RefTradeLicenseDeleteRequest
		resTradeLicenseDelete, err = s.GetSingleTradeLicenseDeleteRequest(ctx, accept.UniqueID)
		if err != nil {
			return errors.New("error in getting the property delete request" + err.Error())
		}
		err = s.Daos.UpdateTradeLicenseStatusDeletedV2(ctx, resTradeLicenseDelete.TradeLicenseID)
		// err = s.Daos.UpdateTradeLicenseEndDate(ctx, resTradeLicenseDelete.TradeLicenseID)
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

// RejectTradeLicenseDeleteRequestUpdate : ""
func (s *Service) RejectTradeLicenseDeleteRequestUpdate(ctx *models.Context, req *models.RejectTradeLicenseDeleteRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectTradeLicenseDeleteRequestUpdate(ctx, req)
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

// FilterTradeLicenseDeleteRequest : ""
func (s *Service) FilterTradeLicenseDeleteRequest(ctx *models.Context, filter *models.TradeLicenseDeleteRequestFilter, pagination *models.Pagination) ([]models.RefTradeLicenseDeleteRequest, error) {
	return s.Daos.FilterTradeLicenseDeleteRequest(ctx, filter, pagination)
}
