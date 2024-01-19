package services

import (
	"errors"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePkgType :""
func (s *Service) SavePkgType(ctx *models.Context, pkg *models.PkgType) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	pkg.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPKG)
	pkg.Status = constants.PRODUCTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PkgType.created")
	pkg.Created = &created
	log.Println("b4 PkgType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePkgType(ctx, pkg)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}

	return nil
}

//UpdatePkgType : ""
func (s *Service) UpdatePkgType(ctx *models.Context, pkg *models.PkgType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePkgType(ctx, pkg)
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

//EnablePkgType : ""
func (s *Service) EnablePkgType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePkgType(ctx, UniqueID)
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

//DisablePkgType : ""
func (s *Service) DisablePkgType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePkgType(ctx, UniqueID)
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

//DeletePkgType : ""
func (s *Service) DeletePkgType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePkgType(ctx, UniqueID)
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

//GetSinglePkgType :""
func (s *Service) GetSinglePkgType(ctx *models.Context, UniqueID string) (*models.RefPkgType, error) {
	PkgType, err := s.Daos.GetSinglePkgType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return PkgType, nil
}

//FilterPkgType :""
func (s *Service) FilterPkgType(ctx *models.Context, filter *models.PkgTypeFilter, pagination *models.Pagination) ([]models.RefPkgType, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPkgType(ctx, filter, pagination)

}

//GetSinglePkgType :""
func (s *Service) GetDefaultPkgType(ctx *models.Context) (*models.RefPkgType, error) {
	pkg, err := s.Daos.GetDefaultPkgType(ctx)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}
