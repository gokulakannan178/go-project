package services

import (
	"errors"
	"fmt"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveExpenseCategory : ""
func (s *Service) SaveExpenseCategory(ctx *models.Context, expensecategory *models.ExpenseCategory) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	expensecategory.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEXPENSECATEGORY)
	expensecategory.Status = constants.EXPENSECATEGORYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ExpenseCategory.created")
	expensecategory.Created = &created
	log.Println("b4 ExpenseCategory.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveExpenseCategory(ctx, expensecategory)
		if dberr != nil {
			return dberr
		}

		for _, v := range expensecategory.SubCategory {
			ExpenseCategoryList := new(models.ExpenseCategoryList)
			ExpenseCategoryListmaster, err := s.Daos.GetSingleExpenseSubcategoryWithActive(ctx, v, constants.ONBOARDINGCHECKLISTMASTERSTATUSACTIVE)
			if err != nil {
				fmt.Println(err)
				return err
			}

			fmt.Println("ExpenseCategoryListmaster=======", ExpenseCategoryListmaster)

			ExpenseCategoryList.CategoryId = expensecategory.UniqueID

			ExpenseCategoryList.SubcategoryId = v
			ExpenseCategoryList.Name = expensecategory.Name

			err = s.SaveExpenseCategoryListWithoutTransaction(ctx, ExpenseCategoryList)
			if err != nil {
				return err
			}

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

// GetSingleExpenseCategory : ""
func (s *Service) GetSingleExpenseCategory(ctx *models.Context, UniqueID string) (*models.RefExpenseCategory, error) {
	expenseCategory, err := s.Daos.GetSingleExpenseCategory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return expenseCategory, nil
}

//UpdateExpenseCategory : ""
func (s *Service) UpdateExpenseCategory(ctx *models.Context, expenseCategory *models.ExpenseCategory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.ExpenseCategoryListRemoveNotPresentValue(ctx, expenseCategory.UniqueID, expenseCategory.SubCategory)
		if err != nil {
			return err
		}
		err = s.Daos.ExpenseCategoryListUpsert(ctx, expenseCategory.UniqueID, expenseCategory.SubCategory, expenseCategory.Name)
		if err != nil {
			return err
		}

		fmt.Println("error==>", err)

		err = s.Daos.UpdateExpenseCategory(ctx, expenseCategory)
		if err != nil {
			return err
		}
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

// EnableExpenseCategory : ""
func (s *Service) EnableExpenseCategory(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableExpenseCategory(ctx, uniqueID)
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

// DisableExpenseCategory : ""
func (s *Service) DisableExpenseCategory(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableExpenseCategory(ctx, uniqueID)
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

//DeleteExpenseCategory : ""
func (s *Service) DeleteExpenseCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteExpenseCategory(ctx, UniqueID)
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

// FilterExpenseCategory : ""
func (s *Service) FilterExpenseCategory(ctx *models.Context, expensecategory *models.FilterExpenseCategory, pagination *models.Pagination) (expensecategorys []models.RefExpenseCategory, err error) {
	err = s.ExpenseCategoryDataAccess(ctx, expensecategory)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterExpenseCategory(ctx, expensecategory, pagination)
}

// func (s *Service) ExpenseCategory(ctx *models.Context, expenseCategory *models.ExpenseCategory) error {
// 	err := s.SaveExpenseCategory(ctx, expenseCategory)
// 	if err != nil {
// 		return err
// 	}
// 	onboardChecklist := new(models.OnboardingCheckList)
// 	refExpenseCategory, err := s.Daos.GetSingleExpenseCategory(ctx, expenseCategory.UniqueID)
// 	if err != nil {
// 		return err
// 	}
// 	// var uniqueID []string
// 	for _, v := range expenseCategory.ExpenseCategory {
// 		if refExpenseCategory != nil {
// 			onboardChecklist.ExpenseCategoryID = refExpenseCategory.UniqueID
// 		}
// 		//onboardChecklist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEXPENSECATEGORY)

// 		onboardChecklist.OnboardingchecklistmasterID = v
// 		err = s.SaveOnboardingCheckList(ctx, onboardChecklist)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
func (s *Service) ExpenseCategoryDataAccess(ctx *models.Context, filter *models.FilterExpenseCategory) (err error) {
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
