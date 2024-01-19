package services

import (
	"bytes"
	"haritv2-service/constants"
	"haritv2-service/models"
	"text/template"

	"context"
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
	resp, err := s.Shared.Get(URL.String(), nil)
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

//SendEmail : ""
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
//SendEmail : ""
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

func (s *Service) SendNotification(title, body, image string, regToken []string) error {
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
		messsages = append(messsages,
			&messaging.Message{
				Notification: &messaging.Notification{
					Title:    title,
					Body:     body,
					ImageURL: apiBaseURL + image,
				},

				Token: v,
			})
	}
	if messsages == nil {
		log.Println("msg is nil ")
		return nil
	}
	if len(messsages) == 0 {
		log.Println("msg is 0 ")
		return nil
	}
	response, err := client.SendAll(ctx, messsages)
	if err != nil {
		fmt.Println(err)
		// log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	s.Shared.BsonToJSONPrintTag("Successfully sent message:", response)
	return nil
}
