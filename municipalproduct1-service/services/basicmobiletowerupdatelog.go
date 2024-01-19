package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateMobileTower : ""
func (s *Service) BasicMobileTowerUpdate(ctx *models.Context, bmtlu *models.BasicMobileTowerUpdateData) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldMobileTowerData, err := s.Daos.GetSinglePreviousMobileTower(ctx, bmtlu.MobileTowerID)
		if err != nil {
			return errors.New("Error in geting old MobileTower" + err.Error())

		}
		if oldMobileTowerData == nil {
			return errors.New("mobile tower Not Found")
		}
		mtul := new(models.BasicMobileTowerUpdateLog)
		mtul.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERUPDATELOG)
		mtul.MobileTowerID = bmtlu.MobileTowerID
		mtul.Previous = *oldMobileTowerData
		mtul.New = bmtlu.UpdateData
		mtul.UserName = bmtlu.UserName
		mtul.UserType = bmtlu.UserType
		mtul.Proof = bmtlu.Proof
		mtul.Status = constants.MOBILETOWERBASICUPDATELOGINIT
		err = s.Daos.SaveBasicMobileTowerUpdate(ctx, mtul)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		// html template path
		templateID := templatePathStart + "MobileTowerUpdateRequestEmail.html"
		templateID = "templates/MobileTowerUpdateRequestEmail.html"

		//sending email
		if err := s.SendEmailWithTemplate("Property Update Request - holding no 1111", []string{"solomon2261993@gmail.com"}, templateID, nil); err != nil {
			log.Println("email not sent - ", err.Error())
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

// AcceptBasicMobileTowerUpdate : ""
func (s *Service) AcceptBasicMobileTowerUpdate(ctx *models.Context, req *models.AcceptBasicMobileTowerUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//Finding the request
		mtbu, err := s.Daos.GetSingleBasicMobileTowerUpdate(ctx, req.UniqueID)
		if err != nil {
			return errors.New("not able to find the request" + err.Error())
		}
		if mtbu == nil {
			return errors.New("request in nil")
		}

		err = s.Daos.BasicMobileTowerUpdate(ctx, &mtbu.BasicMobileTowerUpdateLog)
		if err != nil {
			return errors.New("Error in updating Property" + err.Error())
		}

		//updating the request
		err = s.Daos.AcceptBasicMobileTowerUpdate(ctx, req)
		if err != nil {
			return nil
		}
		// err = s.Daos.BasicPropertyUpdateToPayments(ctx, rbpu)
		// if err != nil {
		// 	return nil
		// }
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// RejectBasicMobileTowerUpdate : ""
func (s *Service) RejectBasicMobileTowerUpdate(ctx *models.Context, req *models.RejectBasicMobileTowerUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectBasicMobileTowerUpdate(ctx, req)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//FilterBasicMobileTowerUpdateLog :""
func (s *Service) FilterBasicMobileTowerUpdateLog(ctx *models.Context, filter *models.FilterBasicMobileTowerUpdateLog, pagination *models.Pagination) ([]models.RefBasicMobileTowerUpdateLog, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterBasicMobileTowerUpdateLog(ctx, filter, pagination)
}

//GetSingleBasicMobileTowerUpdateLogV2 :""
func (s *Service) GetSingleBasicMobileTowerUpdateLogV2(ctx *models.Context, UniqueID string) (*models.BasicMobileTowerUpdateLog, error) {
	tower, err := s.Daos.GetSingleBasicMobileTowerUpdateLogV2(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// BasicMobileTowerUpdateGetPaymentsToBeUpdated : ""
func (s *Service) BasicMobileTowerUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbmtul *models.RefBasicMobileTowerUpdateLogV2) ([]models.RefMobileTowerPayments, error) {
	return s.Daos.BasicMobileTowerUpdateGetPaymentsToBeUpdated(ctx, rbmtul)
}
