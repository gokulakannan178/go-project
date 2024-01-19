package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveTask : ""
func (s *Service) SaveTask(ctx *models.Context, tk *models.Task) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	tk.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTASK)
	tk.Status = constants.TASKSTATUSACTIVE
	t := time.Now()
	tk.Created.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTask(ctx, tk)
		if dberr != nil {
			return dberr
		}
		// calling taskmember daos
		if len(tk.TeamMember) > 0 {
			var tms []models.TaskMember
			for _, v := range tk.TeamMember {
				var tm models.TaskMember
				tm.Status = constants.TASKSTATUSACTIVE
				tm.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTASKMEMBER)
				tm.TaskID = tk.UniqueID
				tm.UserName = v
				tms = append(tms, tm)
			}
			dberr := s.Daos.SaveTaskMembers(ctx, tms)
			if dberr != nil {
				log.Println("err in abort out")
				return errors.New("Transaction Aborted - " + dberr.Error())
			}
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

// SaveTaskTeamMember : ""
func (s *Service) SaveTaskTeamMember(ctx *models.Context, tm *models.TaskMember) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	tm.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTASKMEMBER)
	tm.Status = constants.TASKTEAMMEMBERSTATUSACTIVE
	//user.Password = "#nature32" //Default Password
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	tm.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTaskTeamMember(ctx, tm)
		if dberr != nil {

			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
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

// DisableTaskTeamMember : ""
func (s *Service) DisableTaskTeamMember(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTaskTeamMember(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return err
	}
	return nil
}

// GetSingleTask : ""
func (s *Service) GetSingleTask(ctx *models.Context) (*models.RefTask, error) {
	uniqueID := "1"
	return s.Daos.GetSingleTask(ctx, uniqueID)
}

// EnableTask : ""
func (s *Service) EnableTask(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableTask(ctx, uniqueID)
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

// DisableTask : ""
func (s *Service) DisableTask(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableTask(ctx, uniqueID)
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

// DeleteTask : ""
func (s *Service) DeleteTask(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DeleteTask(ctx, uniqueID)
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

// UpdateTask : ""
func (s *Service) UpdateTask(ctx *models.Context, particular *models.Task) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.UpdateTask(ctx, particular)
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

// FilterTask : ""
func (s *Service) FilterTask(ctx *models.Context, ft *models.FilterTask, pagination *models.Pagination) (property []models.RefTask, err error) {
	return s.Daos.FilterTask(ctx, ft, pagination)
}
