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

//SaveFeedBack :""
func (s *Service) SaveFeedBack(ctx *models.Context, FeedBack *models.FeedBack) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//FeedBack.Code = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFeedBack)

	FeedBack.Status = constants.FEEDBACKSTATUSACTIVE
	FeedBack.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	FeedBack.Date = &t
	created.By = constants.SYSTEM
	log.Println("b4 FeedBack.created")
	FeedBack.Created = created
	log.Println("b4 FeedBack.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFeedBack(ctx, FeedBack)
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

//UpdateFeedBack : ""
func (s *Service) UpdateFeedBack(ctx *models.Context, FeedBack *models.FeedBack) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFeedBack(ctx, FeedBack)
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

//EnableFeedBack : ""
func (s *Service) EnableFeedBack(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFeedBack(ctx, UniqueID)
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

//DisableFeedBack : ""
func (s *Service) DisableFeedBack(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFeedBack(ctx, UniqueID)
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

//DeleteFeedBack : ""
func (s *Service) DeleteFeedBack(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFeedBack(ctx, UniqueID)
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

//GetSingleFeedBack :""
func (s *Service) GetSingleFeedBack(ctx *models.Context, UniqueID string) (*models.RefFeedBack, error) {
	FeedBack, err := s.Daos.GetSingleFeedBack(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return FeedBack, nil
}

//FilterFeedBack :""
func (s *Service) FilterFeedBack(ctx *models.Context, FeedBackfilter *models.FeedBackFilter, pagination *models.Pagination) (FeedBack []models.RefFeedBack, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	if FeedBackfilter != nil {

		dataaccess, err := s.Daos.DataAccess(ctx, &FeedBackfilter.DataAccess)
		if err != nil {
			return nil, err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					FeedBackfilter.ContentOrganisation = append(FeedBackfilter.ContentOrganisation, v.ID)
				}
			}
			if len(dataaccess.Projects) > 0 {
				for _, v := range dataaccess.Projects {
					FeedBackfilter.ContentProject = append(FeedBackfilter.ContentProject, v.ID)
				}
			}
			if len(dataaccess.AccessStates) > 0 {
				for _, v := range dataaccess.AccessStates {
					FeedBackfilter.State = append(FeedBackfilter.State, v.ID)
				}
			}
			if len(dataaccess.AccessDistricts) > 0 {
				for _, v := range dataaccess.AccessDistricts {
					FeedBackfilter.District = append(FeedBackfilter.District, v.ID)
				}
			}
			if len(dataaccess.AccessBlocks) > 0 {
				for _, v := range dataaccess.AccessBlocks {
					FeedBackfilter.Block = append(FeedBackfilter.Block, v.ID)
				}
			}
			if len(dataaccess.AccessVillages) > 0 {
				for _, v := range dataaccess.AccessVillages {
					FeedBackfilter.Village = append(FeedBackfilter.Village, v.ID)

				}
			}
			if len(dataaccess.AccessGrampanchayats) > 0 {
				for _, v := range dataaccess.AccessGrampanchayats {
					FeedBackfilter.GramPanchayat = append(FeedBackfilter.GramPanchayat, v.ID)

				}
			}
		}

	}

	return s.Daos.FilterFeedBack(ctx, FeedBackfilter, pagination)

}

func (s *Service) ConsolidatedFeedBack(ctx *models.Context, FeedBackfilter *models.FeedBackFilter) ([]models.FeedBackRating, error) {
	feedbacks, err := s.Daos.ConsolidatedFeedBack(ctx, FeedBackfilter)
	if err != nil {
		return nil, err
	}
	if len(feedbacks) > 0 {
		var avg float64
		var totalrating float64
		totalrating = float64((feedbacks[0].A1 * 1) + (feedbacks[0].A2 * 2) + (feedbacks[0].A3 * 3) + (feedbacks[0].A4 * 4) + (feedbacks[0].A5 * 5) + (feedbacks[0].A6 * 6) + (feedbacks[0].A7 * 7) + (feedbacks[0].A8 * 8) + (feedbacks[0].A9 * 9) + (feedbacks[0].A10 * 10))

		totalfeedbackfarmer := float64(feedbacks[0].A1 + feedbacks[0].A2 + feedbacks[0].A3 + feedbacks[0].A4 + feedbacks[0].A5 + feedbacks[0].A6 + feedbacks[0].A7 + feedbacks[0].A8 + feedbacks[0].A9 + feedbacks[0].A10)
		avg = float64(totalrating) / totalfeedbackfarmer
		feedbacks[0].Average = avg
		feedbacks[0].AverageStr = fmt.Sprintf("%.2f/10 ", avg)
		fmt.Println("avg===>", feedbacks[0].AverageStr)
	}
	return feedbacks, nil

}
