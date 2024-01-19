package services

import (
	"bytes"
	"context"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"text/template"
	"time"

	"errors"
	"fmt"
	"log"
	"net/smtp"
	"net/url"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

//SendSMS : ""
func (s *Service) SendSMS(mobileNo string, msg string) error {
	smsConfig := s.GetSMSConfig()
	log.Println(smsConfig)
	var URL *url.URL
	URL, err := url.Parse(smsConfig.URL)
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
	URL.RawQuery = parameters.Encode()
	fmt.Println("URL : ", URL.String())
	//s.Daos.GetactiveProductConfig(ctx, true)
	resp, err := s.Shared.Get(URL.String(), nil)
	if err != nil {
		return errors.New("api err - " + err.Error())
	}
	log.Println(resp)
	return nil
}
func (s *Service) SendSMSV2(ctx *models.Context, mobileNo string, msg string) error {

	communicationcredit, err := s.Daos.GetSingleCommunicationCreditWithUniqueId(ctx, constants.CRIDITTYPESMS)
	if err != nil {
		return err
	}

	data := []byte(msg)

	if communicationcredit != nil {
		credit := len(data) / communicationcredit.ChartCountCredit
		if len(data)%communicationcredit.ChartCountCredit != 0 {
			credit++
		}
		fmt.Println("value", credit)
		if communicationcredit.BalanceCredit < float64(credit) {
			return errors.New(constants.INSUFFICIENTBALANCE)
		}

		err = s.Daos.UpdateCommunicationCreditWithBalance(ctx, constants.CRIDITTYPESMS, -float64(credit))
		if err != nil {
			return err
		}
	}
	smsConfig := s.GetSMSConfig()
	log.Println(smsConfig)
	var URL *url.URL
	URL, err = url.Parse(smsConfig.URL)
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
	URL.RawQuery = parameters.Encode()
	fmt.Println("URL : ", URL.String())
	prod, _ := s.Daos.GetactiveProductConfig(ctx, true)
	if prod.IsSms == true {
		resp, err := s.Shared.Get(URL.String(), nil)
		if err != nil {
			return errors.New("api err - " + err.Error())
		}
		log.Println(resp)

	}

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

//SendEmailWithTemplate : ""
func (s *Service) SendEmailWithTemplate(subject string, to []string, templateURL string, data interface{}) error {
	fromEmail := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAIL)
	fromPassword := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAILPASSWORD)
	smtpHost := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPHOST)
	smtpPort := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPPORT)
	// Authentication.
	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

	t, err := template.ParseFiles(templateURL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+subject+" \n%s\n\n", mimeHeaders)))
	t.Execute(&body, data)
	fmt.Println(fromEmail, fromPassword, smtpHost, smtpPort, to)
	// return nil
	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Email Sent!")
	return nil
}

// //SendEmail : ""
// func (s *Service) SendEmail(subject string, to []string, message string) error {
// 	fromEmail := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAIL)
// 	fromPassword := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAILPASSWORD)
// 	smtpHost := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPHOST)
// 	smtpPort := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPPORT)
// 	// Authentication.
// 	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

// 	var body bytes.Buffer
// 	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	body.Write([]byte(fmt.Sprintf("Subject: "+subject+" \n%s\n\n", mimeHeaders)))
// 	fmt.Println(fromEmail, fromPassword, smtpHost, smtpPort, to)
// 	// return nil
// 	// Sending email.
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, to, []byte(message))
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	fmt.Println("Email Sent!")
// 	return nil
// }
func (s *Service) InitiateWhatsAppText(mobileNo string, text []models.WhatsAppText) error {
	swt := new(models.SendWhatsAppText)
	swt.To = mobileNo
	swt.Type = "template"
	swt.Template.NameSpace = "08769e5f_1619_46d7_9c4e_45ba8a266010"
	swt.Template.Name = "farmer_greeting2"
	swt.Template.Language.Policy = "deterministic"
	swt.Template.Language.Code = "en"
	swt.Template.Components = append(swt.Template.Components, models.WhatsAppComponent{"body", text})
	err := s.SendWhatsAppText(swt)
	if err != nil {
		log.Println("error in sending message")
	}
	return err
}

//SendWhatsAppSMS : ""
func (s *Service) SendWhatsAppText(swt *models.SendWhatsAppText) error {
	smsConfig := s.GetSMSConfig()
	log.Println(smsConfig)
	whatsappURL := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.WHATSAPP_URL)
	businessId := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.WHATSAPP_BUSINESSID)
	var URL *url.URL
	URL, err := url.Parse(whatsappURL)
	if err != nil {
		return errors.New("url const err - " + err.Error())
	}
	parameters := url.Values{}
	parameters.Add("businessId", businessId)
	URL.RawQuery = parameters.Encode()
	fmt.Println("URL : ", URL.String())
	resp, err := s.Shared.Post(URL.String(), nil, swt)
	if err != nil {
		return errors.New("api err - " + err.Error())
	}
	log.Println(resp)
	return nil
}

