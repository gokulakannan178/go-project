package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveBillclaimConfig :""
func (s *Service) SaveBillclaimConfig(ctx *models.Context, billclaimconfig *models.BillclaimConfig) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	billclaimconfig.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIMCONFIG)
	billclaimconfig.Status = constants.BILLCLAIMCONFIGSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 BillclaimConfig.created")
	billclaimconfig.Created = created
	log.Println("b4 BillclaimConfig.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		checkbillconfig, err := s.Daos.GetSingleBillclaimConfigWithLevel(ctx, billclaimconfig.Grade, billclaimconfig.Level)
		if err != nil {
			return err
		}
		if checkbillconfig != nil {
			err := s.Daos.DisableBillclaimConfig(ctx, checkbillconfig.UniqueID)
			if err != nil {
				return err
			}
		}
		dberr := s.Daos.SaveBillclaimConfig(ctx, billclaimconfig)
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

//GetSingleBillclaimConfig :""
func (s *Service) GetSingleBillclaimConfig(ctx *models.Context, UniqueID string) (*models.RefBillclaimConfig, error) {
	billclaimconfig, err := s.Daos.GetSingleBillclaimConfig(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return billclaimconfig, nil
}

//UpdateBillclaimConfig : ""
func (s *Service) UpdateBillclaimConfig(ctx *models.Context, billclaimconfig *models.BillclaimConfig) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.UpdateBillclaimConfig(ctx, billclaimconfig)
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

//EnableBillclaimConfig : ""
func (s *Service) EnableBillclaimConfig(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableBillclaimConfig(ctx, UniqueID)
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

//DisableBillclaimConfig : ""
func (s *Service) DisableBillclaimConfig(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableBillclaimConfig(ctx, UniqueID)
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

//DeleteBillclaimConfig : ""
func (s *Service) DeleteBillclaimConfig(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteBillclaimConfig(ctx, UniqueID)
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

//FilterBillclaimConfig :""
func (s *Service) FilterBillclaimConfig(ctx *models.Context, billclaimconfigFilter *models.BillclaimConfigFilter, pagination *models.Pagination) ([]models.RefBillclaimConfig, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err := s.BillclaimConfigDataAccess(ctx, billclaimconfigFilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterBillclaimConfig(ctx, billclaimconfigFilter, pagination)

}
func (s *Service) BillclaimConfigDataAccess(ctx *models.Context, filter *models.BillclaimConfigFilter) (err error) {
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
			// if dataaccess.Branch != "" {
			// 	filter.UniqueID = append(filter.UniqueID, dataaccess.Branch)

			// }

		}

	}
	return err
}
