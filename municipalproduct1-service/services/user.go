package services

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
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
			return dberr
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		log.Println("err in abort out")
		return errors.New("Transaction Aborted - " + err.Error())
	}

	// d := make(map[string]interface{})
	// d["UserName"] = user.UserName
	// d["Password"] = user.Password
	// d["Role"] = func(role string) string {
	// 	switch role {
	// 	case "PM":
	// 		return "Project Manager"
	// 	case "TC":
	// 		return "Tax Collector"
	// 	case "Municipal":
	// 		return "Kochas Commissioner"
	// 	}
	// 	return ""
	// }(user.Type)
	// d["URL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.APIBASEURL) + user.Profile
	// d["LoginURL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.BASEURL) + s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOGINURL)
	// d["ContactUsURL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.BASEURL) + s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.CONTACTUSURL)
	// templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	// //html template path
	// templateID := templatePathStart + "successfull_registration.html"
	// err := s.SendEmailWithTemplate("Kochas Municipality - Registered Successfully", []string{"solomon2261993@gmail.com"}, templateID, d)
	// if err != nil {
	// 	return errors.New("Unable to send email - " + err.Error())
	// }
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
	user, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return "", err
	}
	err = s.ValidateOTP(constants.OTPSCENARIOPASSWORD, user.Mobile, otp)
	if err != nil {
		return "", err
	}
	token, err := s.GenerateOTP(constants.OTPSCENARIOTOKEN, user.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return "", err
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(token))

	fmt.Println(sEnc)
	return sEnc, nil
}

//ForgetPasswordGenerateOTP : ""
func (s *Service) ForgetPasswordGenerateOTP(ctx *models.Context, UniqueID string) error {
	user, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user is nil")
	}
	otp, err := s.GenerateOTP(constants.OTPSCENARIOPASSWORD, user.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return err
	}
	msg := "Use otp " + otp + " for municipal forget password. municipal doesnt ask otp to be shared with anyone"
	err = s.SendSMS(user.Mobile, msg)
	fmt.Println(err)
	return nil
}

//Login :
func (s *Service) LoginGenerateOTP(ctx *models.Context, mobileNo string) error {
	user, err := s.Daos.GetSingleUserWithUniqueID(ctx, mobileNo)
	if err != nil {
		return errors.New("Error in geting user - " + err.Error())
	}
	if user == nil {
		return errors.New("User not available")
	}
	token, err := s.GenerateOTP(constants.USERLOGIN, mobileNo, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return err
	}
	fmt.Println("token =>", token)
	// msg := fmt.Sprintf("Your OTP for Kochas Municipal Corporation Consumer login is %v. This OTP is valid only for 3 minutes. Please do not share OTP to anyone",

	// 	token)
	// msg := fmt.Sprintf("Your OTP is %v KVK App", token)

	msg := fmt.Sprintf("your otp is %v", token)

	if err := s.SendSMS(mobileNo, msg); err != nil {
		log.Println(err)
	}
	return nil
}

//LoginValidateOTP : ""
func (s *Service) LoginValidateOTP(ctx *models.Context, mobileNo string, otp string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleActiveUserWithUniqueID(ctx, mobileNo)
	if err != nil {
		return nil, errors.New("Error in geting user - " + err.Error())
	}
	if user == nil {
		return nil, errors.New("User not available")
	}
	if user.Status == constants.USERSTATUSDISABLED {
		return nil, errors.New("Please contact administrator")
	}
	err = s.ValidateOTP(constants.USERLOGIN, user.Mobile, otp)
	if err != nil {
		return nil, err
	}
	return user, nil
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

func (s *Service) IDCaredPDF(ctx *models.Context, userfilter *models.UserFilter) ([]byte, error) {
	data := make(map[string]interface{})
	r := NewRequestPdfV2("", "Landscape")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "idcardprint2.html"
	config, err := s.GetSingleProductConfiguration(ctx)
	if err != nil {
		return nil, err
	}
	users, err := s.FilterUser(ctx, userfilter, nil)
	if err != nil {
		return nil, err
	}
	data["config"] = config
	data["users"] = users
	if err := r.ParseTemplate(templatePath, data); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())
		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil
}

// UpdateAccessPrivilege : ""
func (s *Service) UpdateAccessPrivilege(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAccessPrivilege(ctx, user)
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

func (s *Service) UpdateAppVersionUser(ctx *models.Context, user *models.AppVersionUser) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateAppVersionUser(ctx, user)
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

func (s *Service) UserMpinRegistration(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UserMpinRegistration(ctx, user)
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

// VerifyUserMpinRegistration : ""
func (s *Service) VerifyUserMpinRegistration(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.VerifyUserMpinRegistration(ctx, user)
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

// GetSingleUserWithDeviceID : ""
func (s *Service) GetSingleUserWithDeviceID(ctx *models.Context, DeviceID string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUserWithDeviceID(ctx, DeviceID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UserMpinLogin(ctx *models.Context, login *models.MpinValidation) (string, bool, error) {
	data, err := s.Daos.GetSingleUserWithDeviceID(ctx, login.DeviceID)
	if err != nil {
		fmt.Println(err)
		return "dal err", false, err
	}
	// log.Println("Data ===================>", data)
	// if data.MobileAuth.MpinStatus == constants.MPINSTATUSINIT {
	// 	return "", false, errors.New("Awaiting Activation")
	// }
	if data.Status != constants.USERSTATUSACTIVE {
		return "", false, errors.New("Please Contact Administrator")
	}
	if data.MobileAuth.MpinStatus != constants.MPINSTATUSACTIVE {
		return "", false, errors.New("Please Contact Administrator")
	}
	if ok := data.MobileAuth.DeviceID == login.DeviceID; !ok {
		fmt.Println("Hi I am Here")
		fmt.Println("ok ====>", ok)
		return "Passs false", false, nil
	}

	// var auth models.Authentication
	// auth.UserID = data.ID
	// auth.UserName = data.UserName

	// auth.Status = data.Status
	// auth.Role = data.Role
	// fmt.Println("auth user ==>", auth, data)
	// token, err := CreateToken(&auth)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return "", false, errors.New("Error in Generating Token - " + err.Error())
	// }
	// data.Token = token
	// data.CurrentLocation = login.Location
	// err = s.Daos.UpdateUserWithUniqueID(data.UserName, data)
	// if err != nil {
	// 	log.Println("Error in saving token - " + err.Error())
	// 	return "", false, errors.New(constants.INTERNALSERVERERROR)
	// }
	return data.UserName, true, nil
}

//RemovedUserToken
func (s *Service) RemovedUserToken(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RemovedUserToken(ctx, UniqueID)
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
