package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveTradeLicenseDashboardDayWise : ""
func (s *Service) SaveTradeLicenseDashboardDayWise(ctx *models.Context, tradeLicense *models.TradeLicenseDashboardDayWise) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	tradeLicense.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDTRADELICENSEDAYWISE)
	tradeLicense.Status = constants.DASHBOARDSHOPRENTDAYWISESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTradeLicenseDashboardDayWise(ctx, tradeLicense)
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

//GetSingleTradeLicenseDashboardDayWise :""
func (s *Service) GetSingleTradeLicenseDashboardDayWise(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseDashboardDayWise, error) {
	tower, err := s.Daos.GetSingleTradeLicenseDashboardDayWise(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateTradeLicenseDashboardDayWise : ""
func (s *Service) UpdateTradeLicenseDashboardDayWise(ctx *models.Context, tradeLicense *models.TradeLicenseDashboardDayWise) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTradeLicenseDashboardDayWise(ctx, tradeLicense)
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

// EnableTradeLicenseDashboardDayWise : ""
func (s *Service) EnableTradeLicenseDashboardDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTradeLicenseDashboardDayWise(ctx, UniqueID)
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

//DisableTradeLicenseDashboardDayWise : ""
func (s *Service) DisableTradeLicenseDashboardDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTradeLicenseDashboardDayWise(ctx, UniqueID)
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

//DeleteTradeLicenseDashboardDayWise : ""
func (s *Service) DeleteDashBoardTradeLicenseDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTradeLicenseDashboardDayWise(ctx, UniqueID)
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

// FilterTradeLicenseDashboardDayWise : ""
func (s *Service) FilterTradeLicenseDashboardDayWise(ctx *models.Context, filter *models.TradeLicenseDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefTradeLicenseDashboardDayWise, error) {
	return s.Daos.FilterTradeLicenseDashboardDayWise(ctx, filter, pagination)

}
