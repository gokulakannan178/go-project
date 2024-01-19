package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveOnboardingCheckList : ""
func (s *Service) SaveOnboardingCheckList(ctx *models.Context, onboardingchecklist *models.OnboardingCheckList) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	onboardingchecklist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONONBOARDINGCHECKLIST)
	onboardingchecklist.Status = constants.ONBOARDINGCHECKLISTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OnboardingCheckList.created")
	onboardingchecklist.Created = &created
	log.Println("b4 OnboardingCheckList.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOnboardingCheckList(ctx, onboardingchecklist)
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

// SaveOnboardingCheckListWithoutTransaction : ""
func (s *Service) SaveOnboardingCheckListWithoutTransaction(ctx *models.Context, onboardingchecklist *models.OnboardingCheckList) error {
	log.Println("transaction start")
	// Start Transaction

	onboardingchecklist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONONBOARDINGCHECKLIST)
	onboardingchecklist.Status = constants.ONBOARDINGCHECKLISTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OnboardingCheckList.created")
	onboardingchecklist.Created = &created
	log.Println("b4 OnboardingCheckList.created")

	dberr := s.Daos.SaveOnboardingCheckListUpdert(ctx, onboardingchecklist)
	if dberr != nil {
		return dberr
	}
	return nil

}

// GetSingleOnboardingCheckList : ""
func (s *Service) GetSingleOnboardingCheckList(ctx *models.Context, UniqueID string) (*models.RefOnboardingCheckList, error) {
	onboardingchecklist, err := s.Daos.GetSingleOnboardingCheckList(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return onboardingchecklist, nil
}

//UpdateOnboardingCheckList : ""
func (s *Service) UpdateOnboardingCheckList(ctx *models.Context, onboardingchecklist *models.OnboardingCheckList) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOnboardingCheckList(ctx, onboardingchecklist)
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

// EnableOnboardingCheckList : ""
func (s *Service) EnableOnboardingCheckList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableOnboardingCheckList(ctx, uniqueID)
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

// DisableOnboardingCheckList : ""
func (s *Service) DisableOnboardingCheckList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableOnboardingCheckList(ctx, uniqueID)
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

//DeleteOnboardingCheckList : ""
func (s *Service) DeleteOnboardingCheckList(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOnboardingCheckList(ctx, UniqueID)
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

// FilterOnboardingCheckList : ""
func (s *Service) FilterOnboardingCheckList(ctx *models.Context, onboardingchecklist *models.FilterOnboardingCheckList, pagination *models.Pagination) (onboardingchecklists []models.RefOnboardingCheckList, err error) {
	return s.Daos.FilterOnboardingCheckList(ctx, onboardingchecklist, pagination)
}
