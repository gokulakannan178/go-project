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

func (s *Service) SaveHelperBeat(ctx *models.Context, helperbeat *models.HelperBeat) ([]models.HelperBeat, error) {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//HelperBeat.IsDefault = true
	//organisationConfig.Activestatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 HelperBeat.created")
	helperbeat.Created = &created
	log.Println("b4 organisation.created")
	var employee []models.HelperBeat
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		beatID, err := s.Daos.GetSingleBeatMaster(ctx, helperbeat.BeatID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		for _, v := range helperbeat.EmployeeUser {
			refEmployee, err := s.Daos.GetSingleGarbaggeCollector(ctx, v)
			if err != nil {
				fmt.Println(err)
				return err
			}
			if refEmployee != nil {
				helperbeat.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONHELPERBEAT)
				helperbeat.Status = "Active"
				helperbeat.Employee.Name = refEmployee.Name
				helperbeat.Employee.Id = refEmployee.UniqueID
				helperbeat.AssignDate = &t
				helperbeat.BeatID = beatID.UniqueID
				dberr := s.Daos.SaveHelperBeatWithUpsert(ctx, helperbeat)
				if dberr != nil {
					return errors.New("HelperBeat Error" + dberr.Error())
				}
				employee = append(employee, *helperbeat)
			}
		}
		err = s.Daos.RemoveInActiveHelperInBeat(ctx, helperbeat.BeatID, helperbeat.EmployeeUser)
		if err != nil {
			return err
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return nil, errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return nil, err
	}
	return employee, nil
}

func (s *Service) UpdateHelperBeat(ctx *models.Context, helperbeat *models.HelperBeat) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateHelperBeat(ctx, helperbeat)
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

func (s *Service) EnableHelperBeat(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableHelperBeat(ctx, UniqueID)
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

func (s *Service) DisableHelperBeat(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableHelperBeat(ctx, UniqueID)
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

func (s *Service) DeleteHelperBeat(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteHelperBeat(ctx, UniqueID)
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

func (s *Service) GetSingleHelperBeat(ctx *models.Context, UniqueID string) (*models.RefHelperBeat, error) {
	helperbeat, err := s.Daos.GetSingleHelperBeat(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return helperbeat, nil
}

func (s *Service) FilterHelperBeat(ctx *models.Context, filter *models.FilterHelperBeat, pagination *models.Pagination) (HelperBeat []models.RefHelperBeat, err error) {
	return s.Daos.FilterHelperBeat(ctx, filter, pagination)
}
