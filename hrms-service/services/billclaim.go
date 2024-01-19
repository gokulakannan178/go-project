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

//SaveBillClaim :""
func (s *Service) SaveBillClaim(ctx *models.Context, billClaim *models.BillClaim) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	billClaim.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIM)
	billClaim.Status = constants.BILLCLAIMSTATUSPENDING
	t := time.Now()
	billClaim.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 BillClaim.created")
	billClaim.Created = created
	log.Println("b4 BillClaim.created")
	var totalamount float64
	for _, v := range billClaim.Bills {
		totalamount = totalamount + v.Amount
		fmt.Println("totalamount==>", totalamount)
	}
	billClaim.TotalAmount = totalamount

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveBillClaim(ctx, billClaim)
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
func (s *Service) SaveBillClaimV2(ctx *models.Context, billClaim *models.BillClaim) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	billClaim.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIM)
	billClaim.Status = constants.BILLCLAIMSTATUSPENDING
	t := time.Now()
	billClaim.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 BillClaim.created")
	billClaim.Created = created
	log.Println("b4 BillClaim.created")
	var totalamount float64
	for _, v := range billClaim.Bills {
		totalamount = totalamount + v.Amount
		fmt.Println("totalamount==>", totalamount)
	}
	billClaim.TotalAmount = totalamount

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		employee, err := s.Daos.GetSingleEmployeeWithUserName(ctx, billClaim.EmployeeId)
		if err != nil {
			return err
		}
		billslevel, err := s.Daos.GetSingleBillclaimApprovalLevels(ctx, billClaim.GradeId, billClaim.TotalAmount)
		if err != nil {
			return err
		}
		if billslevel == nil {
			return errors.New("BillsConfig Not Found")
		}
		linemanager := employee.LineManager
		fmt.Println("bill", billslevel)
		fmt.Println("level", billslevel.Level)
		for i := 0; i < int(billslevel.Level); i++ {
			employees, err := s.Daos.GetSingleEmployeeWithUserName(ctx, linemanager)
			if err != nil {
				return err
			}

			billsapproval := new(models.BillclaimLevels)
			billsapproval.Bill = billClaim.UniqueID
			billsapproval.AssignedBy = linemanager
			billsapproval.EmployeeId = billClaim.EmployeeId
			billsapproval.Grade = billClaim.GradeId
			billsapproval.Level = int64(i + 1)
			billsapproval.NoOfLevel = billslevel.Level
			billsapproval.Organisation = billslevel.Organisation
			billsapproval.Status = constants.BILLCLAIMSTATUSINIT
			if i == 0 {
				billsapproval.Status = constants.BILLCLAIMSTATUSPENDING
			}
			fmt.Println("Status", billsapproval.Status)

			err = s.SaveBillclaimLevelsWithoutTransaction(ctx, billsapproval)
			if err != nil {
				return err
			}
			fmt.Println("Linemanage", i, employees.Name)
			if employees.LineManager != "" {
				linemanager = employees.LineManager
			} else {
				break
			}

		}
		dberr := s.Daos.SaveBillClaim(ctx, billClaim)
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

