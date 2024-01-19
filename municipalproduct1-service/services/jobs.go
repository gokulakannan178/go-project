package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveJob : ""
func (s *Service) SaveJob(ctx *models.Context, job *models.Job) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	job.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOB)
	job.Status = constants.JOBSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	job.Created = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveJob(ctx, job)
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

//GetSingleJob :""
func (s *Service) GetSingleJob(ctx *models.Context, UniqueID string) (*models.RefJob, error) {
	tower, err := s.Daos.GetSingleJob(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateJob : ""
func (s *Service) UpdateJob(ctx *models.Context, job *models.Job) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateJob(ctx, job)
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

// EnableJob : ""
func (s *Service) EnableJob(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableJob(ctx, UniqueID)
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

//DisableJob : ""
func (s *Service) DisableJob(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableJob(ctx, UniqueID)
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

//DeleteJob : ""
func (s *Service) DeleteJob(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteJob(ctx, UniqueID)
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

// FilterJob : ""
func (s *Service) FilterJob(ctx *models.Context, filter *models.JobFilter, pagination *models.Pagination) ([]models.RefJob, error) {
	return s.Daos.FilterJob(ctx, filter, pagination)

}

// ExecuteJob : ""
func (s *Service) ExecuteJob(ctx *models.Context, JobID string) error {
	res, err := s.Daos.GetSingleJob(ctx, JobID)
	if err != nil {
		return err
	}
	joblog := new(models.JobLog)
	joblog.JobID = JobID
	joblog.Title = res.Title
	t := time.Now()
	joblog.StartTime = &t
	joblog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONJOB)
	joblog.Status = constants.JOBLOGSTATUSRUNNING
	created := models.CreatedV2{}
	created.On = &t
	joblog.Created = created
	err = s.Daos.SaveJobLog(ctx, joblog)
	if err != nil {
		return err
	}
	var err1 error
	switch JobID {
	case "UpdateOverallDemandForAll":
		err1 = s.UpdateOverallDemandForAll(ctx, []string{"Active"})
	}
	if err1 != nil {
		joblog.ErrorMsg = err1.Error()
		joblog.Status = "Error"
	} else {
		joblog.Status = constants.JOBLOGSTATUSSUCCESS
	}
	t1 := time.Now()
	joblog.EndTime = &t1
	err = s.Daos.UpdateJobLog(ctx, joblog)
	if err != nil {
		return err
	}
	return nil
}
