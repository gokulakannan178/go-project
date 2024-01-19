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

func (s *Service) SaveContent(ctx *models.Context, content *models.Content) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	content.Status = constants.CREATED

	//market.District = primitive.ObjectID
	t := time.Now()
	var uniqueid string
	if int(t.Month()) < 10 {
		if t.Day() < 10 {
			uniqueid = fmt.Sprintf("%v0%v0%v", t.Year(), int(t.Month()), t.Day())
		} else {
			uniqueid = fmt.Sprintf("%v0%v%v", t.Year(), int(t.Month()), t.Day())
		}
	} else {
		if t.Day() < 0 {
			uniqueid = fmt.Sprintf("%v%v0%v", t.Year(), int(t.Month()), t.Day())
		} else {
			uniqueid = fmt.Sprintf("%v%v%v", t.Year(), int(t.Month()), t.Day())
		}
	}
	fmt.Println("uniqueid====>", uniqueid)
	switch content.Type {
	case constants.CONTENTTYPESMS:
		content.RecordId = fmt.Sprintf("%v%v_%v", "S", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))
	case constants.CONTENTTYPEVIDEO:
		content.RecordId = fmt.Sprintf("%v%v_%v", "U", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))
	case constants.CONTENTTYPEVOICE:
		content.RecordId = fmt.Sprintf("%v%v_%v", "V", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))
	case constants.CONTENTTYPEPOSTER:
		content.RecordId = fmt.Sprintf("%v%v_%v", "P", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))
	case constants.CONTENTTYPEDOCMENT:
		content.RecordId = fmt.Sprintf("%v%v_%v", "D", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))
	default:
		content.RecordId = fmt.Sprintf("%v%v_%v", "CONT", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))

	}
	fmt.Println("RecordId===>", content.RecordId)
	//created := models.Created{}
	//created.On = &t
	//created.By = constants.SYSTEM
	log.Println("b4 organisation.created")
	content.DateCreated = &t
	content.DateReviewed = &t
	log.Println("b4 organisation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveContent(ctx, content)
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

func (s *Service) UpdateContent(ctx *models.Context, content *models.Content) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateContent(ctx, content)
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

func (s *Service) EnableContent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableContent(ctx, UniqueID)
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

func (s *Service) DisableContent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableContent(ctx, UniqueID)
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

func (s *Service) DeleteContent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteContent(ctx, UniqueID)
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

func (s *Service) GetSingleContent(ctx *models.Context, UniqueID string) (*models.RefContent, error) {
	content, err := s.Daos.GetSingleContent(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s *Service) FilterContent(ctx *models.Context, contentfilter *models.ContentFilter, pagination *models.Pagination) (content []models.RefContent, err error) {
	err = s.ContentDataAccess(ctx, contentfilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterContent(ctx, contentfilter, pagination)
}
func (s *Service) ApprovedContent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.ApprovedContent(ctx, UniqueID)
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
func (s *Service) RejectedContent(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectedContent(ctx, UniqueID)
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
func (s *Service) ContentManager(ctx *models.Context, contentfilter *models.ContentFilter, pagination *models.Pagination) (content []models.RefContent, err error) {
	err = s.ContentDataAccess(ctx, contentfilter)
	if err != nil {
		return nil, err
	}
	//contentfilter.DataAccess = *dataAceess
	return s.Daos.ContentManager(ctx, contentfilter, pagination)
}
func (s *Service) EditApprovedContent(ctx *models.Context, content *models.ApprovedContent) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EditApprovedContent(ctx, content)
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
func (s *Service) EditRejectedContent(ctx *models.Context, content *models.RejectedContent) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EditRejectedContent(ctx, content)
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
func (s *Service) ContentDataAccess(ctx *models.Context, contentfilter *models.ContentFilter) (err error) {
	if contentfilter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &contentfilter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					contentfilter.Organisation = append(contentfilter.Organisation, v.ID)
				}
			}
			if len(dataaccess.Projects) > 0 {
				for _, v := range dataaccess.Projects {
					contentfilter.Project = append(contentfilter.Project, v.Project)
				}
			}

			if len(dataaccess.AccessStates) > 0 {
				for _, v := range dataaccess.AccessStates {
					contentfilter.State = append(contentfilter.State, v.ID)
				}
			}
			if len(dataaccess.AccessDistricts) > 0 {
				for _, v := range dataaccess.AccessDistricts {
					contentfilter.District = append(contentfilter.District, v.ID)
				}
			}
			if len(dataaccess.AccessBlocks) > 0 {
				for _, v := range dataaccess.AccessBlocks {
					contentfilter.Block = append(contentfilter.Block, v.ID)
				}
			}
			if len(dataaccess.AccessVillages) > 0 {
				for _, v := range dataaccess.AccessVillages {
					contentfilter.Village = append(contentfilter.Village, v.ID)

				}
			}
			if len(dataaccess.AccessGrampanchayats) > 0 {
				for _, v := range dataaccess.AccessGrampanchayats {
					contentfilter.Gram_panchayat = append(contentfilter.Gram_panchayat, v.ID)

				}
			}
		}

	}
	return err
}
func (s *Service) ContentViewCountIncrement(ctx *models.Context, content *models.ContentViewCount) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	if content.UserType == "Farmer" {
		content.UserType = "farmerViewCount"
	}
	if content.UserType == "User" {
		content.UserType = "usersViewCount"
	}
	if content.UserType == "" || content.UserType == "Guest" {
		content.UserType = "guestUsersViewCount"
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.ContentViewCountIncrement(ctx, content)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		t := time.Now()
		ContentCountLog := new(models.ContentCountLog)
		ContentCountLog.UserType = content.UserType
		ContentCountLog.Date = &t
		ContentCountLog.ContentId = content.ContentId
		ContentCountLog.Status = constants.CONTENTCOUNTLOGSTATUSACTIVE
		if content.UserType == "Farmer" {
			ContentCountLog.FarmerId = content.UserId
		}
		if content.UserType == "User" {
			ContentCountLog.UserId = content.UserId
		}
		err := s.Daos.SaveContentCountLog(ctx, ContentCountLog)
		if err != nil {
			return err
		}
		ContentViewLog := new(models.ContentViewLog)
		ContentViewLog.ContentId = content.ContentId
		ContentViewLog.UniqueId = fmt.Sprintf("%v_%v_%v", t.Day(), t.Month(), t.Year())
		ContentViewLog.Date = &t
		ContentViewLog.Status = constants.CONTENTVIEWLOGSTATUSACTIVE
		err = s.Daos.SaveContentViewLogUpdertInc(ctx, ContentViewLog)
		if err != nil {
			return err
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
