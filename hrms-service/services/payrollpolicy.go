package services

import (
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePayrollPolicy : ""
func (s *Service) SavePayrollPolicy(ctx *models.Context, payrollPolicy *models.PayrollPolicy) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	payrollPolicy.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYROLLPOLICY)
	payrollPolicy.Status = constants.PAYROLLPOLICYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PayrollPolicy.created")
	payrollPolicy.Created = &created
	log.Println("b4 PayrollPolicy.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range payrollPolicy.EarningMaster {
			payrollPolicy.TakeHome = v.Value + payrollPolicy.TakeHome
		}
		for _, v2 := range payrollPolicy.DetectionMaster {
			payrollPolicy.Deduction = v2.Value + payrollPolicy.Deduction
		}
		payrollPolicy.GrossAmount = payrollPolicy.TakeHome - payrollPolicy.Deduction
		payrollPolicy.CTC = payrollPolicy.TakeHome * 12
		dberr := s.Daos.SavePayrollPolicy(ctx, payrollPolicy)
		if dberr != nil {
			return dberr
		}
		for _, v := range payrollPolicy.EarningMaster {
			earningmaster := new(models.PayrollPolicyEarning)
			earningmaster.PayRollPolicyId = payrollPolicy.UniqueID
			earningmaster.EarningMasterId = v.UniqueID
			earningmaster.Amount = v.Value
			earningmaster.Status = constants.PAYROLLPOLICYSTATUSACTIVE
			earningmaster.Created = &created
			earningmaster.Name = payrollPolicy.Name
			earningmaster.Description = payrollPolicy.Description
			earningmaster.OrganisationID = payrollPolicy.OrganisationID
			err := s.Daos.SavePayrollPolicyEarning(ctx, earningmaster)
			if err != nil {
				return err
			}

		}
		for _, v := range payrollPolicy.DetectionMaster {
			DetectionMaster := new(models.PayrollPolicyDetection)
			DetectionMaster.PayRollPolicyId = payrollPolicy.UniqueID
			DetectionMaster.DetectionMasterId = v.UniqueID
			DetectionMaster.Amount = v.Value
			DetectionMaster.Status = constants.PAYROLLPOLICYSTATUSACTIVE
			DetectionMaster.Created = &created
			DetectionMaster.Name = payrollPolicy.Name
			DetectionMaster.Description = payrollPolicy.Description
			DetectionMaster.OrganisationID = payrollPolicy.OrganisationID
			err := s.Daos.SavePayrollPolicyDetection(ctx, DetectionMaster)
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

// GetSinglePayrollPolicy : ""
func (s *Service) GetSinglePayrollPolicy(ctx *models.Context, UniqueID string) (*models.RefPayrollPolicy, error) {
	payrollPolicy, err := s.Daos.GetSinglePayrollPolicy(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return payrollPolicy, nil
}

//UpdatePayrollPolicy : ""
func (s *Service) UpdatePayrollPolicy(ctx *models.Context, payrollPolicy *models.PayrollPolicy) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePayrollPolicy(ctx, payrollPolicy)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		for _, v := range payrollPolicy.EarningMaster {
			payrollPolicyEaring := new(models.PayrollPolicyEarning)
			payrollPolicyEaring.PayRollPolicyId = payrollPolicy.UniqueID
			payrollPolicyEaring.EarningMasterId = v.UniqueID
			payrollPolicyEaring.Amount = v.Value
			payrollPolicyEaring.Status = constants.PAYROLLPOLICYSTATUSACTIVE
			err := s.Daos.UpdatePayrollPolicyEarningWithEaringId(ctx, payrollPolicyEaring)
			if err != nil {
				return err
			}
		}
		for _, v := range payrollPolicy.DetectionMaster {
			payrollPolicyEaring := new(models.PayrollPolicyDetection)
			payrollPolicyEaring.PayRollPolicyId = payrollPolicy.UniqueID
			payrollPolicyEaring.DetectionMasterId = v.UniqueID
			payrollPolicyEaring.Amount = v.Value
			payrollPolicyEaring.Status = constants.PAYROLLPOLICYSTATUSACTIVE
			err := s.Daos.UpdatePayrollPolicyEarningWithDetection(ctx, payrollPolicyEaring)
			if err != nil {
				return err
			}
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

// EnablePayrollPolicy : ""
func (s *Service) EnablePayrollPolicy(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnablePayrollPolicy(ctx, uniqueID)
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

// DisablePayrollPolicy : ""
func (s *Service) DisablePayrollPolicy(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisablePayrollPolicy(ctx, uniqueID)
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

//DeletePayrollPolicy : ""
func (s *Service) DeletePayrollPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePayrollPolicy(ctx, UniqueID)
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

// FilterPayrollPolicy : ""
func (s *Service) FilterPayrollPolicy(ctx *models.Context, payrollPolicy *models.FilterPayrollPolicy, pagination *models.Pagination) (payrollPolicys []models.RefPayrollPolicy, err error) {
	return s.Daos.FilterPayrollPolicy(ctx, payrollPolicy, pagination)
}
func (s *Service) GetSalaryCalc(ctx *models.Context, UniqueID string) (*models.SalaryCalc, error) {
	payrollPolicy := new(models.SalaryCalc)
	f, err := strconv.ParseFloat(UniqueID, 64)
	if err != nil {
		return nil, err
	}
	basic := s.ConfigReader.GetInt(s.Shared.GetCmdArg(constants.ENV) + "." + constants.BASIC)
	hra := s.ConfigReader.GetInt(s.Shared.GetCmdArg(constants.ENV) + "." + constants.HRA)
	conveyanceAllowances := s.ConfigReader.GetFloat(s.Shared.GetCmdArg(constants.ENV) + "." + constants.CONVEYANCE)
	payrollPolicy.Earnings.BasicSalary = f * (float64(basic) / 100)
	payrollPolicy.Earnings.Hra = f * (float64(hra) / 100)
	fmt.Println("conveyanceAllowances ===>", conveyanceAllowances)
	payrollPolicy.Earnings.ConveyanceAllowances = conveyanceAllowances
	payrollPolicy.Earnings.EducationAllowance = f - payrollPolicy.Earnings.BasicSalary - payrollPolicy.Earnings.Hra - payrollPolicy.Earnings.ConveyanceAllowances
	payrollPolicy.GrossAmount = payrollPolicy.Earnings.BasicSalary + payrollPolicy.Earnings.Hra + payrollPolicy.Earnings.ConveyanceAllowances + payrollPolicy.Earnings.EducationAllowance
	pf := s.ConfigReader.GetInt(s.Shared.GetCmdArg(constants.ENV) + "." + constants.PF)
	esic := s.ConfigReader.GetFloat(s.Shared.GetCmdArg(constants.ENV) + "." + constants.ESIC)
	fmt.Println("esic===>", esic)
	payrollPolicy.Detections.PfContribution = payrollPolicy.Earnings.BasicSalary * (float64(pf) / 100)
	payrollPolicy.Detections.ESICContribution = f * (esic / 100)
	fmt.Println("payrollPolicy.Detections.ESICContribution ===>", payrollPolicy.Detections.ESICContribution)

	payrollPolicy.TotalDeduction = payrollPolicy.Detections.PfContribution + payrollPolicy.Detections.ESICContribution
	payrollPolicy.NetSalary = payrollPolicy.GrossAmount - payrollPolicy.TotalDeduction
	return payrollPolicy, nil
}
func (s *Service) GetSalaryCalcV2(ctx *models.Context, UniqueID string, EmployeeType string) (*models.SalaryCalc, error) {
	payrollPolicy := new(models.SalaryCalc)
	f, err := strconv.ParseFloat(UniqueID, 64)
	if err != nil {
		return nil, err
	}
	AutoSalary, err := s.Daos.GetSingleSalaryConfigWithEmployeeType(ctx, EmployeeType)
	if err != nil {
		return nil, err
	}
	if AutoSalary != nil {
		payrollPolicy.SalaryConfigId = AutoSalary.UniqueID
		payrollPolicy.Earnings.BasicSalary = roundFloat(f*(AutoSalary.Earnings.BasicSalary/100), 0)
		payrollPolicy.Earnings.Hra = roundFloat(f*(AutoSalary.Earnings.Hra/100), 0)
		payrollPolicy.Earnings.ConveyanceAllowances = AutoSalary.Earnings.ConveyanceAllowances
		payrollPolicy.Earnings.EducationAllowance = roundFloat(f-payrollPolicy.Earnings.BasicSalary-payrollPolicy.Earnings.Hra-payrollPolicy.Earnings.ConveyanceAllowances, 0)
		payrollPolicy.GrossAmount = roundFloat(payrollPolicy.Earnings.BasicSalary+payrollPolicy.Earnings.Hra+payrollPolicy.Earnings.ConveyanceAllowances+payrollPolicy.Earnings.EducationAllowance, 0)
		payrollPolicy.Detections.PfContribution = roundFloat(payrollPolicy.Earnings.BasicSalary*(AutoSalary.Detections.PfContribution/100), 0)
		payrollPolicy.Detections.ESICContribution = roundFloat(f*(AutoSalary.Detections.ESICContribution/100), 0)
		payrollPolicy.TotalDeduction = payrollPolicy.Detections.PfContribution + payrollPolicy.Detections.ESICContribution
		payrollPolicy.NetSalary = payrollPolicy.GrossAmount - payrollPolicy.TotalDeduction
	} else {
		return nil, errors.New("Plaese Set Salary Configuration")
	}
	return payrollPolicy, nil
}
