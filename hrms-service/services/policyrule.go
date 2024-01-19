package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePolicyRule : ""
func (s *Service) SavePolicyRule(ctx *models.Context, policyrule *models.PolicyRule) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	policyrule.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEAVEPOLICY)
	policyrule.Status = constants.LEAVEPOLICYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PolicyRule.created")
	policyrule.Created = &created
	log.Println("b4 PolicyRule.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePolicyRule(ctx, policyrule)
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

// SavePolicyRuleWithoutTransaction : ""
func (s *Service) SavePolicyRuleWithoutTransaction(ctx *models.Context, policyrule *models.PolicyRule) error {
	log.Println("transaction start")
	// Start Transaction
	policyrule.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEAVEPOLICY)
	policyrule.Status = constants.LEAVEPOLICYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PolicyRule.created")
	policyrule.Created = &created
	log.Println("b4 PolicyRule.created")
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePolicyRule(ctx, policyrule)
		if dberr != nil {
			return dberr
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// GetSinglePolicyRule : ""
func (s *Service) GetSinglePolicyRule(ctx *models.Context, UniqueID string) (*models.RefPolicyRule, error) {
	PolicyRule, err := s.Daos.GetSinglePolicyRule(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return PolicyRule, nil
}

//UpdatePolicyRule : ""
func (s *Service) UpdatePolicyRule(ctx *models.Context, policyrule *models.PolicyRule) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePolicyRule(ctx, policyrule)
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

// EnablePolicyRule : ""
func (s *Service) EnablePolicyRule(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnablePolicyRule(ctx, uniqueID)
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

// DisablePolicyRule : ""
func (s *Service) DisablePolicyRule(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisablePolicyRule(ctx, uniqueID)
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

//DeletePolicyRule : ""
func (s *Service) DeletePolicyRule(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePolicyRule(ctx, UniqueID)
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

// Filterpolicyrule : ""
func (s *Service) FilterPolicyRule(ctx *models.Context, policyrule *models.FilterPolicyRule, pagination *models.Pagination) (policyrules []models.RefPolicyRule, err error) {
	return s.Daos.FilterPolicyRule(ctx, policyrule, pagination)
}
