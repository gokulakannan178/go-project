package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/app"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDissemination :""
func (s *Service) SaveDissemination(ctx *models.Context, dissemination *models.Dissemination) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	dissemination.IsSent = false
	dissemination.Status = constants.DISSEMINATIONSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	dissemination.DateOfDissemination = &t
	created.By = constants.SYSTEM
	log.Println("b4 dissemination.created")
	dissemination.Created = &created
	log.Println("b4 dissemination.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDissemination(ctx, dissemination)
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

//SaveSendNow :""
func (s *Service) SaveSendNow(ctx *models.Context, dissemination *models.Dissemination) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	dissemination.IsSent = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	dissemination.Status = constants.DISSEMINATIONSTATUSACTIVE
	dissemination.DateOfDissemination = &t
	created.By = constants.SYSTEM
	log.Println("b4 dissemination.created")
	dissemination.Created = &created
	log.Println("b4 dissemination.created")
	if len(dissemination.Users) > 0 {
		dissemination.FarmersCount = len(dissemination.Farmers)
	}
	if len(dissemination.Farmers) > 0 {
		dissemination.UsersCount = len(dissemination.Users)
	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDissemination(ctx, dissemination)

		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		err := s.DisseminateContent(ctx, dissemination, false)
		if err != nil {
			fmt.Println("err", err)
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

//SaveSendLater :""
func (s *Service) SaveSendLater(ctx *models.Context, dissemination *models.Dissemination) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	dissemination.IsSent = false
	t := time.Now()
	created := models.Created{}
	created.On = &t
	dissemination.DateCreated = &t
	created.By = constants.SYSTEM
	log.Println("b4 dissemination.created")
	dissemination.Created = &created
	log.Println("b4 dissemination.created")
	if dissemination.DateOfDissemination == nil {
		return errors.New("dateofdissemination not valid")
	}
	// if len(dissemination.Users) == 0 {
	// 	return errors.New("users not valid")
	// }
	// if len(dissemination.Farmers) == 0 {
	// 	return errors.New("farmers not valid")
	// }
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDissemination(ctx, dissemination)

		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		//	err := s.DisseminateContent(ctx, dissemination)
		// if err != nil {
		// 	fmt.Println("err", err)
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

//UpdateDissemination: ""
func (s *Service) UpdateDissemination(ctx *models.Context, dissemination *models.Dissemination) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDissemination(ctx, dissemination)
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

//EnableDissemination : ""
func (s *Service) EnableDissemination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDissemination(ctx, UniqueID)
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

//DisableDissemination : ""
func (s *Service) DisableDissemination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDissemination(ctx, UniqueID)
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

//DeleteDissemination : ""
func (s *Service) DeleteDissemination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDissemination(ctx, UniqueID)
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

//GetSingleDissemination :""
func (s *Service) GetSingleDissemination(ctx *models.Context, UniqueID string) (*models.RefDissemination, error) {
	Dissemination, err := s.Daos.GetSingleDissemination(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Dissemination, nil
}

//FilterDissemination :""
func (s *Service) FilterDissemination(ctx *models.Context, disseminationfilter *models.DisseminationFilter, pagination *models.Pagination) (dissemination []models.RefDissemination, err error) {
	return s.Daos.FilterDissemination(ctx, disseminationfilter, pagination)
}
func (s *Service) DisseminateContent(ctx *models.Context, dissemination *models.Dissemination, isjob bool) error {

	content, err := s.Daos.GetSingleContent(ctx, dissemination.Content.Hex())
	if err != nil {

		return err
	}
	//getusers
	userfilter := new(models.UserFilter)
	if len(dissemination.Users) < 0 {
		userfilter.OrganisationID = []primitive.ObjectID{
			content.Organisation,
		}
		userfilter.Project = []primitive.ObjectID{
			content.Project,
		}
	}
	productconfig, err := s.Daos.GetactiveProductConfig(ctx, true)
	if err != nil {
		return err
	}
	contentview := "please click following link to view the content -" + productconfig.UIURL + "content/" + content.ID.Hex() + " "
	contentview2 := "please click following link to view the content -" + " " + productconfig.UIURL + "content/" + content.ID.Hex() + " "
	//getfarmers
	fmt.Println("contentview", contentview)
	switch dissemination.Mode {
	case constants.DISSEMINATIONMODESMS:
		if len(dissemination.Farmers) > 0 {
			for _, v := range dissemination.Farmers {
				fmt.Println("farmer id==>", v.ID)
				farmer, err := s.Daos.GetSingleFarmer(ctx, v.ID.Hex())
				if err != nil {
					return errors.New("Farmer Not Found")
				}
				//	msg := fmt.Sprintf(constants.COMMONTEMPLATE, v.Name, "NICESSM", "content ("+content.RecordId+")", contentview, "https://nicessm.org/")
				msg := fmt.Sprintf(constants.COMMONTEMPLATE, v.Name, "NICESSM", "content -"+content.RecordId+"", contentview, productconfig.URL)
				fmt.Println("Farmer mobile=====>", v.MobileNumber)
				err = s.SendSMSV2(ctx, v.MobileNumber, msg)
				if err != nil {
					log.Println(v.MobileNumber + " " + err.Error())
				}
				if err == errors.New(constants.INSUFFICIENTBALANCE) {
					return err
				}
				if err == nil {
					smslog := new(models.SmsLog)
					to := models.To{}
					to.No = v.MobileNumber
					to.Name = v.Name
					to.UserName = v.Name
					to.UserType = "Farmer"
					to.UserId = v.ID
					to.State = farmer.Ref.State.Name
					to.StateCode = farmer.State
					to.District = farmer.Ref.District.Name
					to.DistricCode = farmer.District
					to.Block = farmer.Ref.Block.Name
					to.BlockCode = farmer.Block
					to.GramPanchayat = farmer.Ref.GramPanchayat.Name
					to.GramPanchayatCode = farmer.GramPanchayat
					to.Village = farmer.Ref.Village.Name
					to.VillageCode = farmer.Village
					to.Gender = farmer.Gender
					t := time.Now()
					smslog.SentDate = &t
					smslog.Status = constants.SMSLOGSTATUSACTIVE
					smslog.IsJob = isjob
					smslog.Message = msg
					smslog.Content = content.ID
					smslog.ContentRecordId = content.RecordId
					smslog.SentFor = "Content"
					smslog.To = to
					err = s.Daos.SaveSmsLog(ctx, smslog)
					if err != nil {
						return errors.New("contentsisseminatio sms not save")
					}
				}
			}
		}
		if len(dissemination.Users) > 0 {
			for _, v := range dissemination.Users {
				user, err := s.Daos.GetSingleUser(ctx, v.ID.Hex())
				if err != nil {
					return errors.New("User Not Found")
				}
				//	msg := fmt.Sprintf(constants.COMMONTEMPLATE, v.Name, "NICESSM", "content ("+content.RecordId+")", contentview, "https://nicessm.org/")
				msg := fmt.Sprintf(constants.COMMONTEMPLATE, v.Name, "NICESSM", "content -"+content.RecordId+"", contentview, productconfig.URL)
				fmt.Println("user mobile=====>", v.MobileNumber)
				err = s.SendSMSV2(ctx, v.MobileNumber, msg)
				if err != nil {
					log.Println(v.MobileNumber + " " + err.Error())
				}
				if err == errors.New(constants.INSUFFICIENTBALANCE) {
					return err
				}
				if err == nil {
					smslog := new(models.SmsLog)
					to := models.To{}
					to.No = v.MobileNumber
					to.Name = v.Name
					to.UserName = v.Name
					to.UserType = "User"
					to.UserId = v.ID
					to.State = user.Ref.State.Name
					to.StateCode = user.StateCode
					to.District = user.Ref.District.Name
					to.DistricCode = user.DistrictCode
					to.Block = user.Ref.Block.Name
					to.BlockCode = user.BlockCode
					to.GramPanchayat = user.Ref.Grampanchayat.Name
					to.GramPanchayatCode = user.GrampanchayatCode
					to.Village = user.Ref.Village.Name
					to.VillageCode = user.VillageCode
					to.Gender = user.Gender
					t := time.Now()
					smslog.SentDate = &t
					smslog.IsJob = isjob
					smslog.Status = constants.SMSLOGSTATUSACTIVE
					smslog.Message = msg
					smslog.Content = content.ID
					smslog.ContentRecordId = content.RecordId
					smslog.SentFor = "Content"
					smslog.To = to
					err = s.Daos.SaveSmsLog(ctx, smslog)
					if err != nil {
						return errors.New("contentsisseminatio sms not save")
					}
				}
			}
		}
	case constants.DISSEMINATIONMODEWHATSAPP:
		//SendWhatsAppText := new(models.SendWhatsAppText2)
		//	WhatsText := new(WhatsAppText)
		if len(dissemination.Farmers) > 0 {
			for _, v := range dissemination.Farmers {
				farmer, err := s.Daos.GetSingleFarmer(ctx, v.ID.Hex())
				if err != nil {
					return errors.New("Farmer Not Found")
				}
				SendWhatsAppText := new(models.SendWhatsAppText2)
				SendWhatsAppText.MobileNo = v.MobileNumber
				if len(v.MobileNumber) <= 10 {
					SendWhatsAppText.MobileNo = ("91" + v.MobileNumber)
					fmt.Println("mobile===>", SendWhatsAppText.MobileNo)
				}
				SendWhatsAppText.Type = make([]models.WhatsAppText, 0)
				SendWhatsAppText.Type = append(SendWhatsAppText.Type, models.WhatsAppText{"text", content.RecordId})
				SendWhatsAppText.Type = append(SendWhatsAppText.Type, models.WhatsAppText{"text", contentview2})

				err = s.WhatsAppSingle(ctx, SendWhatsAppText)
				if err != nil {
					log.Println(v.MobileNumber + " " + err.Error())
				}
				if err == nil {
					WhatsappLog := new(models.WhatsappLog)
					to := models.ToWhatsappLog{}
					to.No = v.MobileNumber
					to.Name = v.Name
					to.UserName = v.Name
					to.UserType = "Farmer"
					to.UserId = v.ID
					to.State = farmer.Ref.State.Name
					to.StateCode = farmer.State
					to.District = farmer.Ref.District.Name
					to.DistricCode = farmer.District
					to.Block = farmer.Ref.Block.Name
					to.BlockCode = farmer.Block
					to.GramPanchayat = farmer.Ref.GramPanchayat.Name
					to.GramPanchayatCode = farmer.GramPanchayat
					to.Village = farmer.Ref.Village.Name
					to.VillageCode = farmer.Village
					to.Gender = farmer.Gender
					t := time.Now()
					WhatsappLog.SentDate = &t
					WhatsappLog.IsJob = isjob
					WhatsappLog.Message = contentview
					WhatsappLog.SentFor = "Content"
					WhatsappLog.Status = constants.WHATSAPPLOGSTATUSACTIVE
					WhatsappLog.To = to
					err = s.Daos.SaveWhatsappLog(ctx, WhatsappLog)
					if err != nil {
						return errors.New("contentsisseminatio sms not save")
					}
				}
			}
		}
		if len(dissemination.Users) > 0 {
			for _, v := range dissemination.Users {
				user, err := s.Daos.GetSingleUser(ctx, v.ID.Hex())
				if err != nil {
					return errors.New("User Not Found")
				}
				SendWhatsAppText := new(models.SendWhatsAppText2)
				SendWhatsAppText.MobileNo = v.MobileNumber
				if len(v.MobileNumber) <= 10 {
					SendWhatsAppText.MobileNo = ("91" + v.MobileNumber)
					fmt.Println("mobile===>", SendWhatsAppText.MobileNo)
				}
				SendWhatsAppText.Type = make([]models.WhatsAppText, 0)
				SendWhatsAppText.Type = append(SendWhatsAppText.Type, models.WhatsAppText{"text", content.RecordId})
				SendWhatsAppText.Type = append(SendWhatsAppText.Type, models.WhatsAppText{"text", contentview2})

				err = s.WhatsAppSingle(ctx, SendWhatsAppText)
				if err != nil {
					log.Println(v.MobileNumber + " " + err.Error())
				}
				if err == nil {
					WhatsappLog := new(models.WhatsappLog)
					to := models.ToWhatsappLog{}
					to.No = v.MobileNumber
					to.Name = v.Name
					to.UserName = v.Name
					to.UserType = "User"
					to.UserId = v.ID
					to.State = user.Ref.State.Name
					to.StateCode = user.StateCode
					to.District = user.Ref.District.Name
					to.DistricCode = user.DistrictCode
					to.Block = user.Ref.Block.Name
					to.BlockCode = user.BlockCode
					to.GramPanchayat = user.Ref.Grampanchayat.Name
					to.GramPanchayatCode = user.GrampanchayatCode
					to.Village = user.Ref.Village.Name
					to.VillageCode = user.VillageCode
					to.Gender = user.Gender
					t := time.Now()
					WhatsappLog.SentDate = &t
					WhatsappLog.IsJob = isjob
					WhatsappLog.Message = contentview
					WhatsappLog.Status = constants.WHATSAPPLOGSTATUSACTIVE
					WhatsappLog.SentFor = "Content"
					WhatsappLog.To = to
					err = s.Daos.SaveWhatsappLog(ctx, WhatsappLog)
					if err != nil {
						return errors.New("contentsisseminatio sms not save")
					}
				}
			}
		}
	case constants.DISSEMINATIONMODENOTTIFICATION:
		if len(dissemination.Farmers) > 0 {

			for _, v := range dissemination.Farmers {
				if v.AppRegistrationToken != "" {
					farmer, err := s.Daos.GetSingleFarmer(ctx, v.ID.Hex())
					if err != nil {
						return errors.New("Farmer Not Found")
					}
					var token []string
					token = append(token, v.AppRegistrationToken)
					fmt.Println("appToken===>", v.AppRegistrationToken)
					topic := ""
					tittle := "Hi " + v.Name + " check the following content - " + content.RecordId
					fmt.Println("notification==>", tittle)
					Body := content.ContentTitle
					image := ""
					data := make(map[string]string)
					data["notificationType"] = "ViewSingleContent"
					data["id"] = content.ID.Hex()
					err = s.SendNotification(topic, tittle, Body, image, token, data)
					if err != nil {
						log.Println(v.MobileNumber + " " + err.Error())
					}
					if err == nil {
						t := time.Now()
						to := new(models.ToNotificationLog)
						notifylog := new(models.NotificationLog)
						to.AppRegistrationToken = v.AppRegistrationToken
						to.Name = v.Name
						to.UserName = v.ID
						to.UserType = "Farmer"
						to.UserId = v.ID
						to.State = farmer.Ref.State.Name
						to.StateCode = farmer.State
						to.District = farmer.Ref.District.Name
						to.DistricCode = farmer.District
						to.Block = farmer.Ref.Block.Name
						to.BlockCode = farmer.Block
						to.GramPanchayat = farmer.Ref.GramPanchayat.Name
						to.GramPanchayatCode = farmer.GramPanchayat
						to.Village = farmer.Ref.Village.Name
						to.VillageCode = farmer.Village
						to.Gender = farmer.Gender
						notifylog.Body = Body
						notifylog.Tittle = tittle
						notifylog.Topic = topic
						notifylog.Image = image
						notifylog.IsJob = false
						notifylog.Message = Body
						notifylog.SentDate = &t
						notifylog.SentFor = topic
						notifylog.Data = data
						notifylog.Status = "Active"
						notifylog.To = *to
						err = s.Daos.SaveNotificationLog(ctx, notifylog)
						if err != nil {
							return err
						}
					}
				}
			}
		}
		if len(dissemination.Users) > 0 {
			for _, v := range dissemination.Users {
				if v.AppRegistrationToken != "" {
					user, err := s.Daos.GetSingleUser(ctx, v.ID.Hex())
					if err != nil {
						return errors.New("User Not Found")
					}
					var token []string
					token = append(token, v.AppRegistrationToken)
					topic := ""
					tittle := "Hi " + v.Name + " check the following content - " + content.RecordId
					//	tittle := "content -" + content.RecordId + " -ContentDissemination"
					Body := dissemination.Message
					image := ""
					data := make(map[string]string)
					data["notificationType"] = "ViewSingleContent"
					data["id"] = content.ID.Hex()
					err = s.SendNotification(topic, tittle, Body, image, token, data)
					if err != nil {
						log.Println(v.MobileNumber + " " + err.Error())
					}
					if err == nil {
						t := time.Now()
						to := new(models.ToNotificationLog)
						notifylog := new(models.NotificationLog)
						to.AppRegistrationToken = v.AppRegistrationToken
						to.Name = v.Name
						to.UserName = v.ID
						to.UserId = v.ID
						to.UserType = "User"
						to.State = user.Ref.State.Name
						to.StateCode = user.StateCode
						to.District = user.Ref.District.Name
						to.DistricCode = user.DistrictCode
						to.Block = user.Ref.Block.Name
						to.BlockCode = user.BlockCode
						to.GramPanchayat = user.Ref.Grampanchayat.Name
						to.GramPanchayatCode = user.GrampanchayatCode
						to.Village = user.Ref.Village.Name
						to.VillageCode = user.VillageCode
						to.Gender = user.Gender
						notifylog.Body = Body
						notifylog.Tittle = tittle
						notifylog.Topic = topic
						notifylog.Image = image
						notifylog.IsJob = false
						notifylog.Message = Body
						notifylog.Data = data
						notifylog.SentDate = &t
						notifylog.SentFor = v.MobileNumber
						notifylog.Status = "Active"
						notifylog.To = *to
						err = s.Daos.SaveNotificationLog(ctx, notifylog)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

//DisseminationPDF
func (s *Service) DisseminationPDF(ctx *models.Context, disseminationfilter *models.DisseminationFilter, pagination *models.Pagination) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.FilterDissemination(ctx, disseminationfilter, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	//productConfigUniqueID := "6176962a9dac3d102e979b54"
	//productConfig, err := s.Daos.GetactiveProductConfig(ctx, true)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["dissemination"] = data
	m2["currentDate"] = time.Now()
	m2["mod"] = func(a, b int) bool {
		if a%b == 0 {
			return true
		}
		return false
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	//pdfdata.Config = productConfig.ProductConfig

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//fmt.Println("this is yuva", templatePathStart)
	//html template path
	templatePath := templatePathStart + "dissemination.html"
	err = r.ParseTemplate(templatePath, pdfdata)
	if err != nil {
		return nil, err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	fmt.Println(ok, "pdf generated successfully")

	return file, nil
}

//SendLaterCron
func (s *Service) SendLaterCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	//Dissemination := new(models.Dissemination)
	Disseminations, err := s.Daos.SendLaterDissemination(ctx)
	if err != nil {
		log.Println("dissemination not found" + err.Error())
	}
	for _, v := range Disseminations {
		err := s.DisseminateContent(ctx, &v, true)
		if err != nil {
			log.Println("sms not sended" + err.Error())
		}
	}
}
func (s *Service) DisseminationReport(ctx *models.Context, disseminationfilter *models.DisseminationReportFilter) (dissemination []models.RefDisseminationReport, err error) {
	return s.Daos.DisseminationReport(ctx, disseminationfilter)
}
func (s *Service) DisseminationReportExcel(ctx *models.Context, filter *models.DisseminationReportFilter) (*excelize.File, error) {
	t := time.Now()
	data, err := s.DisseminationReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "Dissemination Report"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "State")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "District")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Farmers")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Sms")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Voice")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Video")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Poster")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Docments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Total Disseminated")

	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Ref.State.Name)
		for _, v2 := range v.Districts {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.District.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.Farmers)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.Sms)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v2.Voice)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v2.Video)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v2.Poster)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v2.Document)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v2.Dessiminations)
			rowNo++
		}
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
