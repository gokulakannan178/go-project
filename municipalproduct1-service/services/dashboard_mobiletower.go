package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveMobileTower : ""
func (s *Service) SaveMobileTower(ctx *models.Context, mobileTower *models.PropertyMobileTower) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	mobileTower.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWER)
	mobileTower.Status = constants.PROPERTYMOBILETOWERSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMobileTower(ctx, mobileTower)
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

//GetSingleMobileTower :""
func (s *Service) GetSingleMobileTower(ctx *models.Context, UniqueID string) (*models.RefPropertyMobileTower, error) {
	tower, err := s.Daos.GetSingleMobileTower(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateMobileTower : ""
func (s *Service) UpdateMobileTower(ctx *models.Context, mobileTower *models.PropertyMobileTower) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateMobileTower(ctx, mobileTower)
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

// EnableMobileTower : ""
func (s *Service) EnableMobileTower(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableMobileTower(ctx, UniqueID)
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

//DisableMobileTower : ""
func (s *Service) DisableMobileTower(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableMobileTower(ctx, UniqueID)
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

//DeleteMobileTower : ""
func (s *Service) DeleteMobileTower(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteMobileTower(ctx, UniqueID)
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

// FilterMobileTower : ""
func (s *Service) FilterMobileTower(ctx *models.Context, filter *models.PropertyMobileTowerFilter, pagination *models.Pagination) ([]models.RefPropertyMobileTower, error) {
	return s.Daos.FilterMobileTower(ctx, filter, pagination)

}

// DashboardMobileTowerDemandAndCollection : ""
func (s *Service) DashboardMobileTowerDemandAndCollection(ctx *models.Context, filter *models.DashboardMobileTowerDemandAndCollectionFilter) (*models.DashboardMobileTowerDemandAndCollection, error) {
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
	res, err := s.Daos.DashBoardStatusWiseMobileTowerCollectionAndChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	data, err1 := s.Daos.DashboardMobileTowerDemandAndCollection(ctx, filter)
	if err1 != nil {
		return nil, err1
	}
	data.SAFCount = *res
	return data, nil
}
