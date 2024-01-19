package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveMobileTowerDayWise : ""
func (s *Service) SaveMobileTowerDayWise(ctx *models.Context, mobileTower *models.MobiletowerDashboardDayWise) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	mobileTower.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDMOBILETOWERDAYWISE)
	mobileTower.Status = constants.DASHBOARDMOBILETOWERDAYWISESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMobileTowerDayWise(ctx, mobileTower)
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

//GetSingleMobileTowerDayWise :""
func (s *Service) GetSingleMobileTowerDayWise(ctx *models.Context, UniqueID string) (*models.RefMobileTowerDayWiseDashboard, error) {
	tower, err := s.Daos.GetSingleMobileTowerDayWise(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateMobileTowerDayWise : ""
func (s *Service) UpdateMobileTowerDayWise(ctx *models.Context, mobileTower *models.MobiletowerDashboardDayWise) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateMobileTowerDayWise(ctx, mobileTower)
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

// EnableMobileTowerDayWise : ""
func (s *Service) EnableMobileTowerDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableMobileTowerDayWise(ctx, UniqueID)
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

//DisableMobileTowerDayWise : ""
func (s *Service) DisableMobileTowerDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableMobileTowerDayWise(ctx, UniqueID)
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

//DeleteMobileTowerDayWise : ""
func (s *Service) DeleteMobileTowerDayWise(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteMobileTowerDayWise(ctx, UniqueID)
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

// FilterMobileTowerDayWise : ""
func (s *Service) FilterMobileTowerDayWise(ctx *models.Context, filter *models.MobileTowerDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefMobileTowerDayWiseDashboard, error) {
	return s.Daos.FilterMobileTowerDayWise(ctx, filter, pagination)

}
