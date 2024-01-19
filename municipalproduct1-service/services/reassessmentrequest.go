package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

// BasicUpdateReassessmentRequest : ""
func (s *Service) BasicUpdateReassessmentRequest(ctx *models.Context, request *models.ReassessmentRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldPropertyData, err := s.Daos.GetSingleProperty(ctx, request.PropertyID)
		if err != nil {
			return errors.New("Error in geting old Property" + err.Error())
		}
		if oldPropertyData == nil {
			return errors.New("Property Not Found")
		}
		t := time.Now()
		created := new(models.Created)
		created.On = &t
		// created.By = request.UserName

		request.Created = created

		request.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONREASSESSMENTREQUEST)
		request.Previous.Property = oldPropertyData.Property
		if len(request.Previous.Ref.Floors) > 0 {
			request.Previous.Ref.Floors = oldPropertyData.Ref.Floors
		}
		if len(request.Previous.Ref.PropertyOwner) > 0 {
			request.Previous.Ref.PropertyOwner[0] = oldPropertyData.Ref.PropertyOwner[0]
		}
		//
		filter := new(models.PropertyDocumentsFilter)
		filter.PropertyID = append(filter.PropertyID, oldPropertyData.UniqueID)
		fmt.Println("filter =======>", filter)

		resPropertyDocument, err1 := s.FilterPropertyDocument(ctx, filter, nil)
		if err1 != nil {
			return errors.New("error in getting the property documets list" + err1.Error())
		}
		fmt.Println("resPropertyDocument =======>", resPropertyDocument)
		if len(resPropertyDocument) > 0 {
			request.Previous.Ref.Documents = append(request.Previous.Ref.Documents, resPropertyDocument...)

		}

		//
		request.Status = constants.REASSESSMENTREQUESTSTATUSINIT
		request.Requester.On = &t
		err = s.Daos.SaveReassessmentRequestUpdate(ctx, request)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		// // html template path
		// templateID := templatePathStart + "ReassessmentUpdateRequestEmail.html"
		// templateID = "templates/ReassessmentRequestEmail.html"

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
func (s *Service) RejectReassessmentRequestUpdate(ctx *models.Context, req *models.RejectReassessmentRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectReassessmentRequestUpdate(ctx, req)
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

// AcceptReassessmentRequestUpdate : ""
func (s *Service) AcceptReassessmentRequestUpdate(ctx *models.Context, req *models.AcceptReassessmentRequestUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	client := s.Daos.GetDBV3(context.TODO())
	defer client.Disconnect(context.TODO())

	database := client.Database("municipalproduct1")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		res, err := s.Daos.GetSingleReassessmentRequest(ctx, req.UniqueID)
		if err != nil {
			return errors.New("Error in getting in Reassessment Request" + err.Error())
		}
		fmt.Println("res.PropertyID=====>", res.PropertyID)
		oldPropertyData, err := s.Daos.GetSingleProperty(ctx, res.PropertyID)
		if err != nil {
			return errors.New("Error in getting in old property" + err.Error())
		}

		if oldPropertyData == nil {
			return errors.New("property Not Found")
		}

		for _, v := range res.New.Ref.ReassessmentOwners {
			res.New.Property.Owner = append(res.New.Property.Owner, v.PropertyOwner)
		}
		for _, v := range res.New.Ref.ReassessmentFloors {
			res.New.Property.Floors = append(res.New.Property.Floors, v.PropertyFloor)
		}

		fmt.Println("property floor new b4 saving =====>", res.New.Ref.ReassessmentFloors)
		fmt.Println("property floor new after saving =====>", res.New.Property.Floors)
		//
		for _, v := range res.New.Ref.ReassessmentDocuments {

			res.New.Property.PropertyDocument = append(res.New.Property.PropertyDocument, v.PropertyDocuments)

			dberr := s.Daos.SavePropertyDocumentv2(ctx, database, &sc, res.New.Property.PropertyDocument)
			if dberr != nil {
				return errors.New("Transaction Aborted <property documents> - " + dberr.Error())
			}
		}
		fmt.Println("Helooooooooooooooooooooooooooooooooooooooo", res.New.Property.PropertyDocument)
		//
		err = s.UpdatePropertyWithOutTransaction(ctx, &res.New.Property)
		if err != nil {
			return errors.New("Error in updating in Property" + err.Error())
		}

		err = s.Daos.BasicReassessmentRequestUpdateToPayments(ctx, res)
		if err != nil {
			return errors.New("Error in upating in Reassessment Request Payments" + err.Error())
		}
		err = s.Daos.AcceptReassessmentRequestUpdate(ctx, req)
		if err != nil {
			return errors.New("Error in upating in Reassessment Request" + err.Error())
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

// FilterReassessmentRequest : ""
func (s *Service) FilterReassessmentRequest(ctx *models.Context, filter *models.ReassessmentRequestFilter, pagination *models.Pagination) ([]models.RefReassessmentRequest, error) {
	return s.Daos.FilterReassessmentRequest(ctx, filter, pagination)
}

// GetSingleReassessmentRequest :""
func (s *Service) GetSingleReassessmentRequest(ctx *models.Context, UniqueID string) (*models.RefReassessmentRequest, error) {
	rrr, err := s.Daos.GetSingleReassessmentRequest(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return rrr, nil
}

// FilterReassessmentRequestExcel :""
func (s *Service) FilterReassessmentRequestExcel(ctx *models.Context, filter *models.ReassessmentRequestFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterReassessmentRequest(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Reassessment Report"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "H3")
		excel.MergeCell(sheet1, "A4", "H5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "H3")
		excel.MergeCell(sheet1, "C4", "H5")
	}
	excel.MergeCell(sheet1, "A6", "H6")
	excel.MergeCell(sheet1, "A7", "H7")
	// excel.MergeCell(sheet1, "A8", "H8")

	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Reassessment Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Reassessment Report")
	}
	rowNo++
	rowNo++

	reportFromMsg := "Reassessment Report"
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg3 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg3)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Requested By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Requested On")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Approved By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Approved On")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Remarks")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Status")
	rowNo++

	for k, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.PropertyID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Ref.RequestedUser.Name != "" {
				return v.Ref.RequestedUser.Name
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Requester.On != nil {
				return v.Requester.On.Format("02-January-2006")
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Action.By != "" {
				return v.Action.By
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Action.On != nil {
				return v.Action.On.Format("02-January-2006")
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if v.Action.Remarks != "" {
				return v.Action.Remarks
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Status)

		rowNo++

	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)

	return excel, nil

}

// UpdatePropertyReassessmentRequestPropertyID : ""
func (s *Service) UpdatePropertyReassessmentRequestPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range uniqueIds.UniqueIDs {
			resProperty, err := s.GetSingleProperty(ctx, v)
			if err != nil {
				return errors.New("Not able to get property - " + err.Error())
			}

			uniqueIds.UniqueID = resProperty.OldUniqueID
			uniqueIds.OldUniqueID = resProperty.OldUniqueID
			uniqueIds.NewUniqueID = resProperty.NewUniqueID
			err = s.Daos.UpdatePropertyReassessmentRequestPropertyID(ctx, uniqueIds)
			if err != nil {
				return err
			}
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
