package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveExpenseSubcategory : ""
func (s *Service) SaveExpenseSubcategory(ctx *models.Context, expensesubcategory *models.ExpenseSubcategory) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	expensesubcategory.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEXPENSESUBCATEGORY)
	expensesubcategory.Status = constants.EXPENSESUBCATEGORYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ExpenseSubcategory.created")
	expensesubcategory.Created = &created
	log.Println("b4 ExpenseSubcategory.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveExpenseSubcategory(ctx, expensesubcategory)
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

// GetSingleExpenseSubcategory : ""
func (s *Service) GetSingleExpenseSubcategory(ctx *models.Context, UniqueID string) (*models.RefExpenseSubcategory, error) {
	expensesubcategory, err := s.Daos.GetSingleExpenseSubcategory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return expensesubcategory, nil
}

//UpdateExpenseSubcategory : ""
func (s *Service) UpdateExpenseSubcategory(ctx *models.Context, expensesubcategory *models.ExpenseSubcategory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateExpenseSubcategory(ctx, expensesubcategory)
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

// EnableExpenseSubcategory : ""
func (s *Service) EnableExpenseSubcategory(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableExpenseSubcategory(ctx, uniqueID)
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

// DisableExpenseSubcategory : ""
func (s *Service) DisableExpenseSubcategory(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableExpenseSubcategory(ctx, uniqueID)
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

//DeleteExpenseSubcategory : ""
func (s *Service) DeleteExpenseSubcategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteExpenseSubcategory(ctx, UniqueID)
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

// FilterExpenseSubcategory : ""
func (s *Service) FilterExpenseSubcategory(ctx *models.Context, expensesubcategory *models.FilterExpenseSubcategory, pagination *models.Pagination) (expensesubcategorys []models.RefExpenseSubcategory, err error) {
	err = s.ExpenseSubcategoryDataAccess(ctx, expensesubcategory)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterExpenseSubcategory(ctx, expensesubcategory, pagination)
}
func (s *Service) ExpenseSubcategoryDataAccess(ctx *models.Context, filter *models.FilterExpenseSubcategory) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationID = append(filter.OrganisationID, v.UniqueID)
				}
			}

		}

	}
	return err
}
