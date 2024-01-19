package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveStateWeatherAlertDissimination :""
func (s *Service) SaveStateWeatherAlertDissimination(ctx *models.Context, StateWeatherAlertDissimination *models.StateWeatherAlertDissimination) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	StateWeatherAlertDissimination.Status = constants.STATEWEATHERALERTDISSIMINATIONSTATUSACTIVE
	StateWeatherAlertDissimination.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 StateWeatherAlertDissimination.created")
	StateWeatherAlertDissimination.Created = &created
	log.Println("b4 StateWeatherAlertDissimination.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveStateWeatherAlertDissimination(ctx, StateWeatherAlertDissimination)
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
func (s *Service) SaveStateWeatherAlertDissiminationSendNow(ctx *models.Context, dissemination *models.StateWeatherAlert) (*models.DissseminationUserFarmer, error) {
	DissseminationUserFarmer, err := s.DisseminationWeatherAlert(ctx, dissemination, false)
	if err != nil {
		fmt.Println("err", err)
	}
	return DissseminationUserFarmer, nil
}
func (s *Service) DisseminationWeatherAlert(ctx *models.Context, WeatherAlert *models.StateWeatherAlert, isjob bool) (*models.DissseminationUserFarmer, error) {
	fmt.Println("state===>", WeatherAlert.State.ID)
	Users, err := s.GetWeatherDisseminationUser(ctx, WeatherAlert.State.ID.Hex())
	if err != nil {
		return nil, err
	}
	fmt.Println("no.of.user===>", len(Users))
	Farmers, err := s.GetWeatherDisseminationFarmer(ctx, WeatherAlert.State.ID.Hex())
	if err != nil {
		return nil, err
	}
	fmt.Println("no.of.Farmer===>", len(Farmers))
	weatherDissemination := new(models.WeatherDisseminationChennal)
	weatherDissemination.Users = Users
	weatherDissemination.Farmers = Farmers
	weatherDissemination.WeatherAlert = *WeatherAlert
	s.C <- weatherDissemination
	DissseminationUserFarmer := new(models.DissseminationUserFarmer)
	DissseminationUserFarmer.NoOfFarmers = len(Farmers)
	DissseminationUserFarmer.NoOfUsers = len(Users)
	return DissseminationUserFarmer, nil
}
func (s *Service) WeatheralertDissiminationSendNow(ctx *models.Context, Farmers []models.DissiminateFarmer, Users []models.DissiminateUser, WeatherAlert *models.StateWeatherAlert, isjob bool) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		productconfig, err := s.Daos.GetactiveProductConfig(ctx, true)
		if err != nil {
			return err
		}

		fmt.Println("no.of.farmers===>", len(Farmers))
		fmt.Println("no.of.users===>", len(Users))
		if productconfig.WeatherAlert != false {
			if WeatherAlert.WeatherDataAlert.IsSms == "Yes" {
				if len(Farmers) > 0 {
					for _, v := range Farmers {
						// s.C <- k
						if v.MobileNumber != "" {
							fmt.Println("farmer id==>", v.ID)
							farmer, err := s.Daos.GetSingleFarmer(ctx, v.ID.Hex())
							if err != nil {
								// return nil, errors.New("User Not Found")
								log.Println("farmer not found - " + v.ID.Hex())
								continue
							}
							msg := fmt.Sprintf(constants.COMMONTEMPLATE, v.Name, "NICESSM", "weatherAlert -"+WeatherAlert.SeverityType.Name+"", WeatherAlert.WeatherDataAlert.WeatherAlert, productconfig.URL)
							//	msg := fmt.Sprintf(constants.COMMONTEMPLATE, v.Name, "NICESSM", "content ("+content.RecordId+")", contentview, "https://nicessm.org/")
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
								smslog.WeatherAlert = *WeatherAlert
								smslog.SentFor = "WeatherAlert"
								smslog.To = to
								err = s.Daos.SaveSmsLog(ctx, smslog)
								if err != nil {
									return errors.New("contentsisseminatio sms not save")
								}
							}
						}
					}
				}
				if len(Users) > 0 {
					for _, v := range Users {
						if v.MobileNumber != "" {
							user, err := s.Daos.GetSingleUser(ctx, v.ID.Hex())
							if err != nil {
								// return nil, errors.New("User Not Found")
								log.Println("user not found - " + v.ID.Hex())
								continue
							}
							fmt.Println("user mobile=====>", v.MobileNumber)
							msg := fmt.Sprintf(constants.COMMONTEMPLATE, v.Name, "NICESSM", "weatherAlert -"+WeatherAlert.SeverityType.Name+"", WeatherAlert.WeatherDataAlert.WeatherAlert, productconfig.URL)
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
								//smslog.Content = content.ID
								smslog.WeatherAlert = *WeatherAlert
								smslog.SentFor = "WeatherAlert"
								smslog.To = to
								err = s.Daos.SaveSmsLog(ctx, smslog)
								if err != nil {
									return errors.New("contentsisseminatio sms not save")
								}
							}
						}
					}
				}
			}
			if WeatherAlert.WeatherDataAlert.IsWhatsApp == "Yes" {
				//SendWhatsAppText := new(models.SendWhatsAppText2)
				//	WhatsText := new(WhatsAppText)
				if len(Farmers) > 0 {
					for _, v := range Farmers {
						if v.MobileNumber != "" {
							farmer, err := s.Daos.GetSingleFarmer(ctx, v.ID.Hex())
							if err != nil {
								// return nil, errors.New("User Not Found")
								log.Println("farmer not found - " + v.ID.Hex())
								continue
							}
							SendWhatsAppText := new(models.SendWhatsAppText2)
							SendWhatsAppText.MobileNo = v.MobileNumber
							if len(v.MobileNumber) <= 10 {
								SendWhatsAppText.MobileNo = ("91" + v.MobileNumber)
								fmt.Println("mobile===>", SendWhatsAppText.MobileNo)
							}
							SendWhatsAppText.Type = make([]models.WhatsAppText, 0)
							SendWhatsAppText.Type = append(SendWhatsAppText.Type, models.WhatsAppText{"text", WeatherAlert.SeverityType.Name})
							SendWhatsAppText.Type = append(SendWhatsAppText.Type, models.WhatsAppText{"text", WeatherAlert.WeatherDataAlert.WeatherAlert})

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
								WhatsappLog.WeatherAlert = *WeatherAlert
								WhatsappLog.SentFor = "WeatherAlert"
								WhatsappLog.Status = constants.WHATSAPPLOGSTATUSACTIVE
								WhatsappLog.To = to
								err = s.Daos.SaveWhatsappLog(ctx, WhatsappLog)
								if err != nil {
									return errors.New("Weatherdisseminatio Whatsapp not save")
								}
							}
						}
					}
				}
				if len(Users) > 0 {
					for _, v := range Users {
						if v.MobileNumber != "" {
							user, err := s.Daos.GetSingleUser(ctx, v.ID.Hex())
							if err != nil {
								// return nil, errors.New("User Not Found")
								log.Println("user not found - " + v.ID.Hex())
								continue
							}
							SendWhatsAppText := new(models.SendWhatsAppText2)
							SendWhatsAppText.MobileNo = v.MobileNumber
							if len(v.MobileNumber) <= 10 {
								SendWhatsAppText.MobileNo = ("91" + v.MobileNumber)
								fmt.Println("mobile===>", SendWhatsAppText.MobileNo)
							}
							SendWhatsAppText.Type = make([]models.WhatsAppText, 0)
							SendWhatsAppText.Type = append(SendWhatsAppText.Type, models.WhatsAppText{"text", WeatherAlert.SeverityType.Name})
							SendWhatsAppText.Type = append(SendWhatsAppText.Type, models.WhatsAppText{"text", WeatherAlert.WeatherDataAlert.WeatherAlert})

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
								WhatsappLog.WeatherAlert = *WeatherAlert
								WhatsappLog.Status = constants.WHATSAPPLOGSTATUSACTIVE
								WhatsappLog.SentFor = "WeatherAlert"
								WhatsappLog.To = to
								err = s.Daos.SaveWhatsappLog(ctx, WhatsappLog)
								if err != nil {
									return errors.New("Weatherdisseminatio Whatsapp not save")
								}
							}
						}
					}
				}
			}
		}
		if WeatherAlert.WeatherDataAlert.IsNicessm == "Yes" {
			if len(Farmers) > 0 {
				for _, v := range Farmers {
					if v.AppRegistrationToken != "" {
						farmer, err := s.Daos.GetSingleFarmer(ctx, v.ID.Hex())
						if err != nil {
							// return nil, errors.New("User Not Found")
							log.Println("farmer not found - " + v.ID.Hex())
							continue
						}
						var token []string
						token = append(token, v.AppRegistrationToken)
						fmt.Println("appToken===>", v.AppRegistrationToken)
						topic := ""
						tittle := "Hi " + v.Name + " check the following weather alert - " + WeatherAlert.WeatherDataAlert.WeatherAlert
						fmt.Println("notification==>", tittle)
						body := fmt.Sprintf("%v-%v", WeatherAlert.ParameterId.Name, WeatherAlert.SeverityType.Name)
						Body := body
						image := ""
						data := make(map[string]string)
						data["notificationType"] = "ViewSingleWeatherAlert"
						data["id"] = WeatherAlert.ID.Hex()
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
			if len(Users) > 0 {
				for _, v := range Users {
					if v.AppRegistrationToken != "" {
						user, err := s.Daos.GetSingleUser(ctx, v.ID.Hex())
						if err != nil {
							// return nil, errors.New("User Not Found")
							log.Println("user not found - " + v.ID.Hex())
							continue
						}
						var token []string
						token = append(token, v.AppRegistrationToken)
						topic := ""
						tittle := "Hi " + v.Name + " check the following weather alert - " + WeatherAlert.WeatherDataAlert.WeatherAlert
						//	tittle := "content -" + content.RecordId + " -ContentDissemination"
						body := fmt.Sprintf("%v-%v", WeatherAlert.ParameterId.Name, WeatherAlert.SeverityType.Name)
						Body := body
						image := ""
						data := make(map[string]string)
						data["notificationType"] = "ViewSingleWeatherAlert"
						data["id"] = WeatherAlert.ID.Hex()
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
							notifylog.WeatherAlert = *WeatherAlert
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

		t := time.Now()
		StateWeatherAlertDissimination := new(models.StateWeatherAlertDissimination)
		StateWeatherAlertDissimination.Farmers = Farmers
		StateWeatherAlertDissimination.Users = Users
		StateWeatherAlertDissimination.NoOfFarmers = len(Farmers)
		StateWeatherAlertDissimination.NoOfUsers = len(Users)
		StateWeatherAlertDissimination.Date = &t
		StateWeatherAlertDissimination.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION)
		StateWeatherAlertDissimination.Status = constants.STATEWEATHERALERTDISSIMINATIONSTATUSACTIVE
		StateWeatherAlertDissimination.WeatherAlert = *WeatherAlert
		dberr := s.Daos.SaveStateWeatherAlertDissimination(ctx, StateWeatherAlertDissimination)
		if dberr != nil {
			return errors.New("Db Error" + dberr.Error())
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

//UpdateStateWeatherAlertDissimination : ""
func (s *Service) UpdateStateWeatherAlertDissimination(ctx *models.Context, StateWeatherAlertDissimination *models.StateWeatherAlertDissimination) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateStateWeatherAlertDissimination(ctx, StateWeatherAlertDissimination)
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

//EnableStateWeatherAlertDissimination : ""
func (s *Service) EnableStateWeatherAlertDissimination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableStateWeatherAlertDissimination(ctx, UniqueID)
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

//DisableStateWeatherAlertDissimination : ""
func (s *Service) DisableStateWeatherAlertDissimination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableStateWeatherAlertDissimination(ctx, UniqueID)
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

//DeleteStateWeatherAlertDissimination : ""
func (s *Service) DeleteStateWeatherAlertDissimination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteStateWeatherAlertDissimination(ctx, UniqueID)
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

//GetSingleStateWeatherAlertDissimination :""
func (s *Service) GetSingleStateWeatherAlertDissimination(ctx *models.Context, UniqueID string) (*models.RefStateWeatherAlertDissimination, error) {
	StateWeatherAlertDissimination, err := s.Daos.GetSingleStateWeatherAlertDissimination(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return StateWeatherAlertDissimination, nil
}

//FilterStateWeatherAlertDissimination :""
func (s *Service) FilterStateWeatherAlertDissimination(ctx *models.Context, StateWeatherAlertDissiminationfilter *models.StateWeatherAlertDissiminationFilter, pagination *models.Pagination) (StateWeatherAlertDissimination []models.RefStateWeatherAlertDissimination, err error) {
	return s.Daos.FilterStateWeatherAlertDissimination(ctx, StateWeatherAlertDissiminationfilter, pagination)
}
func (s *Service) GetSingleWeatherAlertFarmerUserCount(ctx *models.Context, UniqueID string) (*models.ContentDissiminateUserAndFarmer, error) {
	Users, err := s.GetWeatherDisseminationUser(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	Farmers, err := s.GetWeatherDisseminationFarmer(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	StateWeatherAlertDissimination := new(models.ContentDissiminateUserAndFarmer)
	StateWeatherAlertDissimination.Farmers = Farmers
	StateWeatherAlertDissimination.Users = Users
	StateWeatherAlertDissimination.FarmersCount = len(Farmers)
	StateWeatherAlertDissimination.UsersCount = len(Users)
	return StateWeatherAlertDissimination, nil
}
func (s *Service) FilterStateWeatherAlertDissiminationReport(ctx *models.Context, StateWeatherAlertDissiminationfilter *models.StateWeatherAlertDissiminationFilter, pagination *models.Pagination) (StateWeatherAlertDissimination []models.RefStateWeatherAlertDissimination, err error) {
	return s.Daos.FilterStateWeatherAlertDissiminationReport(ctx, StateWeatherAlertDissiminationfilter, pagination)
}
func (s *Service) FilterStateWeatherAlertDissiminationReportExcel(ctx *models.Context, filter *models.StateWeatherAlertDissiminationFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterStateWeatherAlertDissiminationReport(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "WeatherAlertDisseminationReport"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Message")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Alert")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Value")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Min")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Max")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "No.Of.Farmers")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "No.Of.Users")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		var date string
		if v.Date != nil {
			date = fmt.Sprintf("%v-%v-%v", v.Date.Day(), v.Date.Month(), v.Date.Year())
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), date)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.WeatherAlert.Tittle)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v-%v", v.WeatherAlert.ParameterId.Name, v.WeatherAlert.SeverityType.Name))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.WeatherAlert.Value)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.WeatherAlert.WeatherDataAlert.Min)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.WeatherAlert.WeatherDataAlert.Max)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.NoOfFarmers)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.NoOfUsers)
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
