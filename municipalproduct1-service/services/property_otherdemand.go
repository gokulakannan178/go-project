package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyOtherDemand : ""
func (s *Service) SavePropertyOtherDemand(ctx *models.Context, propertyotherdemand *models.PropertyOtherDemand) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	propertyotherdemand.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYOTHERDEMAND)
	propertyotherdemand.Status = constants.PROPERTYOTHERDEMANDSTATUSACTIVE
	propertyotherdemand.PaymentStatus = constants.PROPERTYOTHERDEMANDPAYMENTSTATUSNOTPAID
	t := time.Now()
	Created := new(models.CreatedV2)
	Created.On = &t
	//propertyotherdemand.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePropertyOtherDemand(ctx, propertyotherdemand)
		if dberr != nil {
			return dberr
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

//GetSinglePropertyOtherDemand :""
func (s *Service) GetSinglePropertyOtherDemand(ctx *models.Context, UniqueID string) (*models.RefPropertyOtherDemand, error) {
	propertyotherdemand, err := s.Daos.GetSinglePropertyOtherDemand(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return propertyotherdemand, nil
}

// UpdatePropertyOtherDemand : ""
func (s *Service) UpdatePropertyOtherDemand(ctx *models.Context, propertyotherdemand *models.PropertyOtherDemand) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyOtherDemand(ctx, propertyotherdemand)
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

// EnablePropertyOtherDemand : ""
func (s *Service) EnablePropertyOtherDemand(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyOtherDemand(ctx, UniqueID)
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

//DisablePropertyOtherDemand : ""
func (s *Service) DisablePropertyOtherDemand(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyOtherDemand(ctx, UniqueID)
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

//DeletePropertyOtherDemand : ""
func (s *Service) DeletePropertyOtherDemand(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePropertyOtherDemand(ctx, UniqueID)
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

// UpdatePropertyOtherDemandStatus : ""
func (s *Service) UpdatePropertyOtherDemandStatus(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyOtherDemandStatus(ctx, UniqueID)
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

// FilterPropertyOtherDemand : ""
func (s *Service) FilterPropertyOtherDemand(ctx *models.Context, filter *models.PropertyOtherDemandFilter, pagination *models.Pagination) ([]models.RefPropertyOtherDemand, error) {
	return s.Daos.FilterPropertyOtherDemand(ctx, filter, pagination)

}
