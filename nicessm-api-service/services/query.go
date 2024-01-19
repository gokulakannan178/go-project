package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveQuery :""
func (s *Service) SaveQuery(ctx *models.Context, Query *models.Query) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//Query.Code = s.Daos.GetUniqueID(ctx, constants.COLLECTIONQuery)
	t := time.Now()
	var uniqueid string
	if int(t.Month()) < 10 {
		if t.Day() < 10 {
			uniqueid = fmt.Sprintf("%v0%v0%v", t.Year(), int(t.Month()), t.Day())
		} else {
			uniqueid = fmt.Sprintf("%v0%v%v", t.Year(), int(t.Month()), t.Day())
		}
	} else {
		if t.Day() < 0 {
			uniqueid = fmt.Sprintf("%v%v0%v", t.Year(), int(t.Month()), t.Day())
		} else {
			uniqueid = fmt.Sprintf("%v%v%v", t.Year(), int(t.Month()), t.Day())
		}
	}
	fmt.Println("uniqueid====>", uniqueid)
	Query.UniqueID = fmt.Sprintf("%v%v_%v", "Q", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONQUERY))
	fmt.Println("Query.UniqueID ====>", Query.UniqueID)
	Query.Status = constants.QUERYSTATUSCREATED
	Query.ActiveStatus = true
	//t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	Query.Date = &t
	log.Println("b4 Query.created")
	Query.Created = created
	log.Println("b4 Query.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if !Query.KnowledgeDomain.IsZero() && !Query.SubDomain.IsZero() {
			user, _ := s.Daos.GetSingleUserWithQueryCount(ctx, Query.KnowledgeDomain.Hex(), Query.SubDomain.Hex())
			if user != nil {
				fmt.Println("Assinged User====>", user.UserName)
				Query.AssignedTo = user.ID
				Query.AssignedDate = &t
				Query.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONQUERY)
				Query.Status = constants.QUERYSTATUSASSIGN
				Query.ActiveStatus = true

			}
		}
		dberr := s.Daos.SaveQuery(ctx, Query)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		if !Query.Farmer.IsZero() {
			farmer, err := s.Daos.GetSingleFarmer(ctx, Query.Farmer.Hex())
			if err != nil {
				return err
			}
			msg := fmt.Sprintf(constants.COMMONTEMPLATE, farmer.Name, "NICESSM", "The Query you have Raised", "Your Query is Successfully Registered,Ref no is -"+Query.UniqueID+"", "https://nicessm.org/")

			//	msg := fmt.Sprintf(constants.COMMONTEMPLATE, farmer.Name, "NICESSM", "The Query you have Raised", "Your Query is Successfully Registered,Ref no is -"+Query.UniqueID+"", "https://nicessm.org/")
			err = s.SendSMSV2(ctx, farmer.MobileNumber, msg)
			if err != nil {
				return errors.New(farmer.MobileNumber + " " + err.Error())
			}
			if err == errors.New(constants.INSUFFICIENTBALANCE) {
				return err
			}
			if err == nil {
				smslog := new(models.SmsLog)
				to := models.To{}
				to.No = farmer.MobileNumber
				to.Name = farmer.Name
				to.UserType = "Farmer"
				to.UserName = farmer.FarmerID
				t := time.Now()
				smslog.SentDate = &t
				smslog.IsJob = false
				smslog.Message = msg
				smslog.SentFor = "Query"
				smslog.Status = constants.SMSLOGSTATUSACTIVE
				smslog.To = to
				smslog.Query = Query.ID
				smslog.QueryRecordId = Query.UniqueID
				err = s.Daos.SaveSmsLog(ctx, smslog)
				if err != nil {
					return errors.New("query sms not save")
				}
			}
			apptoken, _ := s.Daos.GetSingleApptokenWithUserID(ctx, Query.Farmer.Hex())
			if apptoken != nil {
				fmt.Println("apptoken===>", apptoken.Apptoken.RegistrationToken)
				var token []string
				token = append(token, apptoken.RegistrationToken)

				fmt.Println("appToken===>", apptoken.RegistrationToken)
				topic := ""
				tittle := "Query -" + Query.UniqueID + " Query Submitted Successfully"
				Body := Query.Query
				var image string
				if len(Query.Images) > 0 {
					image = Query.Images[0]
				}
				data := make(map[string]string)
				data["notificationType"] = "ViewSingleQuery"
				data["id"] = Query.ID.Hex()
				err := s.SendNotification(topic, tittle, Body, image, token, data)
				if err != nil {
					log.Println(apptoken.RegistrationToken + " " + err.Error())
				}
				if err == nil {
					t = time.Now()
					ToNotificationLog := new(models.ToNotificationLog)
					notifylog := new(models.NotificationLog)
					ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
					ToNotificationLog.Name = farmer.Name
					ToNotificationLog.UserName = farmer.ID
					ToNotificationLog.UserType = "Farmer"
					notifylog.Body = Body
					notifylog.Tittle = tittle
					notifylog.Topic = topic
					notifylog.Image = image
					notifylog.IsJob = false
					notifylog.Message = Body
					notifylog.SentDate = &t
					notifylog.SentFor = topic
					notifylog.Data = data
					notifylog.Status = "Active"
					notifylog.To = *ToNotificationLog
					err = s.Daos.SaveNotificationLog(ctx, notifylog)
					if err != nil {
						return err
					}
				}
			}
		}
		if !Query.KnowledgeDomain.IsZero() && !Query.SubDomain.IsZero() {
			user, _ := s.Daos.GetSingleUserWithQueryCount(ctx, Query.KnowledgeDomain.Hex(), Query.SubDomain.Hex())
			if !Query.Farmer.IsZero() {
				farmer, err := s.Daos.GetSingleFarmer(ctx, Query.Farmer.Hex())
				if err != nil {
					return err
				}
				if farmer == nil {
					log.Println("farmer not found")
				}
				if user == nil {
					log.Println("user not found")
				}
				msg := fmt.Sprintf(constants.COMMONTEMPLATE, farmer.Name, "NICESSM", "Your Query Assinged to the User", "Your Query is Assigned user,user name is -"+user.UserName+"", "https://nicessm.org/")

				//	msg := fmt.Sprintf(constants.COMMONTEMPLATE, farmer.Name, "NICESSM", "Your Query Assinged to the User", "Your Query is Assigned user,user name is -"+user.UserName+"", "https://nicessm.org/")
				err = s.SendSMSV2(ctx, farmer.MobileNumber, msg)
				if err != nil {
					return errors.New(farmer.MobileNumber + " " + err.Error())
				}
				if err == errors.New(constants.INSUFFICIENTBALANCE) {
					return err
				}
				if err == nil {
					smslog := new(models.SmsLog)
					to := models.To{}
					to.No = farmer.MobileNumber
					to.Name = farmer.Name
					to.UserType = "Farmer"
					to.UserName = farmer.FarmerID
					t := time.Now()
					smslog.SentDate = &t
					smslog.IsJob = false
					smslog.Message = msg
					smslog.SentFor = "Query"
					smslog.Status = constants.SMSLOGSTATUSACTIVE
					smslog.Query = Query.ID
					smslog.QueryRecordId = Query.UniqueID
					smslog.To = to
					err = s.Daos.SaveSmsLog(ctx, smslog)
					if err != nil {
						return errors.New("query sms not save")
					}
				}
				apptoken, _ := s.Daos.GetSingleApptokenWithUserID(ctx, Query.Farmer.Hex())
				if apptoken != nil {
					fmt.Println("apptoken===>", apptoken.Apptoken.RegistrationToken)
					var token []string
					token = append(token, apptoken.RegistrationToken)

					fmt.Println("appToken===>", apptoken.RegistrationToken)
					topic := ""
					tittle := "Query -" + Query.UniqueID + " Query Assigned To User"
					Body := Query.Query
					var image string
					if len(Query.Images) > 0 {
						image = Query.Images[0]
					}
					data := make(map[string]string)
					data["notificationType"] = "ViewSingleQuery"
					data["id"] = Query.ID.Hex()
					err = s.SendNotification(topic, tittle, Body, image, token, data)
					if err != nil {
						log.Println("Printing error for query assign", apptoken.RegistrationToken+" "+err.Error())
					}
					if err == nil {
						t = time.Now()
						ToNotificationLog := new(models.ToNotificationLog)
						notifylog := new(models.NotificationLog)
						ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
						ToNotificationLog.Name = farmer.Name
						ToNotificationLog.UserName = farmer.ID
						ToNotificationLog.UserType = "Farmer"
						notifylog.Body = Body
						notifylog.Tittle = tittle
						notifylog.Topic = topic
						notifylog.Image = image
						notifylog.IsJob = false
						notifylog.Message = Body
						notifylog.SentDate = &t
						notifylog.SentFor = topic
						notifylog.Data = data
						notifylog.Status = "Active"
						notifylog.To = *ToNotificationLog
						err = s.Daos.SaveNotificationLog(ctx, notifylog)
						if err != nil {
							return err
						}
					}
				}

			} else {
				Query.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONQUERY)
				Query.Status = constants.QUERYSTATUSCREATED
				Query.ActiveStatus = true
				t := time.Now()
				created := models.Created{}
				created.On = &t
				created.By = constants.SYSTEM
				Query.Date = &t
			}
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// err := s.DisseminateContent(ctx, dissemination)
		// if err != nil {
		// 	fmt.Println("err", err)
		// }

		// }
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

