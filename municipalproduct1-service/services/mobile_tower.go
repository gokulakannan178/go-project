package services

import (
	"errors"
	"log"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyMobileTower : ""
func (s *Service) SavePropertyMobileTower(ctx *models.Context, mobile *models.PropertyMobileTower) error {
	log.Println("transaction start +++++++++++++")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	// mobile.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYMOBILETOWER)
	// mobile.Status = constants.PROPERTYMOBILETOWERSTATUSACTIVE
	// t := time.Now()
	// mobile.Created.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		s.PreSaveMobileTower(ctx, mobile)
		dberr := s.Daos.SavePropertyMobileTower(ctx, mobile)
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

//GetSinglePropertyMobileTower :""
func (s *Service) GetSinglePropertyMobileTower(ctx *models.Context, UniqueID string) (*models.RefPropertyMobileTower, error) {
	tower, err := s.Daos.GetSinglePropertyMobileTower(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdatePropertyMobileTower : ""
func (s *Service) UpdatePropertyMobileTower(ctx *models.Context, mobile *models.PropertyMobileTower) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyMobileTower(ctx, mobile)
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

// EnableMobileTowerTax : ""
func (s *Service) EnablePropertyMobileTower(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyMobileTower(ctx, UniqueID)
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

//DisablePropertyMobileTower : ""
func (s *Service) DisablePropertyMobileTower(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyMobileTower(ctx, UniqueID)
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

// DeletePropertyMobileTower : ""
func (s *Service) DeletePropertyMobileTower(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePropertyMobileTower(ctx, UniqueID)
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

// RejectPropertyMobileTower : ""
func (s *Service) RejectPropertyMobileTower(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectPropertyMobileTower(ctx, UniqueID)
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

// FilterPropertyMobileTower : ""
func (s *Service) FilterPropertyMobileTower(ctx *models.Context, filter *models.PropertyMobileTowerFilter, pagination *models.Pagination) ([]models.RefPropertyMobileTower, error) {
	return s.Daos.FilterPropertyMobileTower(ctx, filter, pagination)

}

// MobileTowerWithMobileNo :""
func (s *Service) MobileTowerWithMobileNo(ctx *models.Context, filter *models.MobileTowerWithMobileNoReq, pagination *models.Pagination) (property []models.MobileTowerWithMobileNoRes, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.MobileTowerWithMobileNo(ctx, filter, pagination)
}

// MobileTowerPenaltyUpdate : ""
func (s *Service) MobileTowerPenaltyUpdate(ctx *models.Context, mobile *models.MobileTowerPenaltyUpdate) error {
	err := s.Daos.MobileTowerPenaltyUpdate(ctx, mobile)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMobileTowerPropertyID : ""
func (s *Service) UpdateMobileTowerPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
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
			err = s.Daos.UpdateMobileTowerPropertyID(ctx, uniqueIds)
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
