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

// SaveLeavePolicy : ""
func (s *Service) SaveLeavePolicy(ctx *models.Context, leavepolicy *models.LeavePolicy) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	leavepolicy.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEAVEPOLICY)
	leavepolicy.Status = constants.LEAVEPOLICYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 LeavePolicy.created")
	leavepolicy.Created = &created
	log.Println("b4 LeavePolicy.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveLeavePolicy(ctx, leavepolicy)
		if dberr != nil {
			return dberr
		}

		for _, v := range leavepolicy.LeaveMaster {
			policyrule := new(models.PolicyRule)
			leavemaster, err := s.Daos.GetSingleLeaveMasterWithActive(ctx, v.UniqueID, constants.LEAVEPOLICYSTATUSACTIVE)
			if err != nil {
				fmt.Println(err)
				return err
			}
			if leavemaster == nil {
				return errors.New("Leave master id is not available" + err.Error())
			}
			fmt.Println("leavemaster", leavemaster)
			policyrule.LeavePolicyID = leavepolicy.UniqueID
			policyrule.LeaveMasterID = v.UniqueID
			policyrule.Name = leavepolicy.Name
			policyrule.Value = v.Value
			//policyrule.Value = leavemaster.Value
			err = s.SavePolicyRuleWithoutTransaction(ctx, policyrule)
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

// GetSingleLeavePolicy : ""
func (s *Service) GetSingleLeavePolicy(ctx *models.Context, UniqueID string) (*models.RefLeavePolicy, error) {
	LeavePolicy, err := s.Daos.GetSingleLeavePolicy(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return LeavePolicy, nil
}

//UpdateLeavePolicy : ""
func (s *Service) UpdateLeavePolicy(ctx *models.Context, leavepolicy *models.LeavePolicy) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if len(leavepolicy.LeaveMasterId) > 0 {
			for _, v := range leavepolicy.LeaveMaster {
				err := s.Daos.PolicyRuleRemoveNotPresentValue(ctx, leavepolicy.UniqueID, v.UniqueID)
				if err != nil {
					return err
				}
				err = s.Daos.PolicyRuleUpsert(ctx, leavepolicy.UniqueID, v.UniqueID, leavepolicy.Name)
				if err != nil {
					return err
				}
				fmt.Println("error==>", err)

			}
		}

		err := s.Daos.UpdateLeavePolicy(ctx, leavepolicy)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		for _, v := range leavepolicy.LeaveMaster {
			policyrole := new(models.PolicyRule)
			policyrole.LeavePolicyID = leavepolicy.UniqueID
			policyrole.LeaveMasterID = v.UniqueID
			policyrole.Value = v.Value
			err := s.Daos.UpdatePolicyRuleWithUpsert(ctx, policyrole)
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

// EnableLeavePolicy : ""
func (s *Service) EnableLeavePolicy(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableLeavePolicy(ctx, uniqueID)
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

// DisableLeavePolicy : ""
func (s *Service) DisableLeavePolicy(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableLeavePolicy(ctx, uniqueID)
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

//DeleteLeavePolicy : ""
func (s *Service) DeleteLeavePolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteLeavePolicy(ctx, UniqueID)
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

// FilterLeavePolicy : ""
func (s *Service) FilterLeavePolicy(ctx *models.Context, leavepolicy *models.FilterLeavePolicy, pagination *models.Pagination) (leavepolicys []models.RefLeavePolicy, err error) {
	err = s.LeavePolicyDataAccess(ctx, leavepolicy)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterLeavePolicy(ctx, leavepolicy, pagination)
}
func (s *Service) LeavePolicyDataAccess(ctx *models.Context, filter *models.FilterLeavePolicy) (err error) {
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
