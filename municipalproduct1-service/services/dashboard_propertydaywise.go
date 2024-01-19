package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveDashBoardPropertyDayWiseV2 : ""
func (s *Service) SaveDashBoardPropertyDayWiseV2(ctx *models.Context, property *models.PropertyDashboardDayWise) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	property.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDPROPERTYDAYWISE)
	property.Status = constants.DASHBOARDPROPERTYDAYWISESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDashBoardPropertyDayWiseV2(ctx, property)
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

//GetSingleDashBoardPropertyDayWiseV2 :""
func (s *Service) GetSingleDashBoardPropertyDayWiseV2(ctx *models.Context, UniqueID string) (*models.RefPropertyDashboardDayWise, error) {
	tower, err := s.Daos.GetSingleDashBoardPropertyDayWiseV2(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateDashBoardPropertyDayWiseV2 : ""
func (s *Service) UpdateDashBoardPropertyDayWiseV2(ctx *models.Context, property *models.PropertyDashboardDayWise) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDashBoardPropertyDayWiseV2(ctx, property)
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

// EnableDashBoardPropertyDayWiseV2 : ""
func (s *Service) EnableDashBoardPropertyDayWiseV2(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDashBoardPropertyDayWiseV2(ctx, UniqueID)
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

//DisableDashBoardPropertyDayWiseV2 : ""
func (s *Service) DisableDashBoardPropertyDayWiseV2(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDashBoardPropertyDayWiseV2(ctx, UniqueID)
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

//DeleteDashBoardPropertyDayWiseV2 : ""
func (s *Service) DeleteDashBoardPropertyDayWiseV2(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDashBoardPropertyDayWiseV2(ctx, UniqueID)
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

// FilterDashBoardPropertyDayWiseV2 : ""
func (s *Service) FilterDashBoardPropertyDayWiseV2(ctx *models.Context, filter *models.PropertyDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefPropertyDashboardDayWise, error) {
	return s.Daos.FilterDashBoardPropertyDayWiseV2(ctx, filter, pagination)

}
