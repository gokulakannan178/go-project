package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveCircleWiseHouseVisited : ""
func (s *Service) SaveCircleWiseHouseVisited(ctx *models.Context, Dept *models.CircleWiseHouseVisited) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Dept.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCIRCLEWISEHOUSEVISITED)
	Dept.Status = constants.CIRCLEWISEHOUSEVISITEDSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 CircleWiseHouseVisited.created")
	Dept.Created = &created
	log.Println("b4 CircleWiseHouseVisited.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveCircleWiseHouseVisited(ctx, Dept)
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

// GetSingleCircleWiseHouseVisited : ""
func (s *Service) GetSingleCircleWiseHouseVisited(ctx *models.Context, UniqueID string) (*models.RefCircleWiseHouseVisited, error) {
	CircleWiseHouseVisited, err := s.Daos.GetSingleCircleWiseHouseVisited(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return CircleWiseHouseVisited, nil
}

func (s *Service) UpdateCircleWiseHouseVisited(ctx *models.Context, dept *models.CircleWiseHouseVisited) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateCircleWiseHouseVisited(ctx, dept)
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

// EnableCircleWiseHouseVisited : ""
func (s *Service) EnableCircleWiseHouseVisited(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableCircleWiseHouseVisited(ctx, uniqueID)
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

// DisableCircleWiseHouseVisited : ""
func (s *Service) DisableCircleWiseHouseVisited(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableCircleWiseHouseVisited(ctx, uniqueID)
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
func (s *Service) DeleteCircleWiseHouseVisited(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteCircleWiseHouseVisited(ctx, UniqueID)
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

// FilterCircleWiseHouseVisited : ""
func (s *Service) FilterCircleWiseHouseVisited(ctx *models.Context, Filter *models.FilterCircleWiseHouseVisited, pagination *models.Pagination) (Dept []models.RefCircleWiseHouseVisited, err error) {
	return s.Daos.FilterCircleWiseHouseVisited(ctx, Filter, pagination)
}
