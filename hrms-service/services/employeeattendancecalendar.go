package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveEmployeeAttendanceCalendar : ""
func (s *Service) SaveEmployeeAttendanceCalendar(ctx *models.Context, employeeAttendanceCalendar *models.EmployeeAttendanceCalendar) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	employeeAttendanceCalendar.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR)
	employeeAttendanceCalendar.Status = constants.EMPLOYEEATTENDANCECALENDARSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 EmployeeAttendanceCalendar.created")
	employeeAttendanceCalendar.Created = &created
	log.Println("b4 EmployeeAttendanceCalendar.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveEmployeeAttendanceCalendar(ctx, employeeAttendanceCalendar)
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

// GetSingleEmployeeAttendanceCalendar : ""
func (s *Service) GetSingleEmployeeAttendanceCalendar(ctx *models.Context, UniqueID string) (*models.RefEmployeeAttendanceCalendar, error) {
	employeeAttendanceCalendar, err := s.Daos.GetSingleEmployeeAttendanceCalendar(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return employeeAttendanceCalendar, nil
}

//UpdateEmployeeAttendanceCalendar : ""
func (s *Service) UpdateEmployeeAttendanceCalendar(ctx *models.Context, employeeAttendanceCalendar *models.EmployeeAttendanceCalendar) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateEmployeeAttendanceCalendar(ctx, employeeAttendanceCalendar)
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

// EnableEmployeeAttendanceCalendar : ""
func (s *Service) EnableEmployeeAttendanceCalendar(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableEmployeeAttendanceCalendar(ctx, uniqueID)
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

// DisableEmployeeAttendanceCalendar : ""
func (s *Service) DisableEmployeeAttendanceCalendar(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableEmployeeAttendanceCalendar(ctx, uniqueID)
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

//DeleteEmployeeAttendanceCalendar : ""
func (s *Service) DeleteEmployeeAttendanceCalendar(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteEmployeeAttendanceCalendar(ctx, UniqueID)
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

// FilterEmployeeAttendanceCalendar : ""
func (s *Service) FilterEmployeeAttendanceCalendar(ctx *models.Context, employeeAttendanceCalendar *models.FilterEmployeeAttendanceCalendar, pagination *models.Pagination) (employeeAttendanceCalendars []models.RefEmployeeAttendanceCalendar, err error) {
	return s.Daos.FilterEmployeeAttendanceCalendar(ctx, employeeAttendanceCalendar, pagination)
}
func (s *Service) GetSingleEmployeeAttendanceCalendarWithCurrentMonth(ctx *models.Context) (*models.RefEmployeeAttendanceCalendar, error) {
	employeeAttendanceCalendar, err := s.Daos.GetSingleEmployeeAttendanceCalendarWithCurrentMonth(ctx)
	if err != nil {
		return nil, err
	}
	return employeeAttendanceCalendar, nil
}
