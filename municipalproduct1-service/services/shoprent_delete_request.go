package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateShopRentDeleteRequest : ""
func (s *Service) BasicUpdateShopRentDeleteRequest(ctx *models.Context, request *models.ShopRentDeleteRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldShopRentData, err := s.Daos.GetSingleShopRent(ctx, request.ShopRentID)
		if err != nil {
			return errors.New("error in geting old property" + err.Error())
		}
		if oldShopRentData == nil {
			return errors.New("trade license Not Found")
		}
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSHOPRENTDELETEREQUEST)
		request.Status = constants.SHOPRENTDELETEREQUESTSTATUSINIT
		request.Requester.On = &t
		request.Requester.Scenario = "ShopRent Requested for Delete"
		err = s.Daos.SaveShopRentDeleteRequest(ctx, request)
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

// GetSingleShopRentDeleteRequest : ""
func (s *Service) GetSingleShopRentDeleteRequest(ctx *models.Context, UniqueID string) (*models.RefShopRentDeleteRequest, error) {
	pmr, err := s.Daos.GetSingleShopRentDeleteRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return pmr, nil
}

// AcceptShopRentDeleteRequestUpdate : ""
func (s *Service) AcceptShopRentDeleteRequestUpdate(ctx *models.Context, accept *models.AcceptShopRentDeleteRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.AcceptShopRentDeleteRequestUpdate(ctx, accept)
		if err != nil {
			return err
		}
		var resShopRentDelete *models.RefShopRentDeleteRequest
		resShopRentDelete, err = s.GetSingleShopRentDeleteRequest(ctx, accept.UniqueID)
		if err != nil {
			return errors.New("error in getting the property delete request" + err.Error())
		}
		err = s.Daos.UpdateShopRentStatusDeletedV2(ctx, resShopRentDelete.ShopRentID)
		// err = s.Daos.UpdateShopRentEndDate(ctx, resShopRentDelete.ShopRentID)
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

// RejectShopRentDeleteRequestUpdate : ""
func (s *Service) RejectShopRentDeleteRequestUpdate(ctx *models.Context, req *models.RejectShopRentDeleteRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectShopRentDeleteRequestUpdate(ctx, req)
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

// FilterShopRentDeleteRequest : ""
func (s *Service) FilterShopRentDeleteRequest(ctx *models.Context, filter *models.ShopRentDeleteRequestFilter, pagination *models.Pagination) ([]models.RefShopRentDeleteRequest, error) {
	return s.Daos.FilterShopRentDeleteRequest(ctx, filter, pagination)
}
