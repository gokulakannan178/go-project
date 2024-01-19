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

// BasicUpdateShoprentReassessmentRequest : ""
func (s *Service) BasicUpdateShoprentReassessmentRequest(ctx *models.Context, request *models.ShoprentReassessmentRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldPropertyData, err := s.Daos.GetSingleShopRent(ctx, request.ShoprentID)
		if err != nil {
			return errors.New("Error in geting old shoprent" + err.Error())
		}
		if oldPropertyData == nil {
			return errors.New("shoprent Not Found")
		}

		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSHOPRENTREASSESSMENTREQUEST)
		request.Previous.ShopRent = oldPropertyData.ShopRent
		request.Previous.Ref.Address = oldPropertyData.Ref.Address
		request.Previous.ShopRent.OwnerName = oldPropertyData.OwnerName
		request.Status = constants.SHOPRENTREASSESSMENTREQUESTSTATUSINIT
		request.Requester.On = &t

		err = s.Daos.SaveShoprentReassessmentRequestUpdate(ctx, request)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		// // html template path
		// templateID := templatePathStart + "ReassessmentUpdateRequestEmail.html"
		// templateID = "templates/ShoprentReassessmentRequestEmail.html"

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
func (s *Service) RejectShoprentReassessmentRequestUpdate(ctx *models.Context, req *models.RejectShoprentReassessmentRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectShoprentReassessmentRequestUpdate(ctx, req)
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
func (s *Service) AcceptShoprentReassessmentRequestUpdate(ctx *models.Context, req *models.AcceptShoprentReassessmentRequestUpdate) error {
	fmt.Println("shoprent accept service")

	res, err := s.Daos.GetSingleShoprentReassessmentRequest(ctx, req.UniqueID)
	if err != nil {
		return errors.New("Error in getting in shoprentReassessment Request" + err.Error())
	}

	oldPropertyData, err := s.Daos.GetSingleShopRent(ctx, res.ShoprentID)
	if err != nil {
		return errors.New("Error in getting in old shoprent" + err.Error())
	}

	if oldPropertyData == nil {
		return errors.New("shoprent Not Found")
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
	res.New.ShopRent.UniqueID = res.ShoprentID
	err = s.Daos.UpdateShopRent(ctx, &res.New.ShopRent)
	if err != nil {
		return errors.New("Error in updating in shoprent" + err.Error())
	}
	err = s.Daos.BasicShoprentReassessmentRequestUpdateToPayments(ctx, res)
	if err != nil {
		return errors.New("Error in upating in shoprentReassessment Request Payments" + err.Error())
	}
	err = s.Daos.AcceptShoprentReassessmentRequestUpdate(ctx, req)
	if err != nil {
		return errors.New("Error in upating in shoprentReassessment Request" + err.Error())
	}
	// _, err = s.CalcShopRentOverallMonthlyDemand(ctx, res.ShoprentID, true)
	// if err != nil {
	// 	return errors.New("Error in updating in shoprent" + err.Error())
	// }
	return nil
}

// FilterShoprentReassessmentRequest : ""
func (s *Service) FilterShoprentReassessmentRequest(ctx *models.Context, filter *models.ShoprentReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefShoprentReassessmentRequest, error) {
	return s.Daos.FilterShoprentReassessmentRequest(ctx, filter, pagination)

}

//GetSingleShoprentReassessmentRequest :""
func (s *Service) GetSingleShoprentReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefShoprentReassessmentRequest, error) {
	res, err := s.Daos.GetSingleShoprentReassessmentRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
