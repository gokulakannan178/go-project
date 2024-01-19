package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveFarmerLand :""
func (s *Service) SaveFarmerLand(ctx *models.Context, FarmerLand *models.FarmerLand) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//FarmerLand.Code = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFarmerLand)

	FarmerLand.Status = constants.FARMERLANDSTATUSACTIVE
	FarmerLand.ActiveStatus = true
	t := time.Now()
	FarmerLand.KhasraNumber = FarmerLand.ParcelNumber
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 FarmerLand.created")
	FarmerLand.Created = created
	log.Println("b4 FarmerLand.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFarmerLand(ctx, FarmerLand)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		err := s.Daos.SumFarmerLandArea(ctx, FarmerLand.Farmer.Hex())
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

//UpdateFarmerLand : ""
func (s *Service) UpdateFarmerLand(ctx *models.Context, FarmerLand *models.FarmerLand) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFarmerLand(ctx, FarmerLand)
		if err != nil {

			return errors.New("Db Error" + err.Error())
		}
		farmer, err := s.Daos.GetSingleFarmerLand(ctx, FarmerLand.ID.Hex())
		if err != nil {
			return err
		}
		err = s.Daos.SumFarmerLandArea(ctx, farmer.Farmer.Hex())
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

//EnableFarmerLand : ""
func (s *Service) EnableFarmerLand(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFarmerLand(ctx, UniqueID)
		if err != nil {

			return errors.New("Db Error" + err.Error())
		}
		farmer, err := s.Daos.GetSingleFarmerLand(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.SumFarmerLandArea(ctx, farmer.Farmer.Hex())
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

//DisableFarmerLand : ""
func (s *Service) DisableFarmerLand(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFarmerLand(ctx, UniqueID)
		if err != nil {

			return errors.New("Db Error" + err.Error())
		}
		farmer, err := s.Daos.GetSingleFarmerLand(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.SumFarmerLandArea(ctx, farmer.Farmer.Hex())
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

//DeleteFarmerLand : ""
func (s *Service) DeleteFarmerLand(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFarmerLand(ctx, UniqueID)
		if err != nil {

			return errors.New("Db Error" + err.Error())
		}
		farmer, err := s.Daos.GetSingleFarmerLand(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.SumFarmerLandArea(ctx, farmer.Farmer.Hex())
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

//GetSingleFarmerLand :""
func (s *Service) GetSingleFarmerLand(ctx *models.Context, UniqueID string) (*models.RefFarmerLand, error) {
	FarmerLand, err := s.Daos.GetSingleFarmerLand(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return FarmerLand, nil
}

//FilterFarmerLand :""
func (s *Service) FilterFarmerLand(ctx *models.Context, FarmerLandfilter *models.FarmerLandFilter, pagination *models.Pagination) (FarmerLand []models.RefFarmerLand, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterFarmerLand(ctx, FarmerLandfilter, pagination)

}
