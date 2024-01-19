package services

import (
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveNews : ""
func (s *Service) SaveNews(ctx *models.Context, news *models.News) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	news.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONNEWS)
	news.Status = constants.NEWSSTATUSACTIVE
	t := time.Now()
	news.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 News.created")
	news.Created = &created
	log.Println("b4 News.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveNews(ctx, news)
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

// GetSingleNews : ""
func (s *Service) GetSingleNews(ctx *models.Context, UniqueID string) (*models.RefNews, error) {
	news, err := s.Daos.GetSingleNews(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return news, nil
}

//UpdateNews : ""
func (s *Service) UpdateNews(ctx *models.Context, news *models.News) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateNews(ctx, news)
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

// EnableNews : ""
func (s *Service) EnableNews(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableNews(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// DisableNews : ""
func (s *Service) DisableNews(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableNews(ctx, uniqueID)
		if debrr != nil {
			return debrr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}

//DeleteNews : ""
func (s *Service) DeleteNews(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteNews(ctx, UniqueID)
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
func (s *Service) NewsDataAccess(ctx *models.Context, filter *models.FilterNews) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationId = append(filter.OrganisationId, v.UniqueID)
				}
				if dataaccess.SuperAdmin != true {
					if dataaccess.Department != "" {
						filter.DepartmentId = append(filter.DepartmentId, dataaccess.Department)
					}
					if filter.DataAccess.UserName != "" {
						filter.Employee = append(filter.DepartmentId, filter.DataAccess.UserName)
					}

				}
			}

		}

	}
	return err
}

// FilterNews : ""
func (s *Service) FilterNews(ctx *models.Context, news *models.FilterNews, pagination *models.Pagination) (newss []models.RefNews, err error) {
	err = s.NewsDataAccess(ctx, news)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterNews(ctx, news, pagination)
}
func (s *Service) PublishedNews(ctx *models.Context, News *models.News) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	News.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONNEWS)
	News.Status = constants.NEWSSTATUSPUBLISHED
	t := time.Now()
	News.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 News.created")
	News.Created = &created
	log.Println("b4 News.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.SaveNews(ctx, News)
		if dberr != nil {
			return dberr
		}

		filter := new(models.FilterEmployee)
		if News.OrganisationId != "" {
			filter.OrganisationID = append(filter.OrganisationID, News.OrganisationId)

		}
		if len(News.SendTo.DepartmentId) > 0 {
			filter.DepartmentID = append(filter.DepartmentID, News.SendTo.DepartmentId...)
		}
		if len(News.SendTo.Employee) > 0 {
			filter.UniqueID = append(filter.DepartmentID, News.SendTo.Employee...)

		}
		employee, err := s.Daos.FilterEmployee(ctx, filter, nil)
		fmt.Println("no.of.employee===>", len(employee))
		if err != nil {
			return err
		}
		var sendmailto []string
		for _, v := range employee {
			fmt.Println("employee===>", v.Name, v.UniqueID)
			apptoken, _ := s.Daos.GetSingleApptokenWithUserID(ctx, v.UniqueID)
			if apptoken != nil {
				fmt.Println("apptoken===>", apptoken.Apptoken.RegistrationToken)
				var token []string
				token = append(token, apptoken.RegistrationToken)

				fmt.Println("appToken===>", apptoken.RegistrationToken)
				topic := ""
				tittle := "News -" + News.Title
				Body := News.Message
				image := ""
				data := make(map[string]string)
				data["notificationType"] = "ViewSingleNews"
				data["id"] = News.UniqueID
				err := s.SendNotification(topic, tittle, Body, image, token, data)
				if err != nil {
					log.Println(apptoken.RegistrationToken + " " + err.Error())
				}
				if err == nil {
					t := time.Now()
					ToNotificationLog := new(models.ToNotificationLog)
					notifylog := new(models.NotificationLog)
					ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
					ToNotificationLog.Name = v.Name
					ToNotificationLog.UserName = v.UniqueID
					ToNotificationLog.UserType = "Employee"
					notifylog.Body = Body
					notifylog.Tittle = tittle
					notifylog.Topic = topic
					notifylog.Image = image
					notifylog.IsJob = false
					notifylog.Message = Body
					notifylog.SentDate = &t
					notifylog.SentFor = constants.COLLECTIONNEWS
					notifylog.SendForId = News.UniqueID
					notifylog.Data = data
					notifylog.Status = "Active"
					notifylog.To = *ToNotificationLog
					err = s.Daos.SaveNotificationLog(ctx, notifylog)
					if err != nil {
						return err
					}
				}
			}
			if v.OfficialEmail != "" {
				fmt.Println("Employee email======>", v.OfficialEmail)
				sendmailto = append(sendmailto, v.OfficialEmail)

				err = s.SendEmailWithHtml(News.Title, sendmailto, News.Message)
				if err != nil {
					return errors.New("email Sending Error - " + err.Error())
				}
				if err == nil {
					emaillog := new(models.EmailLog)
					to2 := models.ToEmailLog{}
					to2.Email = v.Email
					to2.Name = v.UserName
					to2.UserName = v.UserName
					to2.UserType = "Employee"
					t := time.Now()
					emaillog.SentDate = &t
					emaillog.IsJob = false
					emaillog.Message = News.Message
					emaillog.SendForId = News.UniqueID
					emaillog.SentFor = constants.COLLECTIONNEWS
					emaillog.Status = "Active"
					emaillog.To = to2
					err = s.Daos.SaveEmailLog(ctx, emaillog)
					if err != nil {
						return errors.New("email not save")
					}
				}
				sendmailto = []string{}
				fmt.Println("last Email", sendmailto)
				continue
			} else if v.Email != "" {
				fmt.Println("Employee email======>", v.Email)
				sendmailto = append(sendmailto, v.Email)
				err = s.SendEmailWithHtml(News.Title, sendmailto, News.Message)
				if err != nil {
					return errors.New("email Sending Error - " + err.Error())
				}
				if err == nil {
					emaillog := new(models.EmailLog)
					to2 := models.ToEmailLog{}
					to2.Email = v.Email
					to2.Name = v.UserName
					to2.UserName = v.UserName
					to2.UserType = "Employee"
					t := time.Now()
					emaillog.SentDate = &t
					emaillog.IsJob = false
					emaillog.Message = News.Message
					emaillog.SendForId = News.UniqueID
					emaillog.SentFor = constants.COLLECTIONNEWS
					emaillog.Status = "Active"
					emaillog.To = to2
					err = s.Daos.SaveEmailLog(ctx, emaillog)
					if err != nil {
						return errors.New("email not save")
					}
				}
				sendmailto = []string{}
				fmt.Println("last Email", sendmailto)
				continue
			} else if v.Mobile != "" {
				err = s.SendSMS(v.Mobile, News.Message)
				if err != nil {
					return errors.New(v.Mobile + " " + err.Error())
				}

				if err == nil {
					smslog := new(models.SmsLog)
					to := models.To{}
					to.No = v.Mobile
					to.Name = v.Name
					to.UserType = "Employee"
					to.UserName = v.UserName
					t := time.Now()
					smslog.SentDate = &t
					smslog.IsJob = false
					smslog.Message = News.Message
					smslog.SentFor = constants.COLLECTIONNEWS
					smslog.SendForId = News.UniqueID
					smslog.Status = constants.SMSLOGSTATUSACTIVE
					smslog.To = to
					smslog.SendForId = News.UniqueID
					err = s.Daos.SaveSmsLog(ctx, smslog)
					if err != nil {
						return errors.New("sms not save")
					}
				}
			}
		}

		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}
