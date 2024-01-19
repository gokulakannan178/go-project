package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveWardWiseHouseVisited : ""
func (s *Service) SaveWardWiseHouseVisited(ctx *models.Context, Dept *models.WardWiseHouseVisited) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Dept.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONWARDWISEHOUSEVISITED)
	Dept.Status = constants.WARDWISEHOUSEVISITEDSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 WardWiseHouseVisited.created")
	Dept.Created = &created
	log.Println("b4 WardWiseHouseVisited.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveWardWiseHouseVisited(ctx, Dept)
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

// GetSingleWardWiseHouseVisited : ""
func (s *Service) GetSingleWardWiseHouseVisited(ctx *models.Context, UniqueID string) (*models.RefWardWiseHouseVisited, error) {
	WardWiseHouseVisited, err := s.Daos.GetSingleWardWiseHouseVisited(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return WardWiseHouseVisited, nil
}

func (s *Service) UpdateWardWiseHouseVisited(ctx *models.Context, dept *models.WardWiseHouseVisited) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateWardWiseHouseVisited(ctx, dept)
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

// EnableWardWiseHouseVisited : ""
func (s *Service) EnableWardWiseHouseVisited(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableWardWiseHouseVisited(ctx, uniqueID)
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

// DisableWardWiseHouseVisited : ""
func (s *Service) DisableWardWiseHouseVisited(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableWardWiseHouseVisited(ctx, uniqueID)
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

//DeleteState : ""
func (s *Service) DeleteWardWiseHouseVisited(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteWardWiseHouseVisited(ctx, UniqueID)
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

// FilterWardWiseHouseVisited : ""
func (s *Service) FilterWardWiseHouseVisited(ctx *models.Context, Filter *models.FilterWardWiseHouseVisited, pagination *models.Pagination) (Dept []models.RefWardWiseHouseVisited, err error) {
	return s.Daos.FilterWardWiseHouseVisited(ctx, Filter, pagination)
}
