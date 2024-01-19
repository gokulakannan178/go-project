package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveShopRentDashboardDayWise : ""
func (s *Service) SaveShopRentDashboardDayWise(ctx *models.Context, shopRent *models.ShopRentDashboardDayWise) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	shopRent.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDSHOPRENTDAYWISE)
	shopRent.Status = constants.DASHBOARDSHOPRENTDAYWISESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveShopRentDashboardDayWise(ctx, shopRent)
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

//GetSingleShopRentDashboardDayWise :""
func (s *Service) GetSingleShopRentDashboardDayWise(ctx *models.Context, UniqueID string) (*models.RefShopRentDashboardDayWise, error) {
	tower, err := s.Daos.GetSingleShopRentDashboardDayWise(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateShopRentDashboardDayWise : ""
func (s *Service) UpdateShopRentDashboardDayWise(ctx *models.Context, shopRent *models.ShopRentDashboardDayWise) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateShopRentDashboardDayWise(ctx, shopRent)
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

// EnableShopRentDashboardDayWise : ""
func (s *Service) EnableShopRentDashboardDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableShopRentDashboardDayWise(ctx, UniqueID)
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

//DisableShopRentDashboardDayWise : ""
func (s *Service) DisableShopRentDashboardDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableShopRentDashboardDayWise(ctx, UniqueID)
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

//DeleteShopRentDashboardDayWise : ""
func (s *Service) DeleteShopRentDashboardDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteShopRentDashboardDayWise(ctx, UniqueID)
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

// FilterShopRentDashboardDayWise : ""
func (s *Service) FilterShopRentDashboardDayWise(ctx *models.Context, filter *models.ShopRentDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefShopRentDashboardDayWise, error) {
	return s.Daos.FilterShopRentDashboardDayWise(ctx, filter, pagination)

}
