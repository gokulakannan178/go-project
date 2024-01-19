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

// SaveSalaryConfig : ""
func (s *Service) SaveSalaryConfig(ctx *models.Context, salaryConfig *models.SalaryConfig) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	salaryConfig.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSALARYCONFIG)
	salaryConfig.Status = constants.SALARYCONFIGSTATUSACTIVE
	t := time.Now()
	salaryConfig.StartDate = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 SalaryConfig.created")
	salaryConfig.Created = &created
	log.Println("b4 SalaryConfig.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSalaryConfig(ctx, salaryConfig)
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

// GetSingleSalaryConfig : ""
func (s *Service) GetSingleSalaryConfig(ctx *models.Context, UniqueID string) (*models.RefSalaryConfig, error) {
	SalaryConfig, err := s.Daos.GetSingleSalaryConfig(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return SalaryConfig, nil
}

//UpdateSalaryConfig : ""
func (s *Service) UpdateSalaryConfig(ctx *models.Context, salaryConfig *models.SalaryConfig) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		checksalary, err := s.Daos.GetSingleSalaryConfig(ctx, salaryConfig.UniqueID)
		if err != nil {
			return err
		}
		if checksalary != nil {
			salaryConfigLog := new(models.SalaryConfigLog)
			salaryConfigLog.PreSalaryConfig = checksalary.SalaryConfig
			salaryConfigLog.NewSalaryConfig = *salaryConfig
			salaryConfigLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSALARYCONFIGLOG)
			salaryConfigLog.Status = constants.SALARYCONFIGLOGSTATUSACTIVE
			t := time.Now()
			salaryConfigLog.Date = &t
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			salaryConfigLog.Created = &created
			err = s.Daos.SaveSalaryConfigLog(ctx, salaryConfigLog)
			if err != nil {
				return err
			}
		}
		err = s.Daos.UpdateSalaryConfig(ctx, salaryConfig)
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

// EnableSalaryConfig : ""
func (s *Service) EnableSalaryConfig(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableSalaryConfig(ctx, uniqueID)
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

// DisableSalaryConfig : ""
func (s *Service) DisableSalaryConfig(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableSalaryConfig(ctx, uniqueID)
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

//DeleteSalaryConfig : ""
func (s *Service) DeleteSalaryConfig(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteSalaryConfig(ctx, UniqueID)
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

// FilterSalaryConfig : ""
func (s *Service) FilterSalaryConfig(ctx *models.Context, salaryConfig *models.FilterSalaryConfig, pagination *models.Pagination) (salaryConfigs []models.RefSalaryConfig, err error) {
	return s.Daos.FilterSalaryConfig(ctx, salaryConfig, pagination)
}
func (s *Service) SaveSalaryConfigWithEmployeeType(ctx *models.Context, salaryConfig *models.SalaryConfig) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	salaryConfig.Status = constants.SALARYCONFIGSTATUSACTIVE
	t := time.Now()
	uniq := fmt.Sprintf("%v%v%v", t.Day(), t.Month(), t.Year())
	salaryConfig.UniqueID = uniq + s.Daos.GetUniqueID(ctx, constants.COLLECTIONSALARYCONFIG)
	salaryConfig.StartDate = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 SalaryConfig.created")
	salaryConfig.Created = &created
	salaryConfig.EarningsPercentage = salaryConfig.Earnings.BasicSalary + salaryConfig.Earnings.Hra
	salaryConfig.DeductionPercentage = salaryConfig.Detections.ESICContribution + salaryConfig.Detections.PfContribution
	//salaryConfig.NetPercentage = salaryConfig.Detections.Others + salaryConfig.Detections.ESICContribution + salaryConfig.Detections.PfContribution
	//salaryConfig.GrossPercentage = salaryConfig.Earnings.Others + salaryConfig.Earnings.BasicSalary + salaryConfig.Earnings.Hra + salaryConfig.Earnings.ConveyanceAllowances + salaryConfig.Earnings.EducationAllowance
	//salaryConfig.NetPercentage = salaryConfig.GrossPercentage - salaryConfig.DeductionPercentage
	log.Println("b4 SalaryConfig.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		checksalary, err := s.Daos.GetSingleSalaryConfigWithEmployeeType(ctx, salaryConfig.EmployeeType)
		if err != nil {
			return err
		}
		if checksalary != nil {
			checksalary.Status = constants.SALARYCONFIGSTATUSACHIEVED
			checksalary.EndDate = &t
			err := s.Daos.UpdateSalaryConfig(ctx, &checksalary.SalaryConfig)
			if err != nil {
				return err
			}
			salaryConfigLog := new(models.SalaryConfigLog)
			salaryConfigLog.PreSalaryConfig = checksalary.SalaryConfig
			salaryConfigLog.NewSalaryConfig = *salaryConfig
			salaryConfigLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSALARYCONFIGLOG)
			salaryConfigLog.Status = constants.SALARYCONFIGLOGSTATUSACTIVE
			t := time.Now()
			salaryConfigLog.Date = &t
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			salaryConfigLog.Created = &created
			err = s.Daos.SaveSalaryConfigLog(ctx, salaryConfigLog)
			if err != nil {
				return err
			}
		}
		dberr := s.Daos.SaveSalaryConfig(ctx, salaryConfig)
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