//UpdateBillClaim : ""
func (s *Service) UpdateBillClaim(ctx *models.Context, billClaim *models.BillClaim) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		var totalamount float64
		for _, v := range billClaim.Bills {
			totalamount = totalamount + v.Amount
			fmt.Println("totalamount==>", totalamount)
		}
		billClaim.TotalAmount = totalamount

		dberr := s.Daos.UpdateBillClaim(ctx, billClaim)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		billClaimlog := new(models.BillClaimLog)
		bill, err := s.Daos.GetSingleBillClaim(ctx, billClaim.UniqueID)
		if err != nil {
			return err
		}
		billClaimlog.BillClaim = billClaim.UniqueID
		billClaimlog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIMLOG)
		t := time.Now()
		billClaimlog.Log = billClaim.Updated
		billClaimlog.Status = constants.BILLCLAIMLOGSTATUSACTIVE
		billClaimlog.Log.On = &t
		billClaimlog.Log.Scenario = "Edit a BillClaim"
		billClaimlog.Previous = *bill
		billClaimlog.New = *billClaim
		err = s.Daos.SaveBillClaimLog(ctx, billClaimlog)
		if err != nil {
			return err
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

//EnableBillClaim : ""
func (s *Service) EnableBillClaim(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableBillClaim(ctx, UniqueID)
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

//DisableBillClaim : ""
func (s *Service) DisableBillClaim(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableBillClaim(ctx, UniqueID)
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

//DeleteBillClaim : ""
func (s *Service) DeleteBillClaim(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteBillClaim(ctx, UniqueID)
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

//GetSingleBillClaim :""
func (s *Service) GetSingleBillClaim(ctx *models.Context, UniqueID string) (*models.RefBillClaim, error) {
	billClaim, err := s.Daos.GetSingleBillClaim(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return billClaim, nil
}

//FilterBillClaim :""
func (s *Service) FilterBillClaim(ctx *models.Context, filter *models.FilterBillClaim, pagination *models.Pagination) ([]models.RefBillClaim, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterBillClaim(ctx, filter, pagination)

}
func (s *Service) ApprovedBillClaim(ctx *models.Context, approved *models.ReviewedBillClaim) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.ApprovedBillClaim(ctx, approved)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		bill, err := s.Daos.GetSingleBillClaim(ctx, approved.BillClaim)
		if err != nil {
			return err
		}

		apptoken, err := s.Daos.GetRegTokenWithParticulars(ctx, bill.EmployeeId)
		if err != nil {
			return err
		}
		if apptoken != nil {
			fmt.Println("apptoken===>", apptoken.RegistrationToken)
			var token []string
			token = append(token, apptoken.RegistrationToken)

			fmt.Println("appToken===>", apptoken.RegistrationToken)
			topic := ""
			tittle := "Employee -" + approved.BillClaim + "Bills Approved"
			Body := bill.Title
			//	var image string
			//if len(employeeTimeOff.) > 0 {
			image := ""
			//	}
			data := make(map[string]string)
			data["notificationType"] = "ViewBillclaim"
			data["id"] = bill.UniqueID
			err := s.SendNotification(topic, tittle, Body, image, token, data)
			if err != nil {
				log.Println(apptoken.RegistrationToken + " " + err.Error())
			}
			if err == nil {
				t := time.Now()
				ToNotificationLog := new(models.ToNotificationLog)
				notifylog := new(models.NotificationLog)
				ToNotificationLog.AppRegistrationToken = apptoken.RegistrationToken
				ToNotificationLog.Name = bill.Title
				ToNotificationLog.UserName = bill.EmployeeId
				ToNotificationLog.UserType = "Employee"
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
		billClaimlog := new(models.BillClaimLog)
		billClaimlog.BillClaim = approved.BillClaim
		billClaimlog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIMLOG)
		t := time.Now()
		billClaimlog.Log.By = approved.ReviewedBy
		billClaimlog.Log.Scenario = "Approved BillClaim"
		billClaimlog.Status = constants.BILLCLAIMLOGSTATUSACTIVE
		billClaimlog.Log.On = &t
		err = s.Daos.SaveBillClaimLog(ctx, billClaimlog)
		if err != nil {
			return err
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
func (s *Service) RejectedBillClaim(ctx *models.Context, rejected *models.ReviewedBillClaim) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.RejectedBillClaim(ctx, rejected)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		billClaimlog := new(models.BillClaimLog)
		billClaimlog.BillClaim = rejected.BillClaim
		billClaimlog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBILLCLAIMLOG)
		t := time.Now()
		billClaimlog.Log.By = rejected.ReviewedBy
		billClaimlog.Log.Scenario = "Rejected BillClaim"
		billClaimlog.Status = constants.BILLCLAIMLOGSTATUSACTIVE
		billClaimlog.Log.On = &t
		err := s.Daos.SaveBillClaimLog(ctx, billClaimlog)
		if err != nil {
			return err
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
