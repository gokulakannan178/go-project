package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
)

//SendOTPConsumerLogin : ""
func (s *Service) SendOTPConsumerLogin(ctx *models.Context, mobileNo string) error {
	filter := new(models.PropertyOwnerFilter)
	filter.Mobile = []string{mobileNo}
	owners, err := s.Daos.FilterPropertyOwner(ctx, filter, nil)
	if err != nil {
		return errors.New("Error in geting Owner - " + err.Error())
	}
	if len(owners) == 0 {
		//return errors.New("No Owners available with this mobile no")
	}
	token, err := s.GenerateOTP(constants.CONSUMERLOGIN, mobileNo, constants.TOKENOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return err
	}
	fmt.Println("token =>", token)
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return errors.New("Error in getting product config" + err.Error())
	}
	if productConfig.LocationID == "Munger" {
		return errors.New("Please Contact admin")
	}
	// msg := fmt.Sprintf(constants.PROPERTYTAXDEMANDCONTENT, math.Ceil(resPropertyDemand.Demand.TotalTax), resPropertyDemand.UniqueID, resFYs.Name)

	// premsg := fmt.Sprintf("Your OTP for %v Consumer login is %v. This OTP is valid only for 3 minutes. Please do not share OTP to anyone.", productConfig.Name, token)
	premsg := "Your OTP for Municipal Corporation Consumer login is %v. This OTP is valid only for 3 minutes. Please do not share OTP with anyone. Thanks and Regards BRMNCP"
	msg := fmt.Sprintf(premsg, token)
	fmt.Println("OTP Msg = ", msg)
	if err := s.SendSMS(mobileNo, msg); err != nil {
		log.Println(err)
	}
	return nil
}

//ConsumerLoginValidateOTP : ""
func (s *Service) ConsumerLoginValidateOTP(ctx *models.Context, mobileNo string, otp string) ([]string, error) {
	filter := new(models.PropertyOwnerFilter)
	filter.Mobile = []string{mobileNo}
	owners, err := s.Daos.FilterPropertyOwner(ctx, filter, nil)
	if err != nil {
		return nil, errors.New("Error in geting Owner - " + err.Error())
	}
	if len(owners) == 0 {
		//return nil, errors.New("No Owners available with this mobile no")
	}
	properties := []string{}
	for _, v := range owners {
		properties = append(properties, v.PropertyID)
	}
	return properties, s.ValidateOTP(constants.CONSUMERLOGIN, mobileNo, otp)
}
