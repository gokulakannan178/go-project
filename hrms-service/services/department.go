package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveDepartment : ""
func (s *Service) SaveDepartment(ctx *models.Context, Dept *models.Department) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Dept.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDEPARTMENT)
	Dept.Status = constants.DEPARTMENTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Department.created")
	Dept.Created = &created
	log.Println("b4 Department.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDepartment(ctx, Dept)
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

// GetSingleDepartment : ""
func (s *Service) GetSingleDepartment(ctx *models.Context, UniqueID string) (*models.RefDepartment, error) {
	Department, err := s.Daos.GetSingleDepartment(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Department, nil
}

func (s *Service) UpdateDepartment(ctx *models.Context, dept *models.Department) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDepartment(ctx, dept)
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

// EnableDepartment : ""
func (s *Service) EnableDepartment(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableDepartment(ctx, uniqueID)
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

// DisableDepartment : ""
func (s *Service) DisableDepartment(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableDepartment(ctx, uniqueID)
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

//DeleteState : ""
func (s *Service) DeleteDepartment(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDepartment(ctx, UniqueID)
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

// FilterDepartment : ""
func (s *Service) FilterDepartment(ctx *models.Context, Filter *models.FilterDepartment, pagination *models.Pagination) (Dept []models.RefDepartment, err error) {
	err = s.DepartmentDataAccess(ctx, Filter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterDepartment(ctx, Filter, pagination)
}
func (s *Service) DepartmentDataAccess(ctx *models.Context, filter *models.FilterDepartment) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationId = append(filter.OrganisationId, v.UniqueID)
				}
			}
			if dataaccess.Department != "" {
				filter.UniqueID = append(filter.UniqueID, dataaccess.Department)
			}

		}

	}
	return err
}
func (s *Service) GetSingleDepartmentUniqueCheck(ctx *models.Context, UniqueID string, org string) (string, error) {
	Department, err := s.Daos.GetSingleDepartmentUniqueCheck(ctx, UniqueID, org)
	if err != nil {
		return "", err
	}
	result := ""
	if Department != nil {
		result = "Department Already Exist"
	}
	return result, nil
}
