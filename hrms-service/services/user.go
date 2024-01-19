package services

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUser :""
func (s *Service) SaveUser(ctx *models.Context, user *models.User) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	user.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
	user.Status = constants.USERSTATUSACTIVE
	user.Password = "#nature32" //Default Password
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	user.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUser(ctx, user)
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
	/*
		d := make(map[string]interface{})

		d["UserName"] = user.UserName
		d["Password"] = user.Password
		d["Role"] = func(role string) string {
			switch role {
			case "LM 1":
				return "Line Manager 1"
			case "LM 2":
				return "Line Manager 2"
			case "LM 3":
				return "Line Manager 3"
			case "LM 4":
				return "Line Manager 4"
			case "LM 5":
				return "Line Manager 5"
			case "PD":
				return "Project Director"
			}
			return ""
		}(user.Type)

		d["URL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.APIBASEURL) + user.Profile
		d["LoginURL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.BASEURL) + s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOGINURL)
		d["ContactUsURL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.BASEURL) + s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.CONTACTUSURL)
		templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		//html template path
		templateID := templatePathStart + "successfull_registration.html"
		err := s.SendEmailWithTemplate("HRMS - Registered Successfully", []string{"solomon2261993@gmail.com"}, templateID, d)
		if err != nil {
			return errors.New("Unable to send email - " + err.Error())
		}*/
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
func (s *Service) GetSingleUser(ctx *models.Context, UserName string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUser(ctx, UserName)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *Service) GetSingleUserWithLogin(ctx *models.Context, UserName string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUserWithLogin(ctx, UserName)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//FilterUser :""
func (s *Service) FilterUser(ctx *models.Context, userfilter *models.UserFilter, pagination *models.Pagination) (user []models.RefUser, err error) {
	return s.Daos.FilterUser(ctx, userfilter, pagination)
}

//ResetUserPassword : ""
func (s *Service) ResetUserPassword(ctx *models.Context, userName string) error {
	return s.Daos.ResetUserPassword(ctx, userName, "#nature32")
}

//ChangePassword : ""
func (s *Service) ChangePassword(ctx *models.Context, cp *models.UserChangePassword) (bool, string, error) {
	user, err := s.Daos.GetSingleUser(ctx, cp.UserName)
	if err != nil {
		return false, "", err
	}
	fmt.Println("username===>", user.UserName)
	fmt.Println("Oldpassword===>", user.Password)
	fmt.Println("Newpassword===>", cp.OldPassword)
	if user.Password != cp.OldPassword {
		return false, "Enter Valid Password", nil
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
	err = s.SendEmailWithTemplate("HRMS - Password Changed Successfully", []string{"solomon2261993@gmail.com"}, "templates/"+templateID, d)
	if err != nil {
		log.Println("Email not sent - " + err.Error())
		// return errors.New("Unable to send email - " + err.Error())
	}
	if user.Password == cp.NewPassword {
		return false, "Old Password and New Password Same", nil
	}
	return true, "Success", nil
}

//ChangePassword : ""
func (s *Service) ForgetPasswordNewPassword(ctx *models.Context, np *models.UserNewPassword) (bool, string, error) {
	user, err := s.Daos.GetSingleUserWithLogin(ctx, np.UserName)
	if err != nil {
		return false, "error", err
	}
	fmt.Println("username ==>", user)
	if user.UserName == "" {
		return false, "error", err
	}
	err = s.Daos.ForgetPasswordNewPassword(ctx, np.UserName, np.Password)
	if err != nil {
		return false, "error", err
	}
	err = s.ValidateOTP(constants.OTPSCENARIOPASSWORD, user.Mobile, np.Token)
	if err != nil {
		return true, "", err
	}
	token, _ := s.GenerateOTP(constants.OTPSCENARIOTOKEN, user.Mobile, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return false, "", err
	// }
	sEnc := b64.StdEncoding.EncodeToString([]byte(token))

	fmt.Println(sEnc)

	return true, "success", nil
}

//ForgetPasswordValidateOTP : ""
func (s *Service) ForgetPasswordValidateOTP(ctx *models.Context, UniqueID string, otp string) (string, error) {
	user, err := s.Daos.GetSingleUserWithLogin(ctx, UniqueID)
	if err != nil {
		return "", err
	}
	err = s.ValidateOTP(constants.OTPSCENARIOPASSWORD, user.Mobile, otp)
	if err != nil {
		return "", err
	}
	token, _ := s.GenerateOTP(constants.OTPSCENARIOTOKEN, user.Mobile, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return "", err
	// }
	sEnc := b64.StdEncoding.EncodeToString([]byte(token))

	fmt.Println(sEnc)
	return sEnc, nil
}

//ForgetPasswordValidateOTP : ""
func (s *Service) GetSingleUserWithUserName(ctx *models.Context, userName string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUserWithUserName(ctx, userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//ForgetPasswordGenerateOTP : ""
func (s *Service) ForgetPasswordGenerateOTP(ctx *models.Context, UniqueID string) error {
	user, err := s.Daos.GetSingleUserWithLogin(ctx, UniqueID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user is nil")
	}
	otp, _ := s.GenerateOTP(constants.OTPSCENARIOPASSWORD, user.Mobile, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	// if err != nil {
	// 	return err
	// }
	// // msg := "Use otp " + otp + " for municipal forget password. municipal doesnt ask otp to be shared with anyone"
	// err = s.SendSMS(user.Mobile, msg)
	msg := fmt.Sprintf(constants.COMMONTEMPLATE, user.Name, "NICESSM", "otp for forgot password", "OTP for NICESSM forgot password is-"+otp+"", "https://nicessm.org/")

	err = s.SendSMS(user.Mobile, msg)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}
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
	var sendmailto []string
	sendmailto = append(sendmailto, user.Email)
	err = s.SendEmail("NICESSM-OTP Generation For Forget Password", sendmailto, msg)
	if err != nil {
		return errors.New("email Sending Error - " + err.Error())
	}
	emaillog := new(models.EmailLog)
	to2 := models.ToEmailLog{}
	to2.Email = user.Email
	to2.Name = user.UserName
	t = time.Now()
	emaillog.SentDate = &t
	emaillog.IsJob = false
	emaillog.Message = msg
	emaillog.SentFor = "login"
	emaillog.Status = constants.EMAILLOGSTATUSACTIVE
	emaillog.To = to2
	err = s.Daos.SaveEmailLog(ctx, emaillog)
	if err != nil {
		return errors.New("login email not save")
	}
	fmt.Println(err)
	return nil
}
func (s *Service) UserUniquenessCheckRegistration(ctx *models.Context, OrgID string, Param string, Value string) (*models.UserUniquinessChk, error) {
	user, err := s.Daos.UserUniquenessCheckRegistration(ctx, OrgID, Param, Value)
	if err != nil {
		return nil, err
	}
	return user, nil
}
