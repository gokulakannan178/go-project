package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePayrollPolicyDetection : ""
func (s *Service) SavePayrollPolicyDetection(ctx *models.Context, payrollPolicyDetection *models.PayrollPolicyDetection) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	payrollPolicyDetection.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYROLLPOLICYDETECTION)
	payrollPolicyDetection.Status = constants.PAYROLLPOLICYDETECTIONSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PayrollPolicyDetection.created")
	payrollPolicyDetection.Created = &created
	log.Println("b4 PayrollPolicyDetection.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePayrollPolicyDetection(ctx, payrollPolicyDetection)
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

// GetSinglePayrollPolicyDetection : ""
func (s *Service) GetSinglePayrollPolicyDetection(ctx *models.Context, UniqueID string) (*models.RefPayrollPolicyDetection, error) {
	payrollPolicyDetection, err := s.Daos.GetSinglePayrollPolicyDetection(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return payrollPolicyDetection, nil
}

//UpdatePayrollPolicyDetection : ""
func (s *Service) UpdatePayrollPolicyDetection(ctx *models.Context, payrollPolicyDetection *models.PayrollPolicyDetection) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePayrollPolicyDetection(ctx, payrollPolicyDetection)
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

// EnablePayrollPolicyDetection : ""
func (s *Service) EnablePayrollPolicyDetection(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnablePayrollPolicyDetection(ctx, uniqueID)
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

// DisablePayrollPolicyDetection : ""
func (s *Service) DisablePayrollPolicyDetection(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisablePayrollPolicyDetection(ctx, uniqueID)
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

//DeletePayrollPolicyDetection : ""
func (s *Service) DeletePayrollPolicyDetection(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePayrollPolicyDetection(ctx, UniqueID)
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

// FilterPayrollPolicyDetection : ""
func (s *Service) FilterPayrollPolicyDetection(ctx *models.Context, payrollPolicyDetection *models.FilterPayrollPolicyDetection, pagination *models.Pagination) (payrollPolicyDetections []models.RefPayrollPolicyDetection, err error) {
	return s.Daos.FilterPayrollPolicyDetection(ctx, payrollPolicyDetection, pagination)
}
