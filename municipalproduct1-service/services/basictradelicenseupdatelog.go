package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateTradeLicense : ""
func (s *Service) BasicUpdateTradeLicense(ctx *models.Context, btlu *models.BasicTradeLicenseUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldTradeLicenseData, err := s.Daos.GetSinglePreviousTradeLicense(ctx, btlu.TradeLicenseID)
		if err != nil {
			return errors.New("Error in geting old Trade License" + err.Error())

		}
		if oldTradeLicenseData == nil {
			return errors.New("trade license Not Found")
		}
		btlul := new(models.BasicTradeLicenseUpdateLog)
		btlul.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBASICTRADELICENSEUPDATELOG)
		btlul.TradeLicenseID = btlu.TradeLicenseID
		btlul.Previous = *oldTradeLicenseData
		// if len(oldTradeLicenseData.OwnerName) > 0 {
		// bpul.Previous.Owner = oldTradeLicenseData.OwnerName
		// }
		btlul.New = btlu.UpdateData
		// btlul.New.Owner = btlu.Owner
		btlul.UserName = btlu.UserName
		btlul.UserType = btlu.UserType
		// t := time.Now()
		// btlul.Requester = models.Updated{
		// 	On:       &t,
		// 	By:       btlu.UserName,
		// 	Scenario: "BasicUpdate",
		// 	ByType:   btlu.UserType,
		// 	Remarks:  btlu.Remarks,
		// }
		btlul.Proof = btlu.Proof
		btlul.Status = constants.TRADELICENSEBASICUPDATELOGINIT
		err = s.Daos.SaveBasicTradeLicenseUpdateLog(ctx, btlul)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		// html template path
		templateID := templatePathStart + "TradeLicenseUpdateRequestEmail.html"
		templateID = "templates/TradeLicenseUpdateRequestEmail.html"

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

// AcceptBasicTradeLicenseUpdate : ""
func (s *Service) AcceptBasicTradeLicenseUpdate(ctx *models.Context, req *models.AcceptBasicTradeLicenseUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//Finding the request
		rbpu, err := s.Daos.GetSingleBasicTradeLicenseUpdateLog(ctx, req.UniqueID)
		if err != nil {
			return errors.New("not able to find the request" + err.Error())
		}
		if rbpu == nil {
			return errors.New("request in nil")
		}

		err = s.Daos.BasicUpdateTradeLicense(ctx, &rbpu.BasicTradeLicenseUpdateLog)
		if err != nil {
			return errors.New("Error in updating Property" + err.Error())
		}

		//updating the request
		err = s.Daos.AcceptBasicTradeLicenseUpdate(ctx, req)
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

// RejectBasicTradeLicenseUpdate : ""
func (s *Service) RejectBasicTradeLicenseUpdate(ctx *models.Context, req *models.RejectBasicTradeLicenseUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectBasicTradeLicenseUpdate(ctx, req)
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

//FilterBasicTradeLicenseUpdateLog :""
func (s *Service) FilterBasicTradeLicenseUpdateLog(ctx *models.Context, filter *models.FilterBasicTradeLicenseUpdateLog, pagination *models.Pagination) ([]models.RefBasicTradeLicenseUpdateLog, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterBasicTradeLicenseUpdateLog(ctx, filter, pagination)
}

//GetSingleTradeLicenseV2 :""
func (s *Service) GetSingleBasicTradeLicenseUpdateLogV2(ctx *models.Context, UniqueID string) (*models.BasicTradeLicenseUpdateLog, error) {
	tower, err := s.Daos.GetSingleBasicTradeLicenseUpdateLogV2(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// BasicTradeLicenseUpdateGetPaymentsToBeUpdated : ""
func (s *Service) BasicTradeLicenseUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbtlul *models.RefBasicTradeLicenseUpdateLogV2) ([]models.RefTradeLicensePayments, error) {
	return s.Daos.BasicTradeLicenseUpdateGetPaymentsToBeUpdated(ctx, rbtlul)
}
