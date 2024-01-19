package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePunchIn :""
func (s *Service) SavePunchIn(ctx *models.Context, user *models.UserAttendanceAction) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	punchIn := new(models.UserAttendance)

	t := time.Now()
	date := t.Year()
	month := t.Month()
	year := t.Year()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM

	punchIn.UniqueID = fmt.Sprintf("%v%v%v", date, "-", month, "-", year)
	// punchIn.UniqueID = fmt.Fprintf(("%v%v%v" , t.Date() , "-" , t.Month() , "-" , t.Year())
	user.UserName = punchIn.UserName
	user.Image = punchIn.PunchIn.Image
	user.Location = punchIn.PunchIn.Location
	log.Println("b4 shoprent.created")

	log.Println("b4 shoprent.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePunchIn(ctx, user)
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

//SavePunchOut : ""
func (s *Service) SavePunchOut(ctx *models.Context, user *models.UserAttendance) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.SavePunchOut(ctx, user)
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
