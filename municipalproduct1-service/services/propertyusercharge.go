package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyUserCharge : ""
func (s *Service) SavePropertyUserCharge(ctx *models.Context, propertyusercharge *models.PropertyUserCharge) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	propertyusercharge.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYUSERCHARGE)
	propertyusercharge.Status = constants.PROPERTYUSERCHARGESTATUSACTIVE
	//t := time.Now()
	//PropertyUserCharge.Created = new(models.CreatedV2)
	// propertyusercharge.Created.On = &t
	//propertyusercharge.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		// floor, err := s.Daos.GetPropertyDAOUsingPropertyID(ctx, propertyusercharge.PropertyID)
		// if err != nil {
		// 	return err
		// }
		//t2 := time.Date(2023, 4, 1, 9, 0, 0, 0, t.Location())
		propertyusercharge.IsUserCharge = "Yes"
		propertyusercharge.Createdby.On = &t
		propertyusercharge.Status = constants.PROPERTYUSERCHARGESTATUSINIT
		err := s.Daos.CreateUserChargeForProperty(ctx, propertyusercharge)
		if err != nil {
			return err
		}
		dberr := s.Daos.SavePropertyUserCharge(ctx, propertyusercharge)
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

//GetSinglePropertyUserCharge :""
func (s *Service) GetSinglePropertyUserCharge(ctx *models.Context, UniqueID string) (*models.RefPropertyUserCharge, error) {
	tower, err := s.Daos.GetSinglePropertyUserCharge(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdatePropertyUserCharge : ""
func (s *Service) UpdatePropertyUserCharge(ctx *models.Context, propertyusercharge *models.PropertyUserCharge) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()

		usercharge, err := s.Daos.GetSinglePropertyWithUserCharge(ctx, propertyusercharge.PropertyID)
		if err != nil {
			return err
		}
		fmt.Println("usercharge=================>", usercharge)
		// floor, err := s.Daos.GetPropertyDAOUsingPropertyID(ctx, propertyusercharge.PropertyID)
		// if err != nil {
		// 	return err
		// }
		//t2 := time.Date(2023, 4, 1, 9, 0, 0, 0, t.Location())
		//propertyusercharge.DOA = &t2
		//propertyusercharge.DOA = floor.DateFrom
		propertyusercharge.IsUserCharge = "Yes"
		propertyusercharge.Createdby.On = &t
		propertyusercharge.Status = constants.PROPERTYUSERCHARGESTATUSINIT
		err = s.Daos.CreateUserChargeForProperty(ctx, propertyusercharge)
		if err != nil {
			return err
		}
		userchargeupdatelog := new(models.UserChargeUpdateLog)
		userchargeupdatelog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERCHARGEUPDATELOG)
		userchargeupdatelog.PropertyId = usercharge.UniqueID
		userchargeupdatelog.Date = &t
		userchargeupdatelog.Status = "Active"
		userchargeupdatelog.BeforeUserCharge = usercharge.UserCharge
		userchargeupdatelog.AfterUserCharge.CategoryID = propertyusercharge.CategoryID
		userchargeupdatelog.AfterUserCharge.DOA = propertyusercharge.DOA
		userchargeupdatelog.AfterUserCharge.IsUserCharge = propertyusercharge.IsUserCharge
		userchargeupdatelog.AfterUserCharge.Status = propertyusercharge.Status
		userchargeupdatelog.AfterUserCharge.Createdby.On = &t
		userchargeupdatelog.AfterUserCharge.Createdby.By = propertyusercharge.Createdby.By
		userchargeupdatelog.AfterUserCharge.Createdby.ByType = propertyusercharge.Createdby.ByType
		err = s.Daos.SaveUserChargeUpdateLog(ctx, userchargeupdatelog)
		if err != nil {
			return err
		}
		err = s.Daos.UpdatePropertyUserCharge(ctx, propertyusercharge)
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

// EnablePropertyUserCharge : ""
func (s *Service) EnablePropertyUserCharge(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyUserCharge(ctx, UniqueID)
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

//DisablePropertyUserCharge : ""
func (s *Service) DisablePropertyUserCharge(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyUserCharge(ctx, UniqueID)
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

//DeletePropertyUserCharge : ""
func (s *Service) DeletePropertyUserCharge(ctx *models.Context, rp *models.UserChargeAction) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		usercharge, err := s.Daos.GetSinglePropertyWithUserCharge(ctx, rp.UniqueId)
		if err != nil {
			return err
		}
		t := time.Now()
		//fmt.Println("usercharge=================>", usercharge)
		userchargeupdatelog := new(models.UserChargeLog)
		userchargeupdatelog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERCHARGEUPDATELOG)
		userchargeupdatelog.PropertyId = usercharge.UniqueID
		userchargeupdatelog.Date = &t
		userchargeupdatelog.Status = "Deleted"
		userchargeupdatelog.BeforeUserCharge = usercharge.UserCharge
		userchargeupdatelog.AfterUserCharge.CategoryID = usercharge.UserCharge.CategoryID
		userchargeupdatelog.AfterUserCharge.DOA = usercharge.UserCharge.DOA
		userchargeupdatelog.AfterUserCharge.IsUserCharge = usercharge.UserCharge.IsUserCharge
		userchargeupdatelog.AfterUserCharge.Status = "Deleted"
		userchargeupdatelog.AfterUserCharge.Createdby.On = &t
		userchargeupdatelog.AfterUserCharge.Createdby.By = rp.By
		userchargeupdatelog.AfterUserCharge.Createdby.ByType = rp.ByType
		err = s.Daos.SaveUserChargeLog(ctx, userchargeupdatelog)
		if err != nil {
			return err
		}

		err = s.Daos.DeletePropertyUserCharge(ctx, rp)
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

// FilterPropertyUserCharge : ""
func (s *Service) FilterPropertyUserCharge(ctx *models.Context, filter *models.PropertyUserChargeFilter, pagination *models.Pagination) ([]models.RefPropertyUserCharge, error) {
	return s.Daos.FilterPropertyUserCharge(ctx, filter, pagination)

}

// RejectPayment : ""
func (s *Service) RejectPropertyUserCharge(ctx *models.Context, rp *models.UserChargeAction) (string, error) {
	t := time.Now()
	rp.ActionDate = &t
	if rp.Date == nil {
		rp.Date = &t
	}
	usercharge, err := s.Daos.GetSinglePropertyWithUserCharge(ctx, rp.UniqueId)
	if err != nil {
		return "", err
	}
	fmt.Println("usercharge=================>", usercharge)
	userchargeupdatelog := new(models.UserChargeLog)
	userchargeupdatelog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERCHARGEUPDATELOG)
	userchargeupdatelog.PropertyId = usercharge.UniqueID
	userchargeupdatelog.Date = &t
	userchargeupdatelog.Status = "Rejected"
	userchargeupdatelog.BeforeUserCharge = usercharge.UserCharge
	userchargeupdatelog.AfterUserCharge.CategoryID = usercharge.UserCharge.CategoryID
	userchargeupdatelog.AfterUserCharge.DOA = usercharge.UserCharge.DOA
	userchargeupdatelog.AfterUserCharge.IsUserCharge = usercharge.UserCharge.IsUserCharge
	userchargeupdatelog.AfterUserCharge.Status = "Rejected"
	userchargeupdatelog.AfterUserCharge.Createdby.On = &t
	userchargeupdatelog.AfterUserCharge.Createdby.By = rp.By
	userchargeupdatelog.AfterUserCharge.Createdby.ByType = rp.ByType
	err = s.Daos.SaveUserChargeLog(ctx, userchargeupdatelog)
	if err != nil {
		return "", err
	}
	err = s.Daos.RejectPropertyUserCharge(ctx, rp)
	if err != nil {
		return "", err
	}
	//t := time.Now()

	propertypayment, err := s.Daos.GetSingleProperty(ctx, rp.UniqueId)
	if err != nil {
		return "", err
	}
	return propertypayment.UniqueID, err
}

func (s *Service) VerifyPropertyUserCharge(ctx *models.Context, vp *models.UserChargeAction) (string, error) {
	t := time.Now()
	vp.ActionDate = &t
	if vp.Date == nil {
		vp.Date = &t
	}
	usercharge, err := s.Daos.GetSinglePropertyWithUserCharge(ctx, vp.UniqueId)
	if err != nil {
		return "", err
	}
	fmt.Println("usercharge=================>", usercharge)
	userchargeupdatelog := new(models.UserChargeLog)
	userchargeupdatelog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERCHARGEUPDATELOG)
	userchargeupdatelog.PropertyId = usercharge.UniqueID
	userchargeupdatelog.Date = &t
	userchargeupdatelog.Status = "Active"
	userchargeupdatelog.BeforeUserCharge = usercharge.UserCharge
	userchargeupdatelog.AfterUserCharge.CategoryID = usercharge.UserCharge.CategoryID
	userchargeupdatelog.AfterUserCharge.DOA = usercharge.UserCharge.DOA
	userchargeupdatelog.AfterUserCharge.IsUserCharge = usercharge.UserCharge.IsUserCharge
	userchargeupdatelog.AfterUserCharge.Status = "Active"
	userchargeupdatelog.AfterUserCharge.Createdby.On = &t
	userchargeupdatelog.AfterUserCharge.Createdby.By = vp.By
	userchargeupdatelog.AfterUserCharge.Createdby.ByType = vp.ByType
	err = s.Daos.SaveUserChargeLog(ctx, userchargeupdatelog)
	if err != nil {
		return "", err
	}
	err = s.Daos.VerifyPropertyUserCharge(ctx, vp)
	if err != nil {
		return "", err
	}

	propertypayment, err := s.Daos.GetSingleProperty(ctx, vp.UniqueId)
	if err != nil {
		return "", err
	}
	return propertypayment.UniqueID, err
}
