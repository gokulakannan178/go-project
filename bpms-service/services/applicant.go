package services

import (
	"bpms-service/constants"
	"bpms-service/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveApplicant :""
func (s *Service) SaveApplicant(ctx *models.Context, applicant *models.Applicant) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	applicant.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONAPPLICANT)
	applicant.Status = constants.APPLICANTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 applicant.created")
	applicant.Created = created
	log.Println("b4 applicant.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveApplicant(ctx, applicant)
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

//UpdateApplicant : ""
func (s *Service) UpdateApplicant(ctx *models.Context, applicant *models.Applicant) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateApplicant(ctx, applicant)
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

//EnableApplicant : ""
func (s *Service) EnableApplicant(ctx *models.Context, UserName string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableApplicant(ctx, UserName)
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

//DisableApplicant : ""
func (s *Service) DisableApplicant(ctx *models.Context, UserName string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableApplicant(ctx, UserName)
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

//DeleteApplicant : ""
func (s *Service) DeleteApplicant(ctx *models.Context, UserName string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteApplicant(ctx, UserName)
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

//GetSingleApplicant :""
func (s *Service) GetSingleApplicant(ctx *models.Context, UserName string) (*models.RefApplicant, error) {
	applicant, err := s.Daos.GetSingleApplicant(ctx, UserName)
	if err != nil {
		return nil, err
	}
	return applicant, nil
}

//FilterApplicant :""
func (s *Service) FilterApplicant(ctx *models.Context, applicantfilter *models.ApplicantFilter, pagination *models.Pagination) (applicant []models.RefApplicant, err error) {
	return s.Daos.FilterApplicant(ctx, applicantfilter, pagination)
}

//BlacklistApplicant : ""
func (s *Service) BlacklistApplicant(ctx *models.Context, asc *models.ApplicantStatusChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	t := time.Now()
	asc.Info.On = &t
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.BlacklistApplicant(ctx, asc)
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

//LicenseCancelApplicant : ""
func (s *Service) LicenseCancelApplicant(ctx *models.Context, asc *models.ApplicantStatusChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	t := time.Now()
	asc.Info.On = &t
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.LicenseCancelApplicant(ctx, asc)
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

//ReActivateApplicant : ""
func (s *Service) ReActivateApplicant(ctx *models.Context, asc *models.ApplicantStatusChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	t := time.Now()
	asc.Info.On = &t
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.ReActivateApplicant(ctx, asc)
		if dberr != nil {
			if err := ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}
