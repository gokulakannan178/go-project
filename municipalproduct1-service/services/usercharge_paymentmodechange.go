package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveUserchargePaymentModeChange : ""
func (s *Service) SaveUserchargePaymentModeChange(ctx *models.Context, request *models.UserchargePaymentModeChangeRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldPaymentData, err := s.Daos.GetSingleUserChargePaymentWithTxtID(ctx, request.TnxID)
		if err != nil {
			return errors.New("Error in geting old Property Payment" + err.Error())
		}
		if oldPaymentData == nil {
			return errors.New("Property Payment Not Found")
		}
		// propertyOwner, err := s.GetSingleUserchargePaymentTxtID(ctx, request.TnxID)
		// if err != nil {
		// 	return errors.New("Error in geting old Property Payment" + err.Error())
		// }
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERCHARGEPAYMENTMODECHANGE)
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

		request.Status = constants.USERCHARGEPAYMENTMODECHANGESTATUSPENDING

		err = s.Daos.SaveUserchargePaymentModeChange(ctx, request)
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

// GetSingleUserchargePaymentModeChange :""
func (s *Service) GetSingleUserchargePaymentModeChange(ctx *models.Context, UniqueID string) (*models.RefUserchargePaymentModeChangeRequest, error) {
	ppmc, err := s.Daos.GetSingleUserchargePaymentModeChange(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ppmc, nil
}

// AcceptUserchargePaymentModeChange : ""
func (s *Service) AcceptUserchargePaymentModeChange(ctx *models.Context, req *models.AcceptUserchargePaymentModeChangeRequest) error {

	res, err := s.Daos.GetSingleUserchargePaymentModeChange(ctx, req.UniqueID)
	if err != nil {
		return errors.New("Error in geting old Property Payment" + err.Error())

	}
	if res == nil {
		return errors.New("property payment mode change is nil")
	}
	oldPaymentData, err := s.Daos.GetSingleUserChargePaymentWithTxtID(ctx, res.TnxID)
	if err != nil {
		return errors.New("Error in getting in property payment mode change" + err.Error())
	}

	if oldPaymentData == nil {
		return errors.New("property payment Not Found")
	}
	oldPaymentData.Details = res.New

	err = s.Daos.UpdateUserchargePayments(ctx, oldPaymentData)
	if err != nil {
		return errors.New("Error in upating in Property Payment Mode Change Request" + err.Error())
	}
	err = s.Daos.AcceptUserchargePaymentModeChange(ctx, req)
	if err != nil {
		return errors.New("Error in upating in Property Payment Mode Change Request" + err.Error())
	}
	return nil
}

// RejectBasicTradeLicenseUpdate : ""
func (s *Service) RejectUserchargePaymentModeChange(ctx *models.Context, req *models.RejectUserchargePaymentModeChangeRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		req.On = &t

		err := s.Daos.RejectUserchargePaymentModeChange(ctx, req)
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
func (s *Service) FilterUserchargePaymentModeChange(ctx *models.Context, filter *models.UserchargePaymentModeChangeRequestFilter, pagination *models.Pagination) ([]models.RefUserchargePaymentModeChangeRequest, error) {
	return s.Daos.FilterUserchargePaymentModeChange(ctx, filter, pagination)

}

// UpdateUserchargePaymentModeChangePropertyID : ""
func (s *Service) UpdateUserchargePaymentModeChangePropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
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
			err = s.Daos.UpdateUserchargePaymentModeChangePropertyID(ctx, uniqueIds)
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
