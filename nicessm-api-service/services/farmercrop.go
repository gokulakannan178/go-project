package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveFarmerCrop :""
func (s *Service) SaveFarmerCrop(ctx *models.Context, farmerCrop *models.FarmerCrop) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	// FarmerCrop.Name = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFarmerCrop)
	farmerCrop.Status = constants.FARMERCROPSTATUSWIP
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 FarmerCrop.created")
	farmerCrop.Created = &created
	log.Println("b4 FarmerCrop.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFarmerCrop(ctx, farmerCrop)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		err := s.Daos.GetFarmerCropCount(ctx, farmerCrop.Farmer.Hex())
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

//UpdateFarmerCrop : ""
func (s *Service) UpdateFarmerCrop(ctx *models.Context, farmerCrop *models.FarmerCrop) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFarmerCrop(ctx, farmerCrop)
		if err != nil {
			return err
		}
		farmer, err := s.GetSingleFarmerCrop(ctx, farmerCrop.ID.Hex())
		if err != nil {
			return err
		}
		err = s.Daos.GetFarmerCropCount(ctx, farmer.Farmer.Hex())
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

//EnableFarmerCrop : ""
func (s *Service) EnableFarmerCrop(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFarmerCrop(ctx, UniqueID)
		if err != nil {
			return err
		}
		farmer, err := s.GetSingleFarmerCrop(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.GetFarmerCropCount(ctx, farmer.Farmer.Hex())
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

//DisableFarmerCrop : ""
func (s *Service) DisableFarmerCrop(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFarmerCrop(ctx, UniqueID)
		if err != nil {
			return err
		}
		farmer, err := s.GetSingleFarmerCrop(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.GetFarmerCropCount(ctx, farmer.Farmer.Hex())
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

//DeleteFarmerCrop : ""
func (s *Service) DeleteFarmerCrop(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFarmerCrop(ctx, UniqueID)
		if err != nil {
			return err
		}
		farmer, err := s.GetSingleFarmerCrop(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.GetFarmerCropCount(ctx, farmer.Farmer.Hex())
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

//GetSingleFarmerCrop :""
func (s *Service) GetSingleFarmerCrop(ctx *models.Context, UniqueID string) (*models.RefFarmerCrop, error) {
	farmerCrop, err := s.Daos.GetSingleFarmerCrop(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return farmerCrop, nil
}

//FilterFarmerCrop :""
func (s *Service) FilterFarmerCrop(ctx *models.Context, farmerCropfilter *models.FarmerCropFilter, pagination *models.Pagination) (farmerCrop []models.RefFarmerCrop, err error) {
	return s.Daos.FilterFarmerCrop(ctx, farmerCropfilter, pagination)
}

//UpdateFarmerCropDone : ""
func (s *Service) UpdateFarmerCropDone(ctx *models.Context, farmerCrop *models.FarmerCrop) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFarmerCropDone(ctx, farmerCrop)
		if err != nil {
			return err
		}
		// farmer, err := s.GetSingleFarmerCrop(ctx, farmerCrop.ID.Hex())
		// if err != nil {
		// 	return err
		// }
		err = s.Daos.GetFarmerCropCount(ctx, farmerCrop.Farmer.Hex())
		if err != nil {
			return errors.New("farmer crop count not updated")
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
