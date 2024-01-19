package services

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

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
	res, err := s.Daos.GetSingleFarmerWithMobileNo(ctx, Farmer.MobileNo)
	if err != nil {
		return err
	}
	if res != nil {
		return errors.New("mobileNo already exists")
	}
	Farmer.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFARMERS)
	Farmer.Status = constants.FARMERSSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Farmer.created")
	Farmer.Created = &created
	log.Println("b4 Farmer.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFarmer(ctx, Farmer)
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
func (s *Service) FilterFarmer(ctx *models.Context, filter *models.FarmerFilter, pagination *models.Pagination) ([]models.RefFarmer, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterFarmer(ctx, filter, pagination)
}

//OTPLoginGenerateOTP :
func (s *Service) FarmerLoginGenerateOTP(ctx *models.Context, login *models.Login) error {
	data, err := s.Daos.GetSingleUserWithMobileNoForFarmer(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data == nil {
		return errors.New("User Not Available")
	}
	if data.Status == constants.USERSTATUSINIT {
		return errors.New("Awaiting Activation")
	}
	if data.Status != constants.USERSTATUSACTIVE {
		return errors.New("Please Contact Administrator")
	}
	// if data.Type != constants.USERTYPE {
	// 	return errors.New("User Type is invalid")
	// }

	otp, err := s.GenerateOTP(constants.USERLOGIN, login.UserName, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return errors.New("Otp Generate Error - " + err.Error())
	}
	text := fmt.Sprintf("Dear %v, /n Otp For Harit Farmer login is %v .", data.Name, otp)
	err = s.SendSMS(login.UserName, text)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}

	return nil
}

//FarmerLoginValidateOTP :
func (s *Service) FarmerLoginValidateOTP(ctx *models.Context, login *models.OTPLogin) (*models.RefFarmer, bool, error) {

	data, err := s.Daos.GetSingleUserWithMobileNoForFarmer(ctx, login.Mobile)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	if data == nil {
		return nil, false, errors.New("User Not Available")
	}
	if data.Status == constants.USERSTATUSINIT {
		return nil, false, errors.New("Awaiting Activation")
	}
	if data.Status != constants.USERSTATUSACTIVE {
		return nil, false, errors.New("Please Contact Administrator")
	}

	err = s.ValidateOTP(constants.USERLOGIN, login.Mobile, login.OTP)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}

	var auth models.Authentication
	auth.UserID = data.ID
	auth.UserName = data.UniqueID
	auth.Status = data.Status

	token, err := CreateTokenV2(&auth)
	if err != nil {
		log.Println(err.Error())
		return nil, false, errors.New("Error in Generating Token - " + err.Error())
	}
	data.Token = token

	return data, true, nil
}

//FarmerRegistrationLoginGenerateOTP :
func (s *Service) FarmerRegistrationLoginGenerateOTP(ctx *models.Context, login *models.Login) error {
	data, err := s.Daos.GetSingleUserWithMobileNoForFarmer(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data != nil {
		return errors.New("user already available")
	}
	if err != nil {
		return err
	}

	otp, err := s.GenerateOTP(constants.USERLOGIN, login.UserName, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return errors.New("Otp Generate Error - " + err.Error())
	}
	text := fmt.Sprintf("Dear user, Otp For Harit Farmer registration is %v .", otp)
	err = s.SendSMS(login.UserName, text)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}

	return nil
}

//OTPLoginValidateOTP :
func (s *Service) FarmerRegistrationLoginValidateOTP(ctx *models.Context, login *models.OTPLogin) (*models.RefUser, bool, error) {

	data, err := s.Daos.GetSingleUserWithMobileNoForFarmer(ctx, login.Mobile)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	if data != nil {
		return nil, false, errors.New("user already available")
	}

	err = s.ValidateOTP(constants.USERLOGIN, login.Mobile, login.OTP)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	return nil, true, nil

}
