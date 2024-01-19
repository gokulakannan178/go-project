package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateMobileTowerReassessmentRequest : ""
func (s *Service) BasicUpdateMobileTowerReassessmentRequest(ctx *models.Context, request *models.MobileTowerReassessmentRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldMobileTowerData, err := s.Daos.GetSingleMobileTower(ctx, request.MobileTowerID)
		if err != nil {
			return errors.New("Error in geting old mobileTower" + err.Error())
		}
		if oldMobileTowerData == nil {
			return errors.New("mobileTower Not Found")
		}
		t := time.Now()
		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERREASSESSMENTREQUEST)
		request.Previous.PropertyMobileTower = oldMobileTowerData.PropertyMobileTower
		request.Previous.Ref.Address = oldMobileTowerData.Ref.Address
		request.Previous.PropertyMobileTower.OwnerName = oldMobileTowerData.OwnerName
		request.Status = constants.MOBILETOWERREASSESSMENTREQUESTSTATUSINIT
		request.Requester.By = request.UserName
		request.Requester.ByType = request.UserType
		request.Requester.On = &t

		err = s.Daos.SaveMobileTowerReassessmentRequestUpdate(ctx, request)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		// // html template path
		// templateID := templatePathStart + "ReassessmentUpdateRequestEmail.html"
		// templateID = "templates/MobileTowerReassessmentRequestEmail.html"

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

// RejectBasicMobileTowerUpdate : ""
func (s *Service) RejectMobileTowerReassessmentRequestUpdate(ctx *models.Context, req *models.RejectMobileTowerReassessmentRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectMobileTowerReassessmentRequestUpdate(ctx, req)
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

// AcceptMobileTowerReassessmentRequestUpdate : ""
func (s *Service) AcceptMobileTowerReassessmentRequestUpdate(ctx *models.Context, req *models.AcceptMobileTowerReassessmentRequestUpdate) error {

	res, err := s.Daos.GetSingleMobileTowerReassessmentRequest(ctx, req.UniqueID)
	if err != nil {
		return errors.New("Error in getting in mobileTowerReassessment Request" + err.Error())
	}

	oldMobileTowerData, err := s.Daos.GetSingleMobileTower(ctx, res.MobileTowerID)
	if err != nil {
		return errors.New("Error in getting in old mobileTower" + err.Error())
	}

	if oldMobileTowerData == nil {
		return errors.New("mobileTower Not Found")
	}

	// if len(res.New.Ref.MobileTowerOwner) > 0 {
	// 	for _, v := range res.New.Ref.MobileTowerOwner {
	// 		res.New.Ref.MobileTowerOwner = append(res.New.Ref.MobileTowerOwner, v.MobileTowerOwner)
	// 	}
	// }
	// if len(res.New.Ref.Floors) > 0 {
	// 	for _, v := range res.New.Ref.Floors {
	// 		res.New.Ref.Floors = append(res.New.Ref.Floors, v.MobileTowerFloor)
	// 	}
	// }

	// res.New.MobileTower.Owner = res.New.Ref.ReassessmentOwners
	// res.New.MobileTower.Floors = res.New.Ref.ReassessmentFloors
	// fmt.Println("property floor new b4 saving =====>", res.New.Ref.ReassessmentFloors)
	// fmt.Println("property floor new after saving =====>", res.New.MobileTower.Floors)
	res.New.PropertyMobileTower.UniqueID = res.MobileTowerID
	err = s.Daos.UpdateMobileTower(ctx, &res.New.PropertyMobileTower)
	if err != nil {
		return errors.New("Error in updating in mobileTower" + err.Error())
	}
	// _, err = s.CalcMobileTowerOverallMonthlyDemand(ctx, res.MobileTowerID, true)
	// if err != nil {
	// 	return errors.New("Error in updating in mobileTower" + err.Error())
	// }

	err = s.Daos.BasicMobileTowerReassessmentRequestUpdateToPayments(ctx, res)
	if err != nil {
		return errors.New("Error in upating in mobileTowerReassessment Request Payments" + err.Error())
	}
	err = s.Daos.AcceptMobileTowerReassessmentRequestUpdate(ctx, req)
	if err != nil {
		return errors.New("Error in upating in mobileTowerReassessment Request" + err.Error())
	}
	return nil
}

// FilterMobileTowerReassessmentRequest : ""
func (s *Service) FilterMobileTowerReassessmentRequest(ctx *models.Context, filter *models.MobileTowerReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefMobileTowerReassessmentRequest, error) {
	return s.Daos.FilterMobileTowerReassessmentRequest(ctx, filter, pagination)

}

//GetSingleMobileTowerReassessmentRequest :""
func (s *Service) GetSingleMobileTowerReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefMobileTowerReassessmentRequest, error) {
	res, err := s.Daos.GetSingleMobileTowerReassessmentRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
