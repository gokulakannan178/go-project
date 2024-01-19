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

// SaveOffboardingPolicy : ""
func (s *Service) SaveOffboardingPolicy(ctx *models.Context, offboardingpolicy *models.OffboardingPolicy) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	offboardingpolicy.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONOFFBOARDINGPOLICY)
	offboardingpolicy.Status = constants.OFFBOARDINGPOLICYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OffboardingPolicy.created")
	offboardingpolicy.Created = &created
	log.Println("b4 OffboardingPolicy.created")
	//log.Println("offboardingpolicy", offboardingpolicy)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOffboardingPolicy(ctx, offboardingpolicy)
		if dberr != nil {
			return dberr
		}
		refOffboardingPolicy, err := s.Daos.GetSingleOffboardingPolicy(ctx, offboardingpolicy.UniqueID)
		if err != nil {
			return err
		}
		for _, v := range offboardingpolicy.OffboardCheckListMasterId {
			offboardChecklist := new(models.OffboardingCheckList)
			offboardingchecklistmaster, err := s.Daos.GetSingleOffboardingCheckListMasterWithActive(ctx, v, constants.OFFBOARDINGCHECKLISTMASTERSTATUSACTIVE)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println("Offboardingchecklistmaster", offboardingchecklistmaster)
			if refOffboardingPolicy != nil {
				offboardChecklist.OffboardingpolicyID = refOffboardingPolicy.UniqueID
			}
			offboardChecklist.OffboardingchecklistmasterID = v
			offboardChecklist.Name = offboardingpolicy.Name
			err = s.SaveOffboardingCheckListWithoutTransaction(ctx, offboardChecklist)
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

// GetSingleOffboardingPolicy : ""
func (s *Service) GetSingleOffboardingPolicy(ctx *models.Context, UniqueID string) (*models.RefOffboardingPolicy, error) {
	OffboardingPolicy, err := s.Daos.GetSingleOffboardingPolicy(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return OffboardingPolicy, nil
}

//UpdateOffboardingPolicy : ""
func (s *Service) UpdateOffboardingPolicy(ctx *models.Context, OffboardingPolicy *models.OffboardingPolicy) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.OffboardingCheckListRemoveNotPresentValue(ctx, OffboardingPolicy.UniqueID, OffboardingPolicy.OffboardCheckListMasterId)
		if err != nil {
			return err
		}
		err = s.Daos.OffboardingCheckListUpsert(ctx, OffboardingPolicy.UniqueID, OffboardingPolicy.OffboardCheckListMasterId, OffboardingPolicy.Name)
		if err != nil {
			return err
		}

		fmt.Println("error==>", err)

		err = s.Daos.UpdateOffboardingPolicy(ctx, OffboardingPolicy)
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

// EnableOffboardingPolicy : ""
func (s *Service) EnableOffboardingPolicy(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableOffboardingPolicy(ctx, uniqueID)
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

// DisableOffboardingPolicy : ""
func (s *Service) DisableOffboardingPolicy(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableOffboardingPolicy(ctx, uniqueID)
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

//DeleteOffboardingPolicy : ""
func (s *Service) DeleteOffboardingPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOffboardingPolicy(ctx, UniqueID)
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

// FilterOffboardingPolicy : ""
func (s *Service) FilterOffboardingPolicy(ctx *models.Context, offboardingpolicy *models.FilterOffboardingPolicy, pagination *models.Pagination) (offboardingpolicys []models.RefOffboardingPolicy, err error) {
	err = s.OffboardingPolicyDataAccess(ctx, offboardingpolicy)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterOffboardingPolicy(ctx, offboardingpolicy, pagination)
}
func (s *Service) OffboardingPolicyDataAccess(ctx *models.Context, filter *models.FilterOffboardingPolicy) (err error) {
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
