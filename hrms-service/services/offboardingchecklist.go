package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveOffboardingCheckList : ""
func (s *Service) SaveOffboardingCheckList(ctx *models.Context, offboardingchecklist *models.OffboardingCheckList) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	offboardingchecklist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONOFFBOARDINGCHECKLIST)
	offboardingchecklist.Status = constants.OFFBOARDINGCHECKLISTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OffboardingCheckList.created")
	offboardingchecklist.Created = &created
	log.Println("b4 OffboardingCheckList.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOffboardingCheckList(ctx, offboardingchecklist)
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

// SaveOffboardingCheckListWithoutTransaction : ""
func (s *Service) SaveOffboardingCheckListWithoutTransaction(ctx *models.Context, offboardingchecklist *models.OffboardingCheckList) error {
	log.Println("transaction start")
	// Start Transaction

	offboardingchecklist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONOFFBOARDINGCHECKLIST)
	offboardingchecklist.Status = constants.OFFBOARDINGCHECKLISTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OffboardingCheckList.created")
	offboardingchecklist.Created = &created
	log.Println("b4 OffboardingCheckList.created")
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.SaveOffboardingCheckList(ctx, offboardingchecklist)
		if dberr != nil {
			return dberr
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// GetSingleOffboardingCheckList : ""
func (s *Service) GetSingleOffboardingCheckList(ctx *models.Context, UniqueID string) (*models.RefOffboardingCheckList, error) {
	offboardingchecklist, err := s.Daos.GetSingleOffboardingCheckList(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return offboardingchecklist, nil
}

//UpdateOffboardingCheckList : ""
func (s *Service) UpdateOffboardingCheckList(ctx *models.Context, offboardingchecklist *models.OffboardingCheckList) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOffboardingCheckList(ctx, offboardingchecklist)
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

// EnableOffboardingCheckList : ""
func (s *Service) EnableOffboardingCheckList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableOffboardingCheckList(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// DisableOffboardingCheckList : ""
func (s *Service) DisableOffboardingCheckList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableOffboardingCheckList(ctx, uniqueID)
		if debrr != nil {
			return debrr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}

//DeleteOffboardingCheckList : ""
func (s *Service) DeleteOffboardingCheckList(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOffboardingCheckList(ctx, UniqueID)
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

// FilterOffboardingCheckList : ""
func (s *Service) FilterOffboardingCheckList(ctx *models.Context, offboardingchecklist *models.FilterOffboardingCheckList, pagination *models.Pagination) (offboardingchecklists []models.RefOffboardingCheckList, err error) {
	return s.Daos.FilterOffboardingCheckList(ctx, offboardingchecklist, pagination)
}