func (s *Service) WhatsAppSingle(ctx *models.Context, swt *models.SendWhatsAppText2) error {
	s.InitiateWhatsAppText(swt.MobileNo, swt.Type)
	return nil
}
func (s *Service) LoginSms(ctx *models.Context, mobile string, userName string) error {

	loginurl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.LOGINURLV2)

	msg := fmt.Sprintf(constants.COMMONTEMPLATE, userName, "NICCESM", "userName :("+userName+"),password: #nature32", "please login :"+loginurl+")", "https://nicessm.org/")
	err := s.SendSMSV2(ctx, mobile, msg)
	if err != nil {
		log.Println(mobile + " " + err.Error())
	}
	if err == errors.New(constants.INSUFFICIENTBALANCE) {
		return err
	}
	if err != nil {
		smslog := new(models.SmsLog)
		t := time.Now()
		to := models.To{}
		to.No = mobile
		to.Name = userName
		to.UserType = "Farmer"
		smslog.Status = constants.SMSLOGSTATUSACTIVE
		smslog.IsJob = true
		smslog.SentDate = &t
		smslog.Message = msg
		smslog.SentFor = "login"
		smslog.To = to
		err = s.Daos.SaveSmsLog(ctx, smslog)
		if err != nil {
			return errors.New("login sms not save")
		}
	}
	return err
}
func (s *Service) FarmerRegisterSms(ctx *models.Context, mobile string, userName string) error {

	//loginurl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.LOGINURLV2)

	msg := fmt.Sprintf(constants.COMMONTEMPLATE, userName, "NICCESM", "Registration", " Successfully registered Pls download Nicessm app from play store", "https://nicessm.org/")
	err := s.SendSMS(mobile, msg)
	if err != nil {
		log.Println(mobile + " " + err.Error())
	}
	if err == nil {
		smslog := new(models.SmsLog)
		to := models.To{}
		to.No = mobile
		to.Name = userName
		to.UserType = "Farmer"
		smslog.IsJob = false
		smslog.Status = constants.SMSLOGSTATUSACTIVE
		smslog.Message = msg
		smslog.SentFor = "login"
		smslog.To = to
		err = s.Daos.SaveSmsLog(ctx, smslog)
		if err != nil {
			return errors.New("login sms not save")
		}
	}
	return err
}
func (s *Service) SendEmail(subject string, to []string, message string) error {
	fromEmail := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAIL)
	fromPassword := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAILPASSWORD)
	smtpHost := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPHOST)
	smtpPort := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPPORT)
	// Authentication.
	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)
	headers := make(map[string]string)

	headers["Subject"] = subject
	headers["From"] = fromEmail
	headers["To"] = to[0]
	body := message

	var msg bytes.Buffer
	for k, v := range headers {
		msg.WriteString(k + ": " + v + "\r\n")
	}
	msg.WriteString("\r\n")
	msg.WriteString(body)
	// var body bytes.Buffer
	// mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// body.Write([]byte(fmt.Sprintf("Subject: "+subject+" \n%s\n\n", mimeHeaders)))
	fmt.Println(fromEmail, fromPassword, smtpHost, smtpPort, to)
	// return nil
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, to, msg.Bytes())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Email Sent!")
	return nil
}
func (s *Service) SendNotification(topic, title, body, image string, regToken []string, data map[string]string) error {
	fmt.Println("topic===>", topic)
	fmt.Println("title===>", title)
	fmt.Println("body===>", body)
	fmt.Println("image===>", image)
	fmt.Println("regToken===>", regToken)
	fmt.Println("data===>", data)
	getCRED := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.PUSHNOTIFICATIONCRED)
	apiBaseURL := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.APIBASEURL)
	opt := option.WithCredentialsFile(getCRED)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("error initializing app: %v", err)
		return err
	}

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	messsages := []*messaging.Message{}
	log.Println("total tokens - ", len(regToken))
	for _, v := range regToken {
		if v == "" {
			log.Println("Omited - " + body)
			continue
		}
		msg := messaging.Message{
			Notification: &messaging.Notification{
				Title:    title,
				Body:     body,
				ImageURL: apiBaseURL + image,
			},

			Token: v,
		}
		if data != nil {
			msg.Data = data
		}

		if topic != "" {
			msg.Topic = topic

		}
		messsages = append(messsages,
			&msg)
	}
	if messsages == nil {
		log.Println("msg is nil ")
		return nil
	}
	if len(messsages) == 0 {
		log.Println("msg is 0 ")
		return nil
	}
	fmt.Println("messsages length==>", len(messsages))
	fmt.Println("messsages ==>", messsages)
	response, err := client.SendAll(ctx, messsages)
	if err != nil {
		return (err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	s.Shared.BsonToJSONPrintTag("Successfully sent message:", response)
	return nil
}
