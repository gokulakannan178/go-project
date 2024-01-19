package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveShoprentPaymentModeChange : ""
func (s *Service) SaveShoprentPaymentModeChange(ctx *models.Context, request *models.ShoprentPaymentModeChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldPaymentData, err := s.Daos.GetSingleShoprentPaymentWithTxtID(ctx, request.TnxID)
		if err != nil {
			return errors.New("Error in geting old Property Payment" + err.Error())
		}
		if oldPaymentData == nil {
			return errors.New("Property Payment Not Found")
		}
		// propertyOwner, err := s.GetSingleShoprentPaymentTxtID(ctx, request.TnxID)
		// if err != nil {
		// 	return errors.New("Error in geting old Property Payment" + err.Error())
		// }
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSHOPRENTPAYMENTMODECHANGE)
		request.Previous = oldPaymentData.Details
		request.ReciptNo = oldPaymentData.ReciptNo
		request.ReciptDate = oldPaymentData.CompletionDate
		request.Requested.On = &t

		// if len(propertyOwner.Basic.Owners) > 0 {
		// 	request.OwnerName = propertyOwner.Basic.Owners[0].Name

		// }

		// if len(propertyOwner.Basic.Owners) > 0 {
		// 	request.Mobile = propertyOwner.Basic.Owners[0].Mobile

		// }

		request.Status = constants.SHOPRENTPAYMENTMODECHANGESTATUSPENDING

		err = s.Daos.SaveShoprentPaymentModeChange(ctx, request)
		if err != nil {
			return errors.New("Error in saving property payment mode change" + err.Error())
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

// GetSingleShoprentPaymentModeChange :""
func (s *Service) GetSingleShoprentPaymentModeChange(ctx *models.Context, UniqueID string) (*models.RefShoprentPaymentModeChange, error) {
	ppmc, err := s.Daos.GetSingleShoprentPaymentModeChange(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ppmc, nil
}

// AcceptShoprentPaymentModeChange : ""
func (s *Service) AcceptShoprentPaymentModeChange(ctx *models.Context, req *models.AcceptShoprentPaymentModeChange) error {

	res, err := s.Daos.GetSingleShoprentPaymentModeChange(ctx, req.UniqueID)
	if err != nil {
		return errors.New("Error in geting old Property Payment" + err.Error())

	}
	if res == nil {
		return errors.New("property payment mode change is nil")
	}
	oldPaymentData, err := s.Daos.GetSingleShoprentPaymentWithTxtID(ctx, res.TnxID)
	if err != nil {
		return errors.New("Error in getting in property payment mode change" + err.Error())
	}

	if oldPaymentData == nil {
		return errors.New("property payment Not Found")
	}
	oldPaymentData.Details = res.New

	err = s.Daos.UpdateShoprentPayments(ctx, oldPaymentData)
	if err != nil {
		return errors.New("Error in upating in Property Payment Mode Change Request" + err.Error())
	}
	err = s.Daos.AcceptShoprentPaymentModeChange(ctx, req)
	if err != nil {
		return errors.New("Error in upating in Property Payment Mode Change Request" + err.Error())
	}
	return nil
}

// RejectBasicTradeLicenseUpdate : ""
func (s *Service) RejectShoprentPaymentModeChange(ctx *models.Context, req *models.RejectShoprentPaymentModeChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		req.On = &t

		err := s.Daos.RejectShoprentPaymentModeChange(ctx, req)
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

// FilterReassessmentRequest : ""
func (s *Service) FilterShoprentPaymentModeChange(ctx *models.Context, filter *models.ShoprentPaymentModeChangeFilter, pagination *models.Pagination) ([]models.RefShoprentPaymentModeChange, error) {
	return s.Daos.FilterShoprentPaymentModeChange(ctx, filter, pagination)

}

// UpdateShoprentPaymentModeChangePropertyID : ""
func (s *Service) UpdateShoprentPaymentModeChangePropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range uniqueIds.UniqueIDs {
			resProperty, err := s.GetSingleProperty(ctx, v)
			if err != nil {
				return errors.New("Not able to get property - " + err.Error())
			}

			uniqueIds.UniqueID = resProperty.OldUniqueID
			uniqueIds.OldUniqueID = resProperty.OldUniqueID
			uniqueIds.NewUniqueID = resProperty.NewUniqueID
			err = s.Daos.UpdateShoprentPaymentModeChangePropertyID(ctx, uniqueIds)
			if err != nil {
				return err
			}
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
