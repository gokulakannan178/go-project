package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveExpenseCategoryList : ""
func (s *Service) SaveExpenseCategoryList(ctx *models.Context, expensecategorylist *models.ExpenseCategoryList) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	expensecategorylist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEXPENSECATEGORYLIST)
	expensecategorylist.Status = constants.EXPENSECATEGORYLISTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ExpenseCategoryList.created")
	expensecategorylist.Created = &created
	log.Println("b4 ExpenseCategoryList.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveExpenseCategoryList(ctx, expensecategorylist)
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

// SaveExpenseCategoryListWithoutTransaction : ""
func (s *Service) SaveExpenseCategoryListWithoutTransaction(ctx *models.Context, expensecategorylist *models.ExpenseCategoryList) error {
	log.Println("transaction start")
	// Start Transaction

	expensecategorylist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEXPENSECATEGORYLIST)
	expensecategorylist.Status = constants.EXPENSECATEGORYLISTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ExpenseCategoryList.created")
	expensecategorylist.Created = &created
	log.Println("b4 ExpenseCategoryList.created")

	dberr := s.Daos.SaveExpenseCategoryListUpdert(ctx, expensecategorylist)
	if dberr != nil {
		return dberr
	}
	return nil

}

// GetSingleExpenseCategoryList : ""
func (s *Service) GetSingleExpenseCategoryList(ctx *models.Context, UniqueID string) (*models.RefExpenseCategoryList, error) {
	expensecategorylist, err := s.Daos.GetSingleExpenseCategoryList(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return expensecategorylist, nil
}

//UpdateExpenseCategoryList : ""
func (s *Service) UpdateExpenseCategoryList(ctx *models.Context, expensecategorylist *models.ExpenseCategoryList) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateExpenseCategoryList(ctx, expensecategorylist)
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

// EnableExpenseCategoryList : ""
func (s *Service) EnableExpenseCategoryList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableExpenseCategoryList(ctx, uniqueID)
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

// DisableExpenseCategoryList : ""
func (s *Service) DisableExpenseCategoryList(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableExpenseCategoryList(ctx, uniqueID)
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

//DeleteExpenseCategoryList : ""
func (s *Service) DeleteExpenseCategoryList(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteExpenseCategoryList(ctx, UniqueID)
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

// FilterExpenseCategoryList : ""
func (s *Service) FilterExpenseCategoryList(ctx *models.Context, expensecategorylist *models.FilterExpenseCategoryList, pagination *models.Pagination) (expensecategorylists []models.RefExpenseCategoryList, err error) {
	return s.Daos.FilterExpenseCategoryList(ctx, expensecategorylist, pagination)
}
