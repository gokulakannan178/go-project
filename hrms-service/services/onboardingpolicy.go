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

// SaveOnboardingPolicy : ""
func (s *Service) SaveOnboardingPolicy(ctx *models.Context, onboardingpolicy *models.OnboardingPolicy) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	onboardingpolicy.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONONBOARDINGPOLICY)
	onboardingpolicy.Status = constants.ONBOARDINGPOLICYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OnboardingPolicy.created")
	onboardingpolicy.Created = &created
	log.Println("b4 OnboardingPolicy.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOnboardingPolicy(ctx, onboardingpolicy)
		if dberr != nil {
			return dberr
		}

		for _, v := range onboardingpolicy.OnboardCheckListMasterId {
			onboardChecklist := new(models.OnboardingCheckList)
			onboardingchecklistmaster, err := s.Daos.GetSingleOnboardingCheckListMasterWithActive(ctx, v, constants.ONBOARDINGCHECKLISTMASTERSTATUSACTIVE)
			if err != nil {
				fmt.Println(err)
				return err
			}

			fmt.Println("onboardingchecklistmaster=======", onboardingchecklistmaster)

			onboardChecklist.OnboardingpolicyID = onboardingpolicy.UniqueID

			onboardChecklist.OnboardingchecklistmasterID = v
			onboardChecklist.Name = onboardingpolicy.Name

			err = s.SaveOnboardingCheckListWithoutTransaction(ctx, onboardChecklist)
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

// GetSingleOnboardingPolicy : ""
func (s *Service) GetSingleOnboardingPolicy(ctx *models.Context, UniqueID string) (*models.RefOnboardingPolicy, error) {
	onboardingPolicy, err := s.Daos.GetSingleOnboardingPolicy(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return onboardingPolicy, nil
}

//UpdateOnboardingPolicy : ""
func (s *Service) UpdateOnboardingPolicy(ctx *models.Context, onboardingPolicy *models.OnboardingPolicy) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.OnboardingCheckListRemoveNotPresentValue(ctx, onboardingPolicy.UniqueID, onboardingPolicy.OnboardCheckListMasterId)
		if err != nil {
			return err
		}
		err = s.Daos.OnboardingCheckListUpsert(ctx, onboardingPolicy.UniqueID, onboardingPolicy.OnboardCheckListMasterId, onboardingPolicy.Name)
		if err != nil {
			return err
		}

		fmt.Println("error==>", err)

		err = s.Daos.UpdateOnboardingPolicy(ctx, onboardingPolicy)
		if err != nil {
			return err
		}
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

// EnableOnboardingPolicy : ""
func (s *Service) EnableOnboardingPolicy(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableOnboardingPolicy(ctx, uniqueID)
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

// DisableOnboardingPolicy : ""
func (s *Service) DisableOnboardingPolicy(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableOnboardingPolicy(ctx, uniqueID)
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

//DeleteOnboardingPolicy : ""
func (s *Service) DeleteOnboardingPolicy(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOnboardingPolicy(ctx, UniqueID)
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

// FilterOnboardingPolicy : ""
func (s *Service) FilterOnboardingPolicy(ctx *models.Context, onboardingpolicy *models.FilterOnboardingPolicy, pagination *models.Pagination) (onboardingpolicys []models.RefOnboardingPolicy, err error) {
	err = s.OnboardingPolicyDataAccess(ctx, onboardingpolicy)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterOnboardingPolicy(ctx, onboardingpolicy, pagination)
}

// func (s *Service) OnboardingPolicy(ctx *models.Context, onboardingPolicy *models.OnboardingPolicy) error {
// 	err := s.SaveOnboardingPolicy(ctx, onboardingPolicy)
// 	if err != nil {
// 		return err
// 	}
// 	onboardChecklist := new(models.OnboardingCheckList)
// 	refOnboardingPolicy, err := s.Daos.GetSingleOnboardingPolicy(ctx, onboardingPolicy.UniqueID)
// 	if err != nil {
// 		return err
// 	}
// 	// var uniqueID []string
// 	for _, v := range onboardingPolicy.OnboardingPolicy {
// 		if refOnboardingPolicy != nil {
// 			onboardChecklist.OnboardingpolicyID = refOnboardingPolicy.UniqueID
// 		}
// 		//onboardChecklist.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONONBOARDINGPOLICY)

// 		onboardChecklist.OnboardingchecklistmasterID = v
// 		err = s.SaveOnboardingCheckList(ctx, onboardChecklist)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
func (s *Service) OnboardingPolicyDataAccess(ctx *models.Context, filter *models.FilterOnboardingPolicy) (err error) {
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
