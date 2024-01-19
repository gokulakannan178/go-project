package services

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/app"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveFarmer :""
func (s *Service) SaveFarmer(ctx *models.Context, Farmer *models.Farmer) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//Farmer.Code = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFarmer)

	Farmer.Status = constants.FARMERSTATUSACTIVE
	Farmer.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	Farmer.CreatedDate = &t
	log.Println("b4 Farmer.created")
	Farmer.Created = created
	Farmer.FarmerID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFARMER)
	log.Println("b4 Farmer.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		reffarmer, _ := s.Daos.GetSingleFarmerWithMobileno(ctx, Farmer.MobileNumber)
		fmt.Println("farmerOrg======>", Farmer.FarmerOrg)
		if reffarmer == nil {
			dberr := s.Daos.SaveFarmer(ctx, Farmer)
			if dberr != nil {

				return errors.New("Db Error" + dberr.Error())
			}
		}
		if reffarmer != nil {
			return errors.New("mobile no already registered")
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

//UpdateFarmer : ""
func (s *Service) UpdateFarmer(ctx *models.Context, Farmer *models.Farmer) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFarmer(ctx, Farmer)
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

//EnableFarmer : ""
func (s *Service) EnableFarmer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFarmer(ctx, UniqueID)
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

//DisableFarmer : ""
func (s *Service) DisableFarmer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFarmer(ctx, UniqueID)
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

//DeleteFarmer : ""
func (s *Service) DeleteFarmer(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFarmer(ctx, UniqueID)
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

//GetSingleFarmer :""
func (s *Service) GetSingleFarmer(ctx *models.Context, UniqueID string) (*models.RefFarmer, error) {
	Farmer, err := s.Daos.GetSingleFarmer(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Farmer, nil
}

//FilterFarmer :""
func (s *Service) FilterFarmer(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) (Farmer []models.RefFarmer, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err = s.FarmerDataAccess(ctx, Farmerfilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterFarmer(ctx, Farmerfilter, pagination)

}

//FilterFarmerBasic :""
func (s *Service) FilterFarmerBasic(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) (Farmer []models.RefBasicFarmer, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err = s.FarmerDataAccess(ctx, Farmerfilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterFarmerBasic(ctx, Farmerfilter, pagination)

}

func (s *Service) FarmerUniquenessCheckRegistration(ctx *models.Context, OrgID string, Param string, Value string) (*models.FarmerUniquinessChk, error) {
	farmer, err := s.Daos.FarmerUniquenessCheckRegistration(ctx, OrgID, Param, Value)
	if err != nil {
		return nil, err
	}
	return farmer, nil
}
func (s *Service) GetContentDisseminationFarmer(ctx *models.Context, cda *models.ContentDataAccess) ([]models.DissiminateFarmer, error) {
	FarmerFilter := new(models.FarmerFilter)
	if cda != nil {
		if len(cda.Organisation) > 0 {
			FarmerFilter.FarmerOrg = cda.Organisation
		}
		if len(cda.Project) > 0 {
			FarmerFilter.Project = cda.Project
		}
		if len(cda.State) > 0 {
			FarmerFilter.State = cda.State
		}
		if len(cda.District) > 0 {
			FarmerFilter.District = cda.District
		}
		if len(cda.Block) > 0 {
			FarmerFilter.Block = cda.Block
		}
		if len(cda.GramPanchayat) > 0 {
			FarmerFilter.GramPanchayat = cda.GramPanchayat
		}
		if len(cda.Village) > 0 {
			FarmerFilter.Village = cda.Village
		}
		FarmerFilter.Status = []string{constants.FARMERSTATUSACTIVE}
		FarmerFilter.SortBy = "name"
		FarmerFilter.SortOrder = 1
	}
	farmer, err := s.Daos.GetContentDisseminationFarmer(ctx, FarmerFilter)
	if err != nil {
		return nil, err
	}
	return farmer, nil

}
func (s *Service) FarmerDataAccess(ctx *models.Context, Farmerfilter *models.FarmerFilter) (err error) {
	if Farmerfilter != nil {

		dataaccess, err := s.Daos.DataAccess(ctx, &Farmerfilter.DataAccess)
		if err != nil {
			return err
		}
		s.Shared.BsonToJSONPrintTag("farmer dataaccess query =>", dataaccess)

		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					Farmerfilter.FarmerOrg = append(Farmerfilter.FarmerOrg, v.ID)
				}
			}
			if len(dataaccess.Projects) > 0 {
				for _, v := range dataaccess.Projects {
					Farmerfilter.Project = append(Farmerfilter.Project, v.Project)
				}
			}
			if len(dataaccess.AccessStates) > 0 {
				for _, v := range dataaccess.AccessStates {
					Farmerfilter.State = append(Farmerfilter.State, v.ID)
				}
			}
			if len(dataaccess.AccessDistricts) > 0 {
				for _, v := range dataaccess.AccessDistricts {
					Farmerfilter.District = append(Farmerfilter.District, v.ID)
				}
			}
			if len(dataaccess.AccessBlocks) > 0 {
				for _, v := range dataaccess.AccessBlocks {
					Farmerfilter.Block = append(Farmerfilter.Block, v.ID)
				}
			}
			if len(dataaccess.AccessVillages) > 0 {
				for _, v := range dataaccess.AccessVillages {
					Farmerfilter.Village = append(Farmerfilter.Village, v.ID)

				}
			}
			if len(dataaccess.AccessGrampanchayats) > 0 {
				for _, v := range dataaccess.AccessGrampanchayats {
					Farmerfilter.GramPanchayat = append(Farmerfilter.GramPanchayat, v.ID)

				}
			}
		}

	}
	return err
}

//GetSingleFarmerWithMobilenoAndOrg :""
func (s *Service) GetSingleFarmerWithMobilenoAndOrg(ctx *models.Context, org string, UniqueID string) (*models.RefFarmer, error) {
	Farmer, err := s.Daos.GetSingleFarmerWithMobilenoAndOrg(ctx, org, UniqueID)
	if err != nil {
		return nil, err
	}
	return Farmer, nil
}

func (s *Service) FarmerExcel(ctx *models.Context, filter *models.FarmerFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterFarmer(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "FarmerReport"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "J1")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "organisation")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "FatherName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "UserName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Gender")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "MobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "State")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Distric")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Block")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Grampanchat")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Village")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "CreateDate")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Ref.FarmerOrg.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.FatherName)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.MobileNumber)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Gender)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.MobileNumber)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Ref.State.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Ref.District.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Ref.Block.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Ref.GramPanchayat.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Ref.Village.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.CreatedDate)
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
func (s *Service) FarmerCsv(ctx *models.Context, filter *models.FarmerFilter, pagination *models.Pagination) (*csv.Writer, error) {

	csv := csv.NewWriter(os.Stdout)
	return csv, nil

}

