package services

import (
	"errors"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveBillclaimLevels :""
func (s *Service) SaveBillclaimLevels(ctx *models.Context, billclaimlevels *models.BillclaimLevels) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	billclaimlevels.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIMLEVELS)
	billclaimlevels.Status = constants.BILLCLAIMLEVELSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 BillclaimLevels.created")
	billclaimlevels.Created = created
	log.Println("b4 BillclaimLevels.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveBillclaimLevels(ctx, billclaimlevels)
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
func (s *Service) SaveBillclaimLevelsWithoutTransaction(ctx *models.Context, billclaimlevels *models.BillclaimLevels) error {

	defer ctx.Session.EndSession(ctx.CTX)
	billclaimlevels.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIMLEVELS)
	t := time.Now()
	billclaimlevels.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 BillclaimLevels.created")
	billclaimlevels.Created = created
	log.Println("b4 BillclaimLevels.created")

	dberr := s.Daos.SaveBillclaimLevels(ctx, billclaimlevels)
	if dberr != nil {
		return dberr
	}

	return nil
}

//GetSingleBillclaimLevels :""
func (s *Service) GetSingleBillclaimLevels(ctx *models.Context, UniqueID string) (*models.RefBillclaimLevels, error) {
	billclaimlevels, err := s.Daos.GetSingleBillclaimLevels(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return billclaimlevels, nil
}

//UpdateBillclaimLevels : ""
func (s *Service) UpdateBillclaimLevels(ctx *models.Context, billclaimlevels *models.BillclaimLevels) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.UpdateBillclaimLevels(ctx, billclaimlevels)
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

//EnableBillclaimLevels : ""
func (s *Service) EnableBillclaimLevels(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableBillclaimLevels(ctx, UniqueID)
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

//DisableBillclaimLevels : ""
func (s *Service) DisableBillclaimLevels(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableBillclaimLevels(ctx, UniqueID)
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

//DeleteBillclaimLevels : ""
func (s *Service) DeleteBillclaimLevels(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteBillclaimLevels(ctx, UniqueID)
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

//FilterBillclaimLevels :""
func (s *Service) FilterBillclaimLevels(ctx *models.Context, billclaimlevelsFilter *models.BillclaimLevelsFilter, pagination *models.Pagination) ([]models.RefBillclaimLevels, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err := s.BillclaimLevelsDataAccess(ctx, billclaimlevelsFilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterBillclaimLevels(ctx, billclaimlevelsFilter, pagination)

}
func (s *Service) BillclaimLevelsDataAccess(ctx *models.Context, filter *models.BillclaimLevelsFilter) (err error) {
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
			if dataaccess.SuperAdmin == false {
				filter.AssignedBy = append(filter.AssignedBy, filter.DataAccess.UserName)
			}

		}

	}
	return err
}
func (s *Service) ApprovedBillclaimLevels(ctx *models.Context, billclaimlevels *models.BillclaimLevels) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.ApprovedBillclaimLevels(ctx, billclaimlevels)
		if dberr != nil {
			return dberr
		}
		bills, err := s.Daos.GetSingleBillclaimLevels(ctx, billclaimlevels.UniqueID)
		if err != nil {
			return err
		}
		linemanager, err := s.Daos.GetSingleEmployeeWithUserName(ctx, bills.AssignedBy)
		if err != nil {
			return err
		}
		nextlevel, err := s.Daos.GetSingleBillclaimLevelsWithAssigned(ctx, bills.Bill, linemanager.LineManager, bills.Level+1)
		if err != nil {
			return err
		}
		err = s.Daos.PendingBillclaimLevels(ctx, nextlevel.UniqueID)
		if err != nil {
			return err
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
