package services

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveULB :""
func (s *Service) SaveULB(ctx *models.Context, ULB *models.ULB) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	ULB.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONULB)
	ULB.Status = constants.ULBSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ULB.created")
	ULB.Created = created
	log.Println("b4 ULB.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveULB(ctx, ULB)
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

//UpdateULB : ""
func (s *Service) UpdateULB(ctx *models.Context, ULB *models.ULB) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}

	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if ULB.NodalOfficer.MobileNo != "" {
			err := s.Daos.UpdateULBwithMobileno(ctx, ULB.UniqueID, ULB.NodalOfficer.MobileNo)
			if err != nil {
				return errors.New("please check no - " + err.Error())
			}
			return err
		}
		err := s.Daos.UpdateULB(ctx, ULB)
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

//EnableULB : ""
func (s *Service) EnableULB(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableULB(ctx, UniqueID)
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

//DisableULB : ""
func (s *Service) DisableULB(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableULB(ctx, UniqueID)
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

//DeleteULB : ""
func (s *Service) DeleteULB(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteULB(ctx, UniqueID)
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

//GetSingleULB :""
func (s *Service) GetSingleULB(ctx *models.Context, UniqueID string) (*models.RefULB, error) {
	ULB, err := s.Daos.GetSingleULB(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ULB, nil
}

//FilterULB :""
func (s *Service) FilterULB(ctx *models.Context, ULBfilter *models.ULBFilter, pagination *models.Pagination) (ULB []models.RefULB, err error) {

	return s.Daos.FilterULB(ctx, ULBfilter, pagination)

}

//AddULBTestCert : ""
func (s *Service) AddULBTestCert(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	if ulbTestCert.Status == "" {
		ulbTestCert.Status = constants.ULBTESTCERTSTATUSNEW
	}
	return s.Daos.AddULBTestCert(ctx, UniqueID, ulbTestCert)
}

//ApplyForTestCert : ""
func (s *Service) ApplyForTestCert(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	t := time.Now()
	ulbTestCert.AppliedDate = &t
	if ulbTestCert.Status == "" {
		ulbTestCert.Status = constants.ULBTESTCERTSTATUSNEW

	}
	return s.Daos.ApplyForTestCert(ctx, UniqueID, ulbTestCert)
}
func (s *Service) ReApplyForTestCert(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	t := time.Now()
	ulbTestCert.AppliedDate = &t
	if ulbTestCert.Status == "" {
		ulbTestCert.Status = constants.ULBTESTCERTSTATUSREAPPLIED

	}
	return s.Daos.ReApplyForTestCert(ctx, UniqueID, ulbTestCert)
}

//ULBTestCertStatus : ""
func (s *Service) AcceptTestCert(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	return s.Daos.AcceptTestCert(ctx, UniqueID, ulbTestCert)
}
func (s *Service) RejectTestCert(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	return s.Daos.RejectTestCert(ctx, UniqueID, ulbTestCert)
}

//ULBTestCertStatus : ""
func (s *Service) ULBTestCertStatus(ctx *models.Context, UniqueID string, ulbTestCert *models.ULBTestCert) error {
	return s.Daos.ULBTestCertStatus(ctx, UniqueID, ulbTestCert)
}
func (s *Service) ULBMobileUniqueness(ctx *models.Context, ulb *models.ULB, mobileno string) error {
	res, err := s.Daos.GetSingleMobileNoForULB(ctx, mobileno)
	if err != nil {
		return err
	}
	if res.NodalOfficer.MobileNo == mobileno && res.UniqueID != ulb.UniqueID {
		return s.Daos.UpdateULB(ctx, ulb)
	}
	if res.NodalOfficer.MobileNo == "" && res.UniqueID != ulb.UniqueID {
		return s.Daos.UpdateULB(ctx, ulb)
	} else {
		fmt.Println("success")
	}
	return nil
}
func (s *Service) UlbTestcertExcelForPending(ctx *models.Context, filter *models.ULBFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterULB(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "UlbTestcert Pending Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "G1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S NO")
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "AppliedDate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "NodalOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "NodalOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "ChiefOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "ChiefOfficerMobileNo")

	rowNo++

	//	var totalAmount float64
	for i, v := range data {
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.TestCert.AppliedDate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.NodalOfficer.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.NodalOfficer.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.CO.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.CO.MobileNo)

		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.TestCert.Remarks)
		rowNo++
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) UlbTestcertExcelForApproved(ctx *models.Context, filter *models.ULBFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterULB(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "UlbTestcert Approved Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "I1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S No")
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "NodalOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "NodalOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "ChiefOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "ChiefOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "CertificateDate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "ExpiryDate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Remarks")

	rowNo++

	//	var totalAmount float64
	for i, v := range data {
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.NodalOfficer.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.NodalOfficer.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.CO.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.CO.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.TestCert.REGDate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.TestCert.ExpDate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.TestCert.Remarks)
		rowNo++
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) UlbTestcertExcelForRejected(ctx *models.Context, filter *models.ULBFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterULB(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "UlbTestcert Rejected Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "H1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S No")
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "NodalOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "NodalOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "ChiefOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "ChiefOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Remarks")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "RejectedDate")
	rowNo++

	//	var totalAmount float64
	for i, v := range data {
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.NodalOfficer.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.NodalOfficer.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.CO.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.CO.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.TestCert.Remarks)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.TestCert.RejectedDate)
		rowNo++
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) UlbTestcertExcelForReApplied(ctx *models.Context, filter *models.ULBFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterULB(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "UlbTestcert ReApplied Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "G1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S No")
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "NodalOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "NodalOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "ChiefOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "ChiefOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "ReAppliedDate")
	rowNo++

	//	var totalAmount float64
	for i, v := range data {
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.NodalOfficer.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.NodalOfficer.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.CO.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.CO.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.TestCert.AppliedDate)
		rowNo++
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) UlbTestcertExcelForExpiry(ctx *models.Context, filter *models.ULBFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterULB(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "UlbTestcert Expiry Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "G1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S No")
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "ULBName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "NodalOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "NodalOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "ChiefOfficerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "ChiefOfficerMobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "ExpiryDate")

	rowNo++

	//	var totalAmount float64
	for i, v := range data {
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.NodalOfficer.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.NodalOfficer.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.CO.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.CO.MobileNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.TestCert.ExpDate)
		rowNo++
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
