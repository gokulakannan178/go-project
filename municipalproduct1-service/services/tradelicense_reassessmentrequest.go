package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateTradeLicenseReassessmentRequest : ""
func (s *Service) BasicUpdateTradeLicenseReassessmentRequest(ctx *models.Context, request *models.TradeLicenseReassessmentRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldTradeLicenseData, err := s.Daos.GetSingleTradeLicense(ctx, request.TradeLicenseID)
		if err != nil {
			return errors.New("Error in geting old tradelicense" + err.Error())
		}
		if oldTradeLicenseData == nil {
			return errors.New("tradelicense Not Found")
		}
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSEREASSESSMENTREQUEST)
		request.Previous.TradeLicense = oldTradeLicenseData.TradeLicense
		request.Previous.Ref.Address = oldTradeLicenseData.Ref.Address
		request.Previous.TradeLicense.OwnerName = oldTradeLicenseData.OwnerName
		request.Status = constants.TRADELICENSEREASSESSMENTREQUESTSTATUSINIT
		request.Requester.On = &t

		err = s.Daos.SaveTradeLicenseReassessmentRequestUpdate(ctx, request)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		// // html template path
		// templateID := templatePathStart + "ReassessmentUpdateRequestEmail.html"
		// templateID = "templates/TradeLicenseReassessmentRequestEmail.html"

		// //sending email
		// if err := s.SendEmailWithTemplate("Reassessment Update Request - holding no 1111", []string{"solomon2261993@gmail.com"}, templateID, nil); err != nil {
		// 	log.Println("email not sent - ", err.Error())
		// }
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
func (s *Service) RejectTradeLicenseReassessmentRequestUpdate(ctx *models.Context, req *models.RejectTradeLicenseReassessmentRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectTradeLicenseReassessmentRequestUpdate(ctx, req)
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

// AcceptShoprentReassessmentRequestUpdate : ""
func (s *Service) AcceptTradeLicenseReassessmentRequestUpdate(ctx *models.Context, req *models.AcceptTradeLicenseReassessmentRequestUpdate) error {
	fmt.Println("trade license accept service")

	res, err := s.Daos.GetSingleTradeLicenseReassessmentRequest(ctx, req.UniqueID)
	if err != nil {
		return errors.New("Error in getting in shoprentReassessment Request" + err.Error())
	}

	resTradeLicense, err := s.Daos.GetSingleTradeLicense(ctx, res.TradeLicenseID)
	if err != nil {
		return errors.New("Error in getting in old trade license" + err.Error())
	}

	if resTradeLicense == nil {
		return errors.New("tradelicense Not Found")
	}

	// if len(res.New.Ref.PropertyOwner) > 0 {
	// 	for _, v := range res.New.Ref.PropertyOwner {
	// 		res.New.Ref.PropertyOwner = append(res.New.Ref.PropertyOwner, v.PropertyOwner)
	// 	}
	// }
	// if len(res.New.Ref.Floors) > 0 {
	// 	for _, v := range res.New.Ref.Floors {
	// 		res.New.Ref.Floors = append(res.New.Ref.Floors, v.PropertyFloor)
	// 	}
	// }

	// res.New.Property.Owner = res.New.Ref.ReassessmentOwners
	// res.New.Property.Floors = res.New.Ref.ReassessmentFloors
	// fmt.Println("property floor new b4 saving =====>", res.New.Ref.ReassessmentFloors)
	// fmt.Println("property floor new after saving =====>", res.New.Property.Floors)
	res.New.TradeLicense.UniqueID = res.TradeLicenseID
	err = s.Daos.UpdateTradeLicense(ctx, &res.New.TradeLicense)
	if err != nil {
		return errors.New("Error in updating in shoprent" + err.Error())
	}
	err = s.Daos.BasicTradeLicenseReassessmentRequestUpdateToPayments(ctx, res)
	if err != nil {
		return errors.New("Error in upating in tradelicense Reassessment Request Payments" + err.Error())
	}
	err = s.Daos.AcceptTradeLicenseReassessmentRequestUpdate(ctx, req)
	if err != nil {
		return errors.New("Error in upating in tradelicense Reassessment Request" + err.Error())
	}
	// _, err = s.CalcShopRentOverallMonthlyDemand(ctx, res.ShoprentID, true)
	// if err != nil {
	// 	return errors.New("Error in updating in shoprent" + err.Error())
	// }
	return nil
}

// FilterTradeLicenseReassessmentRequest : ""
func (s *Service) FilterTradeLicenseReassessmentRequest(ctx *models.Context, filter *models.TradeLicenseReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefTradeLicenseReassessmentRequest, error) {
	return s.Daos.FilterTradeLicenseReassessmentRequest(ctx, filter, pagination)

}

//GetSingleTradeLicenseReassessmentRequest :""
func (s *Service) GetSingleTradeLicenseReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseReassessmentRequest, error) {
	res, err := s.Daos.GetSingleTradeLicenseReassessmentRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
