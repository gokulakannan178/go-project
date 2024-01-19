package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveShopRentDashboard : ""
func (s *Service) SaveShopRentDashboard(ctx *models.Context, shopRent *models.ShopRentDashboard) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	shopRent.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDSHOPRENT)
	shopRent.Status = constants.DASHBOARDSHOPRENTSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveShopRentDashboard(ctx, shopRent)
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

//GetSingleShopRentDashboard :""
func (s *Service) GetSingleShopRentDashboard(ctx *models.Context, UniqueID string) (*models.RefShopRentDashboard, error) {
	tower, err := s.Daos.GetSingleShopRentDashboard(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateShopRentDashboard : ""
func (s *Service) UpdateShopRentDashboard(ctx *models.Context, shopRent *models.ShopRentDashboard) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateShopRentDashboard(ctx, shopRent)
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

// EnableShopRentDashboard : ""
func (s *Service) EnableShopRentDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableShopRentDashboard(ctx, UniqueID)
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

//DisableShopRentDashboard : ""
func (s *Service) DisableShopRentDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableShopRentDashboard(ctx, UniqueID)
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

//DeleteShopRentDashboard : ""
func (s *Service) DeleteShopRentDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteShopRentDashboard(ctx, UniqueID)
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

// FilterShopRentDashboard : ""
func (s *Service) FilterShopRentDashboard(ctx *models.Context, filter *models.ShopRentDashboardFilter, pagination *models.Pagination) ([]models.RefShopRentDashboard, error) {
	return s.Daos.FilterShopRentDashboard(ctx, filter, pagination)
}

// DashboardShopRentDemandAndCollection : ""
func (s *Service) DashboardShopRentDemandAndCollection(ctx *models.Context, filter *models.DashboardShopRentDemandAndCollectionFilter) (*models.DashboardShopRentDemandAndCollection, error) {
	//defer ctx.Session.EndSession(ctx.CTX)
	if filter != nil {
		t := time.Now()
		sdt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		edt := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		filter.TodayRange.From = &sdt
		filter.TodayRange.To = &edt
		tYesterday := t.AddDate(0, 0, -1)
		sdtyesterday := time.Date(tYesterday.Year(), tYesterday.Month(), tYesterday.Day(), 0, 0, 0, 0, tYesterday.Location())
		edtyesterday := time.Date(tYesterday.Year(), tYesterday.Month(), tYesterday.Day(), 59, 59, 59, 0, tYesterday.Location())
		filter.YesterdayRange.From = &sdtyesterday
		filter.YesterdayRange.To = &edtyesterday
	}
	res, err := s.Daos.DashBoardStatusWiseShopRentCollectionAndChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	data, err := s.Daos.DashboardShopRentDemandAndCollection(ctx, filter)
	if err != nil {
		return nil, err
	}
	data.SAFCount = *res

	return data, nil
}
