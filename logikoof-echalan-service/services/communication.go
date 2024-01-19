package services

import (
	"errors"
	"fmt"
	"log"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"net/url"
)

//SendSMS : ""
func (s *Service) SendSMS(mobileNo string, msg string) error {
	smsConfig := s.GetSMSConfig()
	log.Println(smsConfig)
	var Url *url.URL
	Url, err := url.Parse(smsConfig.URL)
	if err != nil {
		return errors.New("url const err - " + err.Error())
	}
	parameters := url.Values{}
	parameters.Add("user", smsConfig.User)
	parameters.Add("password", smsConfig.Password)
	parameters.Add("senderid", smsConfig.Senderid)
	parameters.Add("channel", smsConfig.Channel)
	parameters.Add("DCS", fmt.Sprintf("%v", smsConfig.DCS))
	parameters.Add("flashsms", fmt.Sprintf("%v", smsConfig.Flashsms))
	parameters.Add("number", mobileNo)
	parameters.Add("text", msg)
	parameters.Add("route", smsConfig.Route)
	Url.RawQuery = parameters.Encode()
	fmt.Println("URL : ", Url.String())
	resp, err := s.Shared.Get(Url.String(), nil)
	if err != nil {
		return errors.New("api err - " + err.Error())
	}
	log.Println(resp)
	return nil
}

//GetSMSConfig : ""
func (s *Service) GetSMSConfig() *models.SMSConfig {

	smsConfig := new(models.SMSConfig)
	smsConfig.URL = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSURL)
	smsConfig.User = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSUSERNAME)
	smsConfig.Password = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSPASSWORD)
	smsConfig.Senderid = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSSENDERID)
	smsConfig.Channel = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSCHANNEL)
	smsConfig.DCS = s.ConfigReader.GetInt(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSDCS)
	smsConfig.Flashsms = s.ConfigReader.GetInt(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSFLASH)
	smsConfig.Route = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSROUTE)
	return smsConfig
}