//UpdateQuery : ""
func (s *Service) UpdateQuery(ctx *models.Context, Query *models.Query) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Query.Status = constants.QUERYSTATUSMODIFY
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateQuery(ctx, Query)
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

//EnableQuery : ""
func (s *Service) EnableQuery(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableQuery(ctx, UniqueID)
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

//DisableQuery : ""
func (s *Service) DisableQuery(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableQuery(ctx, UniqueID)
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

//DeleteQuery : ""
func (s *Service) DeleteQuery(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteQuery(ctx, UniqueID)
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

//GetSingleQuery :""
func (s *Service) GetSingleQuery(ctx *models.Context, UniqueID string) (*models.RefQuery, error) {
	Query, err := s.Daos.GetSingleQuery(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Query, nil
}

//FilterQuery :""
func (s *Service) FilterQuery(ctx *models.Context, Queryfilter *models.QueryFilter, pagination *models.Pagination) (Query []models.RefQuery, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	err = s.QueryDataAccess(ctx, Queryfilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterQuery(ctx, Queryfilter, pagination)

}

//AssignuserQuery : ""
func (s *Service) AssignuserQuery(ctx *models.Context, Query *models.AssignUserToQuery) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.AssignuserQuery(ctx, Query)
		if err != nil {
			return err
		}
		// if dberr != nil {

		// 	return errors.New("Db Error" + dberr.Error())
		// }
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// err := s.DisseminateContent(ctx, dissemination)
		// if err != nil {
		// 	fmt.Println("err", err)
		// }
		user, err := s.Daos.GetSingleUser(ctx, Query.UserId.Hex())
		fmt.Println("username===>", user.Name)
		if err != nil {
			return err
		}
		queryId, err := s.Daos.GetSingleQuery(ctx, Query.QueryId.Hex())
		if err != nil {
			return err
		}
		fmt.Println("QueryId===>", queryId.UniqueID)
		if !Query.UserId.IsZero() {

			msg := fmt.Sprintf(constants.COMMONTEMPLATE, user.Name, "NICESSM", "Ref no is -"+queryId.Query.UniqueID+"", queryId.Query.Query, "https://nicessm.org/")
			err = s.SendSMSV2(ctx, user.Mobile, msg)
			if err != nil {
				return errors.New(user.Mobile + " " + err.Error())
			}
			if err == errors.New(constants.INSUFFICIENTBALANCE) {
				return err
			}
			if err == nil {
				smslog := new(models.SmsLog)
				to := models.To{}
				to.No = user.Mobile
				to.Name = user.Name
				to.UserType = "User"
				to.UserName = user.UserName
				t := time.Now()
				smslog.SentDate = &t
				smslog.IsJob = false
				smslog.Message = msg
				smslog.SentFor = "Query"
				smslog.To = to
				smslog.Status = constants.SMSLOGSTATUSACTIVE
				smslog.Query = queryId.ID
				smslog.QueryRecordId = queryId.UniqueID
				err = s.Daos.SaveSmsLog(ctx, smslog)
				if err != nil {
					return errors.New("query sms not save")
				}
			}
			// apptoken, err := s.Daos.GetSingleApptokenWithUserID(ctx, Query.UserId.Hex())
			// if err != nil {
			// 	return err
			// }
			// var token []string
			// token = append(token, apptoken.RegistrationToken)
			// fmt.Println("appToken===>", apptoken.RegistrationToken)
			// topic := ""
			// tittle := "Query -" + queryId.UniqueID + ""
			// Body := queryId.Query.Query
			// image := queryId.Images[0]
			// data := make(map[string]string)
			// data["Query Type"] = "ViewSingleQuery"
			// data["id"] = queryId.ID.Hex()
			// err = s.SendNotification(topic, tittle, Body, image, token, data)
			// if err != nil {
			// 	log.Println(apptoken.RegistrationToken + " " + err.Error())
			// }
			// t = time.Now()
			// ToNotificationLog := new(models.ToNotificationLog)
			// notifylog := new(models.NotificationLog)
			// ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
			// to.Name = user.Name
			// to.UserName = user.Name
			// to.UserType = "User"
			// notifylog.Body = Body
			// notifylog.Tittle = tittle
			// notifylog.Topic = topic
			// notifylog.Image = image
			// notifylog.IsJob = false
			// notifylog.Message = Body
			// notifylog.SentDate = &t
			// notifylog.SentFor = topic
			// notifylog.Data = data
			// notifylog.Status = "Active"
			// notifylog.To = *ToNotificationLog
			// err = s.Daos.SaveNotificationLog(ctx, notifylog)
			// if err != nil {
			// 	return err
			// }
		}
		queryid, _ := s.Daos.GetSingleQuery(ctx, Query.QueryId.Hex())

		if !queryid.Farmer.IsZero() {
			farmer, err := s.Daos.GetSingleFarmer(ctx, queryid.Farmer.Hex())
			if err != nil {
				return err
			}
			msg := fmt.Sprintf(constants.COMMONTEMPLATE, farmer.Name, "NICESSM", "Your Query Assinged to the User", "Your Query is Assigned user,user name is -"+user.UserName+"", "https://nicessm.org/")
			err = s.SendSMSV2(ctx, farmer.MobileNumber, msg)
			if err != nil {
				return errors.New(farmer.MobileNumber + " " + err.Error())
			}
			if err == errors.New(constants.INSUFFICIENTBALANCE) {
				return err
			}
			if err == nil {
				smslog := new(models.SmsLog)
				to := models.To{}
				to.No = farmer.MobileNumber
				to.Name = farmer.Name
				to.UserType = "Farmer"
				to.UserName = farmer.FarmerID
				t := time.Now()
				smslog.SentDate = &t
				smslog.IsJob = false
				smslog.Message = msg
				smslog.SentFor = "Query"
				smslog.To = to
				smslog.Status = constants.SMSLOGSTATUSACTIVE
				smslog.Query = queryId.ID
				smslog.QueryRecordId = queryId.UniqueID
				err = s.Daos.SaveSmsLog(ctx, smslog)
				if err != nil {
					return errors.New("query sms not save")
				}

			}
			apptoken, _ := s.Daos.GetSingleApptokenWithUserID(ctx, queryid.Farmer.Hex())
			if apptoken != nil {
				fmt.Println("apptoken===>", apptoken.Apptoken.RegistrationToken)
				var token []string
				token = append(token, apptoken.RegistrationToken)

				fmt.Println("appToken===>", apptoken.RegistrationToken)
				topic := ""
				tittle := "Query -" + queryid.UniqueID + "Your Query is Assigned To User"
				Body := queryid.Query.Query
				var image string
				if len(queryid.Images) > 0 {
					image = queryid.Images[0]
				}
				data := make(map[string]string)
				data["notificationType"] = "ViewSingleQuery"
				data["id"] = queryid.ID.Hex()
				err = s.SendNotification(topic, tittle, Body, image, token, data)
				if err != nil {
					log.Println(apptoken.RegistrationToken + " " + err.Error())
				}
				if err == nil {
					t := time.Now()
					ToNotificationLog := new(models.ToNotificationLog)
					notifylog := new(models.NotificationLog)
					ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
					ToNotificationLog.Name = farmer.Name
					ToNotificationLog.UserName = farmer.ID
					ToNotificationLog.UserType = "Farmer"
					notifylog.Body = Body
					notifylog.Tittle = tittle
					notifylog.Topic = topic
					notifylog.Image = image
					notifylog.IsJob = false
					notifylog.Message = Body
					notifylog.SentDate = &t
					notifylog.SentFor = topic
					notifylog.Data = data
					notifylog.Status = "Active"
					notifylog.To = *ToNotificationLog
					err = s.Daos.SaveNotificationLog(ctx, notifylog)
					if err != nil {
						return err
					}

				}
			}

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

//ResolveuserQuery : ""
func (s *Service) ResolveuserQuery(ctx *models.Context, Query *models.ResolveQuery) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.ResolveuserQuery(ctx, Query)
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// err := s.DisseminateContent(ctx, dissemination)
		// if err != nil {
		// 	fmt.Println("err", err)
		// }
		queryId, err := s.Daos.GetSingleQuery(ctx, Query.QueryId.Hex())
		if err != nil {
			return err
		}
		if !queryId.Farmer.IsZero() {
			farmer, err := s.Daos.GetSingleFarmer(ctx, queryId.Farmer.Hex())
			if err != nil {
				return err
			}
			msg := fmt.Sprintf(constants.COMMONTEMPLATE, farmer.Name, "NICESSM", "Solution for - "+queryId.Query.UniqueID+"", Query.Solution, "https://nicessm.org/")

			//	msg := fmt.Sprintf(constants.COMMONTEMPLATE, farmer.Name, "NICESSM", "Solution for - "+queryId.Query.UniqueID+"", Query.Solution, "https://nicessm.org/")
			err = s.SendSMSV2(ctx, farmer.MobileNumber, msg)
			if err != nil {
				return errors.New(farmer.MobileNumber + " " + err.Error())
			}
			if err == errors.New(constants.INSUFFICIENTBALANCE) {
				return err
			}
			if err == nil {
				smslog := new(models.SmsLog)
				to := models.To{}
				to.No = farmer.MobileNumber
				to.Name = farmer.Name
				to.UserType = "Farmer"
				to.UserName = farmer.FarmerID
				smslog.IsJob = false
				t := time.Now()
				smslog.SentDate = &t
				smslog.Message = msg
				smslog.SentFor = "Query"
				smslog.Query = queryId.ID
				smslog.Status = constants.SMSLOGSTATUSACTIVE
				smslog.QueryRecordId = queryId.UniqueID
				smslog.To = to
				err = s.Daos.SaveSmsLog(ctx, smslog)
				if err != nil {
					return errors.New("query sms not save")
				}
			}
			apptoken, _ := s.Daos.GetSingleApptokenWithUserID(ctx, queryId.Farmer.Hex())
			if apptoken != nil {
				fmt.Println("apptoken===>", apptoken.Apptoken.RegistrationToken)
				var token []string
				token = append(token, apptoken.RegistrationToken)

				fmt.Println("appToken===>", apptoken.RegistrationToken)
				topic := ""
				tittle := "Query -" + queryId.UniqueID + " Your query is Resloved"
				Body := queryId.Query.Query
				var image string
				if len(queryId.Images) > 0 {
					image = queryId.Images[0]
				}
				data := make(map[string]string)
				data["notificationType"] = "ViewSingleQuery"
				data["id"] = queryId.ID.Hex()
				err = s.SendNotification(topic, tittle, Body, image, token, data)
				if err != nil {
					log.Println(apptoken.RegistrationToken + " " + err.Error())
				}
				if err == nil {
					t := time.Now()
					ToNotificationLog := new(models.ToNotificationLog)
					notifylog := new(models.NotificationLog)
					ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
					ToNotificationLog.Name = farmer.Name
					ToNotificationLog.UserName = farmer.ID
					ToNotificationLog.UserType = "Farmer"
					notifylog.Body = Body
					notifylog.Tittle = tittle
					notifylog.Topic = topic
					notifylog.Image = image
					notifylog.IsJob = false
					notifylog.Message = Body
					notifylog.SentDate = &t
					notifylog.SentFor = topic
					notifylog.Data = data
					notifylog.Status = "Active"
					notifylog.To = *ToNotificationLog
					err = s.Daos.SaveNotificationLog(ctx, notifylog)
					if err != nil {
						return err
					}
				}
			}
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
func (s *Service) QueryDataAccess(ctx *models.Context, Queryfilter *models.QueryFilter) (err error) {
	if Queryfilter != nil {

		dataaccess, err := s.Daos.DataAccess(ctx, &Queryfilter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.AccessStates) > 0 {
				for _, v := range dataaccess.AccessStates {
					Queryfilter.State = append(Queryfilter.State, v.ID)
				}
			}
			if len(dataaccess.AccessDistricts) > 0 {
				for _, v := range dataaccess.AccessDistricts {
					Queryfilter.District = append(Queryfilter.District, v.ID)
				}
			}
			if len(dataaccess.AccessBlocks) > 0 {
				for _, v := range dataaccess.AccessBlocks {
					Queryfilter.Block = append(Queryfilter.Block, v.ID)
				}
			}
			if len(dataaccess.AccessVillages) > 0 {
				for _, v := range dataaccess.AccessVillages {
					Queryfilter.Village = append(Queryfilter.Village, v.ID)

				}
			}
			if len(dataaccess.AccessGrampanchayats) > 0 {
				for _, v := range dataaccess.AccessGrampanchayats {
					Queryfilter.GramPanchayat = append(Queryfilter.GramPanchayat, v.ID)

				}
			}
		}

	}
	return err
}
