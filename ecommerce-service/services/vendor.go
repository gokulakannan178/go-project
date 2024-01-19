package services

import (
	"ecommerce-service/constants"
	"errors"
	"fmt"
	"log"
	"time"

	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveVendor : ""
func (s *Service) SaveVendor(ctx *models.Context, vendorInfo *models.Vendor) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	vendorInfo.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVENDOR)
	vendorInfo.Status = constants.VENDORSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 vendor.created")
	vendorInfo.Created = &created
	log.Println("b4 vendor.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveVendor(ctx, vendorInfo)
		if dberr != nil {
			return dberr
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

//GetSingleVendor :""
func (s *Service) GetSingleVendor(ctx *models.Context, UniqueID string) (*models.RefVendor, error) {
	tower, err := s.Daos.GetSingleVendor(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateVendor : ""
func (s *Service) UpdateVendor(ctx *models.Context, vendorInfo *models.Vendor) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVendor(ctx, vendorInfo)
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

// EnableVendor : ""
func (s *Service) EnableVendor(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableVendor(ctx, UniqueID)
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

//DisableVendor : ""
func (s *Service) DisableVendor(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableVendor(ctx, UniqueID)
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

//DeleteVendor : ""
func (s *Service) DeleteVendor(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVendor(ctx, UniqueID)
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

// FilterVendor : ""
func (s *Service) FilterVendor(ctx *models.Context, filter *models.VendorFilter, pagination *models.Pagination) ([]models.RefVendor, error) {
	return s.Daos.FilterVendor(ctx, filter, pagination)

}

//VendorOTPLoginGenerateOTP :
func (s *Service) VendorOTPLoginGenerateOTP(ctx *models.Context, login *models.Login) error {
	data, err := s.Daos.GetSingleVendorWithMobileNo(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data == nil {
		return errors.New("Vendor Not Available")
	}

	otp, err := s.GenerateOTP(constants.USERLOGIN, login.UserName, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return errors.New("Otp Generate Error - " + err.Error())
	}
	text := fmt.Sprintf("Hi %v, /n Otp For Logikoof Reporting App Login is %v .", data.Name, otp)
	err = s.SendSMS(login.UserName, text)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}

	return nil
}

//VendorOTPLoginValidateOTP : ""
func (s *Service) VendorOTPLoginValidateOTP(ctx *models.Context, login *models.OTPLogin) (*models.RefVendor, bool, error) {

	data, err := s.Daos.GetSingleVendorWithMobileNo(ctx, login.Mobile)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	if data == nil {
		return nil, false, errors.New("Vendor Not Available")
	}

	err = s.ValidateOTP(constants.USERLOGIN, login.Mobile, login.OTP)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}

	var auth models.Authentication
	auth.UserID = data.ID
	auth.UserName = data.MobileNo
	auth.Type = "Vendor"
	auth.Status = data.Status

	token, err := CreateTokenV2(&auth)
	if err != nil {
		log.Println(err.Error())
		return nil, false, errors.New("Error in Generating Token - " + err.Error())
	}
	data.Token = token

	return data, true, nil

}

//GetSingleVendorWithMobileNoV2 :""
func (s *Service) GetSingleVendorWithMobileNoV2(ctx *models.Context, MobileNo string) (string, error) {
	res, err := s.Daos.GetSingleVendorWithMobileNoV2(ctx, MobileNo)
	if err != nil {
		return "", err
	}
	if res.MobileNo != "" {
		return "Duplicate User", nil
	} else {
		return "User Available", nil
	}
}