//FarmerExcelCron
func (s *Service) FarmerExcelCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	farmer := new(models.FarmerFilter)
	excel, err := s.FarmerExcel(ctx, farmer, nil)
	if err != nil {
		log.Println("excel not generated" + err.Error())
	}
	docPathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOC)
	excel.SaveAs(docPathStart + "farmer/farmer.xlsx")
}

//GenerateOTPFarmer :
func (s *Service) GenerateotpFarmerRegistration(ctx *models.Context, farmer *models.Farmer) error {
	if farmer.MobileNumber != "" {
		data, err := s.Daos.GetSingleFarmerWithMobileno(ctx, farmer.MobileNumber)

		if data != nil {
			return errors.New("Farmer Already Registered")
		}
		if err != nil {
			if err.Error() != "farmer not found" {
				return err

			}
		}
	} else {
		return errors.New("please enter the mobilenumber")
	}
	if farmer.FarmerOrg.IsZero() {
		prod, _ := s.Daos.GetactiveProductConfig(ctx, true)
		if prod != nil {
			farmer.FarmerOrg = prod.Orgnisation.OrgnisationID
		}
	}
	//if data != nil {
	//otp, err := s.GenerateOTP(constants.FARMERREGISTERATIOM, farmer.MobileNumber, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return errors.New("Otp Generate Error - " + err.Error())
	// }
	key := fmt.Sprintf("%v_%v", constants.FARMERREGISTERATIOM, farmer.MobileNumber)
	var otp models.Otp
	otp.Otp = "9999"
	err := s.SetValueCacheMemory(key, otp, 1000)
	if err != nil {
		return err
	}

	//text := fmt.Sprintf("Hi %v, /n Otp For Logikoof Reporting App Login is %v .", data.Name, otp)
	msg := fmt.Sprintf(constants.COMMONTEMPLATE, farmer.Name, "NICESSM", "OTP for nicessm registration app", "OTP for NICESSM registration is-("+otp.Otp+")", "https://nicessm.org/")

	err = s.SendSMSV2(ctx, farmer.MobileNumber, msg)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}
	if err == errors.New(constants.INSUFFICIENTBALANCE) {
		return err
	}
	if err == nil {
		smslog := new(models.SmsLog)
		to := models.To{}
		to.No = farmer.MobileNumber
		to.Name = farmer.Name
		to.UserType = "Farmer"
		to.UserName = farmer.FarmerID
		t := time.Now()
		smslog.SentDate = &t
		smslog.Status = constants.SMSLOGSTATUSACTIVE
		smslog.IsJob = false
		smslog.Message = msg
		smslog.SentFor = "Otp"
		smslog.To = to
		err = s.Daos.SaveSmsLog(ctx, smslog)
		if err != nil {
			return errors.New("otp sms not save")
		}
	}
	return nil
}

