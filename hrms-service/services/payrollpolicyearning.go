package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePayrollPolicyEarning : ""
func (s *Service) SavePayrollPolicyEarning(ctx *models.Context, payrollPolicyEarning *models.PayrollPolicyEarning) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	payrollPolicyEarning.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYROLLPOLICYEARNING)
	payrollPolicyEarning.Status = constants.PAYROLLPOLICYEARNINGSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PayrollPolicyEarning.created")
	payrollPolicyEarning.Created = &created
	log.Println("b4 PayrollPolicyEarning.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePayrollPolicyEarning(ctx, payrollPolicyEarning)
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

// GetSinglePayrollPolicyEarning : ""
func (s *Service) GetSinglePayrollPolicyEarning(ctx *models.Context, UniqueID string) (*models.RefPayrollPolicyEarning, error) {
	payrollPolicyEarning, err := s.Daos.GetSinglePayrollPolicyEarning(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return payrollPolicyEarning, nil
}

//UpdatePayrollPolicyEarning : ""
func (s *Service) UpdatePayrollPolicyEarning(ctx *models.Context, payrollPolicyEarning *models.PayrollPolicyEarning) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePayrollPolicyEarning(ctx, payrollPolicyEarning)
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

// EnablePayrollPolicyEarning : ""
func (s *Service) EnablePayrollPolicyEarning(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnablePayrollPolicyEarning(ctx, uniqueID)
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

// DisablePayrollPolicyEarning : ""
func (s *Service) DisablePayrollPolicyEarning(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisablePayrollPolicyEarning(ctx, uniqueID)
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

//DeletePayrollPolicyEarning : ""
func (s *Service) DeletePayrollPolicyEarning(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePayrollPolicyEarning(ctx, UniqueID)
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

// FilterPayrollPolicyEarning : ""
func (s *Service) FilterPayrollPolicyEarning(ctx *models.Context, payrollPolicyEarning *models.FilterPayrollPolicyEarning, pagination *models.Pagination) (payrollPolicyEarnings []models.RefPayrollPolicyEarning, err error) {
	return s.Daos.FilterPayrollPolicyEarning(ctx, payrollPolicyEarning, pagination)
}
