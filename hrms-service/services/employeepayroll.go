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

// SaveEmployeePayroll : ""
func (s *Service) SaveEmployeePayroll(ctx *models.Context, employeePayroll *models.EmployeePayroll) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeePayroll.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEPAYROLL)
	employeePayroll.Status = constants.EMPLOYEEPAYROLLSTATUSACTIVE
	t := time.Now()
	employeePayroll.StartDate = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeePayroll.created")
	employeePayroll.Created = &created
	log.Println("b4 EmployeePayroll.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeePayroll(ctx, employeePayroll)
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

// GetSingleEmployeePayroll : ""
func (s *Service) GetSingleEmployeePayroll(ctx *models.Context, UniqueID string) (*models.RefEmployeePayroll, error) {
	employeePayroll, err := s.Daos.GetSingleEmployeePayroll(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeePayroll, nil
}

//UpdateEmployeePayroll : ""
func (s *Service) UpdateEmployeePayroll(ctx *models.Context, employeePayroll *models.EmployeePayroll) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeePayroll(ctx, employeePayroll)
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

// EnableEmployeePayroll : ""
func (s *Service) EnableEmployeePayroll(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeePayroll(ctx, uniqueID)
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

// DisableEmployeePayroll : ""
func (s *Service) DisableEmployeePayroll(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeePayroll(ctx, uniqueID)
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

//DeleteEmployeePayroll : ""
func (s *Service) DeleteEmployeePayroll(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeePayroll(ctx, UniqueID)
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

// FilterEmployeePayroll : ""
func (s *Service) FilterEmployeePayroll(ctx *models.Context, employeePayroll *models.FilterEmployeePayroll, pagination *models.Pagination) (employeePayrolls []models.RefEmployeePayroll, err error) {
	return s.Daos.FilterEmployeePayroll(ctx, employeePayroll, pagination)
}
func (s *Service) EmployeeUpdatePayrollWithNewPayroll(ctx *models.Context, employeePayroll *models.EmployeePayrollWithEarningDeduction) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		payroll, _ := s.Daos.GetSingleEmployeePayrollWithEmployee(ctx, employeePayroll.EmployeeId)
		if payroll == nil {
			log.Println("Payroll not Found")
		}
		if payroll != nil {
			err := s.Daos.ArchivedEmployeePayroll(ctx, payroll.UniqueID)
			if err != nil {
				return err
			}
		}
		NewEmployeePayroll := new(models.EmployeePayroll)
		NewEmployeePayroll.CTC = employeePayroll.NetAmount * 12
		NewEmployeePayroll.NetAmount = employeePayroll.NetAmount
		NewEmployeePayroll.EmployeeId = employeePayroll.EmployeeId
		NewEmployeePayroll.OrganisationId = employeePayroll.OrganisationId
		NewEmployeePayroll.OrganisationId = employeePayroll.OrganisationId
		t := time.Now()
		NewEmployeePayroll.StartDate = &t
		NewEmployeePayroll.Status = constants.EMPLOYEEPAYROLLSTATUSACTIVE
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		log.Println("b4 EmployeePayroll.created")
		NewEmployeePayroll.Created = &created
		NewEmployeePayroll.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEPAYROLL)
		var totalEarningAmount float64
		for _, v := range employeePayroll.Earning {
			earning, _ := s.Daos.GetSingleEmployeeEaringWithEmployee(ctx, employeePayroll.EmployeeId, v.EarningId)
			if earning == nil {
				log.Println("Payroll not Found")
			}
			fmt.Println("earningId===>", earning)
			if earning != nil {
				err := s.Daos.ArchivedEmployeeEarning(ctx, earning.UniqueID)
				if err != nil {
					return err
				}
			}
			NewEmployeeEarning := new(models.EmployeeEarning)
			NewEmployeeEarning.OrganisationId = employeePayroll.OrganisationId
			NewEmployeeEarning.EmployeeId = employeePayroll.EmployeeId
			NewEmployeeEarning.EarningId = v.EarningId
			NewEmployeeEarning.Amount = v.Amount
			NewEmployeeEarning.StartDate = &t
			NewEmployeeEarning.Status = constants.EMPLOYEEEARNINGSTATUSACTIVE
			NewEmployeeEarning.Created = &created
			NewEmployeeEarning.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEEARNING)
			totalEarningAmount = totalEarningAmount + v.Amount
			err := s.Daos.SaveEmployeeEarning(ctx, NewEmployeeEarning)
			if err != nil {
				return err
			}
		}
		var totalDeductionAmount float64
		for _, v := range employeePayroll.Deduction {
			Deduction, _ := s.Daos.GetSingleEmployeeDeductionWithEmployee(ctx, employeePayroll.EmployeeId, v.DeductionId)
			if Deduction == nil {
				log.Println("Payroll not Found")
			}
			fmt.Println("Deduction===>", Deduction)
			if Deduction != nil {
				err := s.Daos.ArchivedEmployeeDeduction(ctx, Deduction.UniqueID)
				if err != nil {
					return err
				}
			}
			NewEmployeeDeduction := new(models.EmployeeDeduction)
			NewEmployeeDeduction.OrganisationId = employeePayroll.OrganisationId
			NewEmployeeDeduction.EmployeeId = employeePayroll.EmployeeId
			NewEmployeeDeduction.DeductionId = v.DeductionId
			NewEmployeeDeduction.Amount = v.Amount
			NewEmployeeDeduction.StartDate = &t
			NewEmployeeDeduction.Status = constants.EMPLOYEEEARNINGSTATUSACTIVE
			NewEmployeeDeduction.Created = &created
			NewEmployeeDeduction.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEEARNING)
			totalDeductionAmount = totalDeductionAmount + v.Amount
			err := s.Daos.SaveEmployeeDeduction(ctx, NewEmployeeDeduction)
			if err != nil {
				return err
			}
		}
		NewEmployeePayroll.GrossAmount = totalEarningAmount
		NewEmployeePayroll.Deduction = totalDeductionAmount
		debrr := s.Daos.SaveEmployeePayroll(ctx, NewEmployeePayroll)
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
func (s *Service) SaveEmployeePayrollWithEaringDeduction(ctx *models.Context, employeePayroll *models.EmployeePayrollWithEarningDeduction) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		NewEmployeePayroll := new(models.EmployeePayroll)
		NewEmployeePayroll.CTC = employeePayroll.NetAmount * 12
		NewEmployeePayroll.NetAmount = employeePayroll.NetAmount
		NewEmployeePayroll.EmployeeId = employeePayroll.EmployeeId
		NewEmployeePayroll.OrganisationId = employeePayroll.OrganisationId
		NewEmployeePayroll.OrganisationId = employeePayroll.OrganisationId
		t := time.Now()
		NewEmployeePayroll.StartDate = &t
		NewEmployeePayroll.Status = constants.EMPLOYEEPAYROLLSTATUSACTIVE
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		log.Println("b4 EmployeePayroll.created")
		NewEmployeePayroll.Created = &created
		NewEmployeePayroll.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEPAYROLL)
		var totalEarningAmount float64
		for _, v := range employeePayroll.Earning {
			NewEmployeeEarning := new(models.EmployeeEarning)
			NewEmployeeEarning.OrganisationId = employeePayroll.OrganisationId
			NewEmployeeEarning.EmployeeId = employeePayroll.EmployeeId
			NewEmployeeEarning.EarningId = v.EarningId
			NewEmployeeEarning.Amount = v.Amount
			NewEmployeeEarning.StartDate = &t
			NewEmployeeEarning.Status = constants.EMPLOYEEEARNINGSTATUSACTIVE
			NewEmployeeEarning.Created = &created
			NewEmployeeEarning.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEEARNING)
			totalEarningAmount = totalEarningAmount + v.Amount
			err := s.Daos.SaveEmployeeEarning(ctx, NewEmployeeEarning)
			if err != nil {
				return err
			}
		}
		var totalDeductionAmount float64
		for _, v := range employeePayroll.Deduction {
			NewEmployeeDeduction := new(models.EmployeeDeduction)
			NewEmployeeDeduction.OrganisationId = employeePayroll.OrganisationId
			NewEmployeeDeduction.EmployeeId = employeePayroll.EmployeeId
			NewEmployeeDeduction.DeductionId = v.DeductionId
			NewEmployeeDeduction.Amount = v.Amount
			NewEmployeeDeduction.StartDate = &t
			NewEmployeeDeduction.Status = constants.EMPLOYEEEARNINGSTATUSACTIVE
			NewEmployeeDeduction.Created = &created
			NewEmployeeDeduction.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEEARNING)
			totalDeductionAmount = totalDeductionAmount + v.Amount
			err := s.Daos.SaveEmployeeDeduction(ctx, NewEmployeeDeduction)
			if err != nil {
				return err
			}
		}
		NewEmployeePayroll.GrossAmount = totalEarningAmount
		NewEmployeePayroll.Deduction = totalDeductionAmount
		debrr := s.Daos.SaveEmployeePayroll(ctx, NewEmployeePayroll)
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