func (s *Service) FarmerNearBy(ctx *models.Context, farmernb *models.NearBy, pagination *models.Pagination) ([]models.RefFarmer, error) {

	farmers, err := s.Daos.FarmerNearBy(ctx, farmernb, pagination)
	if err != nil {
		return nil, err
	}

	return farmers, nil
}
func (s *Service) FarmerReportExcel(ctx *models.Context, filter *models.FarmerFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterFarmer(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()
	fmt.Println("farmerlength", len(data))
	excel := excelize.NewFile()
	sheet1 := "FarmerReport"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "E1")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "FatherName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "MobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Village")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "FarmerId")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.FatherName)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.MobileNumber)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Ref.Village.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.FarmerID)
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

//RegistrationValidateOTPFarmer :
func (s *Service) RegistrationValidateOTPFarmer(ctx *models.Context, login *models.FarmerOTPLogin) error {

	data, err := s.Daos.GetSingleFarmerWithMobileno(ctx, login.MobileNumber)

	if data != nil {
		return errors.New("Farmer Already Registered")
	}
	if err != nil {
		if err.Error() != "farmer not found" {
			return err

		}
	}

	if login.MobileNumber != "" {
		key := fmt.Sprintf("%v_%v", constants.FARMERREGISTERATIOM, login.MobileNumber)
		otp := new(models.Otp)
		err = s.GetValueCacheMemory(key, otp)
		if err != nil {
			return err
		}
		fmt.Println("Otp===>", otp.Otp)
		if otp.Otp != login.OTP {
			return errors.New("Invaild Otp")
		}
	} else {
		return errors.New("please enter the mobile number")
	}
	if login.FarmerOrg.IsZero() {
		prod, _ := s.Daos.GetactiveProductConfig(ctx, true)
		if prod != nil {
			login.FarmerOrg = prod.Orgnisation.OrgnisationID
		}
	}
	err = s.SaveFarmer(ctx, &login.Farmer)
	if err != nil {
		return err
	}
	err = s.FarmerRegisterSms(ctx, login.MobileNumber, login.Name)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) AddProjectFarmer(ctx *models.Context, Farmerfilter *models.FarmerFilter) (Farmer []models.AddProjectFarmer, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err = s.FarmerDataAccess(ctx, Farmerfilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.AddProjectfarmer(ctx, Farmerfilter)

}

//UpdateFarmerProfileImage : ""
func (s *Service) UpdateFarmerProfileImage(ctx *models.Context, Farmer *models.Farmer) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFarmerProfileImage(ctx, Farmer)
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

//FilterFarmerWithLOcation :""
func (s *Service) FilterFarmerWithLocation(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) (Farmer []models.FarmerLocation, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err = s.FarmerDataAccess(ctx, Farmerfilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterFarmerWithLocation(ctx, Farmerfilter, pagination)

}
func (s *Service) GetWeatherDisseminationFarmer(ctx *models.Context, state string) ([]models.DissiminateFarmer, error) {
	FarmerFilter := new(models.FarmerFilter)
	stateid, err := primitive.ObjectIDFromHex(state)
	if err != nil {
		return nil, err
	}
	var Arraystate []primitive.ObjectID
	Arraystate = append(Arraystate, stateid)
	fmt.Println("Arraystate===>", Arraystate)
	FarmerFilter.State = Arraystate
	FarmerFilter.Status = []string{constants.FARMERSTATUSACTIVE}

	farmer, err := s.Daos.GetContentDisseminationFarmer(ctx, FarmerFilter)
	if err != nil {
		return nil, err
	}
	return farmer, nil

}
func (s *Service) GetDistrictWeatherDisseminationFarmer(ctx *models.Context, district string) ([]models.DissiminateFarmer, error) {
	FarmerFilter := new(models.FarmerFilter)
	districtid, err := primitive.ObjectIDFromHex(district)
	if err != nil {
		return nil, err
	}
	var Arraystate []primitive.ObjectID
	Arraystate = append(Arraystate, districtid)
	fmt.Println("Arraystate===>", Arraystate)
	FarmerFilter.District = Arraystate
	FarmerFilter.Status = []string{constants.FARMERSTATUSACTIVE}

	farmer, err := s.Daos.GetContentDisseminationFarmer(ctx, FarmerFilter)
	if err != nil {
		return nil, err
	}
	return farmer, nil

}
