package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveWaterBillDayWiseDashboard : ""
func (s *Service) SaveWaterBillDayWiseDashboard(ctx *models.Context, waterBill *models.WaterBillDashboardDayWise) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	waterBill.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDWATERBILLDAYWISE)
	waterBill.Status = constants.DASHBOARDWATERBILLDAYWISESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveWaterBillDayWiseDashboard(ctx, waterBill)
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

//GetSingleWaterBillDayWiseDashboard :""
func (s *Service) GetSingleWaterBillDayWiseDashboard(ctx *models.Context, UniqueID string) (*models.RefWaterBillDashboardDayWise, error) {
	tower, err := s.Daos.GetSingleWaterBillDayWiseDashboard(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateWaterBillDayWiseDashboard : ""
func (s *Service) UpdateWaterBillDayWiseDashboard(ctx *models.Context, waterBill *models.WaterBillDashboardDayWise) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateWaterBillDayWiseDashboard(ctx, waterBill)
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

// EnableWaterBillDayWiseDashboard : ""
func (s *Service) EnableWaterBillDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableWaterBillDayWiseDashboard(ctx, UniqueID)
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

//DisableWaterBillDayWiseDashboard : ""
func (s *Service) DisableWaterBillDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableWaterBillDayWiseDashboard(ctx, UniqueID)
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

//DeleteWaterBillDayWiseDashboard : ""
func (s *Service) DeleteWaterBillDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteWaterBillDayWiseDashboard(ctx, UniqueID)
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

// FilterWaterBillDayWiseDashboard : ""
func (s *Service) FilterWaterBillDayWiseDashboard(ctx *models.Context, filter *models.WaterBillDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefWaterBillDashboardDayWise, error) {
	return s.Daos.FilterWaterBillDayWiseDashboard(ctx, filter, pagination)

}
