package service

import (
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveWorkSchedule :""
func (s *Service) SaveWorkSchedule(ctx *models.Context, workSchedule *models.WorkSchedule) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	workSchedule.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONWORKSCHEDULE)
	workSchedule.Status = constants.WORKSCHEDULESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 WorkSchedule.created")
	workSchedule.Created = created
	log.Println("b4 WorkSchedule.created")
	daysvalue := s.Shared.BoolsToInt(workSchedule.Monday, workSchedule.Tuesday, workSchedule.Wednesday, workSchedule.Thursday, workSchedule.Friday, workSchedule.Saturday, workSchedule.Sunday)
	workSchedule.WorkingDaysinWeek = float64(daysvalue)
	workSchedule.WorkingHoursinWeek = workSchedule.WorkingHoursinDay * workSchedule.WorkingDaysinWeek
	fmt.Println("WorkingHoursinDay", workSchedule.WorkingHoursinDay)
	fmt.Println("WorkingHoursinWeek", workSchedule.WorkingHoursinWeek)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveWorkSchedule(ctx, workSchedule)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//GetSingleWorkSchedule :""
func (s *Service) GetSingleWorkSchedule(ctx *models.Context, UniqueID string) (*models.RefWorkSchedule, error) {
	workSchedule, err := s.Daos.GetSingleWorkSchedule(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return workSchedule, nil
}

//UpdateWorkSchedule : ""
func (s *Service) UpdateWorkSchedule(ctx *models.Context, workSchedule *models.WorkSchedule) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateWorkSchedule(ctx, workSchedule)
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

//EnableWorkSchedule : ""
func (s *Service) EnableWorkSchedule(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableWorkSchedule(ctx, UniqueID)
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

//DisableWorkSchedule : ""
func (s *Service) DisableWorkSchedule(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableWorkSchedule(ctx, UniqueID)
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

//DeleteWorkSchedule : ""
func (s *Service) DeleteWorkSchedule(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteWorkSchedule(ctx, UniqueID)
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

//FilterWorkSchedule :""
func (s *Service) FilterWorkSchedule(ctx *models.Context, filter *models.WorkScheduleFilter, pagination *models.Pagination) (WorkSchedule []models.RefWorkSchedule, err error) {
	return s.Daos.FilterWorkSchedule(ctx, filter, pagination)
}
