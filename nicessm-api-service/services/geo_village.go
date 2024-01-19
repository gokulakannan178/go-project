package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveVillage :""
func (s *Service) SaveVillage(ctx *models.Context, village *models.Village) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	village.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVILLAGE)
	village.ActiveStatus = true
	village.Status = constants.VILLAGESTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 village.created")
	village.Created = created
	log.Println("b4 village.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveVillage(ctx, village)
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

//UpdateVillage : ""
func (s *Service) UpdateVillage(ctx *models.Context, village *models.Village) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVillage(ctx, village)
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

//EnableVillage : ""
func (s *Service) EnableVillage(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableVillage(ctx, UniqueID)
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

//DisableVillage : ""
func (s *Service) DisableVillage(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableVillage(ctx, UniqueID)
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

//DeleteVillage : ""
func (s *Service) DeleteVillage(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVillage(ctx, UniqueID)
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

//GetSingleVillage :""
func (s *Service) GetSingleVillage(ctx *models.Context, UniqueID string) (*models.RefVillage, error) {
	village, err := s.Daos.GetSingleVillage(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return village, nil
}

//FilterVillage :""
func (s *Service) FilterVillage(ctx *models.Context, villagefilter *models.VillageFilter, pagination *models.Pagination) (village []models.RefVillage, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	if villagefilter != nil {

		dataaccess, err := s.Daos.DataAccess(ctx, &villagefilter.DataAccess)
		if err != nil {
			return nil, err
		}
		if dataaccess != nil {
			if len(dataaccess.AccessVillages) > 0 {
				for _, v := range dataaccess.AccessVillages {
					villagefilter.ID = append(villagefilter.ID, v.ID)
				}
			}

		}
	}
	return s.Daos.FilterVillage(ctx, villagefilter, pagination)
}
func (s *Service) SaveVillageWithGrampanchatBlock(ctx *models.Context, village *models.NeWAddVillage) (*models.NeWAddVillage, error) {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}

	defer ctx.Session.EndSession(ctx.CTX)
	//village.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVILLAGE)
	Villages := new(models.Village)
	ActiveStatus := true
	Status := constants.VILLAGESTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 village.created")

	log.Println("b4 village.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		block := new(models.Block)
		GramPanchayat := new(models.GramPanchayat)
		if village.Block.IsZero() {
			block.District = village.District
			block.Name = village.BlockName
			block.Status = Status
			block.ActiveStatus = ActiveStatus
			block.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBLOCK)
			block.Created = created
			err := s.Daos.SaveBlock(ctx, block)
			if err != nil {
				return err
			}
			village.Block = block.ID
			GramPanchayat.Block = block.ID
		} else {
			GramPanchayat.Block = village.Block
		}
		fmt.Println("blockId===>", GramPanchayat.Block)
		if village.GramPanchayat.IsZero() {
			GramPanchayat.Name = village.GramPanchayatName
			GramPanchayat.Status = Status
			GramPanchayat.ActiveStatus = ActiveStatus
			GramPanchayat.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONGRAMPANCHAYAT)
			err := s.Daos.SaveGramPanchayat(ctx, GramPanchayat)
			if err != nil {
				return err
			}
			village.GramPanchayat = GramPanchayat.ID
			Villages.GramPanchayat = GramPanchayat.ID
		} else {
			Villages.GramPanchayat = village.GramPanchayat

		}
		fmt.Println("GramPanchayatId===>", village.GramPanchayat)

		Villages.ActiveStatus = ActiveStatus
		Villages.Status = Status
		Villages.Name = village.VillageName
		Villages.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVILLAGE)
		dberr := s.Daos.SaveVillage(ctx, Villages)
		if dberr != nil {
			return dberr
		}
		village.Village = Villages.ID
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return nil, errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return nil, err
	}
	return village, nil
}
