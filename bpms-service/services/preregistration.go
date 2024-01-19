package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePreregistration :"Draft"
func (s *Service) SavePreregistration(ctx *models.Context, preregistration *models.Preregistration) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if preregistration.UniqueID == "" {
		preregistration.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPREREGISTRATION)
		t := time.Now()
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		log.Println("b4 preregistration.created")
		preregistration.Created = created
		log.Println("b4 preregistration.created")
	}
	preregistration.Status = constants.PREREGISTRATIONSTATUSDRAFT
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		emailOK, err := s.Daos.ValidateUniquenessAtUpdateV2(ctx, preregistration.MobileNumber, preregistration.UniqueID, "email")
		if err != nil {
			return err
		}
		if !emailOK {
			return errors.New("Duplicate Email")
		}
		mobileOK, err := s.Daos.ValidateUniquenessAtUpdateV2(ctx, preregistration.MobileNumber, preregistration.UniqueID, "mobileNumber")
		if err != nil {
			return err
		}
		if !mobileOK {
			return errors.New("Duplicate mobile")
		}

		dberr := s.Daos.SavePreregistrationV2(ctx, preregistration)
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

//SubmitPreregistration :""
func (s *Service) SubmitPreregistration(ctx *models.Context, preregistration *models.Preregistration) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	if preregistration.UniqueID == "" {
		preregistration.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPREREGISTRATION)
		if preregistration.Created.On == nil {
			log.Println("b4 preregistration.created")
			preregistration.Created = created
			log.Println("b4 preregistration.created")
		}
	}
	preregistration.Status = constants.PREREGISTRATIONSTATUSSUBMITTED

	created.By = preregistration.UniqueID
	preregistration.Submitted = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		emailOK, err := s.Daos.ValidateUniquenessAtUpdateV2(ctx, preregistration.MobileNumber, preregistration.UniqueID, "email")
		if err != nil {
			return err
		}
		if !emailOK {
			return errors.New("Duplicate Email")
		}
		mobileOK, err := s.Daos.ValidateUniquenessAtUpdateV2(ctx, preregistration.MobileNumber, preregistration.UniqueID, "mobileNumber")
		if err != nil {
			return err
		}
		if !mobileOK {
			return errors.New("Duplicate mobile")
		}
		dberr := s.Daos.SubmitPreregistration(ctx, preregistration)
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

//ReapplyPreregistration :""
func (s *Service) ReapplyPreregistration(ctx *models.Context, preregistration *models.Preregistration) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	preregistration.Status = constants.PREREGISTRATIONSTATUSREAPPIED
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = preregistration.UniqueID
	preregistration.Reapplied = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		emailOK, err := s.Daos.ValidateUniquenessAtUpdateV2(ctx, preregistration.MobileNumber, preregistration.UniqueID, "email")
		if err != nil {
			return err
		}
		if !emailOK {
			return errors.New("Duplicate Email")
		}
		mobileOK, err := s.Daos.ValidateUniquenessAtUpdateV2(ctx, preregistration.MobileNumber, preregistration.UniqueID, "mobileNumber")
		if err != nil {
			return err
		}
		if !mobileOK {
			return errors.New("Duplicate mobile")
		}
		dberr := s.Daos.ReapplyPreregistration(ctx, preregistration)
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

//UpdatePreregistration : ""
func (s *Service) UpdatePreregistration(ctx *models.Context, preregistration *models.Preregistration) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePreregistration(ctx, preregistration)
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

