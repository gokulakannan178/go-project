package services

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUser :""
func (s *Service) SaveUserwithtransaction(ctx *models.Context, user *models.User) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.SaveUserwithouttransaction(ctx, user)
		if err != nil {
			return err
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

//SaveUser :""
func (s *Service) SaveUserwithouttransaction(ctx *models.Context, user *models.User) error {
	log.Println("transaction start")
	//Start Transaction

	if user.UserName == "" {
		user.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
	}
	user.Status = constants.USERSTATUSACTIVE
	if user.IsSelfRegistration == true {
		user.Status = constants.USERSTATUSINIT
	}
	user.Password = "#nature32" //Default Password
	t := time.Now()
	user.CreatedDate = &t
	user.Name = user.FirstName + user.Lastname
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	user.Created = created
	log.Println("b4 user.created")
	prod, er2 := s.Daos.GetactiveProductConfig(ctx, true)
	if er2 != nil {
		return er2
	}
	if prod == nil {
		return errors.New("product config is nil")
	}
	if prod.ValidateUserregistration {
		err := s.ValidateUser(ctx, user)
		if err != nil {
			return err
		}

	}
	dberr := s.Daos.SaveUser(ctx, user)
	if dberr != nil {

		return errors.New("Db Error" + dberr.Error())
	}

	if len(user.Project) > 0 {
		dberr = s.Daos.SaveMultipleProjectUser(ctx, user)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
	}
	return nil
}
func (s *Service) ValidateUser(ctx *models.Context, user *models.User) error {
	switch user.Type {
	case constants.USERTYPECALLCENTERAGENT:
		_, err := user.ValidateCallcenterAgent()
		if err != nil {
			return err
		}
	case constants.USERTYPECONTENTCREATOR:
		_, err := user.ValidateContentCreator()
		if err != nil {
			return err
		}
	case constants.USERTYPECONTENTMANAGER:
		_, err := user.ValidateContentManager()
		if err != nil {
			return err
		}
	case constants.USERTYPECONTENTPROVIDER:
		_, err := user.ValidateContentProvider()
		if err != nil {
			return err
		}
	case constants.USERTYPECONTENTDISSEMINATOR:
		_, err := user.ValidateContentDisseminator()
		if err != nil {
			return err
		}
	case constants.USERTYPEFIELDAGENT:
		_, err := user.ValidateFieldAgent()
		if err != nil {
			return err
		}
	case constants.USERTYPELANGUAGETRANSLATOR:
		_, err := user.ValidateLanguageTranslator()
		if err != nil {
			return err
		}
	case constants.USERTYPELANGUAGEAPPROVER:
		_, err := user.ValidateLanguageApprover()
		if err != nil {
			return err
		}
	case constants.USERTYPEMANAGEMENT:
		_, err := user.ValidateManagement()
		if err != nil {
			return err
		}
	case constants.USERTYPEMODERATOR:
		_, err := user.ValidateModerator()
		if err != nil {
			return err
		}
	case constants.USERTYPESUBJECTMATTEREXPERT:
		_, err := user.ValidateSubjectMatterExpert()
		if err != nil {
			return err
		}
	case constants.USERTYPESYSTEMADMIN:
		_, err := user.ValidateSystemAdmin()
		if err != nil {
			return err
		}
	case constants.USERTYPEVISTORVIEWER:
		_, err := user.ValidateSystemAdmin()
		if err != nil {
			return err
		}
	case constants.USERTYPETRAINER:
		_, err := user.ValidateTrainer()
		if err != nil {
			return err
		}
	case constants.USERTYPEFIELDAGENTLEAD:
		_, err := user.ValidateFieldAgentLead()
		if err != nil {
			return err
		}
	// case constants.USERTYPEDISTRICTADMIN:
	// 	_, err := user.ValidateDistrictAdmin()
	// 	if err != nil {
	// 		return err
	// 	}
	default:
		return errors.New("invalid user type")
	}
	return nil
}

//UpdateUser : ""
func (s *Service) UpdateUser(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUser(ctx, user)
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

//EnableUser : ""
func (s *Service) EnableUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUser(ctx, UniqueID)
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

//DisableUser : ""
func (s *Service) DisableUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUser(ctx, UniqueID)
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

//DeleteUser : ""
func (s *Service) DeleteUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUser(ctx, UniqueID)
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

//GetSingleUser :""
func (s *Service) GetSingleUser(ctx *models.Context, UniqueID string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//GetSingleUserWithUserName : ""
func (s *Service) GetSingleUserWithUserName(ctx *models.Context, UniqueID string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUserWithUserName(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//GetMobileValidation :""
func (s *Service) GetMobileValidation(ctx *models.Context, MobileNo string) error {
	Validate := s.Daos.GetMobileValidation(ctx, MobileNo)
	if Validate != nil {
		return Validate
	}

	return nil

}

//FilterUser :""
func (s *Service) FilterUser(ctx *models.Context, userfilter *models.UserFilter, pagination *models.Pagination) (user []models.RefUser, err error) {
	err = s.UserDataAccess(ctx, userfilter)
	if err != nil {
		return nil, err
	}
	fmt.Println("district====>", userfilter.Districts)
	return s.Daos.FilterUser(ctx, userfilter, pagination)
}

//ResetUserPassword : ""
func (s *Service) ResetUserPassword(ctx *models.Context, userName string) error {
	return s.Daos.ResetUserPassword(ctx, userName, "#nature32")
}

//ChangePassword : ""
func (s *Service) ChangePassword(ctx *models.Context, cp *models.UserChangePassword) (bool, string, error) {
	user, err := s.Daos.GetSingleUserWithUserName(ctx, cp.UserName)
	if err != nil {
		return false, "", err
	}
	if user.Password != cp.OldPassword {
		return false, "Wrong Password", nil
	}
	err = s.Daos.ResetUserPassword(ctx, cp.UserName, cp.NewPassword)
	if err != nil {
		return false, "", err
	}

	d := make(map[string]interface{})
	d["UserName"] = user.UserName
	d["Name"] = user.Name
	d["URL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.APIBASEURL) + user.Profile
	d["LoginURL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.BASEURL) + s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOGINURL)
	d["ContactUsURL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.BASEURL) + s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.CONTACTUSURL)

	templateID := "successfull_registration.html"
	err = s.SendEmailWithTemplate("Kochas Municipality - Password Changed Successfully", []string{"solomon2261993@gmail.com"}, "templates/"+templateID, d)
	if err != nil {
		log.Println("Email not sent - " + err.Error())
		// return errors.New("Unable to send email - " + err.Error())
	}
	return true, "", nil
}

//ForgetPasswordValidateOTP : ""
func (s *Service) ForgetPasswordValidateOTP(ctx *models.Context, UniqueID string, otp string) (string, error) {
	user, err := s.Daos.GetSingleUserWithMobileNo(ctx, UniqueID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user is nil")
	}
	token, _ := s.GenerateOTP(constants.OTPSCENARIOTOKEN, user.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return "", err
	// }
	key := fmt.Sprintf("%v_%v", constants.OTPSCENARIOPASSWORD, user.Mobile)
	otps := new(models.Otp)
	err = s.GetValueCacheMemory(key, otps)
	if err != nil {
		return "", err
	}
	fmt.Println("Otp===>", otps.Otp)
	if otps.Otp != otp {
		return "", errors.New("Invaild Otp")
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(token))

	fmt.Println(sEnc)
	return sEnc, nil
}

//ForgetPasswordGenerateOTP : ""
func (s *Service) ForgetPasswordGenerateOTP(ctx *models.Context, UniqueID string) error {
	user, err := s.Daos.GetSingleUserWithMobileNo(ctx, UniqueID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user is nil")
	}
	// otp, err := s.GenerateOTP(constants.OTPSCENARIOPASSWORD, user.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return err
	// }
	key := fmt.Sprintf("%v_%v", constants.OTPSCENARIOPASSWORD, user.Mobile)
	var otp models.Otp
	otp.Otp = "9999"
	err = s.SetValueCacheMemory(key, otp, 1000)
	if err != nil {
		return err
	}
	// msg := "Use otp " + otp + " for municipal forget password. municipal doesnt ask otp to be shared with anyone"
	// err = s.SendSMS(user.Mobile, msg)
	msg := fmt.Sprintf(constants.COMMONTEMPLATE, user.Name, "NICESSM", "otp for forgot password", "OTP for NICESSM forgot password is-"+otp.Otp+"", "https://nicessm.org/")

	err = s.SendSMSV2(ctx, user.Mobile, msg)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}
	if err == errors.New(constants.INSUFFICIENTBALANCE) {
		return err
	}
	if err == nil {
		smslog := new(models.SmsLog)
		to := models.To{}
		to.No = user.Mobile
		to.Name = user.Name
		to.UserType = "user"
		to.UserName = user.UserName
		t := time.Now()
		smslog.SentDate = &t
		smslog.IsJob = false
		smslog.Message = msg
		smslog.Status = constants.SMSLOGSTATUSACTIVE
		smslog.SentFor = "Otp"
		smslog.To = to
		err = s.Daos.SaveSmsLog(ctx, smslog)
		if err != nil {
			return errors.New("otp sms not save")
		}
	}
	var sendmailto []string
	sendmailto = append(sendmailto, user.Email)
	err = s.SendEmail("NICESSM-OTP Generation For Forget Password", sendmailto, msg)
	if err != nil {
		return errors.New("email Sending Error - " + err.Error())
	}
	if err == nil {
		emaillog := new(models.EmailLog)
		to2 := models.ToEmailLog{}
		to2.Email = user.Email
		to2.Name = user.UserName
		to2.UserName = user.UserName
		to2.UserType = "user"
		t := time.Now()
		emaillog.SentDate = &t
		emaillog.IsJob = false
		emaillog.Message = msg
		emaillog.SentFor = "login"
		emaillog.Status = "Active"
		emaillog.To = to2
		err = s.Daos.SaveEmailLog(ctx, emaillog)
		if err != nil {
			return errors.New("login email not save")
		}
	}
	fmt.Println(err)
	return nil
}

//UpdateUser : ""
func (s *Service) PasswordUpdate(ctx *models.Context, user *models.RefPassword) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.PasswordUpdate(ctx, user)
		if err != nil {
			return err
		}
		users, err := s.Daos.GetSingleUserWithMobileNo(ctx, user.UniqueID)
		if err != nil {
			return err
		}
		loginurl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.LOGINURLV2)
		//	msg2 := fmt.Sprintf(constants.COMMONTEMPLATE, users.Name, "NICESSM", "otp for forgot password", "OTP for NICESSM forgot password is-"+otp+"", "https://nicessm.org/")

		msg := fmt.Sprintf(constants.COMMONTEMPLATE, users.Name, "NICESSM", "user password updated", "please login:"+loginurl+"", "https://nicessm.org/")
		err = s.SendSMSV2(ctx, users.Mobile, msg)
		if err != nil {
			log.Println(users.Mobile + " " + err.Error())
		}
		if err == errors.New(constants.INSUFFICIENTBALANCE) {
			return err
		}
		if err == nil {
			smslog := new(models.SmsLog)
			to := models.To{}
			to.No = users.Mobile
			to.Name = users.UserName
			to.UserType = "user"
			smslog.IsJob = false
			t := time.Now()
			smslog.SentDate = &t
			smslog.Message = msg
			smslog.SentFor = "password update"
			smslog.Status = constants.SMSLOGSTATUSACTIVE
			smslog.To = to
			err = s.Daos.SaveSmsLog(ctx, smslog)
			if err != nil {
				return errors.New("login sms not save")
			}
		}
		var sendmailto []string
		sendmailto = append(sendmailto, users.Email)
		err = s.SendEmail("NICESSM-Password Updated", sendmailto, msg)
		if err != nil {
			return errors.New("email Sending Error - " + err.Error())
		}
		if err == nil {
			emaillog := new(models.EmailLog)
			to2 := models.ToEmailLog{}
			to2.Email = users.Email
			to2.Name = users.UserName
			to2.UserName = users.UserName
			to2.UserType = "user"
			t := time.Now()
			emaillog.SentDate = &t
			emaillog.IsJob = false
			emaillog.Message = msg
			emaillog.SentFor = "login"
			emaillog.Status = "Active"
			emaillog.To = to2
			err = s.Daos.SaveEmailLog(ctx, emaillog)
			if err != nil {
				return errors.New("login email not save")
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

//UpdateUser : ""
func (s *Service) UserCollectionLimit(ctx *models.Context, UserName string, user *models.CollectionLimit) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UserCollectionLimit(ctx, UserName, user)
		if err != nil {
			return err
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

//UpdateOrganisation : ""
func (s *Service) UpdateUserTypeV2(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserTypeV2(ctx, user)
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

//UpdateUserPassword : ""
func (s *Service) UpdateUserPassword(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserPassword(ctx, user)
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
func (s *Service) UserUniquenessCheckRegistration(ctx *models.Context, OrgID string, Param string, Value string) (*models.UserUniquinessChk, error) {
	user, err := s.Daos.UserUniquenessCheckRegistration(ctx, OrgID, Param, Value)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//ApprovedUser : ""
func (s *Service) ApprovedUser(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	user.Status = constants.USERSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUser(ctx, user)
		if err != nil {
			return err
		}
		loginurl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.LOGINURLV2)

		msg := fmt.Sprintf(constants.COMMONTEMPLATE, user.UserName, "NICESSM", "Registeration is Approved", "please login :"+loginurl+"", "https://nicessm.org/")

		var sendmailto []string
		sendmailto = append(sendmailto, user.Email)
		err = s.SendEmail("NICESSM-Registeration Approved", sendmailto, msg)
		if err != nil {
			return errors.New("email Sending Error - " + err.Error())
		}
		if err == nil {
			emaillog := new(models.EmailLog)
			to2 := models.ToEmailLog{}
			to2.Email = user.Email
			to2.Name = user.UserName
			to2.UserName = user.UserName
			to2.UserType = "user"
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			emaillog.SentDate = &t
			emaillog.IsJob = false
			emaillog.Message = msg
			emaillog.SentFor = "login"
			emaillog.Status = "Active"
			emaillog.To = to2
			emaillog.Created = &created
			err = s.Daos.SaveEmailLog(ctx, emaillog)
			if err != nil {
				return errors.New("login email not save")
			}
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

//RejectUser : ""
func (s *Service) RejectUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectUser(ctx, UniqueID)
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
func (s *Service) GetContentDisseminationUser(ctx *models.Context, cda *models.ContentDataAccess) ([]models.DissiminateUser, error) {
	userfilter := new(models.UserFilter)
	if cda != nil {
		if len(cda.Organisation) > 0 {
			userfilter.OrganisationID = cda.Organisation
		}
		if len(cda.Project) > 0 {
			userfilter.Project = cda.Project
		}
		if len(cda.State) > 0 {
			userfilter.States = cda.State
		}
		if len(cda.District) > 0 {
			userfilter.Districts = cda.District
		}
		if len(cda.Block) > 0 {
			userfilter.Blocks = cda.Block
		}
		if len(cda.GramPanchayat) > 0 {
			userfilter.Grampanchayats = cda.GramPanchayat
		}
		if len(cda.Village) > 0 {
			userfilter.Villages = cda.Village
		}
		userfilter.Status = []string{constants.USERSTATUSACTIVE}
		userfilter.SortBy = "name"
		userfilter.SortOrder = 1

	}
	user, err := s.Daos.GetContentDisseminationUser(ctx, userfilter)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (s *Service) UserDataAccess(ctx *models.Context, userfilter *models.UserFilter) (err error) {
	if userfilter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &userfilter.DataAccess)
		if err != nil {
			return err
		}
		fmt.Println("user district====>", dataaccess.AccessDistricts)
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					userfilter.OrganisationID = append(userfilter.OrganisationID, v.ID)
				}
			}
			if len(dataaccess.Projects) > 0 {
				for _, v := range dataaccess.Organisation {
					userfilter.Project = append(userfilter.Project, v.ID)
				}
			}

			if len(dataaccess.AccessStates) > 0 {
				for _, v := range dataaccess.AccessStates {
					userfilter.States = append(userfilter.States, v.ID)
				}
			}
			if len(dataaccess.AccessDistricts) > 0 {
				for _, v := range dataaccess.AccessDistricts {
					userfilter.Districts = append(userfilter.Districts, v.ID)
				}
			}
			if len(dataaccess.AccessBlocks) > 0 {
				for _, v := range dataaccess.AccessBlocks {
					userfilter.Blocks = append(userfilter.Blocks, v.ID)
				}
			}
			if len(dataaccess.AccessVillages) > 0 {
				for _, v := range dataaccess.AccessVillages {
					userfilter.Villages = append(userfilter.Villages, v.ID)

				}
			}
			if len(dataaccess.AccessGrampanchayats) > 0 {
				for _, v := range dataaccess.AccessGrampanchayats {
					userfilter.Grampanchayats = append(userfilter.Grampanchayats, v.ID)

				}
			}
		}

	}
	return err
}
func (s *Service) UserReportExcelV2(ctx *models.Context, filter *models.UserFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterUser(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "UserReport"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "UserName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "EmailId")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "MobileNumber")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Type")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.FirstName)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UserName)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Email)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Mobile)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Ref.Type.UniqueID)
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
func (s *Service) GenerateotpUserRegistration(ctx *models.Context, user *models.User) error {
	data, err := s.Daos.GetSingleUserWithMobileNo(ctx, user.Mobile)

	if data != nil {
		return errors.New("user Already Registered")
	}
	if err != nil {
		if err.Error() != "user not found" {
			return err

		}
	}

	//if data != nil {
	//otp, err := s.GenerateOTP(constants.USERREGISTERATION, user.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return errors.New("Otp Generate Error - " + err.Error())
	// }
	key := fmt.Sprintf("%v_%v", constants.USERREGISTERATION, user.Mobile)
	var otp models.Otp
	otp.Otp = "9999"
	err = s.SetValueCacheMemory(key, otp, 1000)
	if err != nil {
		return err
	}

	//text := fmt.Sprintf("Hi %v, /n Otp For Logikoof Reporting App Login is %v .", data.Name, otp)
	msg := fmt.Sprintf(constants.COMMONTEMPLATE, user.Name, "NICESSM", "OTP for nicessm registration app", "OTP for NICESSM registration is-"+otp.Otp+"", "https://nicessm.org/")

	err = s.SendSMSV2(ctx, user.Mobile, msg)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}
	if err == errors.New(constants.INSUFFICIENTBALANCE) {
		return err
	}
	if err == nil {
		smslog := new(models.SmsLog)
		to := models.To{}
		to.No = user.Mobile
		to.Name = user.Name
		to.UserName = user.UserName
		to.UserType = "user"
		to.UserName = user.UserName
		smslog.IsJob = false
		t := time.Now()
		smslog.SentDate = &t
		smslog.Message = msg
		smslog.SentFor = "Otp"
		smslog.Status = "Active"
		smslog.To = to
		err = s.Daos.SaveSmsLog(ctx, smslog)
		if err != nil {
			return errors.New("otp sms not save")
		}
	}
	var sendmailto []string
	sendmailto = append(sendmailto, user.Email)
	err = s.SendEmail("NICESSM-OTP for nicessm registration app", sendmailto, msg)
	if err != nil {
		return errors.New("email Sending Error - " + err.Error())
	}
	if err == nil {
		emaillog := new(models.EmailLog)
		to2 := models.ToEmailLog{}
		to2.Email = user.Email
		to2.Name = user.UserName
		to2.UserName = user.UserName
		to2.UserType = "user"
		t := time.Now()
		emaillog.SentDate = &t
		emaillog.IsJob = false
		emaillog.Message = msg
		emaillog.SentFor = "login"
		emaillog.Status = "Active"
		emaillog.To = to2
		err = s.Daos.SaveEmailLog(ctx, emaillog)
		if err != nil {
			return errors.New("login email not save")
		}
	}
	return nil
}
func (s *Service) RegistrationValidateOTPUser(ctx *models.Context, login *models.UserOTPLogin) error {

	data, err := s.Daos.GetSingleUserWithMobileNo(ctx, login.Mobile)

	if data != nil {
		return errors.New("user Already Registered")
	}
	if err != nil {
		if err.Error() != "user not found" {
			return err

		}
	}

	if login.Mobile != "" {

		key := fmt.Sprintf("%v_%v", constants.USERREGISTERATION, login.Mobile)
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
	//	loginurl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.LOGINURLV2)

	msg := fmt.Sprintf(constants.COMMONTEMPLATE, login.Name, "NICESSM", "Successfully Registered", "userName :"+login.UserName+",password: #nature32 please login:https://nicessm.org/", "https://nicessm.org/")
	err = s.SendSMSV2(ctx, login.Mobile, msg)
	if err != nil {
		log.Println(login.Mobile + " " + err.Error())
	}

	if err == errors.New(constants.INSUFFICIENTBALANCE) {
		return err
	}
	if err == nil {
		smslog := new(models.SmsLog)
		to := models.To{}
		to.No = login.Mobile
		to.Name = login.UserName
		to.UserName = login.UserName
		to.UserType = "user"
		t := time.Now()
		smslog.SentDate = &t
		smslog.IsJob = false
		smslog.Message = msg
		smslog.SentFor = "login"
		smslog.Status = "Active"
		smslog.Status = constants.SMSLOGSTATUSACTIVE
		smslog.To = to
		err = s.Daos.SaveSmsLog(ctx, smslog)
		if err != nil {
			return errors.New("login sms not save")
		}
	}
	var sendmailto []string
	sendmailto = append(sendmailto, login.Email)
	err = s.SendEmail("Successfully Registered", sendmailto, msg)
	if err != nil {
		return errors.New("email Sending Error - " + err.Error())
	}
	if err == nil {
		emaillog := new(models.EmailLog)
		to2 := models.ToEmailLog{}
		to2.Email = login.Email
		to2.Name = login.UserName
		to2.UserName = login.UserName
		to2.UserType = "user"
		t := time.Now()
		emaillog.SentDate = &t
		emaillog.IsJob = false
		emaillog.Message = msg
		emaillog.SentFor = "login"
		emaillog.Status = "Active"
		emaillog.To = to2
		err = s.Daos.SaveEmailLog(ctx, emaillog)
		if err != nil {
			return errors.New("login email not save")
		}
	}
	err = s.SaveUserwithtransaction(ctx, &login.User)
	if err != nil {
		return err
	}

	return nil

}
func (s *Service) GetWeatherDisseminationUser(ctx *models.Context, state string) ([]models.DissiminateUser, error) {
	userfilter := new(models.UserFilter)
	stateid, err := primitive.ObjectIDFromHex(state)
	if err != nil {
		return nil, err
	}
	var Arraystate []primitive.ObjectID
	Arraystate = append(Arraystate, stateid)
	fmt.Println("Arraystate===>", Arraystate)
	userfilter.States = Arraystate
	userfilter.Status = []string{constants.FARMERSTATUSACTIVE}

	user, err := s.Daos.GetContentDisseminationUser(ctx, userfilter)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (s *Service) GetDistrictWeatherDisseminationUser(ctx *models.Context, district string) ([]models.DissiminateUser, error) {
	userfilter := new(models.UserFilter)
	districtid, err := primitive.ObjectIDFromHex(district)
	if err != nil {
		return nil, err
	}
	var Arraystate []primitive.ObjectID
	Arraystate = append(Arraystate, districtid)
	fmt.Println("Arraystate===>", Arraystate)
	userfilter.Districts = Arraystate
	userfilter.Status = []string{constants.FARMERSTATUSACTIVE}

	user, err := s.Daos.GetContentDisseminationUser(ctx, userfilter)
	if err != nil {
		return nil, err
	}
	return user, nil

}
