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

// SaveEmployeeSalary : ""
func (s *Service) SaveEmployeeSalary(ctx *models.Context, employeeSalary *models.EmployeeSalary) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeSalary.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEESALARY)
	employeeSalary.Status = constants.EMPLOYEESALARYSTATUSACTIVE
	t := time.Now()
	employeeSalary.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeSalary.created")
	employeeSalary.Created = &created
	log.Println("b4 EmployeeSalary.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeSalary(ctx, employeeSalary)
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

// GetSingleEmployeeSalary : ""
func (s *Service) GetSingleEmployeeSalary(ctx *models.Context, UniqueID string) (*models.RefEmployeeSalary, error) {
	employeeSalary, err := s.Daos.GetSingleEmployeeSalary(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeSalary, nil
}

//UpdateEmployeeSalary : ""
func (s *Service) UpdateEmployeeSalary(ctx *models.Context, employeeSalary *models.EmployeeSalary) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeSalary(ctx, employeeSalary)
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

// EnableEmployeeSalary : ""
func (s *Service) EnableEmployeeSalary(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeSalary(ctx, uniqueID)
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

// DisableEmployeeSalary : ""
func (s *Service) DisableEmployeeSalary(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeSalary(ctx, uniqueID)
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

//DeleteEmployeeSalary : ""
func (s *Service) DeleteEmployeeSalary(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeSalary(ctx, UniqueID)
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

// FilterEmployeeSalary : ""
func (s *Service) FilterEmployeeSalary(ctx *models.Context, employeeSalary *models.FilterEmployeeSalary, pagination *models.Pagination) (employeeSalarys []models.RefEmployeeSalary, err error) {
	return s.Daos.FilterEmployeeSalary(ctx, employeeSalary, pagination)
}
func (s *Service) SaveEmployeeSalaryWithEmployee(ctx *models.Context, employeeSalary *models.FilterEmployeeSalary) ([]models.SalaryError, error) {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	var Salaryerrors []models.SalaryError
	var Salaryerror models.SalaryError
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		employeeFilter := new(models.FilterEmployee)
		employeeFilter.OrganisationID = employeeSalary.OrganisationId
		employee, err := s.Daos.FilterEmployee(ctx, employeeFilter, nil)
		if err != nil {
			return err
		}
		fmt.Println("no.of.employee==", len(employee))
		for _, v := range employee {
			Salary := new(models.EmployeeSalary)

			employeeSalary.Employee = v.UniqueID

			//employeeSalary.OrganisationId = v.UniqueID
			employeepayroll, err := s.Daos.EmployeePayrollWithEarningDeduction(ctx, employeeSalary)
			if err != nil {
				return err
			}
			if employeepayroll.PayRoll == nil {
				log.Println("employeepayroll not Found", v.Name, v.UniqueID)
				Salaryerror.EmployeeId = v.UniqueID
				Salaryerror.EmployeeName = v.Name
				Salaryerror.Message = "employeepayroll not Found"
				Salaryerrors = append(Salaryerrors, Salaryerror)
				continue
			}
			t := employeeSalary.StartDate
			if v.JoiningDate.Year() == t.Year() && v.JoiningDate.Month() == t.Month() {
				employeeSalary.StartDate = v.JoiningDate
			}
			FilterAttendance := new(models.FilterAttendance)
			FilterAttendance.StartDate = employeeSalary.StartDate
			FilterAttendance.Employee = v.UniqueID
			//sd := employeeSalary.StartDate
			sd := time.Date(employeeSalary.StartDate.Year(), employeeSalary.StartDate.Month(), 1, 0, 0, 0, 0, employeeSalary.StartDate.Location())
			ed := time.Date(sd.Year(), sd.Month()+1, 0, 23, 59, 59, 999999999, sd.Location())
			fmt.Println("sd===>", sd)
			fmt.Println("ed===>", ed)
			monthdays := ed.Day()
			fmt.Println("No.Of.days===>", monthdays)
			payroll, err := s.Daos.EmployeeAttendanceParRoll(ctx, FilterAttendance)
			if err != nil {
				return err
			}
			if payroll == nil {
				log.Println("employeeattendancepayroll not Found", v.Name, v.UniqueID)
				Salaryerror.EmployeeId = v.UniqueID
				Salaryerror.EmployeeName = v.Name
				Salaryerror.Message = "employee attendance payroll not Found"
				Salaryerrors = append(Salaryerrors, Salaryerror)
			}
			var earning models.Earning
			var earnings []models.Earning
			var deduction models.Deduction
			var deductions []models.Deduction
			var earningamount float64
			var deductionamount float64
			var LopDeduction float64
			var ParticalPaidDeduction float64
			var LopPf float64
			var ParticalPaidPf float64
			if employeepayroll != nil {
				if payroll != nil {
					Salary.NoOfDaysFullyPaid = payroll.Paid
					Salary.NoOfDaysLop = payroll.LossOfPay
					Salary.NoOfDaysParticalPaid = payroll.PartialPay
					fmt.Println("payroll.Paid==>", payroll.Paid)
					fmt.Println(" payroll.LossOfPay==>", payroll.LossOfPay)
					fmt.Println("payroll.PartialPay==>", payroll.PartialPay)
					perdaySalary := employeepayroll.PayRoll.GrossAmount / float64(monthdays)
					perdaySalary = s.Shared.RoundFloat(perdaySalary, 0)
					fmt.Println("perdaySalary==>", perdaySalary)
					LopDeduction = perdaySalary * payroll.LossOfPay
					ParticalPaidDeduction = (perdaySalary / 2) * payroll.PartialPay
					deduction.DeductionType = "LOPDeduction"
					deduction.Amount = LopDeduction
					deductions = append(deductions, deduction)
					deduction.DeductionType = "ParticalPaidDeduction"
					deduction.Amount = ParticalPaidDeduction
					deductions = append(deductions, deduction)
					deductionamount = ParticalPaidDeduction + LopDeduction
				}
				fmt.Println("deductionamount1==>", deductionamount)
				for _, v := range employeepayroll.Earning {
					earning.EarningType = v.Name
					earning.Amount = v.Amount
					earningamount = earningamount + earning.Amount
				}
				earnings = append(earnings, earning)
				fmt.Println("no.of.earnings==>", len(earnings))
				for _, v := range employeepayroll.Deduction {
					if v.Name == constants.EMPLOYEEDEDUCTIONTYPEPF {
						fmt.Println("Pf==>", v.Amount)
						perdaypf := v.Amount / float64(monthdays)
						perdaypf = s.Shared.RoundFloat(perdaypf, 0)
						fmt.Println("perdaypf==>", perdaypf)
						LopPf = perdaypf * payroll.LossOfPay
						ParticalPaidPf = (perdaypf / 2) * payroll.PartialPay
					}
					deduction.DeductionType = v.Name
					deduction.Amount = v.Amount - LopPf - ParticalPaidPf
					fmt.Println("Pf==>", deduction.Amount)

					deductionamount = deductionamount + deduction.Amount
					fmt.Println("deductionamount2==>", deductionamount)
				}
				fmt.Println("no.of.deductions==>", len(deductions))
				deductions = append(deductions, deduction)
				fmt.Println("no.of.deductions==>", len(deductions))
				Salary.GrossAmount = employeepayroll.PayRoll.GrossAmount - LopDeduction + LopPf + ParticalPaidPf
			}
			Salary.EmployeeId = v.UniqueID
			Salary.OrganisationId = v.OrganisationID
			Salary.Earning = earnings
			Salary.Deduction = deductions
			Salary.TotalEaringAmount = earningamount
			Salary.TotaldeductiongAmount = deductionamount
			t1 := time.Now()
			employeeSalary.Date = &t1
			created := models.Created{}
			created.On = &t1
			created.By = constants.SYSTEM
			Salary.Created = &created
			Salary.Date = &sd
			Salary.UniqueID = fmt.Sprintf("%v%v%v", Salary.Date.Day(), Salary.Date.Month(), Salary.Date.Year())
			err = s.Daos.SaveEmployeeSalaryWithUpsert(ctx, Salary)
			if err != nil {
				return err
			}
			Salaryerror.EmployeeId = v.UniqueID
			Salaryerror.EmployeeName = v.Name
			Salaryerror.Message = "Success"
			Salaryerrors = append(Salaryerrors, Salaryerror)
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
	return Salaryerrors, nil
}