//EnablePreregistration : ""
func (s *Service) EnablePreregistration(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePreregistration(ctx, UniqueID)
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

//DisablePreregistration : ""
func (s *Service) DisablePreregistration(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePreregistration(ctx, UniqueID)
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

//DeletePreregistration : ""
func (s *Service) DeletePreregistration(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePreregistration(ctx, UniqueID)
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

//GetSinglePreregistration :""
func (s *Service) GetSinglePreregistration(ctx *models.Context, mobileNumber string) (*models.RefPreregistration, error) {
	preregistration, err := s.Daos.GetSinglePreregistration(ctx, mobileNumber)
	if err != nil {
		return nil, err
	}
	return preregistration, nil
}

//FilterPreregistration :""
func (s *Service) FilterPreregistration(ctx *models.Context, preregistrationfilter *models.PreregistrationFilter, pagination *models.Pagination) (preregistration []models.RefPreregistration, err error) {
	return s.Daos.FilterPreregistration(ctx, preregistrationfilter, pagination)
}

//ValidateMobileNumber :""
func (s *Service) ValidateMobileNumber(ctx *models.Context, mobileNumber, uniqueID string) (bool, *models.RefPreregistration, error) {
	isvalid, preregistration, err := s.Daos.ValidateMobileNumber(ctx, mobileNumber, uniqueID)
	if err != nil {
		return isvalid, nil, err
	}
	return isvalid, preregistration, nil
}

//ValidateEmailAtUpdate :""
func (s *Service) ValidateEmailAtUpdate(ctx *models.Context, email, uniqueID string) error {
	err := s.Daos.ValidateEmailAtUpdate(ctx, email, uniqueID)
	if err != nil {
		return err
	}
	return nil
}

//ValidateMobileAtUpdate :""
func (s *Service) ValidateMobileAtUpdate(ctx *models.Context, mobile, uniqueID string) error {
	err := s.Daos.ValidateMobileAtUpdate(ctx, mobile, uniqueID)
	if err != nil {
		return err
	}
	return nil
}

//GetSinglePreregistrationWithUniqueID :""
func (s *Service) GetSinglePreregistrationWithUniqueID(ctx *models.Context, uniqueID string) (*models.RefPreregistration, error) {
	preregistration, err := s.Daos.GetSinglePreregistrationWithUniqueID(ctx, uniqueID)
	if err != nil {
		return nil, err
	}
	return preregistration, nil
}

//PreregistrationStatusChange : ""
func (s *Service) PreregistrationStatusChange(ctx *models.Context, psc *models.PreregistrationStatusChange) error {
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		var err1 error
		psc.Status, err1 = s.PreregistrationStatusValidation(ctx, psc)
		if err1 != nil {
			return err1
		}
		psc.On = &t
		dberr := s.Daos.PreregistrationStatusChange(ctx, psc)
		if dberr != nil {

			return errors.New("Transaction Aborted - " + dberr.Error())
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

//PreregistrationStatusValidation : ""
func (s *Service) PreregistrationStatusValidation(ctx *models.Context, psc *models.PreregistrationStatusChange) (string, error) {
	switch psc.Status {
	case "Accept":
		return constants.PREREGISTRATIONSTATUSACCEPTED, nil
	case "Reject":
		return constants.PREREGISTRATIONSTATUSREJECTED, nil
	default:
		return "", errors.New("Invalid Status")
	}
}

//PaymentPreregistration : ""
func (s *Service) PaymentPreregistration(ctx *models.Context, prp *models.PreregistrationPayment) error {

	psc := new(models.PreregistrationStatusChange)

	psc.ApplicantID = prp.ApplicantID
	psc.Status = constants.PREREGISTRATIONSTATUSPAYMENTSCOMPLETED
	psc.Remarks = prp.Remarks
	psc.On = prp.On
	psc.By = prp.By
	psc.ByType = prp.ByType
	psc.ByName = prp.ByName
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		psc.On = &t
		dberr := s.Daos.PreregistrationStatusChange(ctx, psc)
		if dberr != nil {

			return errors.New("Transaction Aborted - " + dberr.Error())
		}

		refApplicant, err := s.Daos.GetSinglePreregistrationWithUniqueID(ctx, psc.ApplicantID)
		if err != nil {
			return errors.New("Not able to not able to find applicant - " + err.Error())
		}

		//Create user
		user := new(models.User)
		user.Name = refApplicant.Name
		user.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
		user.Mobile = refApplicant.MobileNumber
		user.Email = refApplicant.Email
		user.Password = constants.DEFAULTPASSWORD
		user.Type = constants.USERTYPEAPPLICANT
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		user.Created = created
		user.Status = constants.USERSTATUSACTIVE
		if err := s.Daos.SaveUser(ctx, user); err != nil {
			return errors.New("Not able to create user - " + dberr.Error())
		}

		//Log User Created
		pscUserCreated := new(models.PreregistrationStatusChange)
		pscUserCreated.ApplicantID = prp.ApplicantID
		pscUserCreated.Status = constants.PREREGISTRATIONSTATUSUSERCREATED
		pscUserCreated.Remarks = prp.Remarks
		pscUserCreated.On = prp.On
		pscUserCreated.By = constants.SYSTEM
		pscUserCreated.ByType = constants.SYSTEM
		pscUserCreated.ByName = constants.SYSTEM
		dberr = s.Daos.PreregistrationStatusChange(ctx, psc)
		if dberr != nil {
			return errors.New("User Creation log not created - " + dberr.Error())
		}
		refApplicant.Applicant.UserName = user.UserName
		refApplicant.Applicant.Status = constants.APPLICANTSTATUSACTIVE
		dberr = s.Daos.SaveApplicant(ctx, &refApplicant.Applicant)
		if dberr != nil {
			return errors.New("Applicant Creation not created - " + dberr.Error())
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		emailData := make(map[string]interface{})
		emailData["name"] = user.Name
		emailData["username"] = user.UserName
		emailData["password"] = user.Password
		s.SendEmailWithTemplate("Welcome to BPMS", []string{refApplicant.Applicant.Email}, "emailtemplates/registration-completed-v2.html", emailData)
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

//PaymentPendingNoticeForPreRegistration : ""
func (s *Service) PaymentPendingNoticeForPreRegistration(ctx *models.Context, ID string) error {
	preregistration, err := s.Daos.GetSinglePreregistration(ctx, ID)
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	data["name"] = preregistration.Name
	data["amount"] = "1000"
	s.SendEmailWithTemplate("Registration payment pending", []string{preregistration.Email}, "emailtemplates/pending-payment-email-template.html", data)

	return nil
}
