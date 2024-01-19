package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveUserChargeDayWiseDashboard : ""
func (s *Service) SaveUserChargeDayWiseDashboard(ctx *models.Context, userCharge *models.UserChargeDashboardDayWise) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	userCharge.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDUSERCHARGEDAYWISE)
	userCharge.Status = constants.DASHBOARDUSERCHARGEDAYWISESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUserChargeDayWiseDashboard(ctx, userCharge)
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

//GetSingleUserChargeDayWiseDashboard :""
func (s *Service) GetSingleUserChargeDayWiseDashboard(ctx *models.Context, UniqueID string) (*models.RefUserChargeDashboardDayWise, error) {
	tower, err := s.Daos.GetSingleUserChargeDayWiseDashboard(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateUserChargeDayWiseDashboard : ""
func (s *Service) UpdateUserChargeDayWiseDashboard(ctx *models.Context, userCharge *models.UserChargeDashboardDayWise) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserChargeDayWiseDashboard(ctx, userCharge)
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

// EnableUserChargeDayWiseDashboard : ""
func (s *Service) EnableUserChargeDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUserChargeDayWiseDashboard(ctx, UniqueID)
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

//DisableUserChargeDayWiseDashboard : ""
func (s *Service) DisableUserChargeDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUserChargeDayWiseDashboard(ctx, UniqueID)
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

//DeleteUserChargeDayWiseDashboard : ""
func (s *Service) DeleteUserChargeDayWiseDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUserChargeDayWiseDashboard(ctx, UniqueID)
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

// FilterUserChargeDayWiseDashboard : ""
func (s *Service) FilterUserChargeDayWiseDashboard(ctx *models.Context, filter *models.UserChargeDashboardDayWiseFilter, pagination *models.Pagination) ([]models.RefUserChargeDashboardDayWise, error) {
	return s.Daos.FilterUserChargeDayWiseDashboard(ctx, filter, pagination)

}
